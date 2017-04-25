package daogen

import "time"

/* END OF HEADER */

// DAONameMock is a mock DAO
type DAONameMock struct{
	db map[ReferenceModelIDType]ReferenceModel
	lastID ReferenceModelIDType
}

// NewDAONameMock is a factory function for NewDAONameMock
func NewDAONameMock() *DAONameMock{
	return &DAONameMock{
		db: make(map[ReferenceModelIDType]ReferenceModel),
	}
}

// Create will put a model into mock in-memory DB
func (mock *DAONameMock) Create(m *ReferenceModel) (error) {
	created := false
	m.CreatedAt = time.Now()
	for !created{
		mock.lastID++
		if _, exists := mock.db[mock.lastID]; !exists {
			m.ID = mock.lastID
			mock.db[mock.lastID] = *m
			created = true
		}
	}
	return nil
}
