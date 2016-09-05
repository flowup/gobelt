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

  testCases []ReferenceModel
  db *gorm.DB
  dao *DAOName
}

func (s *TemplateTestSuite) SetupSuite() {
  db, err := gorm.Open("sqlite3", "testing.db")
  if err != nil {
    panic(err)
  }
  db.AutoMigrate(&ReferenceModel{})
  db.AutoMigrate(&AuxModel{})
  s.testCases = []ReferenceModel{
    {FieldPrimitive: 1},
    {FieldPrimitive: 2},
    {FieldPrimitive: 2},
    {FieldPrimitive: 3},
    {FieldPrimitive: 4},
    {FieldPrimitive: 5},
  }

  s.dao = NewDAOName(db)
  s.db = db
  for i := range s.testCases {
    s.db.Create(&s.testCases[i])
  }
}

func (s *TemplateTestSuite) TestCreate() {
  model := &ReferenceModel{
    FieldPrimitive:PrimitiveType(42),
  }
  s.dao.Create(model)
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), model)
  get := &ReferenceModel{}
  s.db.First(&get, model.ID)
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), get)
  assert.Equal(s.T(), PrimitiveType(42), get.FieldPrimitive)
}

func (s *TemplateTestSuite) TestRead() {
  model := &ReferenceModel{
    FieldPrimitive:PrimitiveType(2),
  }
  models := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel)(nil), models)
  assert.Equal(s.T(), 2 , len(models))
}

func (s *TemplateTestSuite) TestReadByID() {
  model := &ReferenceModel{
    FieldPrimitive:PrimitiveType(84),
  }
  s.db.Create(model)
  get := s.dao.ReadByID((model.ID))
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), get)
  assert.Equal(s.T(), PrimitiveType(84), get.FieldPrimitive)
}

func (s *TemplateTestSuite) TestUpdate() {
  model := &ReferenceModel{
    FieldPrimitive:PrimitiveType(1),
  }
  list := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel)(nil), list)
  assert.Equal(s.T(), 1, len(list))
  newVal := &ReferenceModel{
    FieldPrimitive:PrimitiveType(40),
  }
  update := s.dao.Update(newVal, (list[0].ID))
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), update)
  assert.Equal(s.T(), PrimitiveType(40), update.FieldPrimitive)
}

func (s *TemplateTestSuite) TestDelete() {
  model := &ReferenceModel{
    FieldPrimitive:PrimitiveType(3),
  }
  list := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel)(nil), list)
  assert.Equal(s.T(), 1, len(list))
  s.dao.Delete(&list[0])
  list = s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel)(nil), list)
  assert.Equal(s.T(), 0, len(list))
}

func (s *TemplateTestSuite) TestStruct() {
	model := &ReferenceModel{
		FieldPrimitive: PrimitiveType(5),
	}
	list := s.dao.Read(model)
	assert.NotNil(s.T(), list)
	assert.Equal(s.T(), 1, len(list))
	model = &list[0]

	assoc := AuxModel{
		AuxModelField: PrimitiveType(20),
	}
	model = s.dao.SetFieldStruct(model, assoc)
	assoc = model.FieldStruct
	assert.NotNil(s.T(), model)
	assert.Equal(s.T(), PrimitiveType(20), assoc.AuxModelField)
	assert.Equal(s.T(), model.ID, assoc.ReferenceModelID)
	assert.Equal(s.T(), model.FieldStruct, assoc)
}

func (s *TemplateTestSuite) TestAssociations() {
  model := &ReferenceModel{
    FieldPrimitive: PrimitiveType(4),
  }
  list := s.dao.Read(model)
  assert.NotEqual(s.T(), ([]ReferenceModel)(nil), list)
  assert.Equal(s.T(), 1, len(list))
  model = &list[0]
  assoc1 := &AuxModel{
    AuxModelField:PrimitiveType(10),
  }
  assoc2 := &AuxModel{
    AuxModelField:PrimitiveType(10),
  }
  model = s.dao.AddFieldSliceAssociation(model, assoc1)
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), model)
  model = s.dao.AddFieldSliceAssociation(model, assoc2)
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), model)
  assert.Equal(s.T(), 2, len(model.FieldSlice))
  model = s.dao.RemoveFieldSliceAssociation(model, assoc2)
  assert.NotEqual(s.T(), (*ReferenceModel)(nil), model)
  assert.Equal(s.T(), 1, len(model.FieldSlice))
}

func (s *TemplateTestSuite) TearDownSuite() {
  for i := range s.testCases {
    s.db.Unscoped().Delete(&s.testCases[i])
  }
}

func TestTemplateTestSuite(t *testing.T) {
  suite.Run(t, &TemplateTestSuite{})
}
