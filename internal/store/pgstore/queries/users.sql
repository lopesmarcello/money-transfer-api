-- name: CreateUserPessoaFisica :one
INSERT INTO pessoa_fisica (
  renda_mensal,
  nome_completo,
  email,
  idade,
  celular,
  categoria,
  saldo
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;


-- name: CreateUserPessoaJuridica :one
INSERT INTO pessoa_juridica (
  faturamento,
  nome_fantasia,
  email_corporativo,
  idade,
  celular,
  categoria,
  saldo
) VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetUserPessoaFisicaByID :one
SELECT * FROM pessoa_fisica
WHERE id = $1;

-- name: GetUserPessoaFisicaByEmail :one
SELECT * FROM pessoa_fisica
WHERE email = $1;

-- name: GetUserPessoaJuridicaByID :one
SELECT * FROM pessoa_juridica
WHERE id = $1;

-- name: GetUserPessoaJuridicaByEmail :one
SELECT * FROM pessoa_juridica
WHERE email_corporativo = $1;
