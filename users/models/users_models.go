package users

// User represents a user in the system.
type User struct {
	ID        string `json:"id"`        // unique identifier
	FirstName string `json:"firstName"` // first name of the user
	LastName  string `json:"lastName"`  // last name of the user
	Email 	  string `json:"email"`     // email address of the user
}

// Address represents an address associated with a user.
type Address struct {
	ID       int    `json:"id"`       // unique identifier
	Street   string `json:"street"`   // street address
	City     string `json:"city"`     // city
	State    string `json:"state"`    // state
	Country  string `json:"country"`  // country
	ZipCode  string `json:"zipCode"`  // zip code
	UserID   string    `json:"userId"`   // user ID
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
	Email     string `json:"email"`
}

// AddressRequestParams represents the request parameters for creating or updating a new address.
type AddressRequestParams struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	ZipCode  string `json:"zipCode"`
	UserID   string    `json:"userId"`
}

// DeleteUserParams represents the parameters for deleting a user.
type DeleteAddressParams struct {
	AddressID	int `json:"addressId"`
	UserID		string `json:"userId"`
}

// Return type for mutations to user data.
type UserChangeRequestReturn struct {
	UserId string `json:"id"`
}

// Return type for mutations to address data.
type AddressChangeRequestReturn struct {
	AddressId int `json:"id"`
}