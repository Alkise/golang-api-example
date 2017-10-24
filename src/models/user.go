package models

import (
	"database"
	"time"
)

const (
	selectUserSQL     = "SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE id=?"
	selectAllUsersSQL = "SELECT * FROM users"
	insertUserSQL     = "INSERT INTO users(email, first_name, last_name, created_at, updated_at) VALUES(?, ?, ?, ?, ?)"
)

// User model
type User struct {
	Model
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UserCollection Collection of users
type UserCollection []User

// AllUsers Returns all users from database
func AllUsers() (*UserCollection, error) {
	return UserCollection{}.All()
}

// All Returns all users from database
func (users UserCollection) All() (*UserCollection, error) {
	rows, err := database.Pool.Query(selectAllUsersSQL)
	panicOnErr(err)
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Model.ID, &user.Email, &user.FirstName, &user.LastName, &user.Model.CreatedAt, &user.Model.UpdatedAt)
		panicOnErr(err)
		users = append(users, user)
	}
	return &users, err
}

// FindUser Select user by id from database
func FindUser(id string) (*User, error) {
	return User{}.Find(id)
}

// Find Select user by id from database
func (user User) Find(id string) (*User, error) {
	err := database.Pool.QueryRow(selectUserSQL, id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Save Save entry in the database
func (user *User) Save() (*User, error) {
	user.UpdatedAt = time.Now().UTC()
	if user.IsNewRecord() {
		user.CreatedAt = user.UpdatedAt
	}
	stmt, err := database.Pool.Prepare(insertUserSQL)
	panicOnErr(err)
	defer stmt.Close()
	result, err := stmt.Exec(&user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	panicOnErr(err)
	user.ID, err = result.LastInsertId()
	panicOnErr(err)
	return user, nil
}

// Destroy Delete user from database
func (user *User) Destroy() (*User, error) {
	return user, nil
}
