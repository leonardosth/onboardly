import { useAuthStore } from '../stores/authStore';
import { useToastStore } from '../stores/toastStore';
import router from '../router';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080';

export async function apiFetch(endpoint: string, options: RequestInit = {}) {
  const authStore = useAuthStore();
  const toastStore = useToastStore();
  
  const headers = {
    'Content-Type': 'application/json',
    ...options.headers,
  } as any;

  if (authStore.token) {
    headers['Authorization'] = `Bearer ${authStore.token}`;
  }

  try {
    const response = await fetch(`${API_URL}${endpoint}`, {
      ...options,
      headers,
    });

    if (response.status === 401) {
      authStore.logout();
      router.push('/login');
      toastStore.error('Sessão expirada. Por favor, faça login novamente.');
      throw new Error('Unauthorized');
    }

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.error || `Erro na requisição: ${response.statusText}`);
    }

    if (response.status === 204) return null;
    return response.json();
  } catch (error: any) {
    if (error.message !== 'Unauthorized') {
      console.error('API Fetch Error:', error);
    }
    throw error;
  }
}
