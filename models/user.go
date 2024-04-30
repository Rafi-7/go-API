package models

type Address struct {
	State   string
	City    string
	Pincode int
}

type User struct {
	Name    string
	Age     int
	Address Address
}
