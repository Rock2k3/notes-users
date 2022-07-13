package adapters

import (
	"github.com/google/uuid"
	"notes-users/config"
	"notes-users/internal/domain/users"
)

const (
	GetUserByUuidQuery = "select user_id, user_name from users where user_id=$1"
	AddUserQuery       = "insert into users(user_name) values($1) returning user_id"
)

type usersPostgres struct {
	config *config.AppConfig
}

func NewUsersPostgres(c *config.AppConfig) *usersPostgres {
	return &usersPostgres{config: c}
}

func (u usersPostgres) AddUser(userName string) (*users.MyUser, error) {
	db := NewPostgresDb(u.config).GetDb()
	defer db.Close()

	usr := users.MyUser{Name: userName}

	err := db.QueryRow(AddUserQuery, userName).Scan(&usr.UserId)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (u usersPostgres) GetUserById(uuid uuid.UUID) (*users.MyUser, error) {
	db := NewPostgresDb(u.config).GetDb()
	defer db.Close()

	usr := users.MyUser{}
	err := db.QueryRow(GetUserByUuidQuery, uuid).Scan(&usr.UserId, &usr.Name)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
