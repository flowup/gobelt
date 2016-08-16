package __P__

import (
  "github.com/jinzhu/gorm"
  // POSSIBLE IMPORT HERE
)
// DO NOT GENERATE START
// this is reference model, it will not be included in generated files

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

// DO NOT GENERATE END

// BASE DAO OPERATIONS START

// __DAOName is a data access object to a database containing __ReferenceModel__s
type __DAOName__ struct {
  db *gorm.DB
}

// New__DAOName creates a new Data Access Object for the
// __ReferenceModel__ model.
func New__DAOName (db *gorm.DB) *__DAOName__ {
  return &__DAOName__{
    db:db,
  }
}

// Create will create single __ReferenceModel__ in database.
func (dao *__DAOName__) Create(m *__ReferenceModel__) {
  dao.db.Create(m)
}

// Read will find all DB records matching
// values in a model given by parameter
func (dao *__DAOName__) Read(m *__ReferenceModel__) []__ReferenceModel__ {
  retVal := []__ReferenceModel__{}
  dao.db.Where(m).Find(&retVal)

return retVal
}

// ReadByID will find __ReferenceModel__ by ID given by parameter
func (dao *__DAOName__) ReadByID(id uint64) *__ReferenceModel__{
  m := &__ReferenceModel__{}
  if dao.db.First(&m, id).RecordNotFound() {
    return nil
  }

  return m
}


// Update will update a record of __ReferenceModel__ in DB
func (dao *__DAOName__) Update(m *__ReferenceModel__, id uint64) *__ReferenceModel__{
  oldVal := dao.ReadByID(id)
  if oldVal == nil {
    return nil
  }

  dao.db.Model(&oldVal).Updates(m)
  return oldVal
}

// Delete will soft-delete a single __ReferenceModel__
func (dao *__DAOName__) Delete(m *__ReferenceModel__) {
  dao.db.Delete(m)
}

// BASE DAO OPERATIONS END
// PRIMITIVE FIELD OPERATIONS START

// ReadBy__FieldPrimitive__ will find all records
// matching the value given by parameter
func (dao *__DAOName__) ReadBy__FieldPrimitive__ (m __PrimitiveType__) []__ReferenceModel__ {
  retVal := []__ReferenceModel__{}
  dao.db.Where(&__ReferenceModel__{ __FieldPrimitive__ : m }).Find(&retVal)

  return retVal
}

// DeleteBy__FieldPrimitive__ deletes all records in database with
// __FieldPrimitive__ the same as parameter given
func (dao *__DAOName__) DeleteBy__FieldPrimitive__ (m __PrimitiveType__) {
  dao.db.Where(&__ReferenceModel__{ __FieldPrimitive__ : m }).Delete(&__ReferenceModel__{})
}

// EditBy__FieldPrimitive__ will edit all records in database
// with the same __FieldPrimitive__ as parameter given
// using model given by parameter
func (dao *__DAOName__) EditBy__FieldPrimitive__ (m __PrimitiveType__, newVals *__ReferenceModel__) {
  dao.db.Table("{{.TableName}}").Where(&__ReferenceModel__{ __FieldPrimitive__ : m }).Updates(newVals)
}

// Set__FieldPrimitive__ will set __FieldPrimitive__
// to a value given by parameter
func (dao *__DAOName__) Set__FieldPrimitive__ (m *__ReferenceModel__, newVal __PrimitiveType__) *__ReferenceModel__ {
  m.__FieldPrimitive__ = newVal
  record := dao.ReadByID(uint64(m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}

// PRIMITIVE FIELD OPERATIONS END
// SLICE FIELD OPERATIONS END


func (dao *__DAOName__) Add__FieldSlice__Association (m *__ReferenceModel__, asocVal *__AuxModel__) *__ReferenceModel__ {
  dao.db.Model(&m).Association("__FieldSlice__").Append(asocVal)

  return m
}

func (dao *__DAOName__) Remove__FieldSlice__Association (m *__ReferenceModel__, asocVal *__AuxModel__) *__ReferenceModel__ {
  dao.db.Model(&m).Association("__FieldSlice__").Delete(asocVal)

  return m
}

// SLICE FIELD OPERATIONS END