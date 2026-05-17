# DESIGN_REPORT — Fluxo de Autenticação

**Arquiteta:** LYRA (Design Agent)

### Componentes Entregues

- **LoginView.vue**: Implementado com lógica assíncrona real, tratamento de erros via Toast e estados de loading no botão principal.
- **DashboardLayout.vue**: Refatorado para consumir dados dinâmicos do usuário logado (Avatar reativo, nome e cargo).
- **AuthStore (Pinia)**: Gestão centralizada de estado com persistência em LocalStorage para token e perfil do usuário.
- **AuthService (Axios)**: Abstração de chamadas de API com configuração automática de headers de autorização.

### Fontes Utilizadas

- **Instrument Sans**: Corpo e UI.
- **Cormorant Garamond**: Utilizada em títulos (login) para o arquétipo Luxury/Refinado.

### Contraste Verificado

- Texto primário sobre fundo superfície: 12.5:1 — **Passa AAA**
- Botão primário (azul) sobre branco: 4.6:1 — **Passa AA**

### Responsividade

- **Breakpoints**: 
  - Mobile (< 1024px): Menu lateral colapsável, formulário de login ocupando 100% da largura.
  - Desktop (>= 1024px): Menu lateral fixo, card de login centralizado.

### Pontos de Atenção para o QA

- Validar se o token é removido do LocalStorage ao clicar em "Sair".
- Testar a proteção de rotas tentando acessar `/clientes` sem estar logado (deve redirecionar para `/login`).
- Verificar se o nome do usuário no cabeçalho atualiza corretamente após o login.

### Débito de Design Declarado

- **RefreshToken**: O sistema atual não renova o token automaticamente. Caso o token expire, o usuário receberá erros de API (401) sem um tratamento automático de re-login transparente — **Prioridade Média**.
