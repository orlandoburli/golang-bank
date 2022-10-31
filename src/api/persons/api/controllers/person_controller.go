package controllers

import (
	"bank/api/persons/models"
	"bank/api/persons/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

var (
	_controller *PersonController
)

type PersonController struct {
	service *services.PersonService
}

func GetController() *PersonController {
	if _controller == nil {
		_controller = &PersonController{service: services.GetPersonService()}
	}
	return _controller
}

func Routes(group *gin.RouterGroup) {
	controller := GetController()
	persons := group.Group("/persons")
	{
		persons.GET("/", controller.ListPersons)
		persons.GET("/:id", controller.GetPerson)
		persons.POST("/", controller.CreatePerson)
		persons.PUT("/:id", controller.UpdatePerson)
		persons.DELETE("/:id", controller.DeletePerson)
	}
}

func (controller *PersonController) ListPersons(c *gin.Context) {
	var pageParam, ok1 = c.GetQuery("page")
	var sizeParam, ok2 = c.GetQuery("size")

	if !ok1 {
		c.AbortWithStatusJSON(400, models.ErrorResponse{Message: "Invalid query params", Details: []string{"Missing page param"}})
		return
	}
	if !ok2 {
		c.AbortWithStatusJSON(400, models.ErrorResponse{Message: "Invalid query params", Details: []string{"Missing size param"}})
		return
	}
	var page, err1 = strconv.Atoi(pageParam)
	var size, err2 = strconv.Atoi(sizeParam)

	if err1 != nil || err2 != nil {
		var details = make([]string, 0, 0)
		if err1 != nil {
			details = append(details, err1.Error())
		}
		if err2 != nil {
			details = append(details, err2.Error())
		}

		c.AbortWithStatusJSON(400, models.ErrorResponse{Message: "Invalid query params", Details: details})
		return
	}

	var persons, totalRecords, pages = controller.service.GetPersons(page, size)

	c.JSON(http.StatusOK, models.PersonResult{Data: *persons, TotalRecords: *totalRecords, Pages: *pages})
}

func (controller *PersonController) GetPerson(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(404, models.ErrorResponse{Message: err.Error()})
		return
	}
	person, err := controller.service.GetPerson(id)

	if err != nil {
		c.AbortWithStatusJSON(404, models.ErrorResponse{Message: "Invalid param: id"})
		return
	}

	c.JSON(200, person)
}

func (controller *PersonController) CreatePerson(c *gin.Context) {
	_person := new(models.Person)

	errBind := c.BindJSON(_person)

	if errBind != nil {
		c.AbortWithStatusJSON(400, models.ErrorResponse{
			Message: errBind.Error(),
		})
		return
	}

	person, err := controller.service.CreatePerson(_person)

	if err != nil {
		c.AbortWithStatusJSON(400, models.ErrorResponse{
			Message: err.Error(),
		})
		return

	}
	c.Header("Location", fmt.Sprintf("/v1/persons/%v", person.ID))
	c.Status(http.StatusCreated)
}

func (controller *PersonController) UpdatePerson(c *gin.Context) {
	id, err1 := uuid.Parse(c.Param("id"))

	if err1 != nil {
		c.AbortWithStatusJSON(404, models.ErrorResponse{Message: "Invalid param: id"})
		return
	}

	_person := new(models.Person)

	errBind := c.BindJSON(_person)

	if errBind != nil {
		c.AbortWithStatusJSON(400, models.ErrorResponse{
			Message: "Error parsing data",
			Details: []string{errBind.Error()},
		})
		return
	}

	person, err := controller.service.UpdatePerson(id, _person)

	if err != nil {
		c.AbortWithStatusJSON(400, models.ErrorResponse{
			Message: "Error updating person",
			Details: []string{err.Error()},
		})
		return
	}

	c.JSON(200, person)
}

func (controller *PersonController) DeletePerson(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(404, models.ErrorResponse{Message: "Invalid param: id"})
		return
	}

	var err2 = controller.service.DeletePerson(&id)

	if err2 != nil {
		c.AbortWithStatusJSON(400, models.ErrorResponse{
			Message: "Error deleting person",
			Details: []string{err2.Error()},
		})
		return
	}

	c.Status(204)
}
