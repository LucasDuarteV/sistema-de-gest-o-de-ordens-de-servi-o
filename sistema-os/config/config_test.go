package config

import "testing"

func TestCarregar(t *testing.T) {
	config := Carregar()

	if config.PostgresUser == "" {
		t.Error("POSTGRES_USER não foi carregado")
	}

	if config.PostgresPassword == "" {
		t.Error("POSTGRES_PASSWORD não foi carregado")
	}

	if config.PostgresDB == "" {
		t.Error("POSTGRES_DB não foi carregado")
	}

	if config.PostgresHost == "" {
		t.Error("POSTGRES_HOST não foi carregado")
	}

	if config.PostgresPort == "" {
		t.Error("POSTGRES_PORT não foi carregado")
	}
}
