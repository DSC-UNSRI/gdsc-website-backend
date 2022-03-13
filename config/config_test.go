package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigNew(t *testing.T) {
	config := New("../.env")
	require.NotEmpty(t, config)
	require.IsType(t, "", config.PostgresDSN)
}
