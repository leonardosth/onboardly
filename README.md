# Onboardly - CRM para Analistas de Implantação

O **Onboardly** é um CRM especializado no pós-venda, projetado para auxiliar Analistas de Implantação a gerenciarem a jornada do cliente desde a configuração inicial até o Go-Live. Ao contrário de CRMs de vendas tradicionais, o Onboardly foca na fase técnica, garantindo que o treinamento, a configuração e a entrega de software ocorram sem falhas.

## 📊 Status do Projeto
**Status:** 🟢 Ativo (Fase de Estabilização e QA)

## 🎯 Entregas Recentes (Done)
- **Migração Completa para Tabela de Usuários:** Unificação definitiva dos fluxos de 'Projetos' e 'Reuniões' para utilizarem a nova estrutura de usuários, eliminando dependências da tabela legada de analistas.
- **Integridade Referencial:** Atualização de todas as Foreign Keys (FKs) no banco de dados para garantir consistência entre projetos, reuniões e os analistas responsáveis.
- **Vínculo de Analistas em Reuniões:** Cada reunião agora possui um analista responsável vinculado, permitindo rastreabilidade completa das interações com o cliente.
- **Correção de Payloads:** Ajuste na formatação de datas (ISO 8601/RFC3339) para garantir compatibilidade entre o Frontend (Vue) e o Backend (Go `time.Time`).
- **RBAC (Controle de Acesso):** Controle de acesso baseado em cargos (Admin vs Analista) com proteção via Middleware no Backend e Navigation Guards no Frontend.
- **Dashboard Dinâmico:** Painel principal e relatórios consumindo dados reais do PostgreSQL com métricas de performance em tempo real.
- **Plano de Testes Estruturado:** Criação do documento `testes.md` definindo a estratégia de cobertura (80% em Services, 50% no Frontend) e cenários críticos.

## 🚧 Em Execução (Doing)
- Implementação da bateria de testes para o `AuthService` (0% -> 100% de cobertura).
- Configuração do ambiente `Vitest` no Frontend para testes de componentes.
- Refinamento das animações de transição entre views.

## 📋 Backlog e Próximos Passos (To Do)
- [ ] [QA] Atingir 80% de cobertura na camada de Services do Backend.
- [ ] [Feature] Sistema de notificações para prazos de projetos próximos do vencimento.
- [ ] [Feature] Exportação de relatórios de produtividade em PDF.
- [ ] [Infra] Configuração de pipeline de CI/CD básica (GitHub Actions).

---

## 🛠️ Tecnologias e Arquitetura

### Backend
- **Linguagem:** Go 1.26.1 (net/http nativo).
- **Arquitetura:** Clean Architecture / Standard Go Layout (Separation of Concerns).
- **Banco de Dados:** PostgreSQL 15+ (Relacional).
- **Segurança:** JWT (Autenticação) e Bcrypt (Hashing de senhas).

### Frontend
- **Framework:** Vue 3 (Composition API) com TypeScript.
- **Build Tool:** Vite.
- **Estilização:** Tailwind CSS.
- **Estado:** Pinia (Gerenciamento de estado global).
- **Testes:** Vitest (Planejado).

---

## 🚀 Como Executar

### Backend
1. Navegue até `/backend`.
2. Configure o banco de dados PostgreSQL usando o arquivo `schema.sql`.
3. Crie um arquivo `.env` com base no `.env.example`.
4. (Opcional) Popule o banco: `go run seed.go`.
5. Execute: `go run cmd/api/main.go`.

### Frontend
1. Navegue até `/frontend`.
2. Instale as dependências: `npm install`.
3. Execute: `npm run dev`.

---

## 📁 Organização do Repositório
- `/backend`: API REST, lógica de negócio e acesso ao banco.
- `/frontend`: Interface reativa e consumo de serviços.
- `/diagrams`: Documentação técnica e modelagem UML (PlantUML).
- `testes.md`: Planejamento e estratégia de QA.

---

## 👨‍💻 Autor
Desenvolvido por **Leonardo** como um projeto acadêmico de excelência em gestão de implantação de software.
