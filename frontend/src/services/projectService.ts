import axios from 'axios';
import type { Projeto } from '../types';

const api = axios.create({
  baseURL: '/api',
});

export const projectService = {
  async getAll(): Promise<Projeto[]> {
    const { data } = await api.get<Projeto[]>('/projetos');
    return data || [];
  },

  async getById(id: string): Promise<Projeto> {
    const { data } = await api.get<Projeto>(`/projetos/${id}`);
    return data;
  },

  async create(projeto: Partial<Projeto>): Promise<Projeto> {
    const { data } = await api.post<Projeto>('/projetos', projeto);
    return data;
  },

  async update(id: string, projeto: Partial<Projeto>): Promise<Projeto> {
    const { data } = await api.put<Projeto>(`/projetos/${id}`, projeto);
    return data;
  },

  async delete(id: string): Promise<void> {
    await api.delete(`/projetos/${id}`);
  },

  async getStats(): Promise<{ total_projetos: number; por_status: Record<string, number> }> {
    const { data } = await api.get('/dashboard/stats');
    return data;
  }
};
