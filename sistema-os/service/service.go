package service

import (
	"fmt"
	"sistema-os/auditoria"
	"sistema-os/models"
	"sistema-os/repository"
	"time"
)

type OrdemServicoService struct {
	Repositorio repository.BancoRepositorio
}

func (s OrdemServicoService) Salvar(ordem *models.OrdemServico) error {

	if ordem == nil {
		return fmt.Errorf("OS não pode ser nula")
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

	if ordem.DataDeEntrega.IsZero() {
		return fmt.Errorf("data de entrega é obrigatória")
	}

	if dataAnterior(ordem.DataDeEntrega) {
		return fmt.Errorf("data de entrega não pode ser anterior à data atual")
	}

	err := s.Repositorio.Salvar(ordem)
	if err != nil {
		return err
	}

	auditoria.OSCriada(ordem.ID)

	return nil
}

func (s OrdemServicoService) Listar() ([]*models.OrdemServico, error) {
	return s.Repositorio.Listar()
}

func (s OrdemServicoService) BuscarPorID(id int) (*models.OrdemServico, error) {

	if id <= 0 {
		return nil, fmt.Errorf("ID deve ser maior que zero")
	}

	ordem, err := s.Repositorio.BuscarPorID(id)
	if err != nil {
		return nil, fmt.Errorf("OS com ID %d não encontrada", id)
	}

	return ordem, nil
}

func (s OrdemServicoService) Atualizar(ordem *models.OrdemServico) error {

	if ordem == nil {
		return fmt.Errorf("OS não pode ser nula")
	}

	if ordem.ID <= 0 {
		return fmt.Errorf("ID deve ser maior que zero")
	}

	_, err := s.Repositorio.BuscarPorID(ordem.ID)
	if err != nil {
		return fmt.Errorf("OS com ID %d não encontrada", ordem.ID)
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

	if ordem.DataDeEntrega.IsZero() {
		return fmt.Errorf("data de entrega é obrigatória")
	}

	if dataAnterior(ordem.DataDeEntrega) {
		return fmt.Errorf("data de entrega não pode ser anterior à data atual")
	}

	err = s.Repositorio.Atualizar(ordem)
	if err != nil {
		return err
	}

	auditoria.OSAtualizada(ordem.ID)

	return nil
}

func (s OrdemServicoService) Deletar(id int) error {

	if id <= 0 {
		return fmt.Errorf("ID deve ser maior que zero")
	}

	_, err := s.Repositorio.BuscarPorID(id)
	if err != nil {
		return fmt.Errorf("OS com ID %d não encontrada", id)
	}

	err = s.Repositorio.Deletar(id)
	if err != nil {
		return err
	}

	auditoria.OSDeletada(id)

	return nil
}

func dataAnterior(data time.Time) bool {

	hoje := time.Now()

	dataHoje := time.Date(
		hoje.Year(),
		hoje.Month(),
		hoje.Day(),
		0,
		0,
		0,
		0,
		hoje.Location(),
	)

	dataInformada := time.Date(
		data.Year(),
		data.Month(),
		data.Day(),
		0,
		0,
		0,
		0,
		hoje.Location(),
	)

	return dataInformada.Before(dataHoje)
}
