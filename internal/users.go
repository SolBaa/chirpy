package database

import "errors"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ResponseUsers struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

//
//func (db *DB) CreateUser(email string, hash string) (ResponseUsers, error) {
//	dbStructure, err := db.loadDB()
//	if err != nil {
//		return ResponseUsers{}, err
//	}
//
//	id := len(dbStructure.Users) + 1
//	user := User{
//		ID:       id,
//		Email:    email,
//		Password: hash,
//	}
//	dbStructure.Users[email] = user
//
//	err = db.writeDB(dbStructure)
//	if err != nil {
//		return ResponseUsers{}, err
//	}
//
//	return ResponseUsers{
//		ID:    user.ID,
//		Email: user.Email,
//	}, nil
//}
//
//func (db *DB) GetUserByEmail(email string) (User, error) {
//	DBStructure, err := db.loadDB()
//	if err != nil {
//		return User{}, err
//	}
//
//	user, ok := DBStructure.Users[email]
//	if !ok {
//		return User{}, errors.New("User not found")
//	}
//	return user, nil
//}
//
//func (db *DB) UpdateUser(userIDInt int, email string, hashedPassword string) (ResponseUsers, error) {
//	dbStructure, err := db.loadDB()
//	if err != nil {
//		return ResponseUsers{}, err
//	}
//
//	user := User{
//		ID:       userIDInt,
//		Email:    email,
//		Password: hashedPassword,
//	}
//	dbStructure.Users[email] = user
//
//	err = db.writeDB(dbStructure)
//	if err != nil {
//		return ResponseUsers{}, err
//	}
//
//	return ResponseUsers{
//		ID:    user.ID,
//		Email: user.Email,
//	}, nil
//}

var ErrAlreadyExists = errors.New("already exists")

func (db *DB) CreateUser(email, hashedPassword string) (User, error) {
	if _, err := db.GetUserByEmail(email); !errors.Is(err, ErrNotExist) {
		return User{}, ErrAlreadyExists
	}

	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	id := len(dbStructure.Users) + 1
	user := User{
		ID:       id,
		Email:    email,
		Password: hashedPassword,
	}
	dbStructure.Users[id] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *DB) GetUser(id int) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, ok := dbStructure.Users[id]
	if !ok {
		return User{}, ErrNotExist
	}

	return user, nil
}

func (db *DB) GetUserByEmail(email string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	for _, user := range dbStructure.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return User{}, ErrNotExist
}

func (db *DB) UpdateUser(id int, email, hashedPassword string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, ok := dbStructure.Users[id]
	if !ok {
		return User{}, ErrNotExist
	}

	user.Email = email
	user.Password = hashedPassword
	dbStructure.Users[id] = user

	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
