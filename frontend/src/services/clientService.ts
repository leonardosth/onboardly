import axios from 'axios';
import type { Cliente } from '../types';

const api = axios.create({
  baseURL: '/api',
});

export const clientService = {
  async getAll(): Promise<Cliente[]> {
    const { data } = await api.get<Cliente[]>('/clientes');
    return data || [];
  },

  async getById(id: string): Promise<Cliente> {
    const { data } = await api.get<Cliente>(`/clientes/${id}`);
    return data;
  },

  async create(cliente: Partial<Cliente>): Promise<Cliente> {
    const { data } = await api.post<Cliente>('/clientes', cliente);
    return data;
  },

  async update(id: string, cliente: Partial<Cliente>): Promise<Cliente> {
    const { data } = await api.put<Cliente>(`/clientes/${id}`, cliente);
    return data;
  },

  async delete(id: string): Promise<void> {
    await api.delete(`/clientes/${id}`);
  }
};
