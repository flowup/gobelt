package daogen

// Add__FieldSlice__Association will add
// an association to model given by parameter
func (dao *__DAOName__) Add__FieldSlice__Association (m *__ReferenceModel__, asocVal *__AuxModel__) *__ReferenceModel__ {
  dao.db.Model(&m).Association("__FieldSlice__").Append(asocVal)

  return m
}

// Remove__FieldSlice__Association will remove
// an association from model given by parameter
func (dao *__DAOName__) Remove__FieldSlice__Association (m *__ReferenceModel__, asocVal *__AuxModel__) *__ReferenceModel__ {
  dao.db.Model(&m).Association("__FieldSlice__").Delete(asocVal)

  return m
}
