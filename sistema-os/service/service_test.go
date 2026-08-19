package service

import (
	"errors"
	"testing"
	"time"

	"sistema-os/models"
)

type BancoRepositorioMock struct {
	ordens        []*models.OrdemServico
	salvarChamado bool
}

func (m *BancoRepositorioMock) Salvar(ordem *models.OrdemServico) error {
	m.salvarChamado = true

	if ordem.ID == 0 {
		ordem.ID = 1
	}

	m.ordens = append(m.ordens, ordem)

	return nil
}

func (m *BancoRepositorioMock) Listar() ([]*models.OrdemServico, error) {
	return m.ordens, nil
}

func (m *BancoRepositorioMock) BuscarPorID(id int) (*models.OrdemServico, error) {
	for _, ordem := range m.ordens {
		if ordem.ID == id {
			return ordem, nil
		}
	}

	return nil, errors.New("OS não encontrada")
}

func (m *BancoRepositorioMock) Atualizar(ordem *models.OrdemServico) error {
	for i, existente := range m.ordens {
		if existente.ID == ordem.ID {
			m.ordens[i] = ordem
			return nil
		}
	}

	return errors.New("OS não encontrada")
}

func (m *BancoRepositorioMock) Deletar(id int) error {
	for i, ordem := range m.ordens {
		if ordem.ID == id {
			m.ordens = append(m.ordens[:i], m.ordens[i+1:]...)
			return nil
		}
	}

	return errors.New("OS não encontrada")
}

func novaOrdemTeste() *models.OrdemServico {
	return &models.OrdemServico{
		Cliente:       "Lucas",
		Descricao:     "Manutenção de computador",
		ValorEstimado: 350.50,
		DataDeEntrega: time.Now().AddDate(0, 0, 5),
		Status:        "Pendente",
		ValorFinal:    0,
	}
}

func TestSalvarSucesso(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()

	err := servico.Salvar(ordem)

	if err != nil {
		t.Fatalf("esperava nenhum erro, recebeu: %v", err)
	}

	if !repositorio.salvarChamado {
		t.Error("esperava que Salvar fosse chamado")
	}

	if ordem.ID == 0 {
		t.Error("esperava que a OS recebesse um ID")
	}
}

func TestSalvarOSNula(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	err := servico.Salvar(nil)

	if err == nil {
		t.Error("esperava erro ao salvar OS nula")
	}
}

func TestSalvarClienteVazio(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.Cliente = ""

	err := servico.Salvar(ordem)

	if err == nil {
		t.Error("esperava erro para cliente vazio")
	}
}

func TestSalvarDescricaoVazia(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.Descricao = ""

	err := servico.Salvar(ordem)

	if err == nil {
		t.Error("esperava erro para descrição vazia")
	}
}

func TestSalvarValorEstimadoNegativo(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.ValorEstimado = -100

	err := servico.Salvar(ordem)

	if err == nil {
		t.Error("esperava erro para valor estimado negativo")
	}
}

func TestSalvarValorFinalNegativo(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.ValorFinal = -50

	err := servico.Salvar(ordem)

	if err == nil {
		t.Error("esperava erro para valor final negativo")
	}
}

func TestSalvarDataEntregaVazia(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.DataDeEntrega = time.Time{}

	err := servico.Salvar(ordem)

	if err == nil {
		t.Error("esperava erro para data vazia")
	}
}

func TestSalvarDataEntregaAnterior(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.DataDeEntrega = time.Now().AddDate(0, 0, -1)

	err := servico.Salvar(ordem)

	if err == nil {
		t.Error("esperava erro para data anterior")
	}
}

func TestBuscarPorIDInvalido(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	_, err := servico.BuscarPorID(0)

	if err == nil {
		t.Error("esperava erro para ID inválido")
	}
}

func TestBuscarPorID(t *testing.T) {
	repositorio := &BancoRepositorioMock{
		ordens: []*models.OrdemServico{
			{
				ID:      1,
				Cliente: "Lucas",
			},
		},
	}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem, err := servico.BuscarPorID(1)

	if err != nil {
		t.Fatalf("esperava nenhum erro, recebeu: %v", err)
	}

	if ordem.Cliente != "Lucas" {
		t.Errorf("esperava cliente Lucas, recebeu %s", ordem.Cliente)
	}
}

func TestDeletarIDInvalido(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	err := servico.Deletar(0)

	if err == nil {
		t.Error("esperava erro para ID inválido")
	}
}

func TestDeletar(t *testing.T) {
	repositorio := &BancoRepositorioMock{
		ordens: []*models.OrdemServico{
			{
				ID:      1,
				Cliente: "Lucas",
			},
		},
	}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	err := servico.Deletar(1)

	if err != nil {
		t.Fatalf("esperava nenhum erro, recebeu: %v", err)
	}

	if len(repositorio.ordens) != 0 {
		t.Error("esperava que a OS fosse deletada")
	}
}

func TestCalcularDiasRestantes(t *testing.T) {
	amanha := time.Now().AddDate(0, 0, 1)

	dias := CalcularDiasRestantes(amanha)

	if dias != 1 {
		t.Errorf("esperava 1 dia, recebeu %d", dias)
	}
}
