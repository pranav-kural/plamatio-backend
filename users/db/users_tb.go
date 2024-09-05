package users

import (
	"context"
	"errors"

	models "encore.app/users/models"
	utils "encore.app/users/utils"
	"encore.dev/storage/sqldb"
)

type UsersTable struct {
	DB *sqldb.Database
}

const (
	SQL_GET_USER = `
			SELECT first_name, last_name, email FROM users
			WHERE id = $1
	`
	SQL_GET_ALL_USERS = `
			SELECT id, first_name, last_name, email FROM users
	`
	SQL_INSERT_USER = `
			INSERT INTO users (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id
	`
	SQL_UPDATE_USER = `
			UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4
	`
	SQL_DELETE_USER = `
			DELETE FROM users WHERE id = $1
	`
)

// Retrieves a user from the database.
func (tb *UsersTable) GetUser(ctx context.Context, id string) (*models.User, error) {
	u := &models.User{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_USER, id).Scan(&u.FirstName, &u.LastName, &u.Email)
	return u, err
}

// Retrieves all users from the database.
func (tb *UsersTable) GetAllUsers(ctx context.Context) (*models.Users, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_ALL_USERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := &models.Users{}
	for rows.Next() {
		u := &models.User{}
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
		if err != nil {
			return nil, err
		}
		users.Data = append(users.Data, u)
	}
	return users, nil
}

// Inserts a user into the database.
func (tb *UsersTable) InsertUser(ctx context.Context, newUser *models.UserRequestParams) (*models.User, error) {
	// validate user data
	err := utils.ValidateNewUserData(newUser)
	if err != nil {
		return nil, err
	}

	u := &models.User{FirstName: newUser.FirstName, LastName: newUser.LastName, Email: newUser.Email}
	err = tb.DB.QueryRow(ctx, SQL_INSERT_USER, u.FirstName, u.LastName, u.Email).Scan(&u.ID)
	return u, err
}

// Updates a user in the database.
func (tb *UsersTable) UpdateUser(ctx context.Context, updatedUser *models.User) error {
	// validate user data
	err := utils.ValidateUpdateUserData(updatedUser)
	if err != nil {
		return err
	}

	_, err = tb.DB.Exec(ctx, SQL_UPDATE_USER, updatedUser.FirstName, updatedUser.LastName, updatedUser.Email, updatedUser.ID)
	return err
}

// Deletes a user from the database.
func (tb *UsersTable) DeleteUser(ctx context.Context, id string) error {
	// Validate ID
	if id == "" {
		return errors.New("invalid user ID")
	}

	_, err := tb.DB.Exec(ctx, SQL_DELETE_USER, id)
	return err
}
