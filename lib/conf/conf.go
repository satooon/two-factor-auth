package conf

import (
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/conf/yaml"
)

func Load(filename ...string) error {
	if err := env.Load(filename...); err != nil {
		return err
	}
	if err := yaml.LoadProxy(env.App().GetProxyYaml()); err != nil {
		return err
	}
	return nil
}
