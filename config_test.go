package serverconfig_test

import (
	"os"
	"testing"
	"time"

	srvcfg "github.com/kiryu-dev/server-config"
	"github.com/stretchr/testify/assert"
)

func TestLoadYamlCfg(t *testing.T) {
	const cfgPath = "./example/example_config.yaml"
	exp := &srvcfg.ServerConfig{
		Addr:         "localhost:8080",
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  2 * time.Minute,
	}
	cfg, err := srvcfg.LoadYamlCfg(cfgPath)
	assert.NoError(t, err)
	assert.Equal(t, exp, cfg)
}

func TestNonexistingCfg(t *testing.T) {
	const cfgPath = "./non-existing-config.yaml"
	cfg, err := srvcfg.LoadYamlCfg(cfgPath)
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestInvalidYamlCfg(t *testing.T) {
	const pattern = "temp-srv-yamls-cfg"
	testCases := []struct {
		name   string
		cfgStr string
	}{
		{
			name: "string port",
			cfgStr: `
host: localhost
port: ":8080"
read_timeout: 1s
write_timeout: 1s
idle_timeout: 2m
`,
		},
		{
			name: "timeout without measurement unit",
			cfgStr: `
host: localhost
port: 8080
read_timeout: 1123
write_timeout: 1312
idle_timeout: 31232
`,
		},
	}
	for _, test := range testCases {
		file, err := os.CreateTemp("", pattern)
		assert.NoError(t, err, test.name)
		n, err := file.WriteString(test.cfgStr)
		assert.NoError(t, err, test.name)
		assert.Equal(t, len(test.cfgStr), n, test.name)
		cfg, err := srvcfg.LoadYamlCfg(file.Name())
		assert.Error(t, err, test.name)
		assert.Nil(t, cfg, test.name)
		os.Remove(file.Name())
	}
}
