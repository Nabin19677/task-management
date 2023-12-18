# Task Management

This is a Task Management web application powered by Go.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Nabin19677/task-management
   ```

2. Navigate to the project directory:

   ```bash
   cd task-management
   ```

3. Create conf.toml file in conf folder.

   ```bash
    DATABASE_SOURCE = "postgres://postgres:2020@localhost:5432/taskmanager?sslmode=disable"
    SERVER_PORT = "8080"
    JWT_SECRET = "thisisasecret"
    JWT_ISSUER = "leapfrog"
    EMAIL = "###@yahoo.com"
    EMAIL_PASSWORD = "#####"
   ```

   > **Note:** Get Email and Password from Yahoo Mail

4. Build and install the CLI:

   ```bash
   go run cmd/server/main.go
   ```

## Usage

Go to http://localhost:8080
