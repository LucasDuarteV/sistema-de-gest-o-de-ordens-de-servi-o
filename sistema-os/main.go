package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sistema-os/cli"
	"sistema-os/config"
	"sistema-os/database"
	"sistema-os/repository"
	"sistema-os/service"
)

func main() {
	fmt.Println("Iniciando sistema de Ordem de Serviço...")

	cfg := config.Carregar()

	conn, err := database.Conectar(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	fmt.Println("Conectando ao PostgreSQL com sucesso!")

	repo := repository.PostgresRepositorio{
		Conn: conn,
	}

	servico := service.OrdemServicoService{
		Repositorio: repo,
	}

	cli.Menu(servico)

	log.Println("Sistema encerrado.")
}
