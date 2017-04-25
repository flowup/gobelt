package daogen

/* END OF HEADER */

// ReadByID will find ReferenceModel by ID given by parameter
func (dao *DAOName) ReadByID(id ReferenceModelIDType) (*ReferenceModel, error) {
	m := &ReferenceModel{}
	if err := dao.db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return m, nil
}
