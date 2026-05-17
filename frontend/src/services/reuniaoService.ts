import type { Reuniao } from '../types';
import { apiFetch } from './api';

export const reuniaoService = {
  async getAll(): Promise<Reuniao[]> {
    return apiFetch('/reunioes') || [];
  },

  async getById(id: string): Promise<Reuniao> {
    return apiFetch(`/reunioes/${id}`);
  },

  async create(reuniao: Partial<Reuniao>): Promise<Reuniao> {
    return apiFetch('/reunioes', {
      method: 'POST',
      body: JSON.stringify(reuniao)
    });
  },

  async update(id: string, reuniao: Partial<Reuniao>): Promise<Reuniao> {
    return apiFetch(`/reunioes/${id}`, {
      method: 'PUT',
      body: JSON.stringify(reuniao)
    });
  },

  async delete(id: string): Promise<void> {
    return apiFetch(`/reunioes/${id}`, {
      method: 'DELETE'
    });
  }
};
