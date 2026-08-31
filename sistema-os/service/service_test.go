package service

import (
	"errors"
	"testing"
	"time"

	"sistema-os/models"
	"sistema-os/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type BancoRepositorioMock struct {
	ordens        []*models.OrdemServico
	salvarChamado bool
}

var _ repository.Repositorio = (*BancoRepositorioMock)(nil)

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

func (m *BancoRepositorioMock) BuscarPorCliente(cliente string) ([]*models.OrdemServico, error) {
	var ordens []*models.OrdemServico

	for _, ordem := range m.ordens {
		if ordem.Cliente == cliente {
			ordens = append(ordens, ordem)
		}
	}

	return ordens, nil
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

	require.NoError(t, err)
	assert.True(t, repositorio.salvarChamado)
	assert.NotZero(t, ordem.ID)
}

func TestSalvarOSNula(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	err := servico.Salvar(nil)

	assert.Error(t, err)
}

func TestSalvarClienteVazio(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.Cliente = ""

	err := servico.Salvar(ordem)

	assert.Error(t, err)
}

func TestSalvarDescricaoVazia(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.Descricao = ""

	err := servico.Salvar(ordem)

	assert.Error(t, err)
}

func TestSalvarValorEstimadoNegativo(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.ValorEstimado = -100

	err := servico.Salvar(ordem)

	assert.Error(t, err)
}

func TestSalvarValorFinalNegativo(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.ValorFinal = -50

	err := servico.Salvar(ordem)

	assert.Error(t, err)
}

func TestSalvarDataEntregaVazia(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.DataDeEntrega = time.Time{}

	err := servico.Salvar(ordem)

	assert.Error(t, err)
}

func TestSalvarDataEntregaAnterior(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	ordem := novaOrdemTeste()
	ordem.DataDeEntrega = time.Now().AddDate(0, 0, -1)

	err := servico.Salvar(ordem)

	assert.Error(t, err)
}

func TestBuscarPorIDInvalido(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	_, err := servico.BuscarPorID(0)

	assert.Error(t, err)
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

	require.NoError(t, err)
	require.NotNil(t, ordem)

	assert.Equal(t, 1, ordem.ID)
	assert.Equal(t, "Lucas", ordem.Cliente)
}

func TestDeletarIDInvalido(t *testing.T) {
	repositorio := &BancoRepositorioMock{}

	servico := OrdemServicoService{
		Repositorio: repositorio,
	}

	err := servico.Deletar(0)

	assert.Error(t, err)
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

	require.NoError(t, err)
	assert.Empty(t, repositorio.ordens)
}

func TestCalcularDiasRestantes(t *testing.T) {
	amanha := time.Now().AddDate(0, 0, 1)

	dias := CalcularDiasRestantes(amanha)

	assert.Equal(t, 1, dias)
}