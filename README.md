# Onboardly - CRM para Analistas de Implantação

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

1.  **Infra & Backend**
    - Modelagem do banco de dados.
    - Setup do projeto em Go.
    - CRUD básico de Clientes e Projetos (Endpoints).
2.  **Frontend & Integração**
    - Setup do Vue.js (Vue Router/Pinia).
    - Criação das telas de listagem e cadastro.
    - Consumo da API.
3.  **Polimento & Deploy**
    - Refinamento da UI/UX.
    - Tratamento de erros e validações.
    - Documentação final.

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

## ⚙️ Como Executar (Setup Local)

### Pré-requisitos

- Go 1.20+
- Node.js 18+ (para o frontend)
- PostgreSQL rodando localmente ou via Docker

### Backend

1. Acesse o diretório do backend: `cd backend`
2. Crie um arquivo `.env` na raiz da pasta `backend/` e defina as variáveis (utilize o `.env.example` como base):
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=suasenha
   DB_NAME=onboardly_db
   ```
3. Execute a API: `go run cmd/api/main.go`

---

## �‍💻 Autor

- **Leonardo**
