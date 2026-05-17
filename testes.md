# Plano de Testes - Onboardly

Este documento descreve a estratégia, escopo e cenários de testes para garantir a qualidade e estabilidade do sistema Onboardly.

## 1. Estratégia de Testes

A estratégia baseia-se na Pirâmide de Testes:
- **Testes Unitários (Base):** Validar lógica de negócio isolada (Services no Backend, Funções/Componetes no Frontend).
- **Testes de Integração (Meio):** Validar a comunicação entre camadas (Handlers -> Services -> Repository) e integração com o Banco de Dados.
- **Testes E2E (Topo):** Validar fluxos críticos do usuário final no Frontend.

---

## 2. Escopo do Backend (Go)

### 2.1. Testes Unitários (Prioridade Alta)
- **AuthService:**
    - Registro de novo usuário com hash de senha.
    - Login com credenciais válidas/inválidas.
    - Geração e validação de claims do JWT.
- **UsuarioService:**
    - Validação de e-mail único.
    - Atribuição de cargos (Admin vs Analista).

### 2.2. Testes de Integração (Handlers)
- **ProjetoHandler:**
    - `POST /projetos`: Validar criação com payload correto e falha com IDs inexistentes.
    - `GET /projetos`: Validar listagem filtrada por permissões.
- **AuthHandler:**
    - Fluxo completo de login e retorno de token.
- **Middleware:**
    - Validar se o `AuthMiddleware` bloqueia requisições sem token ou com token expirado.

### 2.3. Testes de Repositório (Postgres)
- Uso de banco de dados de teste (Docker ou Transações Rollback) para validar as queries complexas do `GetDashboardStats`.

---

## 3. Escopo do Frontend (Vue 3/Vitest)

### 3.1. Testes de Componentes (Unitários)
- **Forms (ProjetoForm, ReuniaoForm):**
    - Verificar se os campos obrigatórios exibem erro ao tentar submeter vazio.
    - Validar a formatação de datas (ISO) antes da emissão do evento `save`.
- **UI (ToastContainer):**
    - Verificar se mensagens de sucesso/erro aparecem corretamente ao interagir com as stores.

### 3.2. Testes de Store (Pinia)
- **AuthStore:**
    - Persistência do token no localStorage após login.
    - Limpeza de estado no logout.

---

## 4. Cenários Críticos (Caminho Feliz & Exceções)

| ID | Cenário | Tipo | Resultado Esperado |
|:---|:---|:---|:---|
| CT01 | Criar Projeto com Sucesso | Integração | HTTP 201 e registro no BD com status 'Backlog' |
| CT02 | Login com Senha Incorreta | Unitário | Erro "credenciais inválidas" e nenhum token gerado |
| CT03 | Acesso a rota Admin por Analista | Segurança | HTTP 403 Forbidden |
| CT04 | Agendar Reunião em Projeto Inexistente | Integração | Erro 400/404 e transação abortada no BD |
| CT05 | Formatação de Data RFC3339 | Componente | Input `YYYY-MM-DD` convertido para ISO string no payload |

---

## 5. Ferramentas e Execução

- **Backend:** `go test ./... -v -cover`
- **Frontend:** `npm run test:unit` (Vitest)
- **Relatórios:** Cobertura alvo de **80%** na camada de Service e **50%** no Frontend.

---

## 6. Próximos Passos (Backlog de Implementação)
1. [ ] Implementar `internal/service/auth_service_test.go`.
2. [ ] Configurar ambiente Vitest no diretório `frontend/`.
3. [ ] Criar testes para o `AuthMiddleware` no backend.
4. [ ] Implementar teste unitário para `ProjetoForm.vue`.
