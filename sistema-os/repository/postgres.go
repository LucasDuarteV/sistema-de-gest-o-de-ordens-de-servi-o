package repository

import (
	"context"
	"fmt"
	"sistema-os/models"

	"github.com/jackc/pgx/v5"
)

type PostgresRepositorio struct {
	Conn *pgx.Conn
}

func (r PostgresRepositorio) Salvar(ordem *models.OrdemServico) error {
	err := r.Conn.QueryRow(
		context.Background(),
		`INSERT INTO ordens_servico
		(cliente, descricao, valor_estimado, data_entrega, status, valor_final)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`,
		ordem.Cliente,
		ordem.Descricao,
		ordem.ValorEstimado,
		ordem.DataDeEntrega,
		ordem.Status,
		ordem.ValorFinal,
	).Scan(&ordem.ID)

	if err != nil {
		return fmt.Errorf("erro ao salvar ordem de serviço: %w", err)
	}

	return nil
}

func (r PostgresRepositorio) Listar() ([]*models.OrdemServico, error) {
	rows, err := r.Conn.Query(
		context.Background(),
		`
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
		`,
	)

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
			return nil, fmt.Errorf("erro ao escanear ordem de serviço: %w", err)
		}

		ordens = append(ordens, &ordem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao percorrer ordens de serviço: %w", err)
	}

	return ordens, nil
}

func (r PostgresRepositorio) BuscarPorID(id int) (*models.OrdemServico, error) {
	var ordem models.OrdemServico

	err := r.Conn.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			cliente,
			descricao,
			valor_estimado,
			data_entrega,
			status,
			valor_final
		FROM ordens_servico
		WHERE id = $1
		`,
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
		return nil, fmt.Errorf("OS com ID %d não encontrada", id)
	}

	return &ordem, nil
}

func (r PostgresRepositorio) Atualizar(ordem *models.OrdemServico) error {
	result, err := r.Conn.Exec(
		context.Background(),
		`
		UPDATE ordens_servico
		SET
			cliente = $1,
			descricao = $2,
			valor_estimado = $3,
			data_entrega = $4,
			status = $5,
			valor_final = $6
		WHERE id = $7
		`,
		ordem.Cliente,
		ordem.Descricao,
		ordem.ValorEstimado,
		ordem.DataDeEntrega,
		ordem.Status,
		ordem.ValorFinal,
		ordem.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar ordem de serviço: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("OS com ID %d não encontrada", ordem.ID)
	}

	return nil
}

func (r PostgresRepositorio) Deletar(id int) error {
	result, err := r.Conn.Exec(
		context.Background(),
		`
		DELETE FROM ordens_servico
		WHERE id = $1
		`,
		id,
	)

	if err != nil {
		return fmt.Errorf("erro ao deletar ordem de serviço: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("OS com ID %d não encontrada", id)
	}

	return nil
}
