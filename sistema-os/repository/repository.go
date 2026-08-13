package repository

import (
	"sistema-os/models"
)

type Repositorio interface {
	Salvar(ordem models.OrdemServico) error
	BuscarPorID(id int) (*models.OrdemServico, error)
	Listar() ([]models.OrdemServico, error)
	Atualizar(ordem *models.OrdemServico) error
	Deletar(id int) error
}
