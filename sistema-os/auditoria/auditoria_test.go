package auditoria

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestRegistrar(t *testing.T) {
	arquivo := "auditoria.log"

	os.Remove(arquivo)

	Registrar("teste de auditoria")

	conteudo, err := os.ReadFile(arquivo)
	if err != nil {
		t.Fatalf("erro ao ler arquivo de auditoria: %v", err)
	}

	if !strings.Contains(string(conteudo), "teste de auditoria") {
		t.Errorf("mensagem não encontrada no arquivo de auditoria")
	}

	os.Remove(arquivo)
}

func TestOSCriada(t *testing.T) {
	arquivo := "auditoria.log"

	os.Remove(arquivo)

	OSCriada(100)

	time.Sleep(100 * time.Millisecond)

	conteudo, err := os.ReadFile(arquivo)
	if err != nil {
		t.Fatalf("erro ao ler arquivo de auditoria: %v", err)
	}

	if !strings.Contains(string(conteudo), "OS 100 criada") {
		t.Errorf("registro de criação não encontrado")
	}

	os.Remove(arquivo)
}

func TestOSAtualizada(t *testing.T) {
	arquivo := "auditoria.log"

	os.Remove(arquivo)

	OSAtualizada(101)

	time.Sleep(100 * time.Millisecond)

	conteudo, err := os.ReadFile(arquivo)
	if err != nil {
		t.Fatalf("erro ao ler arquivo de auditoria: %v", err)
	}

	if !strings.Contains(string(conteudo), "OS 101 atualizada") {
		t.Errorf("registro de atualização não encontrado")
	}

	os.Remove(arquivo)
}

func TestOSDeletada(t *testing.T) {
	arquivo := "auditoria.log"

	os.Remove(arquivo)

	OSDeletada(102)

	time.Sleep(100 * time.Millisecond)

	conteudo, err := os.ReadFile(arquivo)
	if err != nil {
		t.Fatalf("erro ao ler arquivo de auditoria: %v", err)
	}

	if !strings.Contains(string(conteudo), "OS 102 deletada") {
		t.Errorf("registro de deleção não encontrado")
	}

	os.Remove(arquivo)
}
