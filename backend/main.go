package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

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
		// Setup and start frontend automatically
		setupAndStartFrontend()

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

func setupAndStartFrontend() {
	// Get frontend directory path (relative to backend)
	frontendDir := filepath.Join("..", "frontend")

	// Check if frontend directory exists
	if _, err := os.Stat(frontendDir); os.IsNotExist(err) {
		fmt.Println("⚠️  Frontend directory not found, skipping frontend setup")
		return
	}

	// Check if node_modules exists
	nodeModulesPath := filepath.Join(frontendDir, "node_modules")
	if _, err := os.Stat(nodeModulesPath); os.IsNotExist(err) {
		fmt.Println("📦 Installing frontend dependencies...")
		fmt.Println("   This may take a minute on first run...")

		// Run npm install
		cmd := exec.Command("npm", "install")
		cmd.Dir = frontendDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("⚠️  Failed to install frontend dependencies: %v\n", err)
			fmt.Println("   You can manually run: cd ../frontend && npm install")
			return
		}

		fmt.Println("✅ Frontend dependencies installed!")
		fmt.Println()
	}

	// Start frontend dev server in background
	fmt.Println("🚀 Starting frontend dev server...")
	cmd := exec.Command("npm", "run", "dev")
	cmd.Dir = frontendDir

	// Start in background
	if err := cmd.Start(); err != nil {
		fmt.Printf("⚠️  Failed to start frontend dev server: %v\n", err)
		fmt.Println("   You can manually run: cd ../frontend && npm run dev")
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
