package daogen


// ReadByFieldPrimitive__ will find all records
// matching the value given by parameter
func (dao *__DAOName__) ReadByFieldPrimitive__ (m __PrimitiveType__) []ReferenceModel__ {
  retVal := []ReferenceModel__{}
  dao.db.Where(&ReferenceModel__{ FieldPrimitive__ : m }).Find(&retVal)

  return retVal
}

// DeleteByFieldPrimitive__ deletes all records in database with
// FieldPrimitive__ the same as parameter given
func (dao *__DAOName__) DeleteByFieldPrimitive__ (m __PrimitiveType__) {
  dao.db.Where(&ReferenceModel__{ FieldPrimitive__ : m }).Delete(&ReferenceModel__{})
}

// EditByFieldPrimitive__ will edit all records in database
// with the same FieldPrimitive__ as parameter given
// using model given by parameter
func (dao *__DAOName__) EditByFieldPrimitive__ (m __PrimitiveType__, newVals *ReferenceModel__) {
  dao.db.Table("__reference_model__s").Where(&ReferenceModel__{ FieldPrimitive__ : m }).Updates(newVals)
}

// SetFieldPrimitive__ will set FieldPrimitive__
// to a value given by parameter
func (dao *__DAOName__) SetFieldPrimitive__ (m *ReferenceModel__, newVal __PrimitiveType__) *ReferenceModel__ {
  m.FieldPrimitive__ = newVal
  record := dao.ReadByID(uint64(m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}