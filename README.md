# Control Bank

Realiza o controle do sistema

## Implementação

| Endpoint                    | Método | Descrição                                           | Exemplo de Uso                  |
|-----------------------------|--------|-----------------------------------------------------|---------------------------------|
| `/api/v1/banks`           | GET    | Obter a lista de sistemas                           | `GET /api/users`                |
| `/api/v1/banks/:uuid`     | GET    | Obter detalhes de um sistema específico pelo UUID   | `GET /api/users/123`            |
| `/api/v1/banks/:uuid`     | POST   | Cria um sistema novo a partir de um UUID            | `POST /api/users/123`           |
| `/api/v1/banks/:uuid`     | DELETE | Deleta um sistema a partir de um UUID               | `DELETE /api/users/123`         |
