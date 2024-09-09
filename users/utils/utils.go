package users

import (
	"errors"

	models "encore.app/users/models"
)

/*

For reference, here is the SQL to create the tables in the database:

CREATE TABLE users (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

CREATE TABLE addresses (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    street TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    country TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_user_id_users ON users (id);

CREATE INDEX idx_user_id_addresses ON addresses (user_id);

*/

func ValidateUserData(user *models.User) error {
	if user.ID == "" {
		return errors.New("invalid user ID")
	}
	if user.FirstName == "" {
		return errors.New("missing first name")
	}
	if user.LastName == "" {
		return errors.New("missing last name")
	}
	if user.Email == "" {
		return errors.New("missing user email")
	}
	return nil
}

func ValidateUpdateAddressData(address *models.Address) error {
	// validate id
	if address.ID <= 0 {
		return errors.New("invalid address ID")
	}
	return ValidateNewAddressData(&models.AddressRequestParams{
		Street:   address.Street,
		City:     address.City,
		State:    address.State,
		Country:  address.Country,
		ZipCode:  address.ZipCode,
		UserID:   address.UserID,
	})
}

func ValidateNewAddressData(address *models.AddressRequestParams) error {
	if address.Street == "" {
		return errors.New("missing street")
	}
	if address.City == "" {
		return errors.New("missing city")
	}
	if address.State == "" {
		return errors.New("missing state")
	}
	if address.Country == "" {
		return errors.New("missing country")
	}
	if address.ZipCode == "" {
		return errors.New("missing zip code")
	}
	if address.UserID == "" {
		return errors.New("invalid user ID")
	}
	return nil
}