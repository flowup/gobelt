package __P__

import "github.com/jinzhu/gorm"

type __PrimitiveType__ int
type __SliceType__ []__AuxModel__

type __AuxModel__ struct {
  gorm.Model

  __AuxModelField__ __PrimitiveType__
}

type __ReferenceModel__ struct {
  gorm.Model

  __FieldPrimitive__ __PrimitiveType__
  __FieldSlice__ __SliceType__
}
