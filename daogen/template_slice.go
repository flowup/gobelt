package daogen

/* END OF HEADER */

// AddFieldSliceAssociation will add
// an association to model given by parameter
func (dao *DAOName) AddFieldSliceAssociation(m *ReferenceModel, asocVal *AuxModel) (*ReferenceModel, error) {
	if err := dao.db.Model(&m).Association("FieldSlice").Append(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// RemoveFieldSliceAssociation will remove
// an association from model given by parameter
func (dao *DAOName) RemoveFieldSliceAssociation(m *ReferenceModel, asocVal *AuxModel) (*ReferenceModel, error) {
	if err := dao.db.Model(&m).Association("FieldSlice").Delete(asocVal).Error; err != nil {
		return nil, err
	}

	return m, nil
}

// GetAllAssociatedFieldSlice will get all
// an association from model given by parameter
func (dao *DAOName) GetAllAssociatedFieldSlice(m *ReferenceModel) ([]AuxModel, error) {
	retVal := []AuxModel{}

	if err := dao.db.Model(&m).Related(&retVal).Error; err != nil {
		return nil, err
	}
	return retVal, nil
}
