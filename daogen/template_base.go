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
func NewDAOName(db *gorm.DB) *DAOName {
	return &DAOName{
		db: db,
	}
}

// Create will create single ReferenceModel in database.
func (dao *DAOName) Create(m *ReferenceModel) error {
	if err := dao.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

// Read will find all DB records matching
// values in a model given by parameter
func (dao *DAOName) Read(m *ReferenceModel) ([]ReferenceModel, error) {
	retVal := []ReferenceModel{}
	if err := dao.db.Where(m).Find(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}

func (dao *DAOName) ReadT(m *ReferenceModel) (*gorm.DB, error) {
	retVal := dao.db.Where(m)
	return retVal, retVal.Error
}

// ReadByID will find ReferenceModel by ID given by parameter
func (dao *DAOName) ReadByID(id uint) (*ReferenceModel, error) {
	m := &ReferenceModel{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (dao *DAOName) ReadByIDT(id uint) (*gorm.DB, error) {
	m := &ReferenceModel{}
	retVal := dao.db.First(&m, id)

	return retVal, retVal.Error
}

// Update will update a record of ReferenceModel in DB
func (dao *DAOName) Update(m *ReferenceModel, id uint) (*ReferenceModel, error) {
	oldVal, err := dao.ReadByID(id)
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&oldVal).Updates(m).Error; err != nil {
		return nil, err
	}
	return oldVal, nil
}

// UpdateAllFields will update ALL fields of ReferenceModel in db
// with values given in the ReferenceModel by parameter
func (dao *DAOName) UpdateAllFields(m *ReferenceModel) (*ReferenceModel, error) {
	if err := dao.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Delete will soft-delete a single ReferenceModel
func (dao *DAOName) Delete(m *ReferenceModel) error {
	if err := dao.db.Delete(m).Error; err != nil {
		return err
	}
	return nil
}

// GetUpdatedAfter will return all ReferenceModels that were
// updated after given timestamp
func (dao *DAOName) GetUpdatedAfter(timestamp time.Time) ([]ReferenceModel, error) {
	m := []ReferenceModel{}
	if err := dao.db.Where("updated_at > ?", timestamp).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

// GetAll will return all records of ReferenceModel in database
func (dao *DAOName) GetAll() ([]ReferenceModel, error) {
	m := []ReferenceModel{}
	if err := dao.db.Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (dao *DAOName) ExecuteCustomQueryT(query string) (*gorm.DB, error) {
	retVal := dao.db.Where(query)

	return retVal, retVal.Error
}
