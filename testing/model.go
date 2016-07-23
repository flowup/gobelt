package testing

import "github.com/jinzhu/gorm"

// @observable
type UserTest struct {
	gorm.Model

	Name   string
	Number int
}
