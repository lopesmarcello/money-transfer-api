-- name: GetSaldoPessoaFisicaByID :one 
SELECT saldo FROM pessoa_fisica WHERE id = $1;

-- name: GetSaldoPessoaJuridicaByID :one 
SELECT saldo FROM pessoa_juridica WHERE id = $1;
