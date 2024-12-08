package postgres

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	conn := Connection{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "0000",
		Database: "postgres",
		SSLMode:  SSLMode("disable"),
	}

	err := os.Setenv("CONFIG_ALLOW_MIGRATION", "false")
	if err != nil {
		return
	}

	db := InitDatabase(conn)
	require.NotNil(t, db)
}
