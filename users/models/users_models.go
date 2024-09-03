package users

/*

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

*/

type User struct {
	ID        int    `json:"id"` // unique identifier
	FirstName string `json:"firstName"` // first name of the user
	LastName  string `json:"lastName"` // last name of the user
}

type Address struct {
	ID       int    `json:"id"` // unique identifier
	Street   string `json:"street"` // street address
	City     string `json:"city"` // city
	State    string `json:"state"` // state
	Country  string `json:"country"` // country
	ZipCode  string `json:"zipCode"` // zip code
	UserID   int    `json:"userId"` // user ID
}

type Users struct {
	Data []*User `json:"data"`
}

type Addresses struct {
	Data []*Address `json:"data"`
}

type UserRequestParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AddressRequestParams struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	ZipCode  string `json:"zipCode"`
	UserID   int    `json:"userId"`
}

type UserChangeRequestStatus struct {
	Status string `json:"status"`
}

const (
	UserRequestSuccess = "success" 
	UserRequestFailed  = "failed" 
)