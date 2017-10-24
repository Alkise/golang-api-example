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

// UsersCount Return users records count
func UsersCount() (count int, err error) {
	err = database.Pool.QueryRow("SELECT COUNT(*) as count from users").Scan(&count)
	return count, err
}

// AllUsers Returns all users from database
func AllUsers() (*UserCollection, error) {
	usersCount, err := UsersCount()
	if err != nil {
		return nil, err
	}

	return make(UserCollection, 0, usersCount).All()
}

// All Returns all users from database
func (users UserCollection) All() (*UserCollection, error) {
	rows, err := database.Pool.Query(selectAllUsersSQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

loop:
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Model.ID, &user.Email, &user.FirstName, &user.LastName, &user.Model.CreatedAt, &user.Model.UpdatedAt)
		if err != nil {
			break loop
		}
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
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(&user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Destroy Delete user from database
func (user *User) Destroy() (*User, error) {
	return user, nil
}
