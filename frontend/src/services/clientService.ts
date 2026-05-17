import type { Cliente } from '../types';
import { apiFetch } from './api';

export const clientService = {
  async getAll(): Promise<Cliente[]> {
    return apiFetch('/clientes') || [];
  },

  async getById(id: string): Promise<Cliente> {
    return apiFetch(`/clientes/${id}`);
  },

  async create(cliente: Partial<Cliente>): Promise<Cliente> {
    return apiFetch('/clientes', {
      method: 'POST',
      body: JSON.stringify(cliente)
    });
  },

  async update(id: string, cliente: Partial<Cliente>): Promise<Cliente> {
    return apiFetch(`/clientes/${id}`, {
      method: 'PUT',
      body: JSON.stringify(cliente)
    });
  },

  async delete(id: string): Promise<void> {
    return apiFetch(`/clientes/${id}`, {
      method: 'DELETE'
    });
  }
};
