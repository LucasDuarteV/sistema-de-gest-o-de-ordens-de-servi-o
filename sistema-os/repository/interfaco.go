package repository

import "sistema-os/models"

type BancoRepositorio interface {
	Salvar(ordem *models.OrdemServico) error
	Listar() ([]*models.OrdemServico, error)
	BuscarPorID(id int) (*models.OrdemServico, error)
	Atualizar(ordem *models.OrdemServico) error
	Deletar(id int) error
}