package daogen

import (
  "github.com/jinzhu/gorm"
  // POSSIBLE IMPORT HERE
)

// __DAOName__ is a data access object to a database containing __ReferenceModel__s
type __DAOName__ struct {
  db *gorm.DB
}

// New__DAOName__ creates a new Data Access Object for the
// __ReferenceModel__ model.
func New__DAOName__ (db *gorm.DB) *__DAOName__ {
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