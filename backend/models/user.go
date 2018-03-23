package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/gomniauth/common"
	"strconv"
)

var DefaultUserManager *UserManager

type User struct {
	ID            int64
	FirstName     string `orm:"size(128);"`
	LastName      string `orm:"size(128);"`
	Email         string `orm:"size(100)"`
	Password      string `orm:"size(256)"`
	Nickname      string `orm:"size(50)"`
	AvatarURL     string `orm:"size(200)"`
	ProviderName  string `orm:"size(50)"`
	IDForProvider string `orm:"size(100)"`
	AuthCode      string `orm:"size(100)"`
}

// NewUser creates a new user given a user object and provider name.
func CreateProviderUser(user common.User, ProviderName string) (*User, error) {
	if ProviderName == "" {
		return nil, fmt.Errorf("empty provider name")
	}
	return &User{0, user.Name(), user.Name(), user.Email(), "", user.Nickname(), user.AvatarURL(), ProviderName, user.IDForProvider(ProviderName), user.AuthCode()}, nil
}

func CreateUser(firstName string, lastName string, email string, password string, nickName string) (*User, error) {
	return &User{0, firstName, lastName, email, password, nickName, "", "", "", ""}, nil
}

func Find(ID string) (*User, bool) {
	user, err := DefaultUserManager.Find(converID(ID))
	return user, err
}
func UpdateUser(id string, firstName string, lastName string, email string, password string, nickName string) (*User, error) {
	user, err := Find(id)
	var updateError error
	if err {
		updateError = errors.New("Cannot update the User model")
	}
	return user, updateError
}
func DeleteUser(ID string) (*User, bool) {
	return Find(ID)
}
func converID(id string) int64 {
	ID, errConvert := strconv.ParseInt(id, 10, 64)
	if errConvert != nil {
		return 0
	}
	return ID
}

func AllUsers() ([]*User, error) {
	return DefaultUserManager.All()
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

func (m *UserManager) All() ([]*User, error) {
	var users []*User
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("name", "slene").All(&users)
	return users, err
}

func init() {
	DefaultUserManager = NewUserManager()
}
