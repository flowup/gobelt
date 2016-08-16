package daogen


// ReadBy__FieldPrimitive__ will find all records
// matching the value given by parameter
func (dao *__DAOName__) ReadBy__FieldPrimitive__ (m __PrimitiveType__) []__ReferenceModel__ {
  retVal := []__ReferenceModel__{}
  dao.db.Where(&__ReferenceModel__{ __FieldPrimitive__ : m }).Find(&retVal)

  return retVal
}

// DeleteBy__FieldPrimitive__ deletes all records in database with
// __FieldPrimitive__ the same as parameter given
func (dao *__DAOName__) DeleteBy__FieldPrimitive__ (m __PrimitiveType__) {
  dao.db.Where(&__ReferenceModel__{ __FieldPrimitive__ : m }).Delete(&__ReferenceModel__{})
}

// EditBy__FieldPrimitive__ will edit all records in database
// with the same __FieldPrimitive__ as parameter given
// using model given by parameter
func (dao *__DAOName__) EditBy__FieldPrimitive__ (m __PrimitiveType__, newVals *__ReferenceModel__) {
  dao.db.Table("__reference_model__s").Where(&__ReferenceModel__{ __FieldPrimitive__ : m }).Updates(newVals)
}

// Set__FieldPrimitive__ will set __FieldPrimitive__
// to a value given by parameter
func (dao *__DAOName__) Set__FieldPrimitive__ (m *__ReferenceModel__, newVal __PrimitiveType__) *__ReferenceModel__ {
  m.__FieldPrimitive__ = newVal
  record := dao.ReadByID(uint64(m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}