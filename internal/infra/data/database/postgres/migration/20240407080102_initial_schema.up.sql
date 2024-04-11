CREATE TABLE notas_fiscais (
  id varchar(36) NOT NULL PRIMARY KEY,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  emissao_em timestamp without time zone NOT NULL,
  recebido_em timestamp without time zone NOT NULL,
  cnpj varchar(14) NOT NULL,
  chave varchar(44) NOT NULL
);

CREATE INDEX idx_notas_fiscais_chave ON notas_fiscais(chave);
CREATE INDEX idx_notas_fiscais_cnpj ON notas_fiscais(cnpj);
