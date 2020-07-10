package kv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/distributed_lab/kit/kv"
)

func TestFromEnv(t *testing.T) {
	t.Run("no backends", func(t *testing.T) {
		_, err := kv.FromEnv()
		assert.Equal(t, kv.ErrNoBackends, err)
	})

	t.Run("valid viper", func(t *testing.T) {
		if err := os.Setenv(kv.EnvViperConfigFile, "testdata/empty.yml"); err != nil {
			t.Fatalf("failed to set env: %v", err)
		}
		_, err := kv.FromEnv()
		assert.NoError(t, err)
	})

	t.Run("missing viper file", func(t *testing.T) {
		if err := os.Setenv(kv.EnvViperConfigFile, "404.yml"); err != nil {
			t.Fatalf("failed to set env: %v", err)
		}
		_, err := kv.FromEnv()
		assert.Error(t, err)
	})
}
