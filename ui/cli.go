package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo-app/todo"
)

// StartCLI runs the main command-line interface loop
func StartCLI() {
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

	todo.AddTask(task)
	fmt.Println("✅ Task added successfully!")
}

func listTasks() {
	tasks := todo.GetAll()
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks yet.")
		return
	}

	for i, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, t.Task, status)
	}
}

func markDone(reader *bufio.Reader) {
	listTasks()
	fmt.Print("Enter task number to mark as done: ")
	var num int
	fmt.Scanln(&num)
	if todo.MarkDone(num - 1) {
		fmt.Println("✅ Task marked as done!")
	} else {
		fmt.Println("⚠️ Invalid task number.")
	}
}

func deleteTask(reader *bufio.Reader) {
	listTasks()
	fmt.Print("Enter task number to delete: ")
	var num int
	fmt.Scanln(&num)
	if todo.Delete(num - 1) {
		fmt.Println("🗑️ Task deleted.")
	} else {
		fmt.Println("⚠️ Invalid task number.")
	}
}
