package daogen

import (
	"strconv"
	"time"
)

/* END OF HEADER */

// DAONameStringMock is a mock DAO for model with string ID
type DAONameStringMock struct{
	db map[string]ReferenceModelStringID
	lastID int
}

// NewDAONameStringMock is a factory function for DAONameStringMock
func NewDAONameStringMock() *DAONameStringMock{
	return &DAONameStringMock{
		db: make(map[string]ReferenceModelStringID),
	}
}

// Create will put a model into mock in-memory DB
func (mock *DAONameStringMock) Create(m *ReferenceModelStringID) (error) {
	created := false
	m.CreatedAt = time.Now()
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
