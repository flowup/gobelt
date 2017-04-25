package daogen

/* END OF HEADER */

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

func (mock *DAONameMock) SetFieldStruct(m *ReferenceModel, str AuxModel) (*ReferenceModel, error) {
	edit := mock.db[m.ID]
	edit.FieldStruct = str

	mock.db[m.ID] = edit
	return &edit, nil
}
