package repository

import (
	"context"
	"testing"
	"time"

	"sistema-os/models"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func conectarBancoTeste(t *testing.T) *pgx.Conn {
	t.Helper()

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:123456@127.0.0.1:5433/sistema_os",
	)

	require.NoError(t, err, "erro ao conectar no PostgreSQL")

	t.Cleanup(func() {
		conn.Close(context.Background())
	})

	return conn
}

func criarOrdemTeste() *models.OrdemServico {
	return &models.OrdemServico{
		Cliente:       "TESTE",
		Descricao:     "OS criada pelo teste automatizado",
		ValorEstimado: 100.00,
		DataDeEntrega: time.Now().AddDate(0, 0, 5),
		Status:        "Pendente",
		ValorFinal:    0,
	}
}

func TestPostgresSalvar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	require.NoError(t, err, "erro ao salvar OS")
	assert.Greater(t, ordem.ID, 0, "esperava ID maior que zero")

	t.Cleanup(func() {
		repo.Deletar(ordem.ID)
	})
}

func TestPostgresListar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordens, err := repo.Listar()

	require.NoError(t, err, "erro ao listar OS")
	assert.NotNil(t, ordens, "esperava lista de OS, recebeu nil")
}

func TestPostgresBuscarPorID(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	require.NoError(t, err, "erro ao criar OS para teste")

	t.Cleanup(func() {
		repo.Deletar(ordem.ID)
	})

	resultado, err := repo.BuscarPorID(ordem.ID)

	require.NoError(t, err, "erro ao buscar OS")
	require.NotNil(t, resultado, "resultado não deveria ser nil")

	assert.Equal(t, ordem.ID, resultado.ID)
	assert.Equal(t, "TESTE", resultado.Cliente)
}

func TestPostgresBuscarPorIDInexistente(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	_, err := repo.BuscarPorID(999999)

	assert.Error(t, err, "esperava erro ao buscar OS inexistente")
}

func TestPostgresAtualizar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	require.NoError(t, err, "erro ao criar OS para teste")

	t.Cleanup(func() {
		repo.Deletar(ordem.ID)
	})

	ordem.Cliente = "TESTE ATUALIZADO"
	ordem.Descricao = "Descrição atualizada"
	ordem.ValorEstimado = 250.00
	ordem.Status = "Concluída"
	ordem.ValorFinal = 200.00

	err = repo.Atualizar(ordem)

	require.NoError(t, err, "erro ao atualizar OS")

	atualizada, err := repo.BuscarPorID(ordem.ID)

	require.NoError(t, err, "erro ao buscar OS atualizada")
	require.NotNil(t, atualizada)

	assert.Equal(t, "TESTE ATUALIZADO", atualizada.Cliente)
	assert.Equal(t, "Descrição atualizada", atualizada.Descricao)
	assert.Equal(t, 250.00, atualizada.ValorEstimado)
	assert.Equal(t, "Concluída", atualizada.Status)
	assert.Equal(t, 200.00, atualizada.ValorFinal)
}

func TestPostgresAtualizarInexistente(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()
	ordem.ID = 999999

	err := repo.Atualizar(ordem)

	assert.Error(t, err, "esperava erro ao atualizar OS inexistente")
}

func TestPostgresDeletar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	require.NoError(t, err, "erro ao criar OS para teste")

	err = repo.Deletar(ordem.ID)

	require.NoError(t, err, "erro ao deletar OS")

	_, err = repo.BuscarPorID(ordem.ID)

	assert.Error(t, err, "esperava erro ao buscar OS deletada")
}

func TestPostgresDeletarInexistente(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	err := repo.Deletar(999999)

	assert.Error(t, err, "esperava erro ao deletar OS inexistente")
}
