# Cercle Test DevJR

Guia rápido para rodar o projeto localmente.

## Dependências

- Go 1.25+
- Docker
- Docker Compose
- sqlc (CLI)
- swag (CLI)
- Visual Studio Code
- Extensão Go para VS Code (recomendado para debug)

## Variáveis de ambiente

O projeto possui um arquivo base de variáveis em `.env.example`.

Antes de executar a API, copie esse arquivo para `.env`:

```bash
cp .env.example .env
```

Depois disso, ajuste os valores se necessário para o seu ambiente local.

## Comandos principais

Execute na raiz do projeto:

```bash
make up
make sqlc
make swagger
make reset
```

### O que cada comando faz

- `make up`: sobe os containers do ambiente local (banco e serviços do docker-compose).
- `make sqlc`: regenera os arquivos Go a partir das queries SQL.
- `make swagger`: regenera a documentação Swagger em `cmd/api/docs`.
- `make reset`: remove containers e volumes do ambiente Docker.

Tem outros comandos disponíveis, veja o `Makefile` para detalhes. Mas os acima são os principais para desenvolvimento local.

## Execução da API

Existe o comando:

```bash
make api
```

Ele funciona para subir a API via terminal.

No entanto, o ideal para desenvolvimento é rodar pela ferramenta de debug do VS Code (configuração já presente no projeto), pois facilita:

- breakpoints
- inspeção de variáveis
- step-by-step na execução

Use a configuração de debug **Run API (cmd/api)** no VS Code.

## Endereços do Swagger

Com a API rodando, acesse:

- Swagger UI: http://localhost:8080/swagger/index.html
- Swagger JSON: http://localhost:8080/swagger/doc.json

## Fluxo recomendado

1. `make up`
2. `make sqlc`
3. `make swagger`
4. Rodar a API via debug do VS Code (`Run API (cmd/api)`)

# Em caso de dúvidas

Se você tiver alguma dúvida ou encontrar algum problema, entre em contato com a equipe de desenvolvimento (franklin.salesx@gmail.com) ou abra uma issue no repositório do projeto. Também pode entrar em contato pelo [LinkedIn](https://www.linkedin.com/in/franklindux/).