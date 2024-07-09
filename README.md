```markdown
# Code Challenge Project

This project is designed to demonstrate a comprehensive backend system capable of handling user registrations, deposits, withdrawals, and betting transactions. It utilizes a combination of Python and Go, leveraging Docker for containerization to ensure a consistent and isolated environment for development and deployment.

## Features

- **User Registration**: Allows new users to register with an initial balance.
- **User Deposit**: Enables users to deposit funds into their account.
- **User Withdrawal**: Allows users to withdraw funds from their account.
- **User Betting**: Users can place bets, affecting their balance based on the outcome.

## Prerequisites

- Docker and Docker Compose
- Go (for Go services)
- Python (for Python services)

## Installation

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Use Docker Compose to build and start the services:
   ```bash
   docker-compose up --build
   ```

## Usage

The project includes a Postman collection (`docs/Code-Challenge.postman_collection.json`) for easy interaction with the API endpoints. Import this collection into Postman to test the various functionalities.

### Endpoints

- `GET /health/`: Check the health of the service.
- `POST /user/register/`: Register a new user.
- `PATCH /user/deposit/`: Deposit funds into a user's account.
- `PATCH /user/withdraw/`: Withdraw funds from a user's account.
- `POST /user/bet/`: Place a bet.

## Development

This project uses Go modules for managing dependencies. Ensure you have Go installed and configured correctly. For Python services, use a virtual environment and install dependencies from the `requirements.txt` file.

## Testing

Run tests using the standard testing tools in Go and Python. For Go, navigate to the service directory and run:

```bash
go test ./...
```

For Python, ensure pytest is installed and run:

```bash
pytest
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your proposed changes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
```