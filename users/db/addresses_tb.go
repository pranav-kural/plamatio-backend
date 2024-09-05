package users

import (
	"context"
	"errors"

	models "encore.app/users/models"
	utils "encore.app/users/utils"
	"encore.dev/storage/sqldb"
)

type AddressesTable struct {
	DB *sqldb.Database
}

const (
	SQL_GET_ADDRESS = `
			SELECT street, city, state, country, zip_code, user_id FROM addresses
			WHERE id = $1
	`
	SQL_GET_USER_ADDRESSES = `
			SELECT id, street, city, state, country, zip_code, user_id FROM addresses
			WHERE user_id = $1
	`
	SQL_INSERT_ADDRESS = `
			INSERT INTO addresses (street, city, state, country, zip_code, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`
	SQL_UPDATE_ADDRESS = `
			UPDATE addresses SET street = $1, city = $2, state = $3, country = $4, zip_code = $5, user_id = $6 WHERE id = $7
	`
	SQL_DELETE_ADDRESS = `
			DELETE FROM addresses WHERE id = $1
	`
)

// Retrieves an address from the database.
func (tb *AddressesTable) GetAddress(ctx context.Context, id int) (*models.Address, error) {
	a := &models.Address{ID: id}
	err := tb.DB.QueryRow(ctx, SQL_GET_ADDRESS, id).Scan(&a.Street, &a.City, &a.State, &a.Country, &a.ZipCode, &a.UserID)
	return a, err
}

// Retrieves all addresses for a user from the database.
func (tb *AddressesTable) GetUserAddresses(ctx context.Context, userID string) (*models.Addresses, error) {
	rows, err := tb.DB.Query(ctx, SQL_GET_USER_ADDRESSES, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	addresses := &models.Addresses{}
	for rows.Next() {
		a := &models.Address{}
		err = rows.Scan(&a.ID, &a.Street, &a.City, &a.State, &a.Country, &a.ZipCode, &a.UserID)
		if err != nil {
			return nil, err
		}
		addresses.Data = append(addresses.Data, a)
	}

	return addresses, nil
}

// Inserts an address into the database.
func (tb *AddressesTable) InsertAddress(ctx context.Context, newAddress *models.AddressRequestParams) (*models.Address, error) {
	// validate address data
	err := utils.ValidateNewAddressData(newAddress)
	if err != nil {
		return nil, err
	}

	a := &models.Address{
		Street:  newAddress.Street,
		City:    newAddress.City,
		State:   newAddress.State,
		Country: newAddress.Country,
		ZipCode: newAddress.ZipCode,
		UserID:  newAddress.UserID,
	}
	err = tb.DB.QueryRow(ctx, SQL_INSERT_ADDRESS, newAddress.Street, newAddress.City, newAddress.State, newAddress.Country, newAddress.ZipCode, newAddress.UserID).Scan(&a.ID)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Updates an address in the database.
func (tb *AddressesTable) UpdateAddress(ctx context.Context, newAddress *models.Address) error {
	// validate address data
	err := utils.ValidateUpdateAddressData(newAddress)
	if err != nil {
		return err
	}

	_, err = tb.DB.Exec(ctx, SQL_UPDATE_ADDRESS, newAddress.Street, newAddress.City, newAddress.State, newAddress.Country, newAddress.ZipCode, newAddress.UserID, newAddress.ID)

	return err
}

// Deletes an address from the database.
func (tb *AddressesTable) DeleteAddress(ctx context.Context, id int) error {
	// Validate ID
	if id <= 0 {
		return errors.New("invalid address ID")
	}

	_, err := tb.DB.Exec(ctx, SQL_DELETE_ADDRESS, id)
	return err
}