package users

import "github.com/google/uuid"

type MyUser struct {
	UserId uuid.UUID
	Name   string
}

type myUsers interface {
	GetUserById(uuid.UUID) (*MyUser, error)
	AddUser(string) (*MyUser, error)
}

func GetUserById(users myUsers, userId uuid.UUID) (*MyUser, error) {
	user, err := users.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(users myUsers, name string) (*MyUser, error) {
	user, err := users.AddUser(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
