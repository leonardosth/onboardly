# Onboardly - CRM para Analistas de Implantação

O **Onboardly** é um CRM especializado no pós-venda, projetado para auxiliar Analistas de Implantação a gerenciarem a jornada do cliente desde a configuração inicial até o Go-Live. Ao contrário de CRMs de vendas tradicionais, o Onboardly foca na fase técnica, garantindo que o treinamento, a configuração e a entrega de software ocorram sem falhas.

## 📊 Status do Projeto
**Status:** 🟢 Ativo (Fase de Refinamento e Estabilização)

## 🎯 Entregas Recentes (Done)
- **Unificação de Identidade:** As tabelas 'analistas' e 'usuarios' foram unificadas para simplificar a gestão de perfis e reforçar a segurança do sistema.
- **RBAC (Controle de Acesso):** Implementação de controle de acesso baseado em cargos (Admin vs Analista). Inclui Middleware de proteção no Backend e Navigation Guards no Frontend.
- **Módulo Administrativo:** Nova funcionalidade de gestão de usuários (CRUD completo), permitindo que administradores gerenciem a equipe de analistas.
- **Dashboard Dinâmico:** O painel principal e a aba de relatórios agora consomem dados reais do PostgreSQL, apresentando métricas de performance e histórico de projetos em tempo real.
- **Correções de UI/UX:**
    - Barra lateral (Sidebar) com elementos fixos para melhor usabilidade.
    - Implementação de "Native Feel" através do bloqueio de seleção de texto em elementos de navegação.
    - Refinamentos gerais em componentes de formulários e modais.
- **Segurança e Validação:** Validação rigorosa de payloads no backend e interceptores globais de erro no frontend, garantindo uma experiência robusta e resiliente.

## 🚧 Em Execução (Doing)
- Refinamento das animações de transição entre views.
- Expansão da cobertura de testes unitários na camada de serviços do backend.

## 📋 Backlog e Próximos Passos (To Do)
- [ ] [Feature] Sistema de notificações para prazos de projetos próximos do vencimento.
- [ ] [Feature] Exportação de relatórios de produtividade em PDF.
- [ ] [Infra] Configuração de pipeline de CI/CD básica.

---

## 🛠️ Tecnologias e Arquitetura

### Backend
- **Linguagem:** Go 1.26.1 (net/http nativo)
- **Arquitetura:** Clean Architecture / Standard Go Layout (Separation of Concerns).
- **Banco de Dados:** PostgreSQL 15+ (Relacional).
- **Segurança:** JWT para autenticação e Bcrypt para hashing de senhas.

### Frontend
- **Framework:** Vue 3 (Composition API) com TypeScript.
- **Build Tool:** Vite.
- **Estilização:** Tailwind CSS.
- **Estado:** Pinia (Gerenciamento de estado global).

---

## 🚀 Como Executar

### Backend
1. Navegue até `/backend`.
2. Configure o banco de dados PostgreSQL usando o arquivo `schema.sql`.
3. Crie um arquivo `.env` com base no `.env.example`.
4. Execute: `go run cmd/api/main.go`.

### Frontend
1. Navegue até `/frontend`.
2. Instale as dependências: `npm install`.
3. Execute: `npm run dev`.

---

## 📁 Organização do Repositório
- `/backend`: API REST, lógica de negócio e acesso ao banco.
- `/frontend`: Interface reativa e consumo de serviços.
- `/diagrams`: Documentação técnica e modelagem UML (PlantUML).

---

## 👨‍💻 Autor
Desenvolvido por **Leonardo** como um projeto acadêmico de excelência em gestão de implantação de software.
