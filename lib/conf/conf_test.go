package conf_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/conf/yaml"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	assert.Nil(t, env.Load())

	pwd, err := os.Getwd()
	assert.Nil(t, err)

	path := fmt.Sprintf("%s/yaml/proxy_test.yaml", pwd)
	os.Setenv("APP_PROXY_YAML", path)
	assert.Nil(t, yaml.LoadProxy(path))

	assert.NotEqual(t, len(yaml.Proxy().GetProxyConfSlice()), 0)
}
