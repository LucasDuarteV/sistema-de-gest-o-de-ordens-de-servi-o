package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sistema-os/cli"
	"sistema-os/repository"
	"sistema-os/service"

	"github.com/jackc/pgx/v5"
)

func main() {
	fmt.Println("Iniciando sistema de Ordem de Serviço...")

	conn, err := pgx.Connect(
	context.Background(),
	"postgres://postgres:123456@127.0.0.1:5433/sistema_os",
)
	if err != nil {
		fmt.Println("erro ao conectar no PostgreSQL:", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println("erro ao testar conexão com PostgreSQL:", err)
		os.Exit(1)
	}

	fmt.Println("Conectado ao PostgreSQL com sucesso!")

	repo := repository.PostgresRepositorio{
		Conn: conn,
	}

	servico := service.OrdemServicoService{
		Repositorio: repo,
	}

	cli.Menu(servico)

	log.Println("Sistema encerrado.")
}
