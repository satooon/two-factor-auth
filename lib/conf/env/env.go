package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Load(filename ...string) error {
	return godotenv.Load(filename...)
}

func process(prefix string, spec interface{}) interface{} {
	if err := envconfig.Process(prefix, spec); err != nil {
		panic(err)
	}
	return spec
}
