package users

import (
	"context"
	"time"

	db "encore.app/users/db"
	models "encore.app/users/models"
	rlog "encore.dev/rlog"
	"encore.dev/storage/cache"
)

// ------------------------------------------------------
// Setup Database

// UsersTable instance.
var AddressesTable = &db.AddressesTable{DB: PlamatioDB}

// ------------------------------------------------------
// Setup Caching

// AddressesCluster is the cache cluster for address data.
var AddressesCluster = cache.NewCluster("addresses-cache-cluster", cache.ClusterConfig{
		// Use LRU policy to evict keys when the cache reaches memory limit.
		EvictionPolicy: cache.AllKeysLRU,
})

// Address Cache Keyspace to store address data by ID.
var AddressCacheKeyspace = cache.NewStructKeyspace[int, models.Address](AddressesCluster, cache.KeyspaceConfig{
	KeyPattern:    "address-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// User Addresses Cache Keyspace to store address data by user ID.
var UserAddressesCacheKeyspace = cache.NewStructKeyspace[string, models.Addresses](AddressesCluster, cache.KeyspaceConfig{
	KeyPattern:    "user-addresses-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /users/addresses/get/:id
// Retrieves the address from the database with the given ID.
//encore:api auth method=GET path=/users/addresses/get/:id
func GetAddress(ctx context.Context, id int) (*models.Address, error) {
	// First, try retrieving the address from cache if it exists.
	a, err := AddressCacheKeyspace.Get(ctx, id)
	// if address is found (i.e., no error), return it
	if err == nil {
		return &a, nil
	}
	// If the address is not found in cache, retrieve it from the database.
	r, err := AddressesTable.GetAddress(ctx, id)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the address.
	go func() {
		// Cache the address.
		if err := AddressCacheKeyspace.Set(ctx, id, *r); err != nil {
			// Log the error
			rlog.Error("error caching address", err)
		}
	}()
	// Return the address.
	return r, err
}

// GET: /users/addresses/get/user/:user_id
// Retrieves all addresses for a user from the database with the given user ID.
//encore:api auth method=GET path=/users/addresses/user/:user_id
func GetUserAddresses(ctx context.Context, user_id string) (*models.Addresses, error) {
	// First, try retrieving the user addresses from cache if it exists.
	a, err := UserAddressesCacheKeyspace.Get(ctx, user_id)
	// if user addresses are found (i.e., no error), return them
	if err == nil {
		return &a, nil
	}
	// If the user addresses are not found in cache, retrieve them from the database.
	r, err := AddressesTable.GetUserAddresses(ctx, user_id)
	if err != nil {
		return nil, err
	}
	// Fire a go routine to cache the user addresses.
	go func() {
		// Cache the user addresses.
		if err := UserAddressesCacheKeyspace.Set(ctx, user_id, *r); err != nil {
			// Log the error
			rlog.Error("error caching user addresses", err)
		}
	}()
	// Return the user addresses.
	return r, err
}

// POST: /users/addresses/add
// Inserts an address into the database.
//encore:api auth method=POST path=/users/addresses/add
func AddAddress(ctx context.Context, newAddress *models.AddressRequestParams) (*models.Address, error) {
	// Insert the address into the database.
	r, err := AddressesTable.InsertAddress(ctx, newAddress)
	if err != nil {
		return nil, err
	}
	// Return the address.
	return r, nil
}

// PUT: /users/addresses/update
// Updates an address in the database.
//encore:api auth method=PUT path=/users/addresses/update
func UpdateAddress(ctx context.Context, updatedAddress *models.Address) (*models.AddressChangeRequestReturn, error) {
	// Update the address in the database.
	err := AddressesTable.UpdateAddress(ctx, updatedAddress)
	if err != nil {
		return nil, err
	}
	
	// Fire a go routine to invalidate the cache for the address.
	go func() {
		// Invalidate the cache for the address.
		_, err = AddressCacheKeyspace.Delete(ctx, updatedAddress.ID)
		if err != nil {
			// Log the error
			rlog.Error("error deleting address cache", err)
		}

		// Invalidate the cache for the user addresses.
		_, err = UserAddressesCacheKeyspace.Delete(ctx, updatedAddress.UserID)
		if err != nil {
			// Log the error
			rlog.Error("error deleting user addresses cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.

	// Return request status.
	return &models.AddressChangeRequestReturn{AddressId: updatedAddress.ID}, nil
}

// DELETE: /users/addresses/delete/:address_id/user/:user_id
// Deletes an address from the database.
//encore:api auth method=DELETE path=/users/addresses/delete/:address_id/user/:user_id
func DeleteAddress(ctx context.Context, address_id int, user_id string) (*models.AddressChangeRequestReturn, error) {

	// Delete the address from the database.
	err := AddressesTable.DeleteAddress(ctx, address_id)
	if err != nil {
		return nil, err
	}

	// Fire a go routine to invalidate the cache for the address.
	go func() {
		// Invalidate the cache for the address.
		_, err = AddressCacheKeyspace.Delete(ctx, address_id)
		if err != nil {
			// Log the error
			rlog.Error("error deleting address cache", err)
		}
		// Invalidate the cache for the user addresses.
		_, err = UserAddressesCacheKeyspace.Delete(ctx, user_id)
		if err != nil {
			// Log the error
			rlog.Error("error deleting user addresses cache", err)
		}
	}()

	// TODO: Publish a message to a message broker to notify other services of the change.

	// Return request status.
	return &models.AddressChangeRequestReturn{AddressId: address_id}, nil
}