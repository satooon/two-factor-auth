package env

type _database struct{
	Driver DatabaseDriver `default:"sqlite3"`
	Args string `default:"./auth.db"`
	MaxIdleConns int `default:100,envconfig:"max_idle_conns"`
	MaxOpenConns int `default:200,envconfig:"max_open_conns"`
	Debug bool
	Engine string
}

type DatabaseDriver string

func (s DatabaseDriver) String() string {
	return string(s)
}

const (
	DatabaseDriverMySQL DatabaseDriver = "mysql"
	DatabaseDriverSQLite DatabaseDriver = "sqlite3"
)

var database *_database

func Database() *_database {
	if database == nil {
		database = process("db", &_database{}).(*_database)
	}
	return database
}

func (e *_database) GetDriver() DatabaseDriver {
	return e.Driver
}

func (e *_database) GetArgs() string {
	return e.Args
}

func (e *_database) GetMaxIdleConns() int {
	return e.MaxIdleConns
}

func (e *_database) GetMaxOpenConns() int {
	return e.MaxOpenConns
}

func (e *_database) GetDebug() bool {
	return e.Debug
}

func (e *_database) GetEngine() string {
	return e.Engine
}