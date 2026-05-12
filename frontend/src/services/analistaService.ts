import axios from 'axios';
import type { Analista } from '../types';

const api = axios.create({
  baseURL: '/api',
});

export const analistaService = {
  async getAll(): Promise<Analista[]> {
    const { data } = await api.get<Analista[]>('/analistas');
    return data || [];
  },

  async getById(id: string): Promise<Analista> {
    const { data } = await api.get<Analista>(`/analistas/${id}`);
    return data;
  },

  async create(analista: Partial<Analista>): Promise<Analista> {
    const { data } = await api.post<Analista>('/analistas', analista);
    return data;
  },

  async update(id: string, analista: Partial<Analista>): Promise<Analista> {
    const { data } = await api.put<Analista>(`/analistas/${id}`, analista);
    return data;
  },

  async delete(id: string): Promise<void> {
    await api.delete(`/analistas/${id}`);
  }
};
