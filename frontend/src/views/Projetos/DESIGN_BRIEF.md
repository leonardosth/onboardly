# DESIGN_BRIEF — Projetos View
**Arquiteta:** LYRA (Design Agent)

## Contexto
- Problema: Visualizar e gerenciar o fluxo de implementação de múltiplos clientes.
- Usuário: Analista de Implantação que precisa saber o que está parado e o que está avançando.
- Emoção-alvo: Controle, clareza e progressão.

## Direção Estética
- Arquétipo: Industrial/Utility — Foco em densidade de informação e clareza de status.
- Diferencial: Uso de cores semânticas vibrantes sobre um layout sóbrio e tipografia técnica.
- Anti-padrão evitado: Cards genéricos de dashboard sem hierarquia de data ou status.

## Sistema de Design
| Token | Valor | Uso |
|-------|-------|-----|
| --color-bg | #F8FAFC | Fundo da aplicação |
| --color-card | #FFFFFF | Fundo dos cards |
| --status-backlog | #3B82F6 (brand-blue) | Projetos não iniciados |
| --status-doing | #F59E0B (brand-amber) | Projetos em execução |
| --status-done | #10B981 (brand-emerald) | Projetos finalizados |
| --font-display | 'Inter' | Headings e Labels (Mantendo consistência atual) |
| --font-mono | 'JetBrains Mono' | IDs, Datas e Métricas |
| --radius | 24px | Border-radius padrão para containers |

## Decisões de Design
- **Status Pills**: Tags proeminentes com cores de alto contraste para identificação rápida.
- **Grid de Projetos**: Layout flexível que permite ver o cliente, o analista e as datas críticas sem cliques extras.
- **Empty State**: Ilustração sutil e CTA claro se não houver projetos.
