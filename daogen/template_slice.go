package daogen

// AddFieldSliceAssociation will add
// an association to model given by parameter
func (dao *DAOName) AddFieldSliceAssociation (m *ReferenceModel, asocVal *AuxModel) *ReferenceModel {
  dao.db.Model(&m).Association("FieldSlice").Append(asocVal)

  return m
}

// RemoveFieldSliceAssociation will remove
// an association from model given by parameter
func (dao *DAOName) RemoveFieldSliceAssociation (m *ReferenceModel, asocVal *AuxModel) *ReferenceModel {
  dao.db.Model(&m).Association("FieldSlice").Delete(asocVal)

  return m
}

func (dao *DAOName) GetAllAssociatedFieldSlice (m *ReferenceModel) []AuxModel {
	retVal := []AuxModel{}

	dao.db.Model(&m).Related(&retVal)
	return retVal
}
