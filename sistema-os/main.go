package main

import (
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

	db, err := database.ConectarMySQL(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer db.Close()

	fmt.Println("Conectando ao MySQL com sucesso!")

	err = database.CriarTabelasMySQL(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Tabela ordens_servico verificada com sucesso!")

	repo := repository.MySQLRepositorio{
		DB: db,
	}

	servico := service.OrdemServicoService{
		Repositorio: repo,
	}

	cli.Menu(servico)

	log.Println("Sistema encerrado.")
}