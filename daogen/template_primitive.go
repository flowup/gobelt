package daogen

import (
	"github.com/jinzhu/gorm"
	"time"
)

/* END OF HEADER */

// ReadByFieldPrimitive will find all records
// matching the value given by parameter
func (dao *DAOName) ReadByFieldPrimitive(m PrimitiveType) ([]ReferenceModel, error) {
	retVal := []ReferenceModel{}
	if err := dao.db.Where(&ReferenceModel{FieldPrimitive: m}).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

// ReadByFieldPrimitiveT will return a transaction that
// can be used to find all models matching the value given by parameter
func (dao *DAOName) ReadByFieldPrimitiveT(m PrimitiveType) (*gorm.DB, error) {
	retVal := dao.db.Where(&ReferenceModel{FieldPrimitive: m})

	return retVal, retVal.Error
}

// DeleteByFieldPrimitive deletes all records in database with
// FieldPrimitive the same as parameter given
func (dao *DAOName) DeleteByFieldPrimitive(m PrimitiveType) (error) {
	if err := dao.db.Where(&ReferenceModel{FieldPrimitive: m}).Delete(&ReferenceModel{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByFieldPrimitive will edit all records in database
// with the same FieldPrimitive as parameter given
// using model given by parameter
func (dao *DAOName) EditByFieldPrimitive(m PrimitiveType, newVals *ReferenceModel) (error) {
	if err := dao.db.Table("reference_models").Where(&ReferenceModel{FieldPrimitive: m}).Updates(newVals).Error; err != nil {
		return err
	}
	return nil
}

// SetFieldPrimitive will set FieldPrimitive
// to a value given by parameter
func (dao *DAOName) SetFieldPrimitive(m *ReferenceModel, newVal PrimitiveType) (*ReferenceModel, error) {
	m.FieldPrimitive = newVal
	record, err := dao.ReadByID((m.ID))
	if err != nil {
		return nil, err
	}

	if err := dao.db.Model(&record).Updates(m).Error; err != nil {
		return nil, err
	}

	return record, nil
}

// ReadByFieldPrimitive is a mock implementation of ReadByFieldPrimitive method
func (mock *DAONameMock) ReadByFieldPrimitive(m PrimitiveType) ([]ReferenceModel, error) {
	ret := make([]ReferenceModel, 0, len(mock.db))
	for _, val := range mock.db {
		if val.FieldPrimitive == m {
			ret = append(ret, val)
		}
	}

	return ret, nil
}

// ReadByFieldPrimitiveT is a mock implementation of ReadByFieldPrimitiveT method
func (mock *DAONameMock) ReadByFieldPrimitiveT(m PrimitiveType) (*gorm.DB, error) {
	return nil, nil
}

// DeleteByFieldPrimitive is a mock implementation of DeleteByFieldPrimitive method
func (mock *DAONameMock) DeleteByFieldPrimitive(m PrimitiveType) (error) {
	for _, val := range mock.db {
		if val.FieldPrimitive == m {
			delete(mock.db, val.ID)
		}
	}

	return nil
}

// EditByFieldPrimitive is a mock implementation of EditByFieldPrimitive method
func (mock *DAONameMock) EditByFieldPrimitive(m PrimitiveType, newVals *ReferenceModel) (error) {
	for _, val := range mock.db {
		if val.FieldPrimitive == m {
			id := val.ID
			val = *newVals
			val.ID = id
			val.UpdatedAt = time.Now()
		}
	}

	return nil
}

// SetFieldPrimitive is a mock implementation of SetFieldPrimitive method
func (mock *DAONameMock) SetFieldPrimitive(m *ReferenceModel, newVal PrimitiveType) (*ReferenceModel, error) {
	edit := mock.db[m.ID]
	edit.FieldPrimitive = newVal
	edit.UpdatedAt = time.Now()

	mock.db[m.ID] = edit
	return &edit, nil
}
