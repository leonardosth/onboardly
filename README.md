# Onboardly - CRM para Analistas de Implantação

## 📊 Status do Projeto

### 🔄 **SPRINT 1 - Infra & Backend (EM ANDAMENTO)**

**✅ Concluído:**

- ✅ Modelagem do banco de dados (PostgreSQL)
- ✅ Setup do projeto em Go com arquitetura limpa
- ✅ **CRUD completo de Clientes** (5 endpoints REST funcionais)
- ✅ Testes unitários para clientes (13/13 passando)
- ✅ Conexão e migração do banco de dados

**🔄 Pendente:**

- 🔄 CRUD de Projetos de Implantação
- 🔄 CRUD de Analistas
- 🔄 CRUD de Reuniões
- 🔄 Relacionamentos entre entidades
- 🔄 Validações de negócio avançadas

### 📋 Próximas Sprints

- **Sprint 2:** Completar backend (Projetos, Reuniões, Analistas)
- **Sprint 3:** Frontend Vue.js + Integração com API
- **Sprint 4:** Polimento, validações e deploy

---

## 📌 Domínio do Problema

O projeto consiste em um CRM simplificado focado na dor do **Analista de Implantação**. Diferente de CRMs de vendas, este sistema foca no acompanhamento do cliente _pós-venda_, durante a fase técnica de configuração, treinamento e entrega de software.

O objetivo é evitar o "esquecimento" de etapas críticas, centralizar os contatos de stakeholders e registrar o progresso de cada conta de forma clara e ágil.

### ✅ Requisitos Funcionais (RF)

- **Gestão de Carteira:** Cadastro, edição e visualização de clientes em implantação.
- **Status de Projetos:** Definição de etapas (Ex: Configuração, Treinamento, Homologação, Go-Live).
- **Log de Atividades:** Registro de notas e observações sobre as reuniões com o cliente.
- **Dashboard Simples:** Visualização de quantos projetos estão em cada fase.

### ⚙️ Requisitos Não-Funcionais (RNF)

- **Alta Performance:** Backend desenvolvido em Go para garantir tempos de resposta mínimos.
- **Interface Reativa:** Frontend em Vue.js para uma experiência de usuário fluida.
- **Persistência de Dados:** Uso de banco de dados SQL para integridade das informações.

---

## 🚀 Tecnologias e Justificativas

| Tecnologia      | Papel          | Justificativa                                                                                                                                                 |
| :-------------- | :------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Go (Golang)** | Backend / API  | Escolhida pela alta performance, tipagem forte e por ser a stack padrão do meu ambiente de trabalho atual, permitindo aprimoramento profissional.             |
| **Vue.js**      | Frontend / SPA | Framework progressivo que facilita a criação de interfaces reativas e modulares, além de alinhar meu aprendizado com as ferramentas usadas profissionalmente. |
| **PostgreSQL**  | Banco de Dados | Robustez e confiabilidade para o armazenamento de dados relacionais de clientes.                                                                              |

---

## 📅 Organização de Tarefas (Solo)

Para o desenvolvimento deste MVP, o cronograma será dividido em sprints rápidas:

1.  🔄 **SPRINT 1 - Infra & Backend (EM ANDAMENTO)**
    - ✅ Modelagem do banco de dados (PostgreSQL)
    - ✅ Setup do projeto em Go com arquitetura limpa
    - ✅ CRUD de Clientes (5 endpoints funcionais)
    - 🔄 CRUD de Projetos e Reuniões (pendente)
    - ✅ Testes unitários para clientes (13/13)

2.  📋 **SPRINT 2 - Completar Backend (PENDENTE)**
    - CRUD completo de Analistas
    - CRUD completo de Projetos de Implantação
    - CRUD completo de Reuniões
    - Relacionamentos e constraints de FK

3.  📋 **SPRINT 3 - Frontend & Integração (PENDENTE)**
    - Setup do Vue.js (Vue Router/Pinia)
    - Criação das telas de listagem e cadastro
    - Consumo da API REST completa

4.  📋 **SPRINT 4 - Polimento & Deploy (PENDENTE)**
    - Dashboard com métricas
    - Refinamento da UI/UX
    - Tratamento de erros e validações
    - Documentação final e deploy

---

## 📁 Arquitetura e Estrutura do Projeto

O backend segue uma arquitetura baseada no padrão standard do Go, focada na separação de responsabilidades:

