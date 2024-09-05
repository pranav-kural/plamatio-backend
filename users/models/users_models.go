package users

// User represents a user in the system.
type User struct {
	ID        int    `json:"id"`        // unique identifier
	FirstName string `json:"firstName"` // first name of the user
	LastName  string `json:"lastName"`  // last name of the user
	RefID     string `json:"refId"`     // reference ID
}

// Address represents an address associated with a user.
type Address struct {
	ID       int    `json:"id"`       // unique identifier
	Street   string `json:"street"`   // street address
	City     string `json:"city"`     // city
	State    string `json:"state"`    // state
	Country  string `json:"country"`  // country
	ZipCode  string `json:"zipCode"`  // zip code
	UserID   int    `json:"userId"`   // user ID
}

// Users represents a collection of user objects.
type Users struct {
	Data []*User `json:"data"`
}

// Addresses represents a collection of address objects.
type Addresses struct {
	Data []*Address `json:"data"`
}

// UserRequestParams represents the request parameters for creating or updating a new user.
type UserRequestParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RefID     string `json:"refId"`
}

// AddressRequestParams represents the request parameters for creating or updating a new address.
type AddressRequestParams struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	ZipCode  string `json:"zipCode"`
	UserID   int    `json:"userId"`
}

// UserChangeRequestStatus represents the status of a user change request.
type UserChangeRequestStatus struct {
	Status string `json:"status"`
}

// UserRequestSuccess represents a successful user request status.
const UserRequestSuccess = "success"

// UserRequestFailed represents a failed user request status.
const UserRequestFailed = "failed"