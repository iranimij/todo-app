import { Todo } from '../types/todo';
import { TodoItem } from './TodoItem';

interface TodoListProps {
  todos: Todo[];
  onMarkDone: (id: number) => void;
  onDelete: (id: number) => void;
}

export const TodoList = ({ todos, onMarkDone, onDelete }: TodoListProps) => {
  if (todos.length === 0) {
    return (
      <div className="empty-state">
        <p>No todos yet. Add one above to get started!</p>
      </div>
    );
  }

  const pendingTodos = todos.filter(todo => !todo.done);
  const completedTodos = todos.filter(todo => todo.done);

  return (
    <div className="todo-list">
      {pendingTodos.length > 0 && (
        <div className="todo-section">
          <h3 className="section-title">Pending ({pendingTodos.length})</h3>
          {pendingTodos.map(todo => (
            <TodoItem
              key={todo.id}
              todo={todo}
              onMarkDone={onMarkDone}
              onDelete={onDelete}
            />
          ))}
        </div>
      )}

      {completedTodos.length > 0 && (
        <div className="todo-section">
          <h3 className="section-title">Completed ({completedTodos.length})</h3>
          {completedTodos.map(todo => (
            <TodoItem
              key={todo.id}
              todo={todo}
              onMarkDone={onMarkDone}
              onDelete={onDelete}
            />
          ))}
        </div>
      )}
    </div>
  );
};