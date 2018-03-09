package env

type _app struct {
	Port       int    `default:"8080"`
	AdminPort  int    `default:"8081"`
	AdminMail  string `envconfig:"admin_mail",required:"true"`
	AdminPass  string `envconfig:"admin_pass",required:"true"`
	EncryptKey string `envconfig:"encrypt_key",required:"true"`
	EncryptIV  string `envconfig:"encrypt_iv",required:"true"`
	ProxyYaml  string `envconfig:"proxy_yaml",default:"./proxy.yaml"`
}

var app *_app

func App() *_app {
	if app == nil {
		app = process("app", &_app{}).(*_app)
	}
	return app
}

func (e *_app) GetPort() int {
	return e.Port
}

func (e *_app) GetAdminPort() int {
	return e.AdminPort
}

func (e *_app) GetAdminMail() string {
	return e.AdminMail
}

func (e *_app) GetAdminPass() string {
	return e.AdminPass
}

func (e *_app) GetEncryptKey() string {
	return e.EncryptKey
}

func (e *_app) GetEncryptIV() string {
	return e.EncryptIV
}

func (e *_app) GetProxyYaml() string {
	return e.ProxyYaml
}
