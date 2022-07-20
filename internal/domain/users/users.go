package users

import "github.com/google/uuid"

type MyUser struct {
	UserUUID uuid.UUID
	Name     string
}

type myUsers interface {
	GetUserByUUID(uuid.UUID) (*MyUser, error)
	GetUserByName(string) (*MyUser, error)
	AddUser(string) (*MyUser, error)
}

func GetUserByUUID(users myUsers, userUUID uuid.UUID) (*MyUser, error) {
	user, err := users.GetUserByUUID(userUUID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(users myUsers, userName string) (*MyUser, error) {
	user, err := users.GetUserByName(userName)
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
