package entity

import "time"

type UserType string

const (
	Default UserType = "user"
	Admin            = "administrator"
)

type User struct {
	id          uint64
	userType    UserType
	fullName    string
	displayName string
	email       string
	document    Document
	//address	AddressType
	birthDate time.Time
}
