package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo-app/todo"
)

func StartCLI() {
	err := todo.Initialize()
	if err != nil {
		fmt.Println("❌ Failed to initialize database:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n📋 To-Do List (SQLite)")
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

	if err := todo.AddTask(task); err != nil {
		fmt.Println("❌ Failed to add task:", err)
	} else {
		fmt.Println("✅ Task added successfully!")
	}
}

func listTasks() {
	tasks, err := todo.GetAll()
	if err != nil {
		fmt.Println("❌ Failed to load tasks:", err)
		return
	}
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks yet.")
		return
	}

	for _, t := range tasks {
		status := "❌"
		if t.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", t.ID, t.Task, status)
	}
}

func markDone(reader *bufio.Reader) {
	listTasks()
	fmt.Print("Enter task ID to mark as done: ")
	var id int
	fmt.Scanln(&id)
	if err := todo.MarkDone(id); err != nil {
		fmt.Println("⚠️ Failed to mark task:", err)
	} else {
		fmt.Println("✅ Task marked as done!")
	}
}

func deleteTask(reader *bufio.Reader) {
	listTasks()
	fmt.Print("Enter task ID to delete: ")
	var id int
	fmt.Scanln(&id)
	if err := todo.Delete(id); err != nil {
		fmt.Println("⚠️ Failed to delete task:", err)
	} else {
		fmt.Println("🗑️ Task deleted.")
	}
}
