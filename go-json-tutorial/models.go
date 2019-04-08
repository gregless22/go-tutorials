package main

// Users struct contains a slice of users
type Users struct {
	Users []User
}

// User struct
type User struct {
	Name   string `json:"name"`
	Type   string
	Age    int
	Social Social
}

// Social struct
type Social struct {
	Facebook string
	Twitter  string
}
