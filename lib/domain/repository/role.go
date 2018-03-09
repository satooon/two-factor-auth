package repository

import (
	"github.com/pkg/errors"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
)

type Role struct {
	store entity.RoleSlice
}

type iUser interface {
	GetRoleName() entity.RoleName
}

func NewRole() *Role {
	return &Role{
		store: entity.RoleSlice{
			entity.NewRole(
				entity.RoleAdmin,
				entity.PermissionSlice{entity.AdminUserSignUp},
			),
			entity.NewRole(
				entity.RoleUser,
				entity.PermissionSlice{},
			),
		},
	}
}

func (r *Role) NewEntity(name entity.RoleName, permission ...entity.Permission) (e *entity.Role, err error) {
	err = errors.New("This is not available")
	return
}

func (r *Role) FindByName(name ...entity.RoleName) entity.RoleSlice {
	m := make(map[entity.RoleName]struct{}, len(name))
	for i := range name {
		m[name[i]] = struct{}{}
	}
	return r.store.Where(func(e *entity.Role) bool {
		_, ok := m[e.Name]
		return ok
	})
}

func (r *Role) FindByUser(user ...iUser) entity.RoleSlice {
	name := make([]entity.RoleName, len(user))
	for i := range user {
		name[i] = user[i].GetRoleName()
	}
	return r.FindByName(name...)
}
