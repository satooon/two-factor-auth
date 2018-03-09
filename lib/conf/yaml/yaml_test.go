package yaml_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/satooon/two-factor-auth/lib/conf/yaml"
	"github.com/stretchr/testify/assert"
)

func TestLoadProxy(t *testing.T) {
	pwd, err := os.Getwd()
	assert.Nil(t, err)
	assert.Nil(t, yaml.LoadProxy(fmt.Sprintf("%s/proxy_test.yaml", pwd)))
	assert.NotEqual(t, len(yaml.Proxy().GetProxyConfSlice()), 0)

	for _, v := range yaml.Proxy().GetProxyConfSlice() {
		assert.NotEqual(t, v.Name, "")
	}
}
