package todo

// In-memory list of todos
var todos []Todo

// AddTask creates a new todo item
func AddTask(task string) {
	todos = append(todos, Todo{Task: task})
}

// GetAll returns all tasks
func GetAll() []Todo {
	return todos
}

// MarkDone sets a task as completed by index
func MarkDone(index int) bool {
	if index < 0 || index >= len(todos) {
		return false
	}
	todos[index].Done = true
	return true
}

// Delete removes a task by index
func Delete(index int) bool {
	if index < 0 || index >= len(todos) {
		return false
	}
	todos = append(todos[:index], todos[index+1:]...)
	return true
}

// Count returns number of tasks
func Count() int {
	return len(todos)
}
