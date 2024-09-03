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
			SELECT first_name, last_name FROM users
			WHERE id = $1
	`
	SQL_INSERT_USER = `
			INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id
	`
	SQL_UPDATE_USER = `
			UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3
	`
	SQL_DELETE_USER = `
			DELETE FROM users WHERE id = $1
	`
)

// Retrieves a user from the database.
func (tb *UsersTable) GetUser(ctx context.Context, id int) (*models.User, error) {
	u := &models.User{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_USER, id).Scan(&u.FirstName, &u.LastName)
	return u, err
}

// Inserts a user into the database.
func (tb *UsersTable) InsertUser(ctx context.Context, firstName string, lastName string) (*models.User, error) {
	// validate user data
	err := utils.ValidateNewUserData(&models.UserRequestParams{FirstName: firstName, LastName: lastName})
	if err != nil {
		return nil, err
	}

	u := &models.User{FirstName: firstName, LastName: lastName}
	err = tb.DB.QueryRow(ctx, SQL_INSERT_USER, firstName, lastName).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Updates a user in the database.
func (tb *UsersTable) UpdateUser(ctx context.Context, firstName string, lastName string, id int) error {
	// validate id
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	// validate user data
	err := utils.ValidateNewUserData(&models.UserRequestParams{FirstName: firstName, LastName: lastName})
	if err != nil {
		return err
	}

	_, err = tb.DB.Exec(ctx, SQL_UPDATE_USER, firstName, lastName, id)
	return err
}

// Deletes a user from the database.
func (tb *UsersTable) DeleteUser(ctx context.Context, id int) error {
	// Validate ID
	if id <= 0 {
		return errors.New("invalid user ID")
	}

	_, err := tb.DB.Exec(ctx, SQL_DELETE_USER, id)
	return err
}
