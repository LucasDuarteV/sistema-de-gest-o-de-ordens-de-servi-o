package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sistema-os/models"
)

type JSONRepositorio struct {
	Arquivo string
}

func (r JSONRepositorio) Listar() ([]models.OrdemServico, error) {

	dados, err := os.ReadFile(r.Arquivo)
	if err != nil {
		return nil, err
	}

	var ordens []models.OrdemServico

	err = json.Unmarshal(dados, &ordens)

	if err != nil {
		return nil, err
	}

	return ordens, nil

}

func (r JSONRepositorio) Salvar(ordem models.OrdemServico) error {
	dados, err := os.ReadFile(r.Arquivo)
	if err != nil {
		return err
	}

	var ordens []models.OrdemServico

	err = json.Unmarshal(dados, &ordens)
	if err != nil {
		return err
	}

	ordens = append(ordens, ordem)
	dados, err = json.Marshal(ordens)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.Arquivo, dados, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r JSONRepositorio) BuscarPorID(id int) (*models.OrdemServico, error) {
	dados, err := os.ReadFile(r.Arquivo)
	if err != nil {
		return nil, err
	}

	var ordens []models.OrdemServico

	err = json.Unmarshal(dados, &ordens)
	if err != nil {
		return nil, err
	}

	for i := range ordens {
		if ordens[i].ID == id {
			return &ordens[i], nil
		}
	}
	return nil, fmt.Errorf("OS com ID %d não encontrada", id)
}

func (r JSONRepositorio) Atualizar(ordem *models.OrdemServico) error{
	dados, err := os.ReadFile(r.Arquivo)
	if err != nil{
		return err
	}

	var ordens []models.OrdemServico

	err = json.Unmarshal(dados,&ordens)
	if err != nil{
		return err
	}

	for i := range ordens{
		if ordens[i].ID == ordem.ID{
			ordens[i] = *ordem
		}
	}
	
	dados, err = json.Marshal(ordens)
	if err != nil{
		return err
	}

	err = os.WriteFile(r.Arquivo,dados,0644)
	if err != nil{
		return err
	}

	return nil
}

func (r JSONRepositorio) Deletar(id int) error{
	dados, err := os.ReadFile(r.Arquivo)
		if err != nil{
			return err
		}

	var ordens []models.OrdemServico
	
	err = json.Unmarshal(dados,&ordens)
	if err != nil{
		return err
	}

	for i := range ordens{
		if ordens[i].ID == id{
			ordens = append(ordens[:i], ordens[i+1:]...)
			break
		}
	}

	dados, err = json.Marshal(ordens)
	if err != nil{
		return err
	}

	err = os.WriteFile(r.Arquivo,dados,0644)
	if err != nil{
		return err
	}

	return nil
}
