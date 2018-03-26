package repository

import (
	"fmt"
	"math"
	"strconv"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/satooon/two-factor-auth/lib/crypto/aes"
	"github.com/satooon/two-factor-auth/lib/crypto/scrypt"
	"github.com/satooon/two-factor-auth/lib/database"
	"github.com/satooon/two-factor-auth/lib/domain/entity"
)

type User struct {
	db database.IDatabase
}

type iSecret interface {
	GetEncryptKey() string
	GetEncryptIV() string
}

type iAuth interface {
	GetIssuer() string
	GetPeriod() uint
	GetSecretSize() uint
	GetDigits() otp.Digits
	GetAlgorithm() otp.Algorithm
}

func NewUser(db database.IDatabase) *User {
	return &User{
		db: db,
	}
}

func (r *User) NewEntity(secret iSecret, auth iAuth, mail string, pass string, roleName entity.RoleName) (e *entity.User, err error) {
	m, err := aes.Encrypt(secret.GetEncryptKey(), secret.GetEncryptIV(), mail)
	if err != nil {
		return
	}
	s, err := scrypt.NewSalt()
	if err != nil {
		return
	}
	p, err := scrypt.Hash(pass, s)
	if err != nil {
		return
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      auth.GetIssuer(),
		AccountName: mail,
		Period:      auth.GetPeriod(),
		SecretSize:  auth.GetSecretSize(),
		Digits:      auth.GetDigits(),
		Algorithm:   auth.GetAlgorithm(),
	})
	if err != nil {
		return
	}
	sec, err := aes.Encrypt(secret.GetEncryptKey(), secret.GetEncryptIV(), key.Secret())
	if err != nil {
		return
	}
	e = entity.NewUser(m, s, p, sec, roleName)
	return
}

func (r *User) Count() (int, error) {
	var c int
	err := r.db.Table(entity.NewUser("", "", "", "", entity.RoleAdmin).TableName()).Count(&c).Error()
	return c, err
}

func (r *User) FindByID(id ...int) (entity.UserSlice, error) {
	s := entity.UserSlice{}
	err := r.db.Where("id in (?)", id).Find(&s).Error()
	return s, err
}

func (r *User) FindByMail(secret iSecret, mail ...string) (s entity.UserSlice, err error) {
	args := make([]interface{}, len(mail))
	for i := range mail {
		args[i], err = aes.Encrypt(secret.GetEncryptKey(), secret.GetEncryptIV(), mail[i])
		if err != nil {
			return
		}
	}
	s = entity.UserSlice{}
	err = r.db.Where("mail in (?)", args).Find(&s).Error()
	return
}

func (r *User) Transaction(f func() error) error {
	t := r.db.Begin()
	if err := f(); err != nil {
		t.Rollback()
		return err
	}
	return t.Commit().Error()
}

func (r *User) Save(e *entity.User) error {
	if r.db.NewRecord(e) {
		return r.db.Create(e).Error()
	} else {
		return r.db.Update(e).Error()
	}
}

func (r *User) Delete(e *entity.User) error {
	return r.db.Delete(e).Error()
}

func (r *User) Paginate(limit, offset interface{}) (*userPaginate, error) {
	var slice entity.UserSlice
	if err := r.db.Limit(limit).Offset(offset).Find(&slice).Error(); err != nil {
		return nil, err
	}
	count, err := r.Count()
	if err != nil {
		return nil, err
	}
	return &userPaginate{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Slice:  slice,
	}, nil
}

type userPaginate struct {
	Count  int
	Limit  interface{}
	Offset interface{}
	Slice  entity.UserSlice
}

func (u *userPaginate) offset() int {
	offset, _ := strconv.Atoi(fmt.Sprintf("%v", u.Offset))
	return offset
}

func (u *userPaginate) limit() int {
	limit, _ := strconv.Atoi(fmt.Sprintf("%v", u.Limit))
	return limit
}

func (u *userPaginate) Pages() map[int]int {
	limit := u.limit()
	size := int(math.Ceil(float64(u.Count) / float64(limit)))
	p := make(map[int]int, size)
	for i := 0; i < size; i++ {
		idx := i * limit
		p[idx] = i + 1
	}
	return p
}

func (u *userPaginate) IsCurrent(p int) bool {
	return u.offset() == p
}

func (u *userPaginate) Prev() int {
	offset := u.offset()
	if offset <= 0 {
		return 0
	}
	limit := u.limit()
	return offset - limit
}

func (u *userPaginate) Next() int {
	offset := u.offset()
	limit := u.limit()
	if offset+limit >= u.Count {
		return offset
	}
	return offset + limit
}
