# DESIGN_BRIEF — Sistema de Notificações e Refinamento de Interface
**Arquiteta:** LYRA (Design Agent)

## Contexto
- Problema: Falta de feedback visual para ações do usuário (CRUD) e interface de navegação simplista.
- Usuário: Analistas de Implantação que precisam de confirmação imediata e navegação fluida.
- Emoção-alvo: Confiança, profissionalismo e agilidade.

## Direção Estética
- Arquétipo: Industrial/Utility Refinado — Foco na utilidade técnica com toques de elegância (transições suaves, sombras sutis).
- Diferencial: Toasts flutuantes com micro-interações e estados de transição refinados.
- Anti-padrão evitado: Alertas nativos do navegador ou toasts retangulares brutos sem hierarquia.

## Sistema de Design
| Token | Valor | Uso |
|-------|-------|-----|
| --color-success | #10B981 | Sucesso/Go-Live |
| --color-error | #F43F5E | Erros/Alertas |
| --color-info | #3B82F6 | Informações/Avisos |
| --font-sans | 'Instrument Sans' | Tipografia principal |
| --radius-xl | 12px | Arredondamento de componentes |
| --shadow-lg | 0 10px 15px -3px rgba(0,0,0,0.1) | Elevação de Toasts e Dropdowns |

## Decisões de Design
- **Toasts no canto inferior direito**: Padrão de indústria para não obstruir a visão principal do conteúdo, mas garantir visibilidade imediata.
- **Transição "Scale + Translate"**: Adiciona profundidade e sensação de "objeto físico" entrando na tela.
- **Avatar com fallback**: Uso de `ui-avatars` para garantir que a interface nunca tenha "buracos" visuais.
- **Overlay de Profile**: Implementação de overlay invisível para fechar dropdowns ao clicar fora, melhorando a usabilidade sem bibliotecas extras.
