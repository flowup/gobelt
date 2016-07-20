package testing

import (
  "github.com/stretchr/testify/suite"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserTestDAOSuite struct {
  suite.Suite

  testCases []UserTest
  db *gorm.DB
  dao *UserTestDAO
}

func (s *UserTestDAOSuite) SetupSuite() {
  db, err := gorm.Open("sqlite3", "testing.db")
  if err != nil {
    panic(err)
  }
  db.AutoMigrate(&UserTest{})

  s.testCases = []UserTest{
    {Name: "test1"},
    {Name: "test2"},
    {Name: "test3"},
    {Name: "test4"},
    {Name: "test5"},
    {Name: "test6", Number: 1},
    {Name: "test6", Number: 2},
    {Name: "test6", Number: 3},
  }

  s.db = db
  s.dao = NewUserTestDAO(db)
}

func (s *UserTestDAOSuite) TearDownSuite() {
  /*for i := range s.testCases {
    s.db.Unscoped().Delete(&s.testCases[i])
  }*/
}

func (s *UserTestDAOSuite) SetupTest() {
  for i := range s.testCases {
    user := UserTest{}
    //db.Create(&s.testCases[i])
    s.db.FirstOrCreate(&user, s.testCases[i])
  }
}

func (s *UserTestDAOSuite) TearDownTest() {

}

func (s *UserTestDAOSuite) TestReadByID() {
  users := s.dao.ReadByName("test2")
  assert.Equal(s.T(), 1, len(users))

  user := users[0]
  assert.NotEqual(s.T(), (*UserTest)(nil), user)

  readUser := s.dao.ReadByID((uint64)(user.ID))
  assert.NotEqual(s.T(), (*UserTest)(nil), readUser)
  assert.Equal(s.T(), readUser.Name, user.Name)

}

func (s *UserTestDAOSuite) TestUpdate() {
  users := s.dao.ReadByName("test3")
  assert.Equal(s.T(), 1, len(users))

  user := users[0]
  user.Name = "newTest3"
  s.dao.Update(&user, (uint64)(user.ID))

  users = s.dao.ReadByName("newTest3")
  assert.Equal(s.T(), 1, len(users))
}

func (s *UserTestDAOSuite) TestReadByName() {
  users := s.dao.ReadByName("test6")

  assert.Equal(s.T(), 3, len(users))
}

func (s *UserTestDAOSuite) TestRead() {
  users := s.dao.Read(&UserTest{Name:"test6"})

  assert.Equal(s.T(), 3, len(users))
}

func (s *UserTestDAOSuite) TestDelete() {
  s.dao.Delete(&UserTest{Name:"test4"})

  users := s.dao.ReadByName("test4")
  assert.Equal(s.T(), 0, len(users))
}

func (s *UserTestDAOSuite) TestDeleteByName() {
  s.dao.DeleteByName("test6")

  users := s.dao.ReadByName("test6")
  assert.Equal(s.T(), 0, len(users))
}

func (s *UserTestDAOSuite) TestEditByName() {
  s.dao.EditByName("test5", &UserTest{Name:"newTest5"})

  users := s.dao.ReadByName("newTest5")
  assert.Equal(s.T(), 1, len(users))
}

func (s *UserTestDAOSuite) TestSetName() {
  users := s.dao.ReadByName("test1")
  assert.Equal(s.T(), 1, len(users))

  user := users[0]
  newUser := s.dao.SetName(&user, "newTest1")
  assert.Equal(s.T(), "newTest1", newUser.Name)
}

func TestUserTestDAOSuite(t *testing.T) {
  suite.Run(t, &UserTestDAOSuite{})
}
