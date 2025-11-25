package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
)

var tasks []TodoItem

func main() {
	loadedTasks, err := LoadTasks()
	tasks = loadedTasks
	var command string
	fmt.Println("")
	fmt.Println("----------------------------")
	fmt.Println("Welcome to GOTODO!")
	fmt.Println("What would you like to do?")
	fmt.Println("Available commands: list | add | complete | delete | exit")
	fmt.Println("----------------------------")

outer:
	for {
		fmt.Println("")
		fmt.Print("Your command: ")
		fmt.Scanln(&command)

		if err != nil {
			panic(err)
		}

		switch command {
		case "list":
			listTasks()
		case "add":
			fmt.Print("Task description: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err := scanner.Err()
			desc := scanner.Text()
			if err != nil {
				log.Fatal(err)
			}

			addTask(desc)
		case "complete":
			var idStr string
			fmt.Print("Enter task id to mark complete: ")
			fmt.Scanln(&idStr)

			id, err := strconv.Atoi(idStr)
			if err != nil {
				println(err.Error())
				return
			}

			completeTask(uint8(id))
		case "delete":
			var idStr string
			fmt.Print("Enter task id to delete: ")
			fmt.Scanln(&idStr)

			id, err := strconv.Atoi(idStr)
			if err != nil {
				println(err.Error())
				return
			}
			deleteTask(uint8(id))
		case "exit":
			fmt.Println("")
			fmt.Println("Exiting GOTODO")
			fmt.Println("See you soon!")
			break outer
		default:
			fmt.Println("")
			fmt.Println("Unkwon command!")
			fmt.Println("Available commands: list | add | complete | delete | exit")
			fmt.Println("")
		}
	}

	// save tasks before exiting
	SaveTasks(tasks)
}

func listTasks() {
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 0, ' ', ' ')
	for _, v := range tasks {
		fmt.Fprintln(w, v.ID, "\t", v.Task, "\t", v.Status)
	}
	w.Flush()
}

// func listTasks() {
// 	fmt.Printf("ID  Description  Status\n")
// 	for _, v := range tasks {
// 		fmt.Printf("%v  %v  %v\n", v.ID, v.Task, v.Status)
// 	}
// }

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
