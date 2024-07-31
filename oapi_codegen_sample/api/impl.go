package api

import (
	"fmt"
	"net/http"
	"oapi_codegen_sample/petstore"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

type ServerInterfaceImpl struct {
	petstore.ServerInterface
}

func (s ServerInterfaceImpl) FindPets(ctx echo.Context, params petstore.FindPetsParams) error {
	tag := "pet"
	var results []petstore.Pet = []petstore.Pet{
		{Id: 1, Name: "dog", Tag: &tag},
	}
	return ctx.JSON(200, results)
}

func (s ServerInterfaceImpl) AddPet(ctx echo.Context) error {
	fmt.Println("AddPet")
	return ctx.String(http.StatusOK, "AddPet")
}

func (s ServerInterfaceImpl) DeletePet(ctx echo.Context, id int64) error {
	fmt.Println("DeletePet")
	return ctx.String(http.StatusOK, "DeletePet")
}

func (s ServerInterfaceImpl) FindPetByID(ctx echo.Context, id int64) error {
	fmt.Println("FindPetByID")
	return ctx.String(http.StatusOK, "FindPetByID")
}

func NewServer() petstore.ServerInterface {
	return ServerInterfaceImpl{}
}

func GetSwagger() (swagger *openapi3.T, err error) {
	return petstore.GetSwagger()
}
