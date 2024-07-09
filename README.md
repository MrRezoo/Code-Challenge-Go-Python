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