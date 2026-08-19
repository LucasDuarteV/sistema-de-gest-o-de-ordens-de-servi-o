package database

import (
	"testing"

	"sistema-os/config"
)

func TestConectar(t *testing.T) {
	cfg := config.Carregar()

	conn, err := Conectar(cfg)
	if err != nil {
		t.Fatalf("erro ao conectar no banco: %v", err)
	}

	defer conn.Close(t.Context())

	if err := conn.Ping(t.Context()); err != nil {
		t.Fatalf("erro ao testar conexão: %v", err)
	}
}
