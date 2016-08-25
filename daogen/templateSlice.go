package daogen

// Add__FieldSlice__Association will add
// an association to model given by parameter
func (dao *__DAOName__) AddFieldSlice__Association (m *ReferenceModel__, asocVal *__AuxModel__) *ReferenceModel__ {
  dao.db.Model(&m).Association("FieldSlice__").Append(asocVal)

  return m
}

// Remove__FieldSlice__Association will remove
// an association from model given by parameter
func (dao *__DAOName__) RemoveFieldSlice__Association (m *ReferenceModel__, asocVal *__AuxModel__) *ReferenceModel__ {
  dao.db.Model(&m).Association("FieldSlice__").Delete(asocVal)

  return m
}
