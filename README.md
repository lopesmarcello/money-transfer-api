# Desafio Prático - Formação GO da Rocketseat: Sistema Bancário

Este projeto é a implementação de um desafio prático proposto na formação de GO da Rocketseat. O objetivo é construir uma API RESTful para simular operações bancárias básicas, aplicando conceitos importantes de desenvolvimento backend com Go.

## Requisitos

-   ✅ Criar um servidor HTTP em Golang.
-   ✅ Configurar a conexão com um banco de dados PostgreSQL.
-   ✅ Implementar seis rotas para operações bancárias:
    -   Criar conta
    -   Consultar saldo
    -   Depositar dinheiro
    -   Sacar dinheiro
    -   Transferir dinheiro entre contas
    -   Fechar conta
-   ✅ Utilizar tokens CSRF para proteger operações que modificam dados.
-   ✅ Implementar respostas JSON para as requisições.

## Estrutura do Banco de Dados (PostgreSQL)

O banco de dados foi configurado utilizando migrations para garantir a correta estruturação das tabelas.

### Tabela `pessoa_fisica`

Armazena informações sobre indivíduos.

| Campo          | Tipo                             | Descrição                                         |
| :------------- | :------------------------------- | :-------------------------------------------------- |
| `id`           | `SERIAL PRIMARY KEY`             | Identificador único da pessoa (autoincrementável).  |
| `renda_mensal` | `DECIMAL`                        | Renda mensal da pessoa.                             |
| `idade`        | `INTEGER`                        | Idade da pessoa.                                    |
| `nome_completo`| `VARCHAR(255)`                   | Nome completo da pessoa.                            |
| `celular`      | `VARCHAR(20)`                    | Número de celular.                                  |
| `email`        | `VARCHAR(255)`                   | Endereço de e-mail.                                 |
| `categoria`    | `VARCHAR(50)`                    | Categoria da pessoa.                                |
| `saldo`        | `DECIMAL`                        | Saldo disponível na conta.                          |

### Tabela `pessoa_juridica`

Armazena informações sobre empresas.

| Campo             | Tipo                             | Descrição                                             |
| :---------------- | :------------------------------- | :---------------------------------------------------- |
| `id`              | `SERIAL PRIMARY KEY`             | Identificador único da empresa (autoincrementável).   |
| `faturamento`     | `DECIMAL`                        | Faturamento anual da empresa.                         |
| `idade`           | `INTEGER`                        | Tempo de existência da empresa em anos.               |
| `nome_fantasia`   | `VARCHAR(255)`                   | Nome comercial da empresa.                            |
| `celular`         | `VARCHAR(20)`                    | Número de celular da empresa.                         |
| `email_corporativo` | `VARCHAR(255)`                   | Endereço de e-mail corporativo.                       |
| `categoria`       | `VARCHAR(50)`                    | Categoria da empresa.                                 |
| `saldo`           | `DECIMAL`                        | Saldo disponível na conta da empresa.                 |

## Rotas da API

| Método   | Rota                         | Descrição                                 |
| :------- | :--------------------------- | :---------------------------------------- |
| `POST`   | `/conta`                     | Cria uma nova conta (pessoa física).      |
| `GET`    | `/conta/{id}/saldo`          | Consulta o saldo de uma conta.            |
| `POST`   | `/conta/{id}/deposito`       | Deposita dinheiro em uma conta.           |
| `POST`   | `/conta/{id}/saque`          | Saca dinheiro de uma conta.               |
| `POST`   | `/conta/transferencia`       | Transfere dinheiro entre duas contas.     |
| `DELETE` | `/conta/{id}`                | Fecha (desativa) uma conta.               |

## Proteção CSRF

Para garantir a segurança contra ataques de Cross-Site Request Forgery (CSRF), as rotas que modificam o estado da aplicação (`POST`, `DELETE`) são protegidas. Um middleware valida a presença e a autenticidade de um token CSRF enviado no cabeçalho `X-CSRF-Token` de cada requisição.

## Como Executar

### Pré-requisitos

-   Go (versão 1.18 ou superior)
-   Docker e Docker Compose

### 1. Clonar o Repositório

```bash
git clone <URL_DO_REPOSITORIO>
cd money-transfer
```

### 2. Configurar Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto. Ele é necessário para configurar tanto o contêiner do Docker quanto a aplicação. Use o seguinte exemplo como base:

```env
# Variáveis utilizadas pelo Docker Compose para iniciar o banco de dados
MT_DATABASE_USER=postgres
MT_DATABASE_PASSWORD=postgres
MT_DATABASE_NAME=moneytransfer
MT_DATABASE_PORT=5555

# Variáveis utilizadas pela aplicação Go para se conectar ao banco
MT_DATABASE_HOST=localhost

# Chave de 32 bytes para proteção CSRF
MT_CSRF_KEY=coloque_aqui_uma_chave_csrf_segura_de_32_bytes
```

**Notas:**
- A porta do servidor da API é fixada em `:8080` no código.
- O `MT_CSRF_KEY` deve ser uma string aleatória e segura de 32 bytes. A aplicação não iniciará sem ela.
- Certifique-se de que `MT_DATABASE_PORT` no seu `.env` corresponda à porta exposta no `docker-compose.yml` e que `MT_DATABASE_HOST` seja `localhost` se estiver executando a API localmente.

### 3. Iniciar o Banco de Dados

Utilize o Docker Compose para iniciar um contêiner com o PostgreSQL e aplicar as migrations automaticamente.

```bash
docker-compose up -d
```

### 4. Executar a Aplicação

Com o banco de dados em execução, inicie o servidor Go:

```bash
go run cmd/api/main.go
```

O servidor estará disponível em `http://localhost:8080`.
