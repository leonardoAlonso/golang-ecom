package types

import "time"

type UserStore interface {
	// The interface type that specifies the methods that a UserStore must implement
	// to be able to interact with the database
	// this will helps to create a mock store for testing
	// and also to create a real store that interacts with the database
	GetUserByEmail(email string) (*User, error)
	CreateUser(u *User) error
	GetUserByID(id int) (*User, error)
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	// A tag for a field allows you to attach meta-information to the field
	// which can be acquired using reflection. Usually it is used to provide
	// transformation info on how a struct field is encoded to or decoded from
	// another format
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
