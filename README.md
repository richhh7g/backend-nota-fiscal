# Backend Nota Fiscal API

## Descrição

O Backend Nota Fiscal API é uma aplicação desenvolvida em Golang que fornece uma API para a criação e consulta de notas fiscais. Utilizando práticas de arquitetura limpa e os princípios do SOLID, o projeto busca garantir um código organizado, desacoplado e de fácil manutenção.

## O que foi utilizado no projeto?

No projeto, foram adotadas as seguintes tecnologias e práticas:

**Linguagem**: Foi utilizado o [Golang](https://go.dev/) para desenvolver o backend.

**Web Framework**: [Go-Chi](https://go-chi.io) é um microframework leve e rápido para criar aplicativos da web.

**Hot Reload**: [Air](https://github.com/cosmtrek/air) foi utilizado para acelerar o desenvolvimento.

**Injeção de Dependência**: [Wire](https://github.com/google/wire) foi utilizado para gerenciar as dependências do projeto.

**Tipagem SQL para Go**: [SQLC](https://sqlc.dev/) foi utilizado para gerar tipagens das queries SQL para Go.

**Migrations**: [Golang-migrate](https://github.com/golang-migrate/migrate) Usado para gerenciar migrações de banco de dados.

**Banco de Dados**: Foi utilizado o [PostgreSQL](https://www.postgresql.org/) como banco de dados relacional.

**Task Runner**: [Taskfile](https://taskfile.dev/) Ferramenta de automação de tarefas baseada em Go que permite definir e executar tarefas

**Arquitetura Limpa (Clean Architecture)**: Foram seguidos os princípios da Arquitetura Limpa para garantir um código organizado, desacoplado e de fácil manutenção.

**SOLID**: Foi utilizado os princípios do SOLID (Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation e Dependency Inversion) para desenvolver um código mais coeso, flexível e fácil de manter.

**Datasource**: Responsável por realizar a integração com fontes externas de dados. No contexto do projeto, o datasource é responsável por gerenciar todas as operações de leitura e escrita de dados no banco de dados. Além disso, ele também atua como uma ponte entre a aplicação e o microserviço de validação de documentos

**Swagger**: A API foi documentada utilizando o Swagger, facilitando o entendimento e a utilização.

**Docker**: Foi utilizado para criar um ambiente de desenvolvimento consistente e portátil, permitindo que o projeto seja facilmente executado e distribuído em diferentes ambientes.

## Instalando dependências do projeto

##### Terminal

```sh
go mod download
```

##### Taskfile

```sh
task mod
```

## Executando o projeto

#### Desenvolvimento

##### Terminal

```sh
air
```

##### Taskfile

```sh
task dev
```

#### Produção

##### Terminal

```sh
go build -o bin/server-bin cmd/server/main.go
./bin/server-bin
```

##### Taskfile

```sh
task build
./bin/server-bin
```

## Migrações

### Gerando migrações

##### Terminal

```sh
migrate create --ext sql --dir "internal/infra/data/database/postgres/migration" --tz UTC nameOfMigration
```

##### Taskfile

```sh
task migration:generate -- nameOfMigration
```

### Executando migrações

##### Terminal

```sh
migrate --path="internal/infra/data/database/postgres/migration" --database "$DATABASE_URL" up
```

##### Taskfile

```sh
task migration:up
```

### Revertendo migrações

##### Terminal

```sh
migrate --path="internal/infra/data/database/postgres/migration" --database "$DATABASE_URL" down
```

##### Taskfile

```sh
task migration:rollback
```

### Gerando tipagens

##### Terminal

```sh
sqlc generate
```

##### Taskfile

```sh
task migration:types
```

## Wire Injeção de Dependência

##### Terminal

```sh
wire gen --output_file_prefix injector. internal/infra/service_registry/injector.go
```

##### Taskfile

```sh
task generate:wire
```

## Rotas

### Criar Nota Fiscal

```sh
POST /api/v1/nota-fiscal
Content-Type: application/json

{
  "chave": "12345678901234567890123456789012345678901234",
  "cnpj": "12345678901234",
  "data_emissao": "2022-08-01T10:00:00Z",
  "data_recebimento": "2022-08-01T10:00:00Z"
}
```

### Consultar Nota Fiscal

```sh
GET /api/v1/nota-fiscal/{chave}
Content-Type: application/json
```

## Documentação completa com Swagger

Acesse: [http://localhost:{PORT}/docs](http://localhost:{PORT}/docs)

### Gerando a documentação

##### Terminal

```sh
swag init -g cmd/server/main.go -ot go,yaml
```

##### Taskfile

```sh
task generate:docs
```

## Docker

#### Executando com o Docker

```sh
docker build -t mm-api-nfe:latest .
docker run -p 3333:3333 mm-api-nfe
```

#### Executando com o Docker Compose

##### Terminal

```sh
docker compose -f .docker/compose/local.yml up
```

##### Taskfile

```sh
task compose:local
```

Por favor, crie um arquivo chamado `.env` na raiz do projeto para configurar as variáveis de ambiente necessárias.

## Pendências no Projeto

**Testes**: Não foi implementado nenhum teste de unidade, integração ou end-to-end. Criar testes adequados é crucial para garantir que o código funcione como esperado e seja confiável.

**Centralização de Textos**: A falta de um local central para armazenar todos os textos do projeto, faz com que os textos sejam repetidos em vários lugares, o que não é muito eficiente. Mesmo que seja apenas para o português, implementar uma internacionalização já seria de grande ajuda.
