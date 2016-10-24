package restgen

import "github.com/jinzhu/gorm"


type ReferenceModel struct {
	gorm.Model

	FieldInt int
	FieldInt64 int64
	FieldString string

}
