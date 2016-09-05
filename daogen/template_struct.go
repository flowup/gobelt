package daogen

func (dao *DAOName) SetFieldStruct(m *ReferenceModel, str StructType) *ReferenceModel {
	m.FieldStruct = str
	m = dao.Update(m, m.ID)

	return m
}
