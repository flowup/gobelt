package daogen

import "github.com/jinzhu/gorm"

// PrimitiveType is a placeholder for int type
type PrimitiveType int

// SliceType is a type of slices of auxiliary models
type SliceType []AuxModel

// AuxModel is an auxiliary structure tha is embedded in ReferenceModel
type AuxModel struct {
	gorm.Model

	ReferenceModelID uint
	AuxModelField    PrimitiveType
}

// ReferenceModel is a model upon which is based template
type ReferenceModel struct {
	gorm.Model

	FieldStruct    AuxModel
	FieldPrimitive PrimitiveType
	FieldSlice     SliceType
}
