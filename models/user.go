package models

import (
	"fmt"
	"github.com/stretchr/gomniauth/common"
	"github.com/astaxie/beego/orm"
)

var DefaultUserManager *UserManager

type User struct {
    ID            int64
    Name          string `orm:"size(50)"`
    Email         string `orm:"size(100)"`
    Nickname      string `orm:"size(50)"` 
    AvatarURL     string `orm:"size(200)"`
    ProviderName  string `orm:"size(50)"`
    IDForProvider string `orm:"size(100)"`
    AuthCode      string `orm:"size(100)"`
}

// NewUser creates a new user given a user object and provider name.
func NewUser(user common.User, ProviderName string) (*User, error) {
	if ProviderName == "" {
		return nil, fmt.Errorf("empty provider name")
	}
	return &User{0, user.Name(), user.Email(), user.Nickname(), user.AvatarURL(), ProviderName, user.IDForProvider(ProviderName), user.AuthCode()}, nil
}

// UserManager manages a list of users in memory.
type UserManager struct {
	users  []*User
	lastID int64
}

// NewUserManager returns an empty UserManager.
func NewUserManager() *UserManager {
	return &UserManager{}
}

// Save saves the given User in the UserManager.
func (m *UserManager) Save(user *User) error {
	if user.ID == 0 {
        o := orm.NewOrm() 
        id, err := o.Insert(user)
        if err != nil {
            return err
        }
        user.ID = id
		return nil
	}
    // Update?
	return fmt.Errorf("unknown user")
}

// Find returns the User with the given id in the UserManager and a boolean
// indicating if the id was found.
func (m *UserManager) Find(ID int64) (*User, bool) {
    u := User{ID: ID}
    o := orm.NewOrm()
    err := o.Read(&u)
    if err != nil {
    	return nil, false
    }
	return &u, true
}

func init() {
	DefaultUserManager = NewUserManager()
}
