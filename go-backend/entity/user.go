package entity

import (
	"errors"
	"go-backend/entity/utils"
	"time"
)

type UserType string

const (
	Default UserType = "user"
	Admin   UserType = "administrator"
)

type User struct {
	id          uint64
	userType    UserType
	fullName    string
	displayName string
	email       string
	document    utils.Document
	//address	AddressType
	birthDate time.Time
}

func (u *User) GetName() string {
	return u.fullName
}

func (u *User) GetDisplayName() string {

	if len(u.displayName) <= 1 {
		return u.fullName
	}

	return u.displayName
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetUserType() UserType {
	return u.userType
}

func (u *User) GetDocument() utils.Document {
	return u.document
}

func (u *User) GetID() uint64 {
	return u.id
}

func (u *User) GetBirthDate() time.Time {
	return u.birthDate
}

func (u *User) GetAge() int {
	return int(time.Since(u.birthDate).Hours() / 24 / 365)
	//return int(time.Now().Sub(u.birthDate).Hours()/24/365) //u.birthDate
}

func (u *User) SetBrithDate(d time.Time) error {
	if int(time.Since(u.birthDate).Hours()/24/365) <= 16 {
		return errors.New("user must be older than 16 years")
	}

	u.birthDate = d
	return nil
}

func (u *User) SetName(name string) error {

	if len(name) <= 1 {
		return errors.New("name must be grater than 1 character")
	}

	if len(name) >= 255 {
		return errors.New("name is too long")
	}

	u.fullName = name

	return nil
}

func (u *User) SetDisplayName(name string) error {

	if len(name) <= 1 {
		return errors.New("name must be grater than 1 character")
	}

	if len(name) >= 255 {
		return errors.New("name is too long")
	}

	u.displayName = name

	return nil
}
