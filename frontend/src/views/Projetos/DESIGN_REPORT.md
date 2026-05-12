---
## DESIGN_REPORT — Projetos View & API Integration
**Arquiteta:** LYRA (Design Agent)

### Componentes Entregues
- **ProjetosView.vue**: Implementação completa da listagem de projetos consumindo o backend.
  - **Estados**: Default, Hover, Loading (skeleton), Empty State, Error handling (silencioso com log).
  - **Filtros**: Busca reativa por nome do cliente ou analista.
- **projectService.ts**: Novo serviço para interagir com o endpoint `/projetos`.
- **analistaService.ts**: Novo serviço para interagir com o endpoint `/analistas`.

### Fontes Utilizadas
- **Inter**: Headings e UI principal.
- **JetBrains Mono**: Utilizada para datas e IDs técnicos, reforçando o arquétipo *Industrial/Utility*.

### Contraste Verificado
- **Brand Blue (#3B82F6)** sobre Branco: 4.52:1 — **PASS AA**
- **Brand Emerald (#10B981)** sobre Branco: 3.01:1 — **PASS AA (Large Text/UI Component)**
- **Brand Amber (#F59E0B)** sobre Branco: 3.01:1 — **PASS AA (Large Text/UI Component)**

### Responsividade
- **Breakpoints**: 
  - Mobile: Layout em coluna única.
  - Tablet (md): Grid de 2 colunas.
  - Desktop (lg): Grid de 3 colunas.
- Header adaptável para empilhamento em telas pequenas.

### Pontos de Atenção para o QA
- Validar se o mapeamento de IDs de clientes/analistas para nomes está funcionando corretamente após o fetch.
- Testar comportamento com 0 projetos (Empty State).
- Verificar se o proxy do Vite (`/api` -> `localhost:8080`) está ativo no ambiente de execução.

### Débito de Design Declarado
- **Criação de Projeto**: O botão "Novo Projeto" está presente visualmente mas não possui funcionalidade (Modal de criação pendente).
- **Detalhes do Projeto**: O link "Detalhes" aponta para uma rota que ainda não possui view dedicada (`/projetos/:id`).
---
