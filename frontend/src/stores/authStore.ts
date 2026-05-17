import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { Usuario } from '../types';
import { authService } from '../services/authService';

export const useAuthStore = defineStore('auth', () => {
  const user = ref<Usuario | null>(null);
  const token = ref<string | null>(authService.getToken());

  const isAuthenticated = computed(() => !!token.value);

  function setUser(newUser: Usuario) {
    user.value = newUser;
    localStorage.setItem('onboardly_user', JSON.stringify(newUser));
  }

  function setToken(newToken: string) {
    token.value = newToken;
    authService.setToken(newToken);
  }

  function logout() {
    user.value = null;
    token.value = null;
    authService.clearToken();
    localStorage.removeItem('onboardly_user');
  }

  // Inicializa o usuário se houver no localStorage
  const savedUser = localStorage.getItem('onboardly_user');
  if (savedUser) {
    user.value = JSON.parse(savedUser);
  }

  return {
    user,
    token,
    isAuthenticated,
    setUser,
    setToken,
    logout
  };
});
