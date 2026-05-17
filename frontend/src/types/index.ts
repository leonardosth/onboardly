export interface Cliente {
  id: string;
  nome: string;
  cnpj: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
}

export interface Analista {
  id: string;
  nome: string;
  email: string;
  created_at: string;
  updated_at: string;
}

export interface Projeto {
  id: string;
  cliente_id: string;
  analista_id: string;
  data_contratacao: string;
  data_ativacao?: string;
  status_ativacao: boolean;
  status_projeto: 'Backlog' | 'Em_Andamento' | 'Concluido';
  created_at: string;
  updated_at: string;
}

export interface Usuario {
  id: string;
  nome: string;
  email: string;
  cargo: string;
  created_at: string;
  updated_at: string;
}

export interface AuthResponse {
  token: string;
  user: Usuario;
}

export interface LoginRequest {
  email: string;
  senha: string;
}

export interface Reuniao {
  id: string;
  projeto_id: string;
  data_agendada: string;
  status: 'Agendada' | 'Realizada' | 'Remarcada' | 'No_Show';
  observacoes: string;
  created_at: string;
  updated_at: string;
}
