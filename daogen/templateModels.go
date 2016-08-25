package daogen

import "github.com/jinzhu/gorm"

type __PrimitiveType__ int
type __SliceType__ []__AuxModel__

type __AuxModel__ struct {
  gorm.Model

  ReferenceModel__ID uint
  AuxModelField__ __PrimitiveType__
}

type ReferenceModel__ struct {
  gorm.Model

  FieldPrimitive__ __PrimitiveType__
  FieldSlice__ __SliceType__
}
