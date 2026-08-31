package database

import (
	"testing"

	"sistema-os/config"

	"github.com/stretchr/testify/require"
)

func TestConectar(t *testing.T) {
	cfg := config.Carregar()

	conn, err := Conectar(cfg)

	require.NoError(t, err, "Erro ao conectar ao banco de dados")
	require.NotNil(t, conn, "Conexão com o banco de dados é nula")

	defer conn.Close(t.Context())

	err = conn.Ping(t.Context())

	require.NoError(t, err, "Erro ao testar conexão com o banco de dados")

}
