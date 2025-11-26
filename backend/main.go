package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"todo-app/api"
	"todo-app/todo"
	"todo-app/ui"
)

func main() {
	// Initialize database
	if err := todo.Initialize(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Check if running in an interactive terminal
	fileInfo, _ := os.Stdin.Stat()
	isInteractive := (fileInfo.Mode() & os.ModeCharDevice) != 0

	if isInteractive {
		// Start API server in background goroutine with ready channel
		ready := make(chan string)
		go startAPIServerAsync(ready)

		// Wait for API server to be ready
		<-ready

		// Display web interface URL
		fmt.Println("\n🌐 Web Interface: http://localhost:3000")
		fmt.Println()

		// Start CLI in main goroutine (interactive mode)
		ui.StartCLI()
	} else {
		// Just run API server (non-interactive mode)
		fmt.Println("Running in non-interactive mode (API only)")
		startAPIServer()
	}
}

func startAPIServerAsync(ready chan string) {
	http.HandleFunc("/todos", api.HandleTodos)
	http.HandleFunc("/todos/", api.HandleTodoByID)

	// Try to find an available port
	startPort := 8080
	maxAttempts := 10

	for i := 0; i < maxAttempts; i++ {
		currentPort := startPort + i
		addr := fmt.Sprintf(":%d", currentPort)

		// Check if port is available
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			continue
		}

		// Signal that server is ready with the port
		ready <- addr

		// Start serving
		if err := http.Serve(listener, nil); err != nil {
			log.Fatal("Failed to start server:", err)
		}
		return
	}

	log.Fatal("Failed to start server after trying multiple ports")
}

func startAPIServer() {
	http.HandleFunc("/todos", api.HandleTodos)
	http.HandleFunc("/todos/", api.HandleTodoByID)

	// Try to find an available port
	startPort := 8080
	maxAttempts := 10

	for i := 0; i < maxAttempts; i++ {
		currentPort := startPort + i
		addr := fmt.Sprintf(":%d", currentPort)

		// Check if port is available
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Printf("⚠️  Port %d is in use, trying next port...\n", currentPort)
			continue
		}

		fmt.Printf("🚀 Todo API server starting on http://localhost:%d\n", currentPort)
		fmt.Println("📝 Endpoints:")
		fmt.Println("  GET    /todos          - Get all todos")
		fmt.Println("  POST   /todos          - Create a new todo")
		fmt.Println("  PUT    /todos/{id}/done - Mark todo as done")
		fmt.Println("  DELETE /todos/{id}      - Delete a todo")
		fmt.Println()

		// Start serving
		if err := http.Serve(listener, nil); err != nil {
			log.Fatal("Failed to start server:", err)
		}
		return
	}

	log.Fatal("Failed to start server after trying multiple ports")
}
