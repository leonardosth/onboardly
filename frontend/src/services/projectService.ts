import type { Projeto } from '../types';
import { apiFetch } from './api';

export const projectService = {
  async getAll(): Promise<Projeto[]> {
    return apiFetch('/projetos') || [];
  },

  async getById(id: string): Promise<Projeto> {
    return apiFetch(`/projetos/${id}`);
  },

  async create(projeto: Partial<Projeto>): Promise<Projeto> {
    return apiFetch('/projetos', {
      method: 'POST',
      body: JSON.stringify(projeto)
    });
  },

  async update(id: string, projeto: Partial<Projeto>): Promise<Projeto> {
    return apiFetch(`/projetos/${id}`, {
      method: 'PUT',
      body: JSON.stringify(projeto)
    });
  },

  async delete(id: string): Promise<void> {
    return apiFetch(`/projetos/${id}`, {
      method: 'DELETE'
    });
  },

  async getStats(): Promise<{ 
    total_projetos: number; 
    total_clientes: number;
    reunioes_hoje: number;
    por_status: Record<string, number>;
    historico_mensal: Array<{ mes: string; total: number }>;
    atividades_recentes: Array<{ tipo: string; descricao: string; status: string; data: string }>;
  }> {
    return apiFetch('/dashboard/stats');
  }
};
