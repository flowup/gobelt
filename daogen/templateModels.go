package daogen

import "github.com/jinzhu/gorm"

type __PrimitiveType__ int
type __SliceType__ []__AuxModel__

type __AuxModel__ struct {
  gorm.Model

  __AuxModelField__ __PrimitiveType__
}

type __ReferenceModel__ struct {
  gorm.Model

  FieldPrimitive__ __PrimitiveType__
  FieldSlice__ __SliceType__
}
