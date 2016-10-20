package restgen

import (
	"github.com/labstack/echo"
	"net/http"
)

type ReferenceModelArray struct {
	array []ReferenceModel
}

type ReferenceModelController struct {
	config *ReferenceModelControllerConfig
}

type ReferenceModelControllerConfig struct {
	dao DAOInterface
}

type DAOInterface interface {
	Create(m *ReferenceModel)
	Read(m *ReferenceModel) []ReferenceModel
	Update(m *ReferenceModel, id uint) *ReferenceModel
	Delete(m *ReferenceModel)
}

func (a ReferenceModelArray) contains(m *ReferenceModel) bool {
	for i := range a.array {
		if a.array[i] == m {
			return true
		}
	}

	return false
}

func NewReferenceModelController(c *ReferenceModelControllerConfig) *ReferenceModelController {
	return &ReferenceModelController{
		config: c,
	}
}

func (c *ReferenceModelController) Create(e echo.Context) error {
	input := ReferenceModelArray{}
	e.Bind(&input)

	for i := range input.array {
		c.config.dao.Create(&input.array[i])
	}

	return e.JSON(http.StatusOK, &input)
}

func (c *ReferenceModelController) Read(e echo.Context) error {
	input := ReferenceModelArray{}
	output := ReferenceModelArray{}
	auxArray := []ReferenceModel{}
	e.Bind(&input)

	for i := range input.array {
		auxArray = c.config.dao.Read(&input.array[i])
		for j := range auxArray {
			if !output.contains(auxArray[j]) {
				output = append(output, auxArray[j])
			}
		}
	}

	return e.JSON(http.StatusOK, &output)
}

func (c *ReferenceModelController) Update(e echo.Context) error {
	input := ReferenceModelArray{}
	output := []ReferenceModel{}
	var auxModel *ReferenceModel
	e.Bind(&input)

	for i := range input.array {
		auxModel = c.config.dao.Update(&input.array[i], input.array[i].ID)
		output = append(output, *auxModel)
	}

	return e.JSON(http.StatusOK, &output)
}

func (c *ReferenceModelController) Delete(e echo.Context) error {
	input := ReferenceModelArray{}
	e.Bind(&input)

	for i := range input.array {
		c.config.dao.Delete(&input.array[i])
	}

	return e.String(http.StatusOK, "ReferenceModels deleted")
}
