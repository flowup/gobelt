package daogen

import "github.com/jinzhu/gorm"


type DAONameEmbedded struct {
	db *gorm.DB
}

/* END OF HEADER */


// ReadByID will find ReferenceModel by ID given by parameter
func (dao *DAONameEmbedded) ReadByID(id ReferenceModelIDType) (*ReferenceModelEmbedded, error) {
	m := &ReferenceModelEmbedded{AuxModelEmbedded: AuxModelEmbedded{ID: id}}
	if err := dao.db.First(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
