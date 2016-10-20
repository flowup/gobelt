package restgen

import (
	"github.com/kataras/iris"
)


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

type response struct {
	Message string `json:"error"`
}

func contains(a []ReferenceModel, m ReferenceModel) bool {
	for i := range a {
		if a[i] == *m {
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

func (c *ReferenceModelController) Create(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}

	for _, item := range *input {
		c.config.dao.Create(&item)
	}

	ctx.JSON(iris.StatusOK, input)
	return
}

func (c *ReferenceModelController) Read(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	output := []ReferenceModel{}
	var auxArray []ReferenceModel
	if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}

	for _, item := range *input {
		auxArray = c.config.dao.Read(&item)
		for j := range auxArray {
			if contains(output, auxArray[j]) {
				output = append(output, auxArray[j])
			}
		}
	}

	ctx.JSON(iris.StatusOK, output)
	return
}

func (c *ReferenceModelController) Update(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	output := []ReferenceModel{}
	var auxModel *ReferenceModel
	if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}

	for _, item := range *input {
		auxModel = c.config.dao.Update(&item, item.ID)
		output = append(output, *auxModel)
	}

	ctx.JSON(iris.StatusOK, output)
	return
}

func (c *ReferenceModelController) Delete(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}

	for _, item := range *input {
		c.config.dao.Delete(&item)
	}

	response := []response{{Message: "Items deleted"}}
	ctx.JSON(iris.StatusOK, response)
	return
}
