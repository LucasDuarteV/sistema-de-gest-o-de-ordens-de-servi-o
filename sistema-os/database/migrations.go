package database

import (
	"database/sql"
	"fmt"
)

func CriarTabelasMySQL(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS ordens_servico (
		id INT AUTO_INCREMENT PRIMARY KEY,
		cliente VARCHAR(255) NOT NULL,
		descricao TEXT NOT NULL,
		valor_estimado DECIMAL(10,2) NOT NULL,
		data_entrega DATE NOT NULL,
		status VARCHAR(50) NOT NULL,
		valor_final DECIMAL(10,2) DEFAULT 0
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("erro ao criar tabela ordens_servico: %w", err)
	}

	return nil
}