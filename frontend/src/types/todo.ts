export interface Todo {
  id: number;
  task: string;
  done: boolean;
}

export interface CreateTodoRequest {
  task: string;
}

export interface ApiResponse {
  message: string;
}

export interface ErrorResponse {
  error: string;
}