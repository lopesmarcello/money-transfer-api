-- name: GetSaldoPessoaFisicaByID :one 
SELECT saldo FROM pessoa_fisica WHERE id = $1;

-- name: GetSaldoPessoaJuridicaByID :one 
SELECT saldo FROM pessoa_juridica WHERE id = $1;

-- name: UpdateSaldoFromPessoaFisica :one
UPDATE pessoa_fisica
SET saldo = $1
WHERE id = $2
RETURNING *;

-- name: UpdateSaldoFromPessoaJuridica :one
UPDATE pessoa_fisica
SET saldo = $1
WHERE id = $2
RETURNING *;
