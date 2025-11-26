import { useState, useEffect } from 'react';
import { todoApi } from './services/api';
import { Todo } from './types/todo';
import { AddTodoForm } from './components/AddTodoForm';
import { TodoList } from './components/TodoList';

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchTodos = async () => {
    try {
      setError(null);
      const data = await todoApi.getAll();
      setTodos(data);
    } catch (err) {
      setError('Failed to load todos');
      console.error(err);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const handleAddTodo = async (task: string) => {
    try {
      setIsLoading(true);
      setError(null);
      await todoApi.create({ task });
      await fetchTodos();
    } catch (err) {
      setError('Failed to add todo');
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleMarkDone = async (id: number) => {
    try {
      setError(null);
      await todoApi.markDone(id);
      await fetchTodos();
    } catch (err) {
      setError('Failed to mark todo as done');
      console.error(err);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      setError(null);
      await todoApi.delete(id);
      await fetchTodos();
    } catch (err) {
      setError('Failed to delete todo');
      console.error(err);
    }
  };

  return (
    <div className="app">
      <div className="container">
        <header className="app-header">
          <h1>Todo App</h1>
          <p className="subtitle">Stay organized and productive</p>
        </header>

        {error && (
          <div className="error-message">
            {error}
          </div>
        )}

        <AddTodoForm onAdd={handleAddTodo} isLoading={isLoading} />

        <TodoList
          todos={todos}
          onMarkDone={handleMarkDone}
          onDelete={handleDelete}
        />

        <footer className="app-footer">
          <p>{todos.length} total tasks</p>
        </footer>
      </div>
    </div>
  );
}

export default App;