# Desafio Técnico — Módulo Resale

## Bem-vindo ao desafio técnico para Dev Fullstack Jr na Cercle! 

Este exercício simula uma tarefa real do dia a dia do time de desenvolvimento.

> PODE UTILIZAR IA PARA RESOLVER O DESAFIO!

No dia a dia, o time de desenvolvimento tem acesso a ferramentas de IA para auxiliar na implementação, ou seja, não faz sentido proibir o uso de IA para resolver o desafio.

---

## Contexto

A Cercle é uma plataforma de revendas.
Foi requisitado a você a implementação de funcionalidades para o módulo de Resale, que é responsável por gerenciar os pedidos de revenda.

## O desafio

Você deverá implementar dois endpoints para o módulo de Resale seguindo a arquitetura DDD + Clean Architecture já adotada no projeto.

O primeiro endpoint é um `GET` para listar os itens de um pedido específico.
O segundo endpoint é um `PUT` para cancelar (mudar o status para "RETURNED") um item específico de um pedido.

O módulo **Resale** já possui a estrutura de pastas criada e dois endpoints registrados no router, mas os handlers(core/modules/resale/presentation/handler.go) retornam apenas `TODO`. Sua missão é implementar a lógica completa seguindo os padrões já estabelecidos no projeto.

O time de negócio decidiu criar uma interface para o cliente pedir devolução de um ou mais itens do pedido. O próprio usuário fará isso através de uma interface na plataforma. Mas para a interface funcionar é necessário ter dois endpoints na API: Um para listar os itens do pedido e outro para cancelar um item específico.

Você deve implementar a lógica de negócio para ambos os endpoints, garantindo que as regras de negócio sejam respeitadas e que a arquitetura do projeto seja seguida corretamente.

---

## Endpoints a implementar

### Status possíveis de shipping_status

Considere os seguintes valores para o campo `shipping_status` no contexto deste desafio:

- `LABEL_GENERATED` (ou "etique gerada", quando a etiqueta de envio é gerada, mas o item ainda não foi postado)
- `POSTED` (ou "postado", quando o item foi postado, mas ainda não foi entregue)
- `DELIVERED` (ou "entregue", quando o item foi entregue ao comprador)
- `RETURNED` (ou "devolvido", quando o item foi devolvido pelo comprador)
- `CANCELLED` (ou "cancelado", quando o item foi cancelado antes do envio ou por algum motivo administrativo)

(os status de pedido aqui não importam para esse teste, somente os dos itens do pedido, que estão na tabela `resale_order_item`)

Observações:
- Para o endpoint de cancelamento deste desafio, o status final esperado é `RETURNED`.
- O status `CANCELLED` existe, mas somente para itens que foram cancelados por motivos administrativos ou antes do envio. Ele não deve ser usado para marcar itens como cancelados pelos compradores, que devem usar o status `RETURNED`.

### `GET /v1/app/users/:cpf/orders/:order_id/items`

User Story: Como um comprador da plataforma, quero poder visualizar os itens de pedidos específicos para acompanhar minhas compras e entregas, assim posso ter clareza sobre o status de cada item.

Constraints:
- O comprador deve fornecer seu CPF e o ID do pedido para acessar os itens.

Critérios de aceitação:
- O endpoint deve retornar uma lista de itens associados ao pedido e ao CPF do comprador.
- Todos os itens devem ser retornados(inclusive os cancelados), os itens retornados devem conter o campo `shipping_status` para indicar o status da entrega.
- Se o pedido não existir ou não pertencer ao comprador, deve retornar um erro 404.
- Se o `order_id` não for um UUID válido, deve retornar um erro 400.
- O endpoint deve seguir as convenções REST e retornar os status HTTP apropriados.

**Regras de negócio:**
- O `cpf` deve ser um parâmetro de path (não vem no body)
- O `order_id` deve ser um UUID válido e também um parâmetro de path
- Itens com `deleted_at` preenchido **não** devem ser retornados
- Se o pedido não existir, retorne `404`
- Se o `order_id` não for um UUID válido, retorne `400`

