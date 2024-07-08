import requests

from config.app import BASE_URL
from handler.response import print_response


def register():
    name = input("Enter user name: ")
    balance = float(input("Enter initial balance: "))
    response = requests.post(f"{BASE_URL}/user/register", json={"name": name, "balance": balance})
    print_response(response)


def deposit():
    user_id = int(input("Enter user ID: "))
    amount = float(input("Enter deposit amount: "))
    response = requests.post(f"{BASE_URL}/user/deposit", json={"id": user_id, "amount": amount})
    print_response(response)


def withdraw():
    user_id = int(input("Enter user ID: "))
    amount = float(input("Enter withdrawal amount: "))
    response = requests.post(f"{BASE_URL}/user/withdraw", json={"id": user_id, "amount": amount})
    print_response(response)


def bet():
    user_id = int(input("Enter user ID: "))
    amount = float(input("Enter bet amount: "))
    response = requests.post(f"{BASE_URL}/user/bet", json={"id": user_id, "amount": amount})
    print_response(response)
