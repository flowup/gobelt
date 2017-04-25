package daogen

import "time"

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

// AddFieldSliceAssociation is a mock implementation of AddFieldSliceAssociation method
func (mock *DAONameMock) AddFieldSliceAssociation(m *ReferenceModel, asocVal *AuxModel) (*ReferenceModel, error) {
	edit := mock.db[m.ID]
	edit.FieldSlice = append(edit.FieldSlice, *asocVal)
	edit.UpdatedAt = time.Now()
	mock.db[m.ID] = edit

	return &edit, nil
}

// RemoveFieldSliceAssociation is a mock implementation of RemoveFieldSliceAssociation method
func (mock *DAONameMock) RemoveFieldSliceAssociation(m *ReferenceModel, asocVal *AuxModel) (*ReferenceModel, error) {
	a := m.FieldSlice
	m.UpdatedAt = time.Now()
	deletedIndex := 0
	for j, val := range a {
		if val.ID == asocVal.ID {
			deletedIndex = j
		}
	}
	a[deletedIndex] = a[len(a)-1]
	a = a[:len(a)-1]

	return m, nil
}

// GetAllAssociatedFieldSlice is a mock implementation of GetAllAssociatedFieldSlice method
func (mock *DAONameMock) GetAllAssociatedFieldSlice(m *ReferenceModel) ([]AuxModel, error) {
	return m.FieldSlice, nil
}
