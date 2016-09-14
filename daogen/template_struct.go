package daogen

// SetFieldStruct will set a FieldStruct property of a model
// to value given by parameter
func (dao *DAOName) SetFieldStruct(m *ReferenceModel, str AuxModel) *ReferenceModel {
	m.FieldStruct = str
	m = dao.Update(m, m.ID)

	return m
}


