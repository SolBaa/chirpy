package database

import (
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type responseUsers struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (db *DB) CreateUser(email string, hash string) (responseUsers, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return responseUsers{}, err
	}

	id := len(dbStructure.Users) + 1
	user := User{
		ID:       id,
		Email:    email,
		Password: hash,
	}
	dbStructure.Users[email] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return responseUsers{}, err
	}

	return responseUsers{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (db *DB) GetUserByEmail(email string) (User, error) {
	DBStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, ok := DBStructure.Users[email]
	if !ok {
		return User{}, errors.New("User not found")
	}
	return user, nil
}
