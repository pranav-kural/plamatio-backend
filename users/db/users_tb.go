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
			SELECT first_name, last_name, ref_id FROM users
			WHERE id = $1
	`
	SQL_GET_USER_BY_REF_ID = `
			SELECT id, first_name, last_name FROM users
			WHERE ref_id = $1
	`
	SQL_GET_ALL_USERS = `
			SELECT id, first_name, last_name, ref_id FROM users
	`
	SQL_INSERT_USER = `
			INSERT INTO users (first_name, last_name, ref_id) VALUES ($1, $2, $3) RETURNING id
	`
	SQL_UPDATE_USER = `
			UPDATE users SET first_name = $1, last_name = $2, ref_id = $3 WHERE id = $4
	`
	SQL_DELETE_USER = `
			DELETE FROM users WHERE id = $1
	`
)

// Retrieves a user from the database.
func (tb *UsersTable) GetUser(ctx context.Context, id int) (*models.User, error) {
	u := &models.User{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_USER, id).Scan(&u.FirstName, &u.LastName, &u.RefID)
	return u, err
}

// Retrieves a user from the database by reference ID.
func (tb *UsersTable) GetUserByRefID(ctx context.Context, refID string) (*models.User, error) {
	u := &models.User{RefID: refID}
	err := tb.DB.QueryRow(ctx, SQL_GET_USER_BY_REF_ID, refID).Scan(&u.ID, &u.FirstName, &u.LastName)
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
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.RefID)
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

	u := &models.User{FirstName: newUser.FirstName, LastName: newUser.LastName, RefID: newUser.RefID}
	err = tb.DB.QueryRow(ctx, SQL_INSERT_USER, u.FirstName, u.LastName, u.RefID).Scan(&u.ID)
	return u, err
}

// Updates a user in the database.
func (tb *UsersTable) UpdateUser(ctx context.Context, updatedUser *models.User) error {
	// validate user data
	err := utils.ValidateUpdateUserData(updatedUser)
	if err != nil {
		return err
	}

	_, err = tb.DB.Exec(ctx, SQL_UPDATE_USER, updatedUser.FirstName, updatedUser.LastName, updatedUser.RefID, updatedUser.ID)
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