**Resposta esperada (200):**
```json
{
  "data": [
    {
      "id": "uuid",
      "fk_resale_order_id": "uuid",
      "sku": "LAB-CAMISETA-001",
      "name": "Camiseta Dry Fit Feminina",
      "quantity": 1,
      "amount_value": "219.00",
      "shipping_code": "BRTRK1001",
      "shipping_status": "DELIVERED"
    }
  ]
}
```

---

### `PUT /v1/app/users/:cpf/orders/:order_id/items/:item_id/cancel`

User Story: Como um comprador da plataforma, quero poder cancelar um item específico do meu pedido para solicitar uma devolução.
Realiza o cancelamento de um item específico de um pedido, marcando-o como cancelado (mudando o status do shipping_status para "RETURNED").

Constraints:
- O client deve fornecer seu CPF, o ID do pedido e o ID do item para cancelar um item específico.
- O endpoint deve validar que o pedido pertence ao comprador e que o item pertence ao pedido antes de realizar o cancelamento.
- Somente itens com 7 dias ou menos após a data de entrega (delivered_at) podem ser cancelados. Itens com mais de 7 dias desde a data de entrega não podem ser cancelados e devem retornar um erro 400.

Critérios de aceitação:
- O endpoint deve cancelar o item específico do pedido, marcando-o como cancelado (mudando o status do shipping_status para "RETURNED").
- Se o item já estiver cancelado, o endpoint deve retornar um status 204 (idempotente).
- Se o pedido ou item não existirem, ou se o pedido não pertencer ao comprador, o endpoint deve retornar um erro 404.
- Se o `order_id` ou `item_id` não forem UUIDs válidos, o endpoint deve retornar um erro 400.
- O endpoint deve seguir as convenções REST e retornar os status HTTP apropriados.

**Regras de negócio:**
- O `cpf`, `order_id` e `item_id` devem vir via path
- O `order_id` e `item_id` devem ser UUIDs válidos
- O pedido deve pertencer ao usuário com o CPF informado — caso contrário retorne `404`
- Se o item já estiver cancelado (`shipping_status` igual a "RETURNED"), o endpoint deve retornar `204 No Content` (idempotente)
- Se o item não pertencer ao pedido informado, retorne `404`
- Em caso de sucesso, retorne `204 No Content`

---

## O que esperamos

- Seguir a arquitetura **DDD + Clean Architecture** já adotada no projeto
  - `domain/` — agregados, interfaces de repositório, erros de domínio
  - `application/` — DTOs de entrada/saída, use cases, mapper
  - `infrastructure/` — implementação dos repositórios usando as queries SQLC em `core/database/postgres/query/sqlc`
  - `presentation/http/` — handler que delega para os use cases
- O `setup.go` deve instanciar as dependências via injeção manual (sem frameworks de DI)
- Usar o pacote de resposta compartilhado em `core/modules/shared/presentation/http/response`
- Erros de validação devem retornar `400`, recurso não encontrado `404`, sucesso sem body `204`, sucesso com body `200`, e erros de servidor `500`, seguindo as convenções REST.

## Referência

O módulo `core/modules/retailer` está **100% implementado** e segue exatamente o padrão esperado. Use-o como guia para estrutura de arquivos, nomenclatura e fluxo de dados.

O schema do banco está em `core/database/postgres/migration/00001_initial_schema.up.sql`.

Dados de exemplo já populados estão em `core/database/postgres/samples/`.

---

## Critérios de avaliação

| Critério | Peso |
|---|---|
| Corretude das regras de negócio | Alto |
| Aderência à arquitetura do projeto | Alto |
| Qualidade e clareza do código | Médio |
| Tratamento adequado de erros e status HTTP | Médio |
| Não quebrar o que já está funcionando | Alto |

---

> **Dica:** Antes de começar, rode `make up` para subir o banco com dados de exemplo. Leia o `README.md` para instruções de ambiente.
