# Go Learning Projects

A collection of projects built while learning Go, focusing on fundamental concepts and practical applications.

## Projects

### 1. Calculator API
An HTTP API built with Go's standard library demonstrating RESTful endpoint design.

**Key Concepts:**
- HTTP servers and routing with `http.ServeMux`
- JSON parsing and validation
- Error handling and HTTP status codes
- Handler organization

**Run:** `go run ./cmd/*`

**Endpoints:**
- `POST /add` - Add two numbers
- `POST /subtract` - Subtract two numbers
- `POST /multiply` - Multiply two numbers
- `POST /divide` - Divide two numbers
- `POST /sum` - Sum an array of numbers

---

### 2. Go Todo App
A command-line todo application with persistent JSON storage.

**Key Concepts:**
- Interactive CLI applications
- JSON serialization and file I/O
- Struct types and enums
- User input handling and command processing
- Data persistence

**Run:** `go run ./cmd/*`

**Commands:**
- `list` - Display all tasks
- `add` - Create a new task
- `complete` - Mark task as complete
- `delete` - Remove a task
- `exit` - Quit the application

---

### 3. Go Tutorial
Basic Go fundamentals and syntax examples.

**Key Concepts:**
- Hello World program
- Go basics and syntax

---

### 4. Guess the Number
A simple number guessing game.

**Key Concepts:**
- Random number generation
- User input and loops
- Game logic

---

## Learning Path

These projects progress from basic HTTP APIs to more interactive CLI applications, building understanding of:
1. Go's standard library and HTTP capabilities
2. JSON handling and data persistence
3. CLI interaction patterns
4. Code organization and package structure
