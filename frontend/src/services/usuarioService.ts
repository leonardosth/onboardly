import { Usuario } from '../types';
import { apiFetch } from './api';

export const usuarioService = {
  async getAnalistas(): Promise<Usuario[]> {
    return apiFetch('/analistas');
  },

  async createAnalista(analista: Partial<Usuario>): Promise<Usuario> {
    return apiFetch('/analistas', {
      method: 'POST',
      body: JSON.stringify(analista)
    });
  },

  async updateAnalista(id: string, analista: Partial<Usuario>): Promise<void> {
    return apiFetch(`/analistas/${id}`, {
      method: 'PUT',
      body: JSON.stringify(analista)
    });
  },

  async deleteAnalista(id: string): Promise<void> {
    return apiFetch(`/analistas/${id}`, {
      method: 'DELETE'
    });
  }
};
