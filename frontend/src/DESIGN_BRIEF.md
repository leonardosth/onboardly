# DESIGN_BRIEF — Native App Feel & Interface Immersiveness

**Arquiteta:** LYRA (Design Agent)

## Contexto

- **Problema:** O cursor de texto (caret) e a capacidade de seleção aparecem em elementos de interface não-interativos (h1, p, span), quebrando a percepção de "produto acabado" e assemelhando-se a um documento web genérico.
- **Usuário:** Analistas de implementação que buscam uma ferramenta de trabalho robusta e focada.
- **Emoção-alvo:** Imersão e Precisão. O usuário deve sentir que interage com uma ferramenta nativa, não com uma página de texto.

## Direção Estética

- **Arquétipo:** Minimalismo Radical / Apple-like. O silêncio visual inclui a ausência de artefatos de edição onde eles não pertencem.
- **Diferencial:** Controle total sobre os estados de interação. O texto é tratado como rótulo (label) de interface por padrão, e não como conteúdo editável.
- **Anti-padrão evitado:** "Blue highlight" acidental ao clicar rápido em títulos ou botões e o cursor I-beam aparecendo sobre headings.

## Sistema de Design

| Token          | Valor | Uso                    |
| -------------- | ----- | ---------------------- |
| --cursor-ui    | default | Cursor padrão para a app |
| --cursor-text  | text  | Cursor para inputs/texto |
| --selection-bg | var(--color-primary-soft) | Highlight sutil se necessário |

## Decisões de Design

- **Seleção Bloqueada Globalmente:** Aplicado no `:root` para garantir que toda a árvore de DOM herde o comportamento de "interface imutável".
- **Whitelisting de Formulários:** Restauração explícita para `input`, `textarea` e elementos `[contenteditable]`.
- **Caret Transparente:** Removido via CSS global para evitar que cliques em áreas de texto gerem o marcador piscante, mesmo se o navegador tentar forçá-lo.
- **Cursor Sântico:** Forçar `cursor: default` em todo o `body` para prevenir o I-beam em parágrafos e spans.
