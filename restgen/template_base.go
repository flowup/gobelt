package restgen

import (
	"github.com/kataras/iris"
)


type ReferenceModelController struct {
	config *ReferenceModelControllerConfig
}

type ReferenceModelControllerConfig struct {
	DAO DAOInterface
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
		if a[i] == m {
			return true
		}
	}

	return false
}

func loadintQueryParam(ctx *iris.Context, fieldName string) int {
	val, err := ctx.URLParamInt(fieldName)
	if err != nil {
		return 0
	}

	return val
}

func loadint64QueryParam(ctx *iris.Context, fieldName string) int64 {
	val, err := ctx.URLParamInt64(fieldName)
	if err != nil {
		return 0
	}

	return val
}

func loadstringQueryParam(ctx *iris.Context, fieldName string) string {
	return ctx.URLParam(fieldName)
}


func NewReferenceModelController(c *ReferenceModelControllerConfig) *ReferenceModelController {
	return &ReferenceModelController{
		config: c,
	}
}

func (c *ReferenceModelController) Create(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	output := []ReferenceModel{}
	if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}

	for _, item := range *input {
		c.config.DAO.Create(&item)
		output = append(output, item)
	}

	ctx.JSON(iris.StatusOK, output)
	return
}

func (c *ReferenceModelController) Read(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	output := []ReferenceModel{}
	var auxArray []ReferenceModel
	inputJSON := false
	if err := ctx.ReadJSON(input); err == nil {
		inputJSON = true
	}

	if inputJSON {
		for _, item := range *input {
			auxArray = c.config.DAO.Read(&item)
			for j := range auxArray {
				if !contains(output, auxArray[j]) {
					output = append(output, auxArray[j])
				}
			}
		}
	} else {
		m := &ReferenceModel{
			//GENERATE_FIELDS
		}
		output = c.config.DAO.Read(m)
	}

	ctx.JSON(iris.StatusOK, output)
	return
}

func (c *ReferenceModelController) Update(ctx *iris.Context) {
	input := &[]ReferenceModel{}
	output := [](*ReferenceModel){}
	var auxModel *ReferenceModel
	if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(iris.StatusBadRequest, response)
		return
	}

	for _, item := range *input {
		auxModel = c.config.DAO.Update(&item, item.ID)
		output = append(output, auxModel)
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
		c.config.DAO.Delete(&item)
	}

	response := []response{{Message: "Items deleted"}}
	ctx.JSON(iris.StatusOK, response)
	return
}
