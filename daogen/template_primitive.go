package daogen

// ReadByFieldPrimitive will find all records
// matching the value given by parameter
func (dao *DAOName) ReadByFieldPrimitive(m PrimitiveType) ([]ReferenceModel, error) {
	retVal := []ReferenceModel{}
	if err := dao.db.Where(&ReferenceModel{FieldPrimitive: m }).Find(&retVal).Error; err != nil {
		return nil, err
	}

	return retVal, nil
}

func (dao *DAOName) ReadByFieldPrimitiveT(m PrimitiveType) (*gorm.DB, error) {
	retVal := dao.db.Where(&ReferenceModel{FieldPrimitive: m})

	return retVal, retVal.Error
}

// DeleteByFieldPrimitive deletes all records in database with
// FieldPrimitive the same as parameter given
func (dao *DAOName) DeleteByFieldPrimitive(m PrimitiveType) error {
	if err := dao.db.Where(&ReferenceModel{FieldPrimitive: m }).Delete(&ReferenceModel{}).Error; err != nil {
		return err
	}
	return nil
}

// EditByFieldPrimitive will edit all records in database
// with the same FieldPrimitive as parameter given
// using model given by parameter
func (dao *DAOName) EditByFieldPrimitive(m PrimitiveType, newVals *ReferenceModel) error {
	if err := dao.db.Table("reference_models").Where(&ReferenceModel{FieldPrimitive: m }).Updates(newVals).Error; err != nil {
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
