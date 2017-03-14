package daogen

import (
	"time"
)

// If your model does not use uint IDs rewrite this definition
// ReferenceModelIDType is type of ID
type ReferenceModelIDType uint

// PrimitiveType is a placeholder for int type
type PrimitiveType int

// SliceType is a type of slices of auxiliary models
type SliceType []AuxModel

// AuxModel is an auxiliary structure tha is embedded in ReferenceModel
type AuxModel struct {
	ID ReferenceModelIDType `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	ReferenceModelID uint
	AuxModelField    PrimitiveType
}

// ReferenceModel is a model upon which is based template
type ReferenceModel struct {
	ID ReferenceModelIDType `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	FieldStruct    AuxModel
	FieldPrimitive PrimitiveType
	FieldSlice     SliceType
}

type AuxModelEmbedded struct {
	ID ReferenceModelIDType `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type ReferenceModelEmbedded struct {
	AuxModelEmbedded
}

type ReferenceModelStringID struct {
	ID string
}
