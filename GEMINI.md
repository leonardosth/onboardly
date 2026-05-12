# Onboardly - Contexto do Projeto

## 📌 Visão Geral
O **Onboardly** é um CRM especializado para **Analistas de Implantação**. Ao contrário de CRMs de vendas tradicionais, o Onboardly foca na fase de pós-venda, acompanhando a configuração técnica, treinamento e entrega de software (Go-Live).

O objetivo principal é centralizar stakeholders, registrar o progresso de cada conta e evitar o esquecimento de etapas críticas durante a implementação.

## 🚀 Tecnologias e Stack
- **Backend:** Go 1.26.1 (net/http nativo)
- **Frontend:** Vue 3 (Composition API) + Vite + TypeScript
- **Banco de Dados:** PostgreSQL 15+
- **Documentação:** PlantUML (arquivos `.wsd` no diretório `diagrams/`)
- **Arquitetura:** Clean Architecture / Standard Go Layout

## 🏗️ Estrutura do Projeto
- `/backend`: API REST em Go.
  - `cmd/api/main.go`: Ponto de entrada.
  - `internal/handlers`: Camada de transporte (HTTP).
  - `internal/service`: Regras de negócio.
  - `internal/repository`: Camada de persistência (Postgres).
  - `internal/models`: Entidades de domínio.
- `/frontend`: Interface reativa em Vue 3.
- `/diagrams`: Modelagem C4 e diagramas de classe/sequência.

## 🛠️ Comandos Úteis

### Backend
- **Rodar a API:** `cd backend && go run cmd/api/main.go`
- **Rodar Testes:** `cd backend && go test ./internal/service/... -v`
- **Migrações:** O esquema inicial está em `backend/schema.sql`.

### Frontend
- **Rodar Dev:** `cd frontend && npm run dev`
- **Build:** `cd frontend && npm run build`

## 📋 Convenções e Padrões
- **Código Limpo:** Seguir princípios SOLID e Clean Architecture no Go.
- **Injeção de Dependência:** Utilizada no `main.go` para compor handlers, services e repositories.
- **Testes:** Priorizar testes unitários na camada de `service` utilizando mocks para os repositórios.
- **Frontend:** Preferir Vue 3 Composition API com `<script setup>`.

## 🤖 Agentes Especializados (Sub-Agentes)
- `cto`: Arquiteto e Planejador Técnico. Use para decisões estratégicas e planejamento.
- `dev`: Desenvolvedor Sênior. Use para implementação técnica no backend (Go) e testes.
- `design`: Engenheira de Interface. Use para frontend (Vue 3/Tailwind) e UX.
- `qa`: Engenheiro de Qualidade. Use para revisão de código e automação de testes.

---
*Nota: Este arquivo é atualizado automaticamente para refletir o estado do projeto.*
