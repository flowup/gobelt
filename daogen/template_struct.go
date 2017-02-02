package daogen

// SetFieldStruct will set a FieldStruct property of a model
// to value given by parameter
func (dao *DAOName) SetFieldStruct(m *ReferenceModel, str AuxModel) (*ReferenceModel, error) {
	m.FieldStruct = str
	var err error
	m, err = dao.Update(m, m.ID)
	if err != nil {
		return nil, err
	}

	return m, nil
}
