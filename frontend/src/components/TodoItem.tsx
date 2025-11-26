import { Todo } from '../types/todo';

interface TodoItemProps {
  todo: Todo;
  onMarkDone: (id: number) => void;
  onDelete: (id: number) => void;
}

export const TodoItem = ({ todo, onMarkDone, onDelete }: TodoItemProps) => {
  return (
    <div className={`todo-item ${todo.done ? 'done' : ''}`}>
      <div className="todo-content">
        <span className="todo-id">#{todo.id}</span>
        <span className={`todo-task ${todo.done ? 'completed' : ''}`}>
          {todo.task}
        </span>
        <span className={`todo-status ${todo.done ? 'status-done' : 'status-pending'}`}>
          {todo.done ? '✓' : '○'}
        </span>
      </div>
      <div className="todo-actions">
        {!todo.done && (
          <button
            className="btn btn-done"
            onClick={() => onMarkDone(todo.id)}
          >
            Mark Done
          </button>
        )}
        <button
          className="btn btn-delete"
          onClick={() => onDelete(todo.id)}
        >
          Delete
        </button>
      </div>
    </div>
  );
};