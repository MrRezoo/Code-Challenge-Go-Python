package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MrRezoo/code-challenge/db"
)

var ctx = context.Background()

type User struct {
	ID      int     `json:"id" validate:"required"`
	Balance float64 `json:"balance" validate:"required,gte=0"`
}

func (u *User) Save() error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return db.RedisClient.Set(ctx, fmt.Sprintf("user:%d", u.ID), data, 0).Err()
}

func GetNextUserID() (int, error) {
	id, err := db.RedisClient.Incr(ctx, "next_user_id").Result()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func GetUserByID(id int) (*User, error) {
	data, err := db.RedisClient.Get(ctx, fmt.Sprintf("user:%d", id)).Result()
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	iter := db.RedisClient.Scan(ctx, 0, "user:*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		data, err := db.RedisClient.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		var user User
		err = json.Unmarshal([]byte(data), &user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
