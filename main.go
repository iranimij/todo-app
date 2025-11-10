package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Task   string
	Done   bool
}

var todos []Todo

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n📋 Simple To-Do List App")
		fmt.Println("----------------------------")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			listTasks()
		case "3":
			markDone(reader)
		case "4":
			deleteTask(reader)
		case "5":
			fmt.Println("👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice, please try again.")
		}
	}
}

func addTask(reader *bufio.Reader) {
	fmt.Print("Enter new task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	if task == "" {
		fmt.Println("⚠️ Task cannot be empty!")
		return
	}

	todos = append(todos, Todo{Task: task})
	fmt.Println("✅ Task added successfully!")
}

func listTasks() {
	if len(todos) == 0 {
		fmt.Println("📭 No tasks yet.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for i, todo := range todos {
		status := "❌"
		if todo.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, todo.Task, status)
	}
}

func markDone(reader *bufio.Reader) {
	if len(todos) == 0 {
		fmt.Println("📭 No tasks to mark.")
		return
	}

	listTasks()
	fmt.Print("Enter task number to mark as done: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(todos) {
		fmt.Println("⚠️ Invalid task number!")
		return
	}

	todos[num-1].Done = true
	fmt.Println("✅ Task marked as done!")
}

func deleteTask(reader *bufio.Reader) {
	if len(todos) == 0 {
		fmt.Println("📭 No tasks to delete.")
		return
	}

	listTasks()
	fmt.Print("Enter task number to delete: ")
	var num int
	fmt.Scanln(&num)

	if num < 1 || num > len(todos) {
		fmt.Println("⚠️ Invalid task number!")
		return
	}

	todos = append(todos[:num-1], todos[num:]...)
	fmt.Println("🗑️ Task deleted.")
}
