package repository

import (
	"database/sql"
	"fmt"

	"sistema-os/models"
)

type MySQLRepositorio struct {
	DB *sql.DB
}

func (r MySQLRepositorio) Salvar(ordem *models.OrdemServico) error {
	query := `
		INSERT INTO ordens_servico
		(cliente, descricao, valor_estimado, data_entrega, status, valor_final)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		ordem.Cliente,
		ordem.Descricao,
		ordem.ValorEstimado,
		ordem.DataDeEntrega,
		ordem.Status,
		ordem.ValorFinal,
	)

	if err != nil {
		return fmt.Errorf("erro ao salvar ordem de serviço: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("erro ao obter ID da ordem de serviço: %w", err)
	}

	ordem.ID = int(id)

	return nil
}

func (r MySQLRepositorio) Listar() ([]*models.OrdemServico, error) {
	query := `
		SELECT
			id,
			cliente,
			descricao,
			valor_estimado,
			data_entrega,
			status,
			valor_final
		FROM ordens_servico
		ORDER BY id
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar ordens de serviço: %w", err)
	}

	defer rows.Close()

	var ordens []*models.OrdemServico

	for rows.Next() {
		var ordem models.OrdemServico

		err := rows.Scan(
			&ordem.ID,
			&ordem.Cliente,
			&ordem.Descricao,
			&ordem.ValorEstimado,
			&ordem.DataDeEntrega,
			&ordem.Status,
			&ordem.ValorFinal,
		)

		if err != nil {
			return nil, fmt.Errorf("erro ao ler ordem de serviço: %w", err)
		}

		ordens = append(ordens, &ordem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao percorrer ordens de serviço: %w", err)
	}

	return ordens, nil
}

func (r MySQLRepositorio) BuscarPorID(id int) (*models.OrdemServico, error) {
	query := `
		SELECT
			id,
			cliente,
			descricao,
			valor_estimado,
			data_entrega,
			status,
			valor_final
		FROM ordens_servico
		WHERE id = ?
	`

	var ordem models.OrdemServico

	err := r.DB.QueryRow(
		query,
		id,
	).Scan(
		&ordem.ID,
		&ordem.Cliente,
		&ordem.Descricao,
		&ordem.ValorEstimado,
		&ordem.DataDeEntrega,
		&ordem.Status,
		&ordem.ValorFinal,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("OS com ID %d não encontrada", id)
		}

		return nil, fmt.Errorf("erro ao buscar OS: %w", err)
	}

	return &ordem, nil
}

func (r MySQLRepositorio) BuscarPorCliente(cliente string) ([]*models.OrdemServico, error) {
	query := `
		SELECT
			id,
			cliente,
			descricao,
			valor_estimado,
			data_entrega,
			status,
			valor_final
		FROM ordens_servico
		WHERE cliente LIKE ?
		ORDER BY id
	`

	rows, err := r.DB.Query(query, "%"+cliente+"%")
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar OS por cliente: %w", err)
	}

	defer rows.Close()

	var ordens []*models.OrdemServico

	for rows.Next() {
		var ordem models.OrdemServico

		err := rows.Scan(
			&ordem.ID,
			&ordem.Cliente,
			&ordem.Descricao,
			&ordem.ValorEstimado,
			&ordem.DataDeEntrega,
			&ordem.Status,
			&ordem.ValorFinal,
		)

		if err != nil {
			return nil, fmt.Errorf("erro ao ler OS: %w", err)
		}

		ordens = append(ordens, &ordem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao percorrer resultados: %w", err)
	}

	return ordens, nil
}

func (r MySQLRepositorio) Atualizar(ordem *models.OrdemServico) error {
	query := `
		UPDATE ordens_servico
		SET
			cliente = ?,
			descricao = ?,
			valor_estimado = ?,
			data_entrega = ?,
			status = ?,
			valor_final = ?
		WHERE id = ?
	`

	result, err := r.DB.Exec(
		query,
		ordem.Cliente,
		ordem.Descricao,
		ordem.ValorEstimado,
		ordem.DataDeEntrega,
		ordem.Status,
		ordem.ValorFinal,
		ordem.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar OS: %w", err)
	}

	linhas, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualização: %w", err)
	}

	if linhas == 0 {
		return fmt.Errorf("OS com ID %d não encontrada", ordem.ID)
	}

	return nil
}

func (r MySQLRepositorio) Deletar(id int) error {
	query := `
		DELETE FROM ordens_servico
		WHERE id = ?
	`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar OS: %w", err)
	}

	linhas, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar exclusão: %w", err)
	}

	if linhas == 0 {
		return fmt.Errorf("OS com ID %d não encontrada", id)
	}

	return nil
}