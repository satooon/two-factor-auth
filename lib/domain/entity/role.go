//go:generate gen
package entity

// +gen slice:"Where,All,Any,First"
type Permission string

const (
	AdminUserSignUp Permission = "admin.user.sign_up"
)

type RoleName string

const (
	RoleAdmin RoleName = "admin"
	RoleUser  RoleName = "user"
)

func (r RoleName) String() string {
	return string(r)
}

// +gen * slice:"Where,All,Any,First"
type Role struct {
	Name        RoleName
	Permissions PermissionSlice
}

func NewRole(name RoleName, permissions PermissionSlice) *Role {
	return &Role{
		Name:        name,
		Permissions: permissions,
	}
}

func (r *Role) ContainsPermission(ps ...Permission) bool {
	m := make(map[Permission]struct{}, len(ps))
	for i := range ps {
		m[ps[i]] = struct{}{}
	}
	return r.Permissions.Any(func(p Permission) bool {
		_, ok := m[p]
		return ok
	})
}

func (r *Role) PermissionAdminUserSignUp() bool {
	return r.ContainsPermission(AdminUserSignUp)
}
