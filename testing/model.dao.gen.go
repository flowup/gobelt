package testing

import (
  "github.com/jinzhu/gorm"
)

/*
@Init
*/

// UserTestDAO is a data access object to a database containing UserTests
type UserTestDAO struct {
  db *gorm.DB
}

// NewUserTestDAO creates a new Data Access Object for the
// UserTest model.
func NewUserTestDAO(db *gorm.DB) *UserTestDAO {
  return &UserTestDAO{
    db: db,
  }
}

/*
@CRUD
*/

// Create will create single UserTest in database.
func (dao *UserTestDAO) Create(m *UserTest) {
  dao.db.Create(m)
}

// Read will find all DB records matching
// values in a model given by parameter
func (dao *UserTestDAO) Read(m *UserTest) []UserTest {
  retVal := []UserTest{}
  dao.db.Where(m).Find(&retVal)

  return retVal
}

// ReadByID will find UserTest by ID given by parameter
func (dao *UserTestDAO) ReadByID(id uint64) *UserTest {
  m := &UserTest{}
  if dao.db.First(&m, id).RecordNotFound() {
    return nil
  }

  return m
}

// Update will update a record of UserTest in DB
func (dao *UserTestDAO) Update(m *UserTest, id uint64) *UserTest {
  oldVal := dao.ReadByID(id)
  if oldVal == nil {
    return nil
  }

  dao.db.Model(&oldVal).Updates(m)
  return oldVal
}

// Delete will soft-delete a single UserTest
func (dao *UserTestDAO) Delete(m *UserTest) {
  dao.db.Delete(m)
}

/*
@Name
*/

// ReadByName will find all records
// matching the value given by parameter
func (dao *UserTestDAO) ReadByName(m string) []UserTest {
  retVal := []UserTest{}
  dao.db.Where(&UserTest{Name: m}).Find(&retVal)

  return retVal
}

// DeleteByName deletes all records in database with
// Name the same as parameter given
func (dao *UserTestDAO) DeleteByName(m string) {
  dao.db.Where(&UserTest{Name: m}).Delete(&UserTest{})
}

// EditByName will edit all records in database
// with the same Name as parameter given
// using model given by parameter
func (dao *UserTestDAO) EditByName(m string, newVals *UserTest) {
  dao.db.Table("user_tests").Where(&UserTest{Name: m}).Updates(newVals)
}

// SetName will set Name
// to a value given by parameter
func (dao *UserTestDAO) SetName(m *UserTest, newVal string) *UserTest {
  m.Name = newVal
  record := dao.ReadByID(uint64(m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}

/*
@Number
*/

// ReadByNumber will find all records
// matching the value given by parameter
func (dao *UserTestDAO) ReadByNumber(m int) []UserTest {
  retVal := []UserTest{}
  dao.db.Where(&UserTest{Number: m}).Find(&retVal)

  return retVal
}

// DeleteByNumber deletes all records in database with
// Number the same as parameter given
func (dao *UserTestDAO) DeleteByNumber(m int) {
  dao.db.Where(&UserTest{Number: m}).Delete(&UserTest{})
}

// EditByNumber will edit all records in database
// with the same Number as parameter given
// using model given by parameter
func (dao *UserTestDAO) EditByNumber(m int, newVals *UserTest) {
  dao.db.Table("user_tests").Where(&UserTest{Number: m}).Updates(newVals)
}

// SetNumber will set Number
// to a value given by parameter
func (dao *UserTestDAO) SetNumber(m *UserTest, newVal int) *UserTest {
  m.Number = newVal
  record := dao.ReadByID(uint64(m.ID))

  dao.db.Model(&record).Updates(m)

  return record
}
