-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS pessoa_fisica (

id SERIAL PRIMARY KEY,
renda_mensal FLOAT NOT NULL, 
idade INT NOT NULL, 
nome_completo VARCHAR(255) NOT NULL,
celular VARCHAR(20) NOT NULL,
email VARCHAR(255) NOT NULL,
categoria VARCHAR(50) NOT NULL,
saldo FLOAT NOT NULL
);

---- create above / drop below ----
DROP TABLE IF EXISTS pessoa_fisica;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
