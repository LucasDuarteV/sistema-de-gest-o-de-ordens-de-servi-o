package auditoria

import (
	"fmt"
	"os"
	"time"
)

func Registrar(mensagem string) {
	arquivo, err := os.OpenFile(
		"auditoria.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		fmt.Println("Erro ao abrir arquivo de auditoria:", err)
		return
	}

	defer arquivo.Close()

	data := time.Now().Format("02/01/2006 15:04:05")
	fmt.Fprintf(arquivo, "%s - %s\n", data, mensagem)
}

func OSCriada(id int) {
	go Registrar(fmt.Sprintf("OS %d criada", id))
}

func OSAtualizada(id int) {
	go Registrar(fmt.Sprintf("OS %d atualizada", id))
}

func OSDeletada(id int) {
	go Registrar(fmt.Sprintf("OS %d deletada", id))
}
