package service

import (
	"fmt"
	"sistema-os/auditoria"
	"sistema-os/models"
	"sistema-os/repository"
)

type OrdemServicoService struct {
	Repositorio repository.Repositorio
}

func (s OrdemServicoService) Salvar(ordem models.OrdemServico) error {

	_, err := s.Repositorio.BuscarPorID(ordem.ID)

	if err == nil {
		return fmt.Errorf("já existe uma OS com o ID %d", ordem.ID)
	}

	if ordem.Cliente == "" {
		return fmt.Errorf("cliente é obrigatório")
	}

	if ordem.Descricao == "" {
		return fmt.Errorf("descrição é obrigatória")
	}

	if ordem.ValorEstimado < 0 {
		return fmt.Errorf("valor estimado não pode ser negativo")
	}

	if ordem.ValorFinal < 0 {
		return fmt.Errorf("valor final não pode ser negativo")
	}

	err = s.Repositorio.Salvar(ordem)

	if err != nil{
		return err
	}

	auditoria.OSCriada(ordem.ID)
	
	return nil
}

func (s OrdemServicoService) Listar() ([]models.OrdemServico, error) {
	return s.Repositorio.Listar()
}

func (s OrdemServicoService) BuscarPorID(id int) (*models.OrdemServico, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID deve ser maior que zero!")
	}

	return s.Repositorio.BuscarPorID(id)
}

func (s OrdemServicoService) Atualizar(ordem *models.OrdemServico) error {
	if ordem == nil {
		return fmt.Errorf("OS não pode ser nula!")
	}

	if ordem.ID <= 0 {
		return fmt.Errorf("ID deve ser maior que zero!")
	}

	_, err := s.Repositorio.BuscarPorID(ordem.ID)
	if err != nil {
		return fmt.Errorf("OS com ID %d não encontrada", ordem.ID)
	}

	if ordem.Cliente == "" {
		return fmt.Errorf("cliente obrigatorio")
	}

	if ordem.Descricao == "" {
		return fmt.Errorf("descricao obrigatoria")
	}

	if ordem.ValorEstimado < 0 {
		return fmt.Errorf("valor estimado não pode ser negativo!")
	}

	if ordem.ValorFinal < 0 {
		return fmt.Errorf("valor final não pode ser menos que zero")
	}

	err = s.Repositorio.Atualizar(ordem)

	if err != nil{
		return err
	}

	auditoria.OSAtualizada(ordem.ID)

	return nil

}

func (s OrdemServicoService) Deletar(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID tem que ser maior que zero!")
	}

	_, err := s.Repositorio.BuscarPorID(id)
	if err != nil {
		return fmt.Errorf("OS com ID %d não encontrada!", id)
	}

	err = s.Repositorio.Deletar(id)

	if err != nil{
		return err
	}

	auditoria.OSDeletada(id)

	return nil
}
