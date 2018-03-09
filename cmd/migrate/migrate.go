package main

import (
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/database"
)

func main() {
	if err := env.Load(); err != nil {
		panic(err)
	}
	db, err := database.Open(env.Database())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(database.Migrate...)
}
