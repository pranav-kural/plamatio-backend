package users

import (
	"context"
	"time"

	db "encore.app/users/db"
	models "encore.app/users/models"
	"encore.dev/storage/cache"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Database

// ProductDB instance.
var ProductsDB = sqldb.Named("products")

// UsersTable instance.
var UsersTable = &db.UsersTable{DB: ProductsDB}

// ------------------------------------------------------
// Setup Caching

// UsersCluster is the cache cluster for user data.
var UsersCluster = cache.NewCluster("users-cache-cluster", cache.ClusterConfig{
    // Use LRU policy to evict keys when the cache reaches memory limit.
    EvictionPolicy: cache.AllKeysLRU,
})

// User Cache Keyspace to store user data by ID.
var UserCacheKeyspace = cache.NewStructKeyspace[int, models.User](UsersCluster, cache.KeyspaceConfig{
	KeyPattern:    "user-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// Ref User Cache Keyspace to store user data by reference ID.
var RefUserCacheKeyspace = cache.NewStructKeyspace[string, models.User](UsersCluster, cache.KeyspaceConfig{
	KeyPattern:    "ref-user-cache/:key",
	DefaultExpiry: cache.ExpireIn(2 * time.Hour),
})

// ------------------------------------------------------
// Setup API

// GET: /users/get/:id
// Retrieves the user from the database with the given ID.
//encore:api auth method=GET path=/users/get/:id
func GetUser(ctx context.Context, id int) (*models.User, error) {
	// First, try retrieving the user from cache if it exists.
	u, err := UserCacheKeyspace.Get(ctx, id)
	// if user is found (i.e., no error), return it
	if err == nil {
		return &u, nil
	}
	// If the user is not found in cache, retrieve it from the database.
	r, err := UsersTable.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cache the user.
	if err := UserCacheKeyspace.Set(ctx, id, *r); err != nil {
		return nil, err
	}
	// Return the user.
	return r, err
}

// GET: /users/ref/:id
// Retrieves the user from the database with the given reference ID.
//encore:api auth method=GET path=/users/ref/:id
func GetUserByRefID(ctx context.Context, id string) (*models.User, error) {
	// First, try retrieving the user from cache if it exists.
	u, err := RefUserCacheKeyspace.Get(ctx, id)
	// if user is found (i.e., no error), return it
	if err == nil {
		return &u, nil
	}
	// If the user is not found in cache, retrieve it from the database.
	r, err := UsersTable.GetUserByRefID(ctx, id)
	if err != nil {
		return nil, err
	}
	// Cache the user.
	if err := RefUserCacheKeyspace.Set(ctx, id, *r); err != nil {
		return nil, err
	}
	// Return the user.
	return r, err
}

// GET: /users/all
// Retrieves all users from the database.
//encore:api private method=GET path=/users/all
func GetAllUsers(ctx context.Context) (*models.Users, error) {
	// Retrieve all users from the database.
	r, err := UsersTable.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	// Return the users.
	return r, nil
}

// POST: /users/add
// Inserts a user into the database.
//encore:api auth method=POST path=/users/add
func AddUser(ctx context.Context, newUser *models.UserRequestParams) (*models.User, error) {
	// Insert the user into the database.
	r, err := UsersTable.InsertUser(ctx, newUser)
	if err != nil {
		return nil, err
	}
	// Return the user.
	return r, nil
}

// PUT: /users/update
// Updates a user in the database.
//encore:api auth method=PUT path=/users/update
func UpdateUser(ctx context.Context, updatedUser *models.User) (*models.UserChangeRequestStatus, error) {
	// Update the user in the database.
	err := UsersTable.UpdateUser(ctx, updatedUser)
	if err != nil {
		return &models.UserChangeRequestStatus{Status: models.UserRequestFailed}, err
	}
	// Invalidate the cache for the user.
	_, err = UserCacheKeyspace.Delete(ctx, updatedUser.ID)
	if err != nil {
		return &models.UserChangeRequestStatus{Status: models.UserRequestFailed}, err
	}

	// TODO: Publish a message to a message broker to notify other services of the change.

	// Return the user.
	return &models.UserChangeRequestStatus{Status: models.UserRequestSuccess}, nil
}

// DELETE: /users/delete/:id
// Deletes a user from the database.
//encore:api auth method=DELETE path=/users/delete/:id
func DeleteUser(ctx context.Context, id int) (*models.UserChangeRequestStatus, error) {
	// Delete the user from the database.
	err := UsersTable.DeleteUser(ctx, id)
	if err != nil {
		return &models.UserChangeRequestStatus{Status: models.UserRequestFailed}, err
	}
	// Invalidate the cache for the user.
	_, err = UserCacheKeyspace.Delete(ctx, id)
	if err != nil {
		return &models.UserChangeRequestStatus{Status: models.UserRequestFailed}, err
	}

	// TODO: Publish a message to a message broker to notify other services of the change.

	// Return the user.
	return &models.UserChangeRequestStatus{Status: models.UserRequestSuccess}, nil
}