package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"gopulse/backend/models"
)

type userResolver struct {
	entity *models.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.entity.ID)
}

func (r *userResolver) FirstName() string {
	return r.entity.FirstName
}

func (r *userResolver) LastName() string {
	return r.entity.LastName
}

func (r *userResolver) Email() string {
	return r.entity.Email
}

func (r *userResolver) Password() string {
	return "********"
}

func (r *userResolver) NickName() string {
	return r.entity.Nickname
}

func ResolveUsers() (result []*userResolver) {
	users, err := models.AllUsers()
	if err != nil {
		return
	}
	for _, user := range users {
		result = append(result, &userResolver{entity: user})
	}
	return
}

func ResolveUser(id string) (result *userResolver) {
	user, err := models.Find(id)
	if err {
		return nil
	}
	result = &userResolver{entity: user}
	return
}

func ResolveCreateUser(firstName string, lastName string, email string, password string, nickname string) (result *userResolver) {
	user, err := models.CreateUser(firstName, lastName, email, password, nickname)
	if err != nil {
		return nil
	}
	result = &userResolver{entity: user}
	return
}

func ResolveUpdateUser(id string, firstName string, lastName string, email string, password string, nickname string) (result *userResolver) {
	user, err := models.UpdateUser(id, firstName, lastName, email, password, nickname)
	if err != nil {
		return nil
	}
	result = &userResolver{entity: user}
	return
}

func ResolveDeleteUser(id string) *string {
	// TODO need to implement Delete user functionality
	_, err := models.DeleteUser(id)
	if err {
		return nil
	}

	return &id
}
