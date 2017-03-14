package daogen

import (
	"time"
)

// ReferenceModelIDType is placeholder type of ID
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

// AuxModelEmbedded is an auxiliary structure tha is embedded in ReferenceModel
type AuxModelEmbedded struct {
	ID ReferenceModelIDType `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// ReferenceModelEmbedded is a model upon which is based template
// with embedded ID declaration
type ReferenceModelEmbedded struct {
	AuxModelEmbedded
}

// ReferenceModelStringID is a model upon which is based template
// with string id
type ReferenceModelStringID struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
