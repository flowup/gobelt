package daogen

import (
  "github.com/stretchr/testify/suite"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "testing"
  "github.com/stretchr/testify/assert"
)

type TemplateTestSuite struct {
  suite.Suite

  testCases []ReferenceModel__
  db *gorm.DB
  dao *__DAOName__
}

func (s *TemplateTestSuite) SetupSuite() {
  db, err := gorm.Open("sqlite3", "testing.db")
  if err != nil {
    panic(err)
  }
  db.AutoMigrate(&ReferenceModel__{})
  db.AutoMigrate(&__AuxModel__{})
  s.testCases = []ReferenceModel__{
    {FieldPrimitive__: 1},
    {FieldPrimitive__: 2},
    {FieldPrimitive__: 2},
    {FieldPrimitive__: 3},
    {FieldPrimitive__: 4},
    {FieldPrimitive__: 5},
  }

  s.dao = New__DAOName__(db)
  s.db = db
  for i := range s.testCases {
    s.db.Create(&s.testCases[i])
  }
}

func (s *TemplateTestSuite) TestCreate() {
  model := &ReferenceModel__{
    FieldPrimitive__:__PrimitiveType__(42),
  }
  s.dao.Create(model)
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), model)
  get := &ReferenceModel__{}
  s.db.First(&get, model.ID)
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), get)
  assert.Equal(s.T(), __PrimitiveType__(42), get.FieldPrimitive__)
}

func (s *TemplateTestSuite) TestRead() {
  model := &ReferenceModel__{
    FieldPrimitive__:__PrimitiveType__(2),
  }
  models := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel__)(nil), models)
  assert.Equal(s.T(), 2 , len(models))
}

func (s *TemplateTestSuite) TestReadByID() {
  model := &ReferenceModel__{
    FieldPrimitive__:__PrimitiveType__(84),
  }
  s.db.Create(model)
  get := s.dao.ReadByID((model.ID))
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), get)
  assert.Equal(s.T(), __PrimitiveType__(84), get.FieldPrimitive__)
}

func (s *TemplateTestSuite) TestUpdate() {
  model := &ReferenceModel__{
    FieldPrimitive__:__PrimitiveType__(1),
  }
  list := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel__)(nil), list)
  assert.Equal(s.T(), 1, len(list))
  newVal := &ReferenceModel__{
    FieldPrimitive__:__PrimitiveType__(40),
  }
  update := s.dao.Update(newVal, (list[0].ID))
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), update)
  assert.Equal(s.T(), __PrimitiveType__(40), update.FieldPrimitive__)
}

func (s *TemplateTestSuite) TestDelete() {
  model := &ReferenceModel__{
    FieldPrimitive__:__PrimitiveType__(3),
  }
  list := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel__)(nil), list)
  assert.Equal(s.T(), 1, len(list))
  s.dao.Delete(&list[0])
  list = s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel__)(nil), list)
  assert.Equal(s.T(), 0, len(list))
}

func (s *TemplateTestSuite) TestAssociations() {
  model := &ReferenceModel__{
    FieldPrimitive__: __PrimitiveType__(4),
  }
  list := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel__)(nil), list)
  assert.Equal(s.T(), 1, len(list))
  model = &list[0]
  assoc1 := &__AuxModel__{
    AuxModelField__:__PrimitiveType__(10),
  }
  assoc2 := &__AuxModel__{
    AuxModelField__:__PrimitiveType__(10),
  }
  model = s.dao.AddFieldSlice__Association(model, assoc1)
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), model)
  model = s.dao.AddFieldSlice__Association(model, assoc2)
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), model)
  assert.Equal(s.T(), 2, len(model.FieldSlice__))
  model = s.dao.RemoveFieldSlice__Association(model, assoc2)
  assert.NotEqual(s.T(), (*ReferenceModel__)(nil), model)
  assert.Equal(s.T(), 1, len(model.FieldSlice__))
}

func (s *TemplateTestSuite) TearDownSuite() {
  for i := range s.testCases {
    s.db.Unscoped().Delete(&s.testCases[i])
  }
}

func TestTemplateTestSuite(t *testing.T) {
  suite.Run(t, &TemplateTestSuite{})
}