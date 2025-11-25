package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var tasks []TodoItem

func main() {
	command := os.Args[1]

	loadedTasks, err := LoadTasks()
	if err != nil {
		panic(err)
	}
	tasks = loadedTasks

	switch command {
	case "list":
		listTasks()
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error!")
			fmt.Println("Usage: main.go add <task description>")
			break
		}

		var sb strings.Builder
		for i := 2; i < len(os.Args); i++ {
			if i > 2 {
				sb.WriteString(" ")
			}
			sb.WriteString(os.Args[i])
		}

		task := sb.String()
		addTask(task)
		fmt.Println(tasks)
	case "complete":
		if (len(os.Args)) != 3 {
			fmt.Println("Error!")
			fmt.Println("Usage: main.go complete <taskId>")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println(err.Error())
			return
		}

		completeTask(uint8(id))
	case "delete":
		if (len(os.Args)) != 3 {
			fmt.Println("Error!")
			fmt.Println("Usage: main.go delete <taskId>")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println(err.Error())
			return
		}
		deleteTask(uint8(id))

	default:
		fmt.Println("Usage: main.go list | add | complete")
	}

	// save tasks before exiting
	SaveTasks(tasks)
}

func listTasks() {
	fmt.Printf("ID  Description  Status\n")
	for _, v := range tasks {
		fmt.Printf("%v  %v  %v\n", v.ID, v.Task, v.Status)
	}
}

func addTask(task string) {
	id := uint8(rand.Intn(100000))
	todoItem := TodoItem{Task: task, Status: StatusPending, ID: id}
	tasks = append(tasks, todoItem)
}

func completeTask(id uint8) {
	found := false

	for i := range tasks {
		if tasks[i].ID == id {
			found = true
			tasks[i].Status = StatusComplete
		}
	}

	if !found {
		fmt.Printf("Tasks with id %v not found", id)
	}
}

func deleteTask(id uint8) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}
