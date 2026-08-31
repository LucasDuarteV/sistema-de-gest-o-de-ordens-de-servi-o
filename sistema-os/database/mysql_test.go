package database

import (
	"testing"

	"sistema-os/config"

	"github.com/stretchr/testify/require"
)

func TestConectarMySQL(t *testing.T) {
	cfg := config.Carregar()

	db, err := ConectarMySQL(cfg)

	require.NoError(t, err)
	require.NotNil(t, db)

	defer db.Close()

	t.Log("Conexão com MySQL realizada com sucesso!")
}