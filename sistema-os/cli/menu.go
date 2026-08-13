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

		opcao, _ := leitor.ReadString('\n')
		opcao = strings.TrimSpace(opcao)

		switch opcao {
		case "1":
			fmt.Println("\n===== CRIAR OS =====")

			fmt.Print("Digite o ID: ")
			idTexto, _ := leitor.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			id, err := strconv.Atoi(idTexto)
			if err != nil {
				fmt.Println("ID inválido!")
				break
			}

			fmt.Print("Digite o cliente: ")
			cliente, _ := leitor.ReadString('\n')
			cliente = strings.TrimSpace(cliente)

			fmt.Print("Digite a descrição: ")
			descricao, _ := leitor.ReadString('\n')
			descricao = strings.TrimSpace(descricao)

			fmt.Print("Digite o valor estimado: ")
			valorTexto, _ := leitor.ReadString('\n')
			valorTexto = strings.TrimSpace(valorTexto)

			valorEstimado, err := strconv.ParseFloat(valorTexto, 64)
			if err != nil {
				fmt.Println("Valor inválido!")
				break
			}

			fmt.Print("Digite a data de entrega (02/01/2006): ")
			dataTexto, _ := leitor.ReadString('\n')
			dataTexto = strings.TrimSpace(dataTexto)

			dataEntrega, err := time.Parse("02/01/2006", dataTexto)
			if err != nil {
				fmt.Println("Data inválida! Use o formato 02/01/2006.")
				break
			}

			ordem := models.OrdemServico{
				ID:            id,
				Cliente:       cliente,
				Descricao:     descricao,
				ValorEstimado: valorEstimado,
				DataDeEntrega: dataEntrega,
				Status:        "Pendente",
				ValorFinal:    0,
			}

			err = servico.Salvar(ordem)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("OS criada com sucesso!")

		case "2":
			fmt.Println("Listar OS")

			ordens, err := servico.Listar()
			if err != nil {
				fmt.Println(err)
				break
			}
			for _, ordem := range ordens {
				fmt.Println("ID:", ordem.ID)
				fmt.Println("Cliente:", ordem.Cliente)
				fmt.Println("Descrição:", ordem.Descricao)
				fmt.Println("Valor estimado:", ordem.ValorEstimado)

				fmt.Println("Data de entrega:", ordem.DataDeEntrega.Format("02/01/2006"))

				dias := service.CalcularDiasRestantes(ordem.DataDeEntrega)

				fmt.Println("Dias restantes:", dias)

				fmt.Println("Status:", ordem.Status)
				fmt.Println("Valor final:", ordem.ValorFinal)
				fmt.Println("-----------------------------")
			}

		case "3":

			fmt.Println("Buscar OS")

			fmt.Println("Digite o ID da OS:")
			var id int
			fmt.Fscan(leitor, &id)

			leitor.ReadString('\n')

			ordem, err := servico.BuscarPorID(id)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("\nOS encontrada:")
			fmt.Println("ID:", ordem.ID)
			fmt.Println("Cliente:", ordem.Cliente)
			fmt.Println("Descrição:", ordem.Descricao)
			fmt.Println("Valor estimado:", ordem.ValorEstimado)
			fmt.Println("Status:", ordem.Status)
			fmt.Println("Valor final:", ordem.ValorFinal)
			fmt.Println("-----------------------------")

		case "4":
			fmt.Println("\n===== ATUALIZAR OS =====")

			fmt.Print("Digite o ID da OS: ")

			var id int
			fmt.Fscan(leitor, &id)
			leitor.ReadString('\n')

			ordem, err := servico.BuscarPorID(id)

			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("Os encontrada:", ordem)

			fmt.Print("Digite o novo cliente:")
			cliente, _ := leitor.ReadString('\n')
			cliente = strings.TrimSpace(cliente)

			fmt.Print("Digite a nova descrição:")
			descricao, _ := leitor.ReadString('\n')
			descricao = strings.TrimSpace(descricao)

			fmt.Print("Digite o valor estimado:")
			valorTexto, _ := leitor.ReadString('\n')
			valorTexto = strings.TrimSpace(valorTexto)

			valorEstimado, err := strconv.ParseFloat(valorTexto, 64)
			if err != nil {
				fmt.Println("Valor inválido")
				break
			}

			fmt.Print("Digite o novo status:")
			status, _ := leitor.ReadString('\n')
			status = strings.TrimSpace(status)

			fmt.Print("Digite o valor final:")
			valorFinalTexto, _ := leitor.ReadString('\n')
			valorFinalTexto = strings.TrimSpace(valorFinalTexto)

			valorFinal, err := strconv.ParseFloat(valorFinalTexto, 64)
			if err != nil {
				fmt.Println("Valor final inválido!")
				break
			}

			ordem.Cliente = cliente
			ordem.Descricao = descricao
			ordem.ValorEstimado = valorEstimado
			ordem.Status = status
			ordem.ValorFinal = valorFinal

			err = servico.Atualizar(ordem)

			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("OS atualizada com sucesso!")

		case "5":
			fmt.Println("\n===== DELETAR OS =====")

			fmt.Print("Digite o ID da OS: ")

			var id int
			fmt.Fscan(leitor, &id)

			leitor.ReadString('\n')

			ordem, err := servico.BuscarPorID(id)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("\nOS encontrada:")
			fmt.Println("ID:", ordem.ID)
			fmt.Println("Cliente:", ordem.Cliente)
			fmt.Println("Descrição:", ordem.Descricao)

			fmt.Print("Tem certeza que deseja deletar? (s/n): ")
			resposta, _ := leitor.ReadString('\n')
			resposta = strings.TrimSpace(resposta)

			if resposta != "S" && resposta != "s" {
				fmt.Println("Opção cancelada.")
				break
			}

			err = servico.Deletar(id)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("OS deletada com sucesso!")

		case "0":
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
