package testing

import "github.com/jinzhu/gorm"

type UserTest struct {
  gorm.Model

  Name string
  Number int
}