package database

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satooon/two-factor-auth/lib/conf/env"
)

const (
	key = "db"
)

type DB struct {
	db *gorm.DB
}

type IDatabase interface {
	Where(query interface{}, args ...interface{}) IDatabase
	Or(query interface{}, args ...interface{}) IDatabase
	Not(query interface{}, args ...interface{}) IDatabase
	Limit(limit interface{}) IDatabase
	Offset(offset interface{}) IDatabase
	Order(value interface{}, reorder ...bool) IDatabase
	Select(query interface{}, args ...interface{}) IDatabase
	Attrs(attrs ...interface{}) IDatabase
	First(out interface{}, where ...interface{}) IDatabase
	Take(out interface{}, where ...interface{}) IDatabase
	Last(out interface{}, where ...interface{}) IDatabase
	Find(out interface{}, where ...interface{}) IDatabase
	Count(value interface{}) IDatabase
	FirstOrCreate(out interface{}, where ...interface{}) IDatabase
	Update(attrs ...interface{}) IDatabase
	Updates(values interface{}, ignoreProtectedAttrs ...bool) IDatabase
	Save(value interface{}) IDatabase
	Create(value interface{}) IDatabase
	Delete(value interface{}, where ...interface{}) IDatabase
	Raw(sql string, values ...interface{}) IDatabase
	Exec(sql string, values ...interface{}) IDatabase
	Table(name string) IDatabase
	Begin() IDatabase
	Commit() IDatabase
	Rollback() IDatabase
	NewRecord(value interface{}) bool
	GetErrors() []error
	Error() error
}

type context interface {
	Set(key string, value interface{})
	Get(key string) (value interface{}, exists bool)
	MustGet(key string) interface{}
}

type iEnvDatabase interface {
	GetDriver() env.DatabaseDriver
	GetArgs() string
	GetMaxIdleConns() int
	GetMaxOpenConns() int
	GetDebug() bool
	GetEngine() string
}

func Open(e iEnvDatabase) (*DB, error) {
	db, err := gorm.Open(e.GetDriver().String(), e.GetArgs())
	if err != nil {
		return nil, err
	}
	d := newDB(db)
	d.db.LogMode(e.GetDebug())
	if len(e.GetEngine()) > 0 {
		db.Set("gorm:table_options", fmt.Sprintf("ENGINE=%s", e.GetEngine()))
	}
	d.db.DB().SetMaxIdleConns(e.GetMaxIdleConns())
	d.db.DB().SetMaxOpenConns(e.GetMaxOpenConns())
	return d, nil
}

func newDB(d *gorm.DB) *DB {
	return &DB{db: d}
}

func SetContext(c context, db *DB) {
	c.Set(key, db)
}

func Default(c context) *DB {
	return c.MustGet(key).(*DB)
}

func (d *DB) DB() *sql.DB {
	return d.db.DB()
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) AutoMigrate(values ...interface{}) *DB {
	d.db = d.db.AutoMigrate(values...)
	return d
}

func (d *DB) Where(query interface{}, args ...interface{}) IDatabase {
	d.db = d.db.Where(query, args...)
	return d
}

func (d *DB) Or(query interface{}, args ...interface{}) IDatabase {
	d.db = d.db.Or(query, args...)
	return d
}

func (d *DB) Not(query interface{}, args ...interface{}) IDatabase {
	d.db = d.db.Not(query, args...)
	return d
}

func (d *DB) Limit(limit interface{}) IDatabase {
	d.db = d.db.Limit(limit)
	return d
}

func (d *DB) Offset(offset interface{}) IDatabase {
	d.db = d.db.Offset(offset)
	return d
}

func (d *DB) Order(value interface{}, reorder ...bool) IDatabase {
	d.db = d.db.Order(value, reorder...)
	return d
}

func (d *DB) Select(query interface{}, args ...interface{}) IDatabase {
	d.db = d.db.Select(query, args...)
	return d
}
func (d *DB) Attrs(attrs ...interface{}) IDatabase {
	d.db = d.db.Attrs(attrs...)
	return d
}

func (d *DB) First(out interface{}, where ...interface{}) IDatabase {
	d.db = d.db.First(out, where...)
	return d
}

func (d *DB) Take(out interface{}, where ...interface{}) IDatabase {
	d.db = d.db.Take(out, where...)
	return d
}

func (d *DB) Last(out interface{}, where ...interface{}) IDatabase {
	d.db = d.db.Last(out, where...)
	return d
}

func (d *DB) Find(out interface{}, where ...interface{}) IDatabase {
	d.db = d.db.Find(out, where...)
	return d
}

func (d *DB) Count(value interface{}) IDatabase {
	d.db = d.db.Count(value)
	return d
}

func (d *DB) FirstOrCreate(out interface{}, where ...interface{}) IDatabase {
	d.db = d.db.FirstOrCreate(out, where...)
	return d
}

func (d *DB) Update(attrs ...interface{}) IDatabase {
	d.db = d.db.Update(attrs)
	return d
}

func (d *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) IDatabase {
	d.db = d.db.Updates(values, ignoreProtectedAttrs...)
	return d
}

func (d *DB) Save(value interface{}) IDatabase {
	d.db = d.db.Save(value)
	return d
}

func (d *DB) Create(value interface{}) IDatabase {
	d.db = d.db.Create(value)
	return d
}

func (d *DB) Delete(value interface{}, where ...interface{}) IDatabase {
	d.db = d.db.Delete(value)
	return d
}

func (d *DB) Raw(sql string, values ...interface{}) IDatabase {
	d.db = d.db.Raw(sql, values...)
	return d
}

func (d *DB) Exec(sql string, values ...interface{}) IDatabase {
	d.db = d.db.Exec(sql, values...)
	return d
}

func (d *DB) Table(name string) IDatabase {
	d.db = d.db.Table(name)
	return d
}

func (d *DB) Begin() IDatabase {
	return newDB(d.db.Begin())
}

func (d *DB) Commit() IDatabase {
	d.db.Commit()
	return d
}

func (d *DB) Rollback() IDatabase {
	d.db.Rollback()
	return d
}

func (d *DB) NewRecord(value interface{}) bool {
	return d.db.NewRecord(value)
}

func (d *DB) GetErrors() []error {
	return d.db.GetErrors()
}

func (d *DB) Error() error {
	return d.db.Error
}
