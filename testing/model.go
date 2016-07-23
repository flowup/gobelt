package testing

import "github.com/jinzhu/gorm"

// @observable --force
type UserTest struct {
	gorm.Model

	Name   string
	Number int
}
