package cli

import (
	"bufio"
	"fmt"
	"os"
	"sistema-os/models"
	"sistema-os/service"
	"strconv"
	"strings"
	"time"
)

func Menu(servico service.OrdemServicoService) {
	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== SISTEMA DE ORDEM DE SERVIÇO =====")
		fmt.Println("1 - Criar OS")
		fmt.Println("2 - Listar OS")
		fmt.Println("3 - Buscar OS por ID")
		fmt.Println("4 - Atualizar OS")
		fmt.Println("5 - Deletar OS")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		opcao := lerTexto(leitor)

		switch opcao {
		case "1":
			criarOS(leitor, servico)

		case "2":
			listarOS(servico)

		case "3":
			buscarOS(leitor, servico)

		case "4":
			atualizarOS(leitor, servico)

		case "5":
			deletarOS(leitor, servico)

		case "0":
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}

// =========================
// CRIAR OS
// =========================

func criarOS(leitor *bufio.Reader, servico service.OrdemServicoService) {
	fmt.Println("\n===== CRIAR OS =====")

	// Cliente
	fmt.Print("Digite o cliente: ")
	cliente := lerTexto(leitor)

	if cliente == "" {
		fmt.Println("Cliente é obrigatório.")
		return
	}

	// Descrição
	fmt.Print("Digite a descrição: ")
	descricao := lerTexto(leitor)

	if descricao == "" {
		fmt.Println("Descrição é obrigatória.")
		return
	}

	// Valor estimado
	fmt.Print("Digite o valor estimado: ")
	valorTexto := lerTexto(leitor)

	valorEstimado, err := strconv.ParseFloat(valorTexto, 64)
	if err != nil {
		fmt.Println("Valor estimado inválido!")
		return
	}

	if valorEstimado < 0 {
		fmt.Println("Valor estimado não pode ser negativo.")
		return
	}

	// Data de entrega
	fmt.Print("Digite a data de entrega (02/01/2006): ")
	dataTexto := lerTexto(leitor)

	dataEntrega, err := time.Parse("02/01/2006", dataTexto)
	if err != nil {
		fmt.Println("Data inválida! Use o formato 02/01/2006.")
		return
	}

	if dataEntrega.Before(inicioDoDia(time.Now())) {
		fmt.Println("Data de entrega não pode ser anterior à data atual.")
		return
	}

	// ID NÃO É INFORMADO AQUI.
	// O PostgreSQL irá gerar automaticamente.
	ordem := &models.OrdemServico{
		Cliente:       cliente,
		Descricao:     descricao,
		ValorEstimado: valorEstimado,
		DataDeEntrega: dataEntrega,
		Status:        "Pendente",
		ValorFinal:    0,
	}

	err = servico.Salvar(ordem)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Printf("OS criada com sucesso! ID: %d\n", ordem.ID)
}

// =========================
// LISTAR OS
// =========================

func listarOS(servico service.OrdemServicoService) {
	fmt.Println("\n===== LISTAR OS =====")

	ordens, err := servico.Listar()
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	if len(ordens) == 0 {
		fmt.Println("Nenhuma OS cadastrada.")
		return
	}

	for _, ordem := range ordens {
		fmt.Println("-----------------------------")
		fmt.Println("ID:", ordem.ID)
		fmt.Println("Cliente:", ordem.Cliente)
		fmt.Println("Descrição:", ordem.Descricao)

		fmt.Printf(
			"Valor estimado: %.2f\n",
			ordem.ValorEstimado,
		)

		fmt.Println(
			"Data de entrega:",
			ordem.DataDeEntrega.Format("02/01/2006"),
		)

		dias := service.CalcularDiasRestantes(
			ordem.DataDeEntrega,
		)

		fmt.Println("Dias restantes:", dias)
		fmt.Println("Status:", ordem.Status)

		fmt.Printf(
			"Valor final: %.2f\n",
			ordem.ValorFinal,
		)
	}

	fmt.Println("-----------------------------")
}

// =========================
// BUSCAR OS
// =========================

func buscarOS(
	leitor *bufio.Reader,
	servico service.OrdemServicoService,
) {
	fmt.Println("\n===== BUSCAR OS =====")

	fmt.Print("Digite o ID da OS: ")

	idTexto := lerTexto(leitor)

	id, err := strconv.Atoi(idTexto)
	if err != nil {
		fmt.Println("ID inválido!")
		return
	}

	if id <= 0 {
		fmt.Println("ID deve ser maior que zero.")
		return
	}

	ordem, err := servico.BuscarPorID(id)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("\nOS encontrada:")
	mostrarOS(ordem)
}

// =========================
// ATUALIZAR OS
// =========================

func atualizarOS(
	leitor *bufio.Reader,
	servico service.OrdemServicoService,
) {
	fmt.Println("\n===== ATUALIZAR OS =====")

	// PRIMEIRO PEDE O ID
	fmt.Print("Digite o ID da OS: ")

	idTexto := lerTexto(leitor)

	id, err := strconv.Atoi(idTexto)
	if err != nil {
		fmt.Println("ID inválido!")
		return
	}

	if id <= 0 {
		fmt.Println("ID deve ser maior que zero.")
		return
	}

	// PRIMEIRO PROCURA A OS.
	// Se não encontrar, PARA AQUI.
	ordem, err := servico.BuscarPorID(id)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// Só chega aqui se a OS existir.
	fmt.Println("\nOS encontrada:")
	mostrarOS(ordem)

	// Cliente
	fmt.Print("\nDigite o novo cliente: ")
	cliente := lerTexto(leitor)

	if cliente == "" {
		fmt.Println("Cliente é obrigatório.")
		return
	}

	// Descrição
	fmt.Print("Digite a nova descrição: ")
	descricao := lerTexto(leitor)

	if descricao == "" {
		fmt.Println("Descrição é obrigatória.")
		return
	}

	// Valor estimado
	fmt.Print("Digite o novo valor estimado: ")
	valorTexto := lerTexto(leitor)

	valorEstimado, err := strconv.ParseFloat(valorTexto, 64)
	if err != nil {
		fmt.Println("Valor estimado inválido!")
		return
	}

	if valorEstimado < 0 {
		fmt.Println("Valor estimado não pode ser negativo.")
		return
	}

	// Data
	fmt.Print("Digite a nova data de entrega (02/01/2006): ")
	dataTexto := lerTexto(leitor)

	dataEntrega, err := time.Parse("02/01/2006", dataTexto)
	if err != nil {
		fmt.Println("Data inválida! Use o formato 02/01/2006.")
		return
	}

	if dataEntrega.Before(inicioDoDia(time.Now())) {
		fmt.Println("Data de entrega não pode ser anterior à data atual.")
		return
	}

	// Status
	fmt.Print("Digite o novo status: ")
	status := lerTexto(leitor)

	if status == "" {
		fmt.Println("Status é obrigatório.")
		return
	}

	// Valor final
	fmt.Print("Digite o novo valor final: ")
	valorFinalTexto := lerTexto(leitor)

	valorFinal, err := strconv.ParseFloat(valorFinalTexto, 64)
	if err != nil {
		fmt.Println("Valor final inválido!")
		return
	}

	if valorFinal < 0 {
		fmt.Println("Valor final não pode ser negativo.")
		return
	}

	// Atualiza a OS encontrada
	ordem.Cliente = cliente
	ordem.Descricao = descricao
	ordem.ValorEstimado = valorEstimado
	ordem.DataDeEntrega = dataEntrega
	ordem.Status = status
	ordem.ValorFinal = valorFinal

	err = servico.Atualizar(ordem)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("OS atualizada com sucesso!")
}

// =========================
// DELETAR OS
// =========================

func deletarOS(
	leitor *bufio.Reader,
	servico service.OrdemServicoService,
) {
	fmt.Println("\n===== DELETAR OS =====")

	fmt.Print("Digite o ID da OS: ")

	idTexto := lerTexto(leitor)

	id, err := strconv.Atoi(idTexto)
	if err != nil {
		fmt.Println("ID inválido!")
		return
	}

	if id <= 0 {
		fmt.Println("ID deve ser maior que zero.")
		return
	}

	// Procura antes de deletar.
	// Se não existir, para imediatamente.
	ordem, err := servico.BuscarPorID(id)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("\nOS encontrada:")
	mostrarOS(ordem)

	fmt.Print("\nTem certeza que deseja deletar? (s/n): ")

	resposta := lerTexto(leitor)

	if strings.ToLower(resposta) != "s" {
		fmt.Println("Opção cancelada.")
		return
	}

	err = servico.Deletar(id)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("OS deletada com sucesso!")
}

// =========================
// MOSTRAR OS
// =========================

func mostrarOS(ordem *models.OrdemServico) {
	fmt.Println("-----------------------------")
	fmt.Println("ID:", ordem.ID)
	fmt.Println("Cliente:", ordem.Cliente)
	fmt.Println("Descrição:", ordem.Descricao)

	fmt.Printf(
		"Valor estimado: %.2f\n",
		ordem.ValorEstimado,
	)

	fmt.Println(
		"Data de entrega:",
		ordem.DataDeEntrega.Format("02/01/2006"),
	)

	fmt.Println("Status:", ordem.Status)

	fmt.Printf(
		"Valor final: %.2f\n",
		ordem.ValorFinal,
	)

	fmt.Println("-----------------------------")
}

// =========================
// LER TEXTO
// =========================

func lerTexto(leitor *bufio.Reader) string {
	texto, _ := leitor.ReadString('\n')
	return strings.TrimSpace(texto)
}

// =========================
// INÍCIO DO DIA
// =========================

func inicioDoDia(data time.Time) time.Time {
	return time.Date(
		data.Year(),
		data.Month(),
		data.Day(),
		0,
		0,
		0,
		0,
		data.Location(),
	)
}