package restgen

import (
	//"github.com/kataras/iris"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// ReferenceModelController is a controller for ReferenceModels
type ReferenceModelController struct {
	config *ReferenceModelControllerConfig
}

// ReferenceModelControllerConfig is a controller config
// for ReferenceModelConroller
type ReferenceModelControllerConfig struct {
	DAO DAOInterface
}

// DAOInterface defines interface of used DAO.
// This interface is implemented by DAOs generated by gobelt.
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

func convertFieldName(input string) string {
	// these strings will be converted to "Id", "Url" e.t.c.
	// this is done because this is the convention in JSON field naming
	transformedStrings := []string{"ID", "UID", "URL"}

	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return strings.ToLower(input)
	}
	input = strings.ToLower(string(input[0])) + input[1:]
	for _, s := range transformedStrings {
		input = strings.Replace(input, s, strings.Title(strings.ToLower(s)), -1)
	}

	return input
}

func loadintQueryParam(ctx echo.Context, fieldName string) int {
	var val int
	var err error
	strVal := ctx.QueryParam(fieldName)
	if strVal == "" {
		return 0
	} else if val, err = strconv.Atoi(strVal); err != nil {
		return 0
	}

	return val
}

func loadint64QueryParam(ctx echo.Context, fieldName string) int64 {
	var val int64
	var err error
	strVal := ctx.QueryParam(fieldName)
	if strVal == "" {
		return 0
	} else if val, err = strconv.ParseInt(strVal, 10, 64); err != nil {
		return 0
	}

	return val
}

func loadfloat64QueryParam(ctx echo.Context, fieldName string) float64 {
	var val float64
	var err error
	strVal := ctx.QueryParam(fieldName)
	if strVal == "" {
		return 0
	} else if val, err = strconv.ParseFloat(strVal, 64); err != nil {
		return 0
	}

	return val
}

func loadstringQueryParam(ctx echo.Context, fieldName string) string {
	return ctx.QueryParam(fieldName)
}

// NewReferenceModelController is a factory method for ReferenceModelController
func NewReferenceModelController(c *ReferenceModelControllerConfig) *ReferenceModelController {
	return &ReferenceModelController{
		config: c,
	}
}

// Create will read an array of ReferenceModels
// from input JSON and create it in database
func (c *ReferenceModelController) Create(ctx echo.Context) error {
	input := &[]ReferenceModel{}
	output := []ReferenceModel{}
	ctx.Bind(input)
	/*if err := ctx.ReadJSON(input); err != nil {
		response := []response{{Message: "Could not read input JSON"}}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}*/

	for _, item := range *input {
		c.config.DAO.Create(&item)
		output = append(output, item)
	}

	//ctx.JSON(iris.StatusOK, output)
	return ctx.JSON(http.StatusOK, output)
}

// Read will read an array of ReferenceModels
// from input JSON and return all models matching the criteria
func (c *ReferenceModelController) Read(ctx echo.Context) error {
	input := &[]ReferenceModel{}
	output := []ReferenceModel{}
	var auxArray []ReferenceModel
	inputJSON := false
	ctx.Bind(input)
	if len(*input) > 0 {
		inputJSON = true
	}
	/*
		if err := ctx.ReadJSON(input); err == nil {
			inputJSON = true
		}*/

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

	return ctx.JSON(http.StatusOK, output)
}

// Update will read an array of ReferenceModels
// from input JSON and update them all to given values
func (c *ReferenceModelController) Update(ctx echo.Context) error {
	input := &[]ReferenceModel{}
	output := [](*ReferenceModel){}
	var auxModel *ReferenceModel
	ctx.Bind(input)
	/*
		if err := ctx.ReadJSON(input); err != nil {
			response := []response{{Message: "Could not read input JSON"}}
			ctx.JSON(iris.StatusBadRequest, response)
			return
		}*/

	for _, item := range *input {
		auxModel = c.config.DAO.Update(&item, item.ID)
		output = append(output, auxModel)
	}

	//ctx.JSON(iris.StatusOK, output)
	return ctx.JSON(http.StatusOK, output)
}

// Delete will read an array of ReferenceModels
// from input JSON and delete them
func (c *ReferenceModelController) Delete(ctx echo.Context) error {
	input := &[]ReferenceModel{}
	ctx.Bind(input)
	/*
		if err := ctx.ReadJSON(input); err != nil {
			response := []response{{Message: "Could not read input JSON"}}
			ctx.JSON(iris.StatusBadRequest, response)
			return
		}*/

	for _, item := range *input {
		c.config.DAO.Delete(&item)
	}

	response := []response{{Message: "Items deleted"}}
	//ctx.JSON(iris.StatusOK, response)
	return ctx.JSON(http.StatusOK, response)
}
