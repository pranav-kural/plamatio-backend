package core

import (
	"context"

	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

// ------------------------------------------------------
// Setup Databases

var PlamatioDB = sqldb.NewDatabase("plamatio_db", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

// ------------------------------------------------------
// Setup Authentication

// secrets struct for API-key authentication.
var secrets struct {
    PlamatioWebFrontendApiKey string    // API key for the Plamatio Web Frontend
}

// AuthHandler - authentication handler to validate API key for authenticated endpoints.
//encore:authhandler
func AuthHandler(ctx context.Context, token string) (auth.UID, error) {
    // Validate the token - confirm it matches the API key.
		if token == secrets.PlamatioWebFrontendApiKey {
			// Return nil if the token is valid.
			return auth.UID("authenticated"), nil
		}
		// Return an error if API key is invalid.
		return "", &errs.Error{
        Code: errs.Unauthenticated,
        Message: "invalid API key",	
    }
}

// ------------------------------------------------------
// Setup API

type StringResponse struct {
	Data string `json:"data"`
}

// GET: /core/info
// Returns information about the core services.
//encore:api public method=GET path=/core/info
func Get(ctx context.Context) (*StringResponse, error) {
		return &StringResponse{
			Data: "This is Plamatio Backend REST API. For more information, check: https://github.com/pranav-kural/plamatio-backend",
		}, nil
}