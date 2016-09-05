package templateManager

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"github.com/stretchr/testify/assert"
)

type ManagerSuite struct {
	suite.Suite

	testFile string
	manager *Manager
}

func (s *ManagerSuite)SetupSuite() {
	s.manager = GetInstance()
	s.testFile = "template_manager/test_fixtures/template_test"
}

func (s *ManagerSuite) TestCachingTemplates(){
	readData := s.manager.LoadTemplate(s.testFile)
	assert.Equal(s.T(), "THIS IS A TEST TEMPLATE\n", string(readData))
}

func TestManagerSuite(t *testing.T) {
	suite.Run(t, &ManagerSuite{})
}
