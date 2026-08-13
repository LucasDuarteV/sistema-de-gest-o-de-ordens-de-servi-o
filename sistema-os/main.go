package main

import (
	"sistema-os/cli"
	"sistema-os/repository"
	"sistema-os/service"
)

func main() {

	repo := repository.JSONRepositorio{
		Arquivo: "dados.json",
	}

	servico := service.OrdemServicoService{
		Repositorio: repo,
	}

	cli.Menu(servico)
}
