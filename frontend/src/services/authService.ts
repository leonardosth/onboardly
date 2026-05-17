import axios from 'axios';
import type { AuthResponse, LoginRequest } from '../types';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export const authService = {
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    const response = await axios.post<AuthResponse>(`${API_URL}/auth/login`, credentials);
    return response.data;
  },

  setToken(token: string) {
    localStorage.setItem('onboardly_token', token);
    axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  },

  clearToken() {
    localStorage.removeItem('onboardly_token');
    delete axios.defaults.headers.common['Authorization'];
  },

  getToken() {
    return localStorage.getItem('onboardly_token');
  }
};
