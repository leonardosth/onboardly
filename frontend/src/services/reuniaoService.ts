import axios from 'axios';
import type { Reuniao } from '../types';

const api = axios.create({
  baseURL: '/api',
});

export const reuniaoService = {
  async getAll(): Promise<Reuniao[]> {
    const { data } = await api.get<Reuniao[]>('/reunioes');
    return data || [];
  },

  async getById(id: string): Promise<Reuniao> {
    const { data } = await api.get<Reuniao>(`/reunioes/${id}`);
    return data;
  },

  async create(reuniao: Partial<Reuniao>): Promise<Reuniao> {
    const { data } = await api.post<Reuniao>('/reunioes', reuniao);
    return data;
  },

  async update(id: string, reuniao: Partial<Reuniao>): Promise<Reuniao> {
    const { data } = await api.put<Reuniao>(`/reunioes/${id}`, reuniao);
    return data;
  },

  async delete(id: string): Promise<void> {
    await api.delete(`/reunioes/${id}`);
  }
};
