package restgen

import "github.com/jinzhu/gorm"

type PrimitiveType int

type ReferenceModel struct {
	gorm.Model

	FieldPrimitive PrimitiveType
}
