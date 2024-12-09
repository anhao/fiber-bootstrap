# Fiber Bootstrap
Fiber bootstrap for rapid development using Go-Fiber / Gorm / Validator.

## Components

- fiber
    - Html Engine Template
    - Logger
    - Monitoring
- Gorm
    - MYSQL Driver
- Validator
- Env File

## Router

API Router `/api` with rate limiter middleware Http Router `/` with CORS and CSRF middleware

## Setup

-  Copy the example env file over:

```shell
cp .env.example .env
```

- Modify the env file you just copied .env with the correct credentials for your database. Make sure the database you
   entered in DB_NAME has been created.

- Run the API:

```shell
go run main.go
```

Your api should be running at http://localhost:3000/ if the port is in use you may modify it in the .env you just
created.