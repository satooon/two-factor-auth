package env

type _session struct{
	Name string `default:"SessionName"`
	Store SessionStore `default:"cookie"`
	Option map[string]string
	Secret string `default:"secret"`
}

type SessionStore string

const (
	SessionStoreCookie SessionStore = "cookie"
	SessionStoreMemcache SessionStore = "memcache"
	SessionStoreRedis SessionStore = "redis"
)

var session *_session

func Session() *_session {
	if session == nil {
		session = process("session", &_session{}).(*_session)
	}
	return session
}

func (e *_session) GetName() string {
	return e.Name
}

func (e *_session) GetStore() SessionStore {
	return e.Store
}

func (e *_session) GetOption() map[string]string {
	return e.Option
}

func (e *_session) GetSecret() string {
	return e.Secret
}