```text
onboardly/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go       # Ponto de entrada da aplicação (API)
│   ├── internal/
│   │   ├── database/         # Configuração e conexão com o banco (PostgreSQL)
│   │   ├── models/           # Entidades do domínio (Cliente, Projeto, etc.)
│   │   ├── repository/       # Camada de acesso a dados (SQL)
│   │   └── handlers/         # Controladores HTTP (Rotas)
│   ├── .env.example          # Template das variáveis de ambiente
│   └── go.mod
├── frontend/                 # Projeto Vue.js (SPA)
└── README.md
```

---

## 🔌 API REST - Endpoints Implementados

### ✅ Clientes (`/clientes`) - IMPLEMENTADO

| Método | Endpoint         | Descrição                | Status |
| ------ | ---------------- | ------------------------ | ------ |
| POST   | `/clientes`      | Criar novo cliente       | ✅     |
| GET    | `/clientes`      | Listar todos os clientes | ✅     |
| GET    | `/clientes/{id}` | Buscar cliente por ID    | ✅     |
| PUT    | `/clientes/{id}` | Atualizar cliente        | ✅     |
| DELETE | `/clientes/{id}` | Remover cliente          | ✅     |

### 📋 Próximas Implementações

- **Projetos** (`/projetos`) - CRUD pendente
- **Analistas** (`/analistas`) - CRUD pendente
- **Reuniões** (`/reunioes`) - CRUD pendente

#### Exemplo de Uso

**Criar Cliente:**

```bash
curl -X POST http://localhost:8080/clientes \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Empresa Teste S.A.",
    "cnpj": "12.345.678/0001-90"
  }'
```

**Listar Clientes:**

```bash
curl -X GET http://localhost:8080/clientes
```

---

## ⚙️ Como Executar (Setup Local)

### Pré-requisitos

- Go 1.20+
- Node.js 18+ (para o frontend)
- PostgreSQL rodando localmente ou via Docker

### Backend

1. **Configurar Banco de Dados:**

   ```bash
   # Criar banco de dados
   createdb onboardly_db

   # Executar migração (schema.sql)
   psql -U postgres -d onboardly_db -f schema.sql
   ```

2. **Configurar Ambiente:**

   ```bash
   cd backend
   cp .env.example .env  # Editar com suas credenciais
   ```

3. **Executar API:**

   ```bash
   go run cmd/api/main.go
   ```

4. **Executar Testes:**

   ```bash
   # Todos os testes
   go test ./internal/service/... -v

   # Testes específicos
   go test -run TestCreateClient ./internal/service/...
   ```

### Testes Automatizados

- ✅ **13 testes unitários** para clientes (100% cobertura)
- 📋 Testes para outras entidades (pendente)
- Uso de mocks para isolamento de dependências

---

## 🏗️ Arquitetura Técnica Implementada

### Padrões Utilizados

- **Clean Architecture:** Separação clara entre camadas (Handlers → Service → Repository)
- **Dependency Injection:** Injeção de dependências para facilitar testes
- **Repository Pattern:** Abstração do acesso a dados
- **SOLID Principles:** Interface segregation e single responsibility

### Estrutura de Código

```
internal/
├── handlers/     # Controladores HTTP (JSON, validações)
├── service/      # Regras de negócio e lógica de aplicação
├── repository/   # Acesso a dados (SQL queries)
├── models/       # Estruturas de dados do domínio
└── database/     # Configuração de conexão DB
```

### Tecnologias Backend

- **Framework:** Go 1.26.1 (net/http nativo)
- **Banco:** PostgreSQL com driver `lib/pq`
- **Testes:** Testes unitários com mocks (atualmente apenas para clientes)
- **Configuração:** Variáveis de ambiente (.env)

### Entidades Implementadas

- ✅ **Clientes** - CRUD completo com testes
- 📋 **Analistas** - Pendente
- 📋 **Projetos** - Pendente
- 📋 **Reuniões** - Pendente

## 🚀 Próximos Passos

### Sprint 1 - Completar Backend

1. Implementar CRUD de Analistas
2. Implementar CRUD de Projetos de Implantação
3. Implementar CRUD de Reuniões
4. Configurar relacionamentos entre entidades
5. Adicionar validações de negócio

### Sprint 2 - Frontend Vue.js

1. Inicializar projeto Vue.js com Vite
2. Configurar Pinia para gerenciamento de estado
3. Criar componentes para CRUD completo
4. Implementar chamadas para a API REST
5. Estilizar com Tailwind CSS

### Funcionalidades Planejadas

- **Dashboard:** Métricas de projetos por status
- **Gestão de Projetos:** CRUD completo com relacionamentos
- **Reuniões:** Log de atividades e follow-ups
- **Relatórios:** Exportação de dados

---

## 👨‍💻 Autor

- **Leonardo**
