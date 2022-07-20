package adapters

import (
	"github.com/Rock2k3/notes-users/config"
	"github.com/Rock2k3/notes-users/internal/domain/users"
	"github.com/google/uuid"
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

func (u usersPostgres) GetUserByUUID(uuid uuid.UUID) (*users.MyUser, error) {
	db := NewPostgresDb(u.config).GetDb()
	defer db.Close()

	usr := users.MyUser{}
	err := db.QueryRow(GetUserByUuidQuery, uuid).Scan(&usr.UserUUID, &usr.Name)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (u usersPostgres) GetUserByName(name string) (*users.MyUser, error) {
	db := NewPostgresDb(u.config).GetDb()
	defer db.Close()

	usr := users.MyUser{}
	//err := db.QueryRow(GetUserByUuidQuery, uuid).Scan(&usr.UserUUID, &usr.Name)
	//if err != nil {
	//	return nil, err
	//}

	return &usr, nil
}

func (u usersPostgres) AddUser(userName string) (*users.MyUser, error) {
	db := NewPostgresDb(u.config).GetDb()
	defer db.Close()

	usr := users.MyUser{Name: userName}

	err := db.QueryRow(AddUserQuery, userName).Scan(&usr.UserUUID)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
