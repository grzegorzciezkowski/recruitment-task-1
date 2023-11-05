package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	cfg := Load()

	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, true, cfg.ReleaseMode)
	assert.Equal(t, "input.txt", cfg.InputFile)
	assert.Equal(t, "debug", cfg.LogLevel)
	assert.Equal(t, "logfile.log", cfg.LogFile)
}
