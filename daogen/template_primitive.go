package daogen


// ReadByFieldPrimitive will find all records
// matching the value given by parameter
func (dao *DAOName) ReadByFieldPrimitive (m PrimitiveType) []ReferenceModel {
  retVal := []ReferenceModel{}
  dao.db.Where(&ReferenceModel{ FieldPrimitive : m }).Find(&retVal)

  return retVal
}

// DeleteByFieldPrimitive deletes all records in database with
// FieldPrimitive the same as parameter given
func (dao *DAOName) DeleteByFieldPrimitive (m PrimitiveType) {
  dao.db.Where(&ReferenceModel{ FieldPrimitive : m }).Delete(&ReferenceModel{})
}

// EditByFieldPrimitive will edit all records in database
// with the same FieldPrimitive as parameter given
// using model given by parameter
func (dao *DAOName) EditByFieldPrimitive (m PrimitiveType, newVals *ReferenceModel) {
  dao.db.Table("reference_models").Where(&ReferenceModel{ FieldPrimitive : m }).Updates(newVals)
}

// SetFieldPrimitive will set FieldPrimitive
// to a value given by parameter
func (dao *DAOName) SetFieldPrimitive (m *ReferenceModel, newVal PrimitiveType) *ReferenceModel {
  m.FieldPrimitive = newVal
  record := dao.ReadByID((m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}