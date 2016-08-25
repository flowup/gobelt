package daogen

import "github.com/jinzhu/gorm"

type PrimitiveType int
type SliceType []AuxModel

type AuxModel struct {
  gorm.Model

  ReferenceModelID uint
  AuxModelField PrimitiveType
}

type ReferenceModel struct {
  gorm.Model

  FieldPrimitive PrimitiveType
  FieldSlice SliceType
}
