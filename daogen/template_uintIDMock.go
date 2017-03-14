package daogen

/* END OF HEADER */

type DAONameMock struct{
	db map[ReferenceModelIDType]ReferenceModel
	lastID ReferenceModelIDType
}

func NewDAONameMock() *DAONameMock{
	return &DAONameMock{
		db: make(map[ReferenceModelIDType]ReferenceModel),
	}
}

func (mock *DAONameMock) Create(m *ReferenceModel) (error) {
	created := false
	for !created{
		mock.lastID++
		if _, exists := mock.db[mock.lastID]; !exists {
			m.ID = mock.lastID
			mock.db[mock.lastID] = *m
		}
	}
	return nil
}
