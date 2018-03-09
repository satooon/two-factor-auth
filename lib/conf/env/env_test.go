package env_test

import (
	"testing"
	"github.com/satooon/two-factor-auth/lib/conf/env"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestApp(t *testing.T) {
	mail := "test_app@mail.com"
	os.Setenv("APP_ADMIN_MAIL", mail)
	assert.Equal(t, env.App().GetAdminMail(), mail)
}
