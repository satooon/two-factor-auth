package lib_test

import (
	"testing"

	"github.com/satooon/two-factor-auth/lib"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	if err := env.Load(); err != nil {
		panic(err)
	}

	db, err := database.Open(env.Database())
	assert.Nil(t, err)
	defer db.Close()
	db.AutoMigrate(database.Migrate...)

	assert.Nil(t, lib.Init(repository.NewUser(db)))
}
