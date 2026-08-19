package repository

import (
	"context"
	"testing"
	"time"

	"sistema-os/models"

	"github.com/jackc/pgx/v5"
)

func conectarBancoTeste(t *testing.T) *pgx.Conn {
	t.Helper()

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:123456@127.0.0.1:5433/sistema_os",
	)

	if err != nil {
		t.Fatalf("erro ao conectar no PostgreSQL: %v", err)
	}

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

	if err != nil {
		t.Fatalf("erro ao salvar OS: %v", err)
	}

	if ordem.ID <= 0 {
		t.Errorf("esperava ID maior que zero, recebeu %d", ordem.ID)
	}

	t.Cleanup(func() {
		repo.Deletar(ordem.ID)
	})
}

func TestPostgresListar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordens, err := repo.Listar()

	if err != nil {
		t.Fatalf("erro ao listar OS: %v", err)
	}

	if ordens == nil {
		t.Error("esperava lista de OS, recebeu nil")
	}
}

func TestPostgresBuscarPorID(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	if err != nil {
		t.Fatalf("erro ao criar OS para teste: %v", err)
	}

	t.Cleanup(func() {
		repo.Deletar(ordem.ID)
	})

	resultado, err := repo.BuscarPorID(ordem.ID)

	if err != nil {
		t.Fatalf("erro ao buscar OS: %v", err)
	}

	if resultado.ID != ordem.ID {
		t.Errorf(
			"esperava ID %d, recebeu %d",
			ordem.ID,
			resultado.ID,
		)
	}

	if resultado.Cliente != "TESTE" {
		t.Errorf(
			"esperava cliente TESTE, recebeu %s",
			resultado.Cliente,
		)
	}
}

func TestPostgresBuscarPorIDInexistente(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	_, err := repo.BuscarPorID(999999)

	if err == nil {
		t.Error("esperava erro ao buscar OS inexistente")
	}
}

func TestPostgresAtualizar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	if err != nil {
		t.Fatalf("erro ao criar OS para teste: %v", err)
	}

	t.Cleanup(func() {
		repo.Deletar(ordem.ID)
	})

	ordem.Cliente = "TESTE ATUALIZADO"
	ordem.Descricao = "Descrição atualizada"
	ordem.ValorEstimado = 250.00
	ordem.Status = "Concluída"
	ordem.ValorFinal = 200.00

	err = repo.Atualizar(ordem)

	if err != nil {
		t.Fatalf("erro ao atualizar OS: %v", err)
	}

	atualizada, err := repo.BuscarPorID(ordem.ID)

	if err != nil {
		t.Fatalf("erro ao buscar OS atualizada: %v", err)
	}

	if atualizada.Cliente != "TESTE ATUALIZADO" {
		t.Errorf(
			"esperava cliente TESTE ATUALIZADO, recebeu %s",
			atualizada.Cliente,
		)
	}

	if atualizada.ValorEstimado != 250.00 {
		t.Errorf(
			"esperava valor 250.00, recebeu %.2f",
			atualizada.ValorEstimado,
		)
	}

	if atualizada.Status != "Concluída" {
		t.Errorf(
			"esperava status Concluída, recebeu %s",
			atualizada.Status,
		)
	}
}

func TestPostgresAtualizarInexistente(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()
	ordem.ID = 999999

	err := repo.Atualizar(ordem)

	if err == nil {
		t.Error("esperava erro ao atualizar OS inexistente")
	}
}

func TestPostgresDeletar(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	ordem := criarOrdemTeste()

	err := repo.Salvar(ordem)

	if err != nil {
		t.Fatalf("erro ao criar OS para teste: %v", err)
	}

	err = repo.Deletar(ordem.ID)

	if err != nil {
		t.Fatalf("erro ao deletar OS: %v", err)
	}

	_, err = repo.BuscarPorID(ordem.ID)

	if err == nil {
		t.Error("esperava erro ao buscar OS deletada")
	}
}

func TestPostgresDeletarInexistente(t *testing.T) {
	conn := conectarBancoTeste(t)
	repo := PostgresRepositorio{Conn: conn}

	err := repo.Deletar(999999)

	if err == nil {
		t.Error("esperava erro ao deletar OS inexistente")
	}
}
