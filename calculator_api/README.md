# Calculator API

A simple HTTP API built with Go's standard library to perform basic arithmetic operations.

## Learning Objectives

This project demonstrates:
- Building HTTP servers using Go's `net/http` package
- Creating RESTful endpoints with the `http.ServeMux` router
- Parsing and validating JSON request bodies
- Handling errors and returning structured JSON responses
- Organizing handlers into separate files

## Endpoints

The API runs on `localhost:3000` and supports the following POST endpoints:

- `/add` - Add two numbers
- `/subtract` - Subtract two numbers
- `/multiply` - Multiply two numbers
- `/divide` - Divide two numbers
- `/sum` - Sum multiple numbers

## Request Format

Most endpoints accept two numbers:

```json
{
  "number1": 10,
  "number2": 5
}
```

The `/sum` endpoint accepts an array of numbers:

```json
{
  "numbers": [1, 2, 3, 4, 5]
}
```

## Getting Started

```bash
go run ./cmd/*
```

The server will start on port 3000.
