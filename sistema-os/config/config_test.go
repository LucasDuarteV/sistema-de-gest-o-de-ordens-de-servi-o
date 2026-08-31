package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarregar(t *testing.T) {
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "123456")
	os.Setenv("POSTGRES_DB", "sistema_os")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5433")

	cfg := Carregar()

	assert.Equal(t, "postgres", cfg.PostgresUser)
	assert.Equal(t, "123456", cfg.PostgresPassword)
	assert.Equal(t, "sistema_os", cfg.PostgresDB)
	assert.Equal(t, "localhost", cfg.PostgresHost)
	assert.Equal(t, "5433", cfg.PostgresPort)
}
