-- name: CreateInvoice :one
INSERT INTO notas_fiscais(id, emissao_em, recebido_em, cnpj, chave) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: FindInvoiceKeyExists :one
SELECT id FROM notas_fiscais WHERE chave = $1;
