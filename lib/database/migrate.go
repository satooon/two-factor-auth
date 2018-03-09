package database

import "github.com/satooon/two-factor-auth/lib/domain/entity"

var Migrate = []interface{}{
	&entity.User{},
}
