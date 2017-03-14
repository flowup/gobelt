package daogen

import "strconv"

/* END OF HEADER */

type DAONameStringMock struct{
	db map[string]ReferenceModelStringID
	lastID int
}

func NewDAONameStringMock() *DAONameStringMock{
	return &DAONameStringMock{
		db: make(map[string]ReferenceModelStringID),
	}
}

func (mock *DAONameStringMock) Create(m *ReferenceModelStringID) (error) {
	created := false
	for !created{
		mock.lastID++
		id := strconv.Itoa(mock.lastID)
		if _, exists := mock.db[id]; !exists {
			m.ID = id
			mock.db[id] = *m
		}
	}
	return nil
}
