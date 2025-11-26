import { useState, FormEvent } from 'react';

interface AddTodoFormProps {
  onAdd: (task: string) => void;
  isLoading: boolean;
}

export const AddTodoForm = ({ onAdd, isLoading }: AddTodoFormProps) => {
  const [task, setTask] = useState('');

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (task.trim()) {
      onAdd(task.trim());
      setTask('');
    }
  };

  return (
    <form className="add-todo-form" onSubmit={handleSubmit}>
      <input
        type="text"
        className="todo-input"
        placeholder="What needs to be done?"
        value={task}
        onChange={(e) => setTask(e.target.value)}
        disabled={isLoading}
      />
      <button
        type="submit"
        className="btn btn-add"
        disabled={isLoading || !task.trim()}
      >
        Add Task
      </button>
    </form>
  );
};