# Go Todo App

A command-line todo application built with Go that stores tasks in a JSON file.

## Learning Objectives

This project demonstrates:
- Building interactive CLI applications with Go
- Reading and parsing JSON files
- Writing data to persistent JSON storage
- Organizing code into separate files and packages
- Working with structs and enums in Go
- Handling user input and command processing

## Features

- **List** - Display all tasks with their status
- **Add** - Create a new task
- **Complete** - Mark a task as complete
- **Delete** - Remove a task
- **Exit** - Quit the application

## Task Status

Tasks can have the following statuses:
- `pending` - Not yet started
- `active` - Currently in progress
- `complete` - Finished

## Getting Started

```bash
go run ./cmd/*
```

This will start the interactive CLI. The app stores all tasks in `todo.json` for persistence across sessions.

## Usage Example

```
Your command: add
Task description: Learn Go basics
(Task added and saved)

Your command: list
(Shows all tasks in a formatted table)

Your command: complete
(Prompts to select a task to mark as complete)

Your command: delete
(Prompts to select a task to delete)

Your command: exit
(Saves and exits)
```
