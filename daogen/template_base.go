package daogen

import (
  "github.com/jinzhu/gorm"
  // POSSIBLE IMPORT HERE
	"time"
)

// DAOName is a data access object to a database containing ReferenceModels
type DAOName struct {
  db *gorm.DB
}

// NewDAOName creates a new Data Access Object for the
// ReferenceModel model.
func NewDAOName (db *gorm.DB) *DAOName {
  return &DAOName{
    db:db,
  }
}

// Create will create single ReferenceModel in database.
func (dao *DAOName) Create(m *ReferenceModel) {
  dao.db.Create(m)
}

// Read will find all DB records matching
// values in a model given by parameter
func (dao *DAOName) Read(m *ReferenceModel) []ReferenceModel {
  retVal := []ReferenceModel{}
  dao.db.Where(m).Find(&retVal)

  return retVal
}

// ReadByID will find ReferenceModel by ID given by parameter
func (dao *DAOName) ReadByID(id uint) *ReferenceModel{
  m := &ReferenceModel{}
  if dao.db.First(&m, id).RecordNotFound() {
    return nil
  }

  return m
}

// Update will update a record of ReferenceModel in DB
func (dao *DAOName) Update(m *ReferenceModel, id uint) *ReferenceModel{
  oldVal := dao.ReadByID(id)
  if oldVal == nil {
    return nil
  }

  dao.db.Model(&oldVal).Updates(m)
  return oldVal
}

// Delete will soft-delete a single ReferenceModel
func (dao *DAOName) Delete(m *ReferenceModel) {
  dao.db.Delete(m)
}

// GetUpdatedAfter will return all ReferenceModels that were
// updated after given timestamp
func (dao *DAOName) GetUpdatedAfter(timestamp time.Time) []ReferenceModel {
	m := []ReferenceModel{}
	dao.db.Where("updated_at > ?", timestamp).Find(&m)
	return m
}

// GetAll will return all records of ReferenceModel in database
func (dao *DAOName) GetAll() []ReferenceModel {
	m := []ReferenceModel{}
	dao.db.Find(&m)

	return m
}
