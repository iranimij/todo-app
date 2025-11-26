import { Todo, CreateTodoRequest, ApiResponse } from '../types/todo';

const API_BASE_URL = '/api';

export const todoApi = {
  async getAll(): Promise<Todo[]> {
    const response = await fetch(`${API_BASE_URL}/todos`);
    if (!response.ok) {
      throw new Error('Failed to fetch todos');
    }
    return response.json();
  },

  async create(data: CreateTodoRequest): Promise<ApiResponse> {
    const response = await fetch(`${API_BASE_URL}/todos`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      throw new Error('Failed to create todo');
    }
    return response.json();
  },

  async markDone(id: number): Promise<ApiResponse> {
    const response = await fetch(`${API_BASE_URL}/todos/${id}/done`, {
      method: 'PUT',
    });
    if (!response.ok) {
      throw new Error('Failed to mark todo as done');
    }
    return response.json();
  },

  async delete(id: number): Promise<ApiResponse> {
    const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) {
      throw new Error('Failed to delete todo');
    }
    return response.json();
  },
};