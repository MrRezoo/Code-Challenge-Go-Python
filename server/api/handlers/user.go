package handlers

import (
	"encoding/json"
	"errors"
	"github.com/MrRezoo/code-challenge/api/helpers"
	"github.com/MrRezoo/code-challenge/models"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type UserHandler struct{ validate *validator.Validate }

func NewUserHandler() *UserHandler {
	return &UserHandler{
		validate: validator.New(),
	}
}

func (h *UserHandler) Users(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := helpers.GenerateBaseResponse(users, true, http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.validate.Struct(user)
	if err != nil {
		response := helpers.GenerateBaseResponseWithValidationErrors(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	user.ID, err = models.GetNextUserID()
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = user.Save()
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := helpers.GenerateBaseResponse(user, true, http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID     int     `json:"id" validate:"required"`
		Amount float64 `json:"amount" validate:"required,gt=0"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		response := helpers.GenerateBaseResponseWithValidationErrors(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := models.GetUserByID(req.ID)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusNotFound, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Balance += req.Amount

	err = user.Save()
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := helpers.GenerateBaseResponse(user, true, http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID     int     `json:"id" validate:"required"`
		Amount float64 `json:"amount" validate:"required,gt=0"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		response := helpers.GenerateBaseResponseWithValidationErrors(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := models.GetUserByID(req.ID)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusNotFound, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.Balance < req.Amount {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, errors.New("insufficient balance"))
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Balance -= req.Amount

	err = user.Save()
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := helpers.GenerateBaseResponse(user, true, http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) Bet(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		ID     int     `json:"id" validate:"required"`
		Amount float64 `json:"amount" validate:"required,gt=0"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = h.validate.Struct(req)
	if err != nil {
		response := helpers.GenerateBaseResponseWithValidationErrors(nil, false, http.StatusBadRequest, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := models.GetUserByID(req.ID)
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusNotFound, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.Balance < req.Amount {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusBadRequest, errors.New("insufficient balance"))
		json.NewEncoder(w).Encode(response)
		return
	}

	// Assume the betting logic here. For simplicity, let's assume the user always loses.
	user.Balance -= req.Amount

	err = user.Save()
	if err != nil {
		response := helpers.GenerateBaseResponseWithError(nil, false, http.StatusInternalServerError, err)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := helpers.GenerateBaseResponse(user, true, http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
