# DESIGN_BRIEF — Fluxo de Autenticação e Gestão de Estado

**Arquiteta:** LYRA (Design Agent)

## Contexto

- Problema: Transição de um mock de autenticação para um sistema real com persistência de token e segurança.
- Usuário: Analista que precisa de um acesso rápido, seguro e uma interface que confirme quem ele é.
- Emoção-alvo: Segurança, controle e "welcome home".

## Direção Estética

- Arquétipo: Luxury/Refinado — O login é o portal. Deve ser limpo, focado e transmitir solidez.
- Diferencial: Transições suaves entre estados de autenticação e feedback visual imediato para erros de credenciais.
- Anti-padrão evitado: Redirecionamentos bruscos sem feedback de carregamento ou mensagens de erro genéricas ("Something went wrong").

## Sistema de Design

| Token          | Valor | Uso                    |
| -------------- | ----- | ---------------------- |
| --color-auth-bg| #FAF8F5 | Fundo da tela de login |
| --radius-apple | 24px  | Arredondamento premium |
| --shadow-premium| 0 25px 50px -12px rgba(0,0,0,0.08) | Profundidade do Card |

## Decisões de Design

- **State Persistence (Pinia + LocalStorage)**: O token JWT será armazenado de forma persistente para evitar re-logins desnecessários, mas com limpeza imediata em caso de expiração (401).
- **Navigation Guard Síncrono**: Verificação de token antes de renderizar qualquer rota protegida para evitar "flicker" de conteúdo privado.
- **Identity Display**: O cabeçalho deve exibir o nome do usuário com um fallback de avatar para criar senso de propriedade sobre o workspace.
- **Loading State no Botão**: O botão de login não apenas desabilita, mas transforma-se para mostrar progresso, mantendo a affordance.
