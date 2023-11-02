package models

import "github.com/go-faker/faker/v4"

type User struct {
	Model
	FirstName string `json:"first_name" faker:"first_name"`
	LastName  string `json:"last_name" faker:"last_name"`
}

func (u *User) Fake() error {
	return faker.FakeData(u)
}
