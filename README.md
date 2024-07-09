# Code Challenge API

This project is a simple HTTP server built with Go that provides user account management functionalities, such as
registering users, depositing, withdrawing, and placing bets. It uses Redis for data storage and includes a Python CLI
client for interacting with the server.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
- [Running Tests](#running-tests)
- [Docker Support](#docker-support)
- [Contributing](#contributing)
- [License](#license)

## Features

- User registration
- Deposit funds to user accounts
- Withdraw funds from user accounts
- Place bets
- Retrieve all users
- Health check endpoint

## Prerequisites

- Go 1.18+
- Redis
- Python 3.12+
- Docker (optional)

## Installation

### Clone the repository

```sh
git clone https://github.com/MrRezoo/code-challenge.git
cd code-challenge
```

### Server Setup

1. **Install Go dependencies**

    ```sh
    go mod tidy
    ```

2. **Run the server**

    ```sh
    go run cmd/main.go
    ```

### Client Setup

1. **Navigate to the client directory**

    ```sh
    cd client
    ```

2. **Install Python dependencies**

    ```sh
    pip install -r requirements.txt
   ```
   or you can use poetry to install the dependencies
    ```sh
    poetry install
    ```

## Usage

### Running the Server

```sh
go run cmd/main.go
```

### Running the Python CLI

```sh
python main.py
╭─────────────────  Challenge CLI  ──────────────────╮
│ Welcome to the Challenge!                          │
│ Commands available:                                │
│ 1. register - Register a new user                  │
│ 2. deposit - Deposit funds into a user account     │
│ 3. withdraw - Withdraw funds from a user account   │
│ 4. bet - Perform a betting operation (placeholder) │
│ 5. all - Retrieve all users                        │
╰────────────────────────────────────────────────────╯
```

### Available CLI Commands

- `register`: Register a new user
- `deposit`: Deposit funds to a user's account
- `withdraw`: Withdraw funds from a user's account
- `bet`: Place a bet for a user
- `list`: Retrieve and display all users

## API Endpoints

### Health Check

- **URL:** `/health`
- **Method:** `GET`
- **Responses:**
    - `200 OK`

### Retrieve All Users

- **URL:** `/user/all`
- **Method:** `GET`
- **Responses:**
    - `200 OK`

### Register User

- **URL:** `/user/register`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "id": 1,
    "balance": 100.0
  }
  ```
- **Responses:**
    - `200 OK`
    - `400 Bad Request`

### Deposit to User

- **URL:** `/user/deposit`
- **Method:** `PATCH`
- **Request Body:**
  ```json
  {
    "id": 1,
    "amount": 50.0
  }
  ```
- **Responses:**
    - `200 OK`
    - `400 Bad Request`

### Withdraw from User

- **URL:** `/user/withdraw`
- **Method:** `PATCH`
- **Request Body:**
  ```json
  {
    "id": 1,
    "amount": 20.0
  }
  ```
- **Responses:**
    - `200 OK`
    - `400 Bad Request`

### Place a Bet

- **URL:** `/user/bet`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "id": 1,
    "amount": 10.0
  }
  ```
- **Responses:**
    - `200 OK`
    - `400 Bad Request`

## Configuration

Configuration files are located in the `config` directory. Modify the `.yaml` files according to your environment.

- `config-development.yaml`
- `config-docker.yaml`

## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## Docker Support

To run the server using Docker, use the provided `docker-compose` file.

### Build and Run

```sh
docker-compose up --build
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

```

This `README.md` provides comprehensive documentation for your project, detailing its features, installation steps, usage, API endpoints, configuration, and more. Adjust any specific details as needed.