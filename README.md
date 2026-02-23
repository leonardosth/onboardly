# ImplantaFlow - CRM para Analistas de Implantação

## 📌 Domínio do Problema
O projeto consiste em um CRM simplificado focado na dor do **Analista de Implantação**. Diferente de CRMs de vendas, este sistema foca no acompanhamento do cliente *pós-venda*, durante a fase técnica de configuração, treinamento e entrega de software.

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

| Tecnologia | Papel | Justificativa |
| :--- | :--- | :--- |
| **Go (Golang)** | Backend / API | Escolhida pela alta performance, tipagem forte e por ser a stack padrão do meu ambiente de trabalho atual, permitindo aprimoramento profissional. |
| **Vue.js** | Frontend / SPA | Framework progressivo que facilita a criação de interfaces reativas e modulares, além de alinhar meu aprendizado com as ferramentas usadas profissionalmente. |
| **PostgreSQL** | Banco de Dados | Robustez e confiabilidade para o armazenamento de dados relacionais de clientes. |

---

## 📅 Organização de Tarefas (Solo)

Para o desenvolvimento deste MVP, o cronograma será dividido em sprints rápidas:

1.  **Infra & Backend**
    * Modelagem do banco de dados.
    * Setup do projeto em Go.
    * CRUD básico de Clientes e Projetos (Endpoints).
2.  **Frontend & Integração**
    * Setup do Vue.js (Vue Router/Pinia).
    * Criação das telas de listagem e cadastro.
    * Consumo da API.
3.  **Polimento & Deploy** te amo <3
    * Refinamento da UI/UX.
    * Tratamento de erros e validações.
    * Documentação final.

---

## 👨‍💻 Autor
- **Leonardo**
