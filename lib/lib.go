package lib

import (
	"encoding/gob"

	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/satooon/two-factor-auth/lib/conf/yaml"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
	"github.com/satooon/two-factor-auth/lib/domain/repository"
)

func Init(repo *repository.User) error {
	gob.Register(&entity.User{})
	gob.Register(&yaml.ProxyConf{})

	// admin user find or create
	return repo.Transaction(func() error {
		s, err := repo.FindByMail(env.App(), env.App().GetAdminMail())
		if err != nil {
			return err
		}
		if len(s) > 0 {
			return nil
		}
		u, err := repo.NewEntity(env.App(), env.Auth(), env.App().GetAdminMail(), env.App().GetAdminPass(), entity.RoleAdmin)
		if err != nil {
			return err
		}
		return repo.Save(u)
	})
}
