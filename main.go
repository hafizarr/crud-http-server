package main

import (
	"crud-http-server/docs"
	"strconv"

	"github.com/gin-gonic/gin"
	// "github.com/go-project-name/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

var listPersonGlobal []*Person

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiV1 := r.Group("/api/v1")
	{
		person := apiV1.Group("/person")
		{
			person.GET("/", GetPersons)
			person.POST("/", StorePerson)
			person.PUT("/:id", UpdatePerson)
			person.DELETE("/:id", DeletePerson)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}

// @BasePath /api/v1
// GetPerson godoc
// @Summary get person
// @Schemes
// @Description do get person
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Success 400 {object} Response
// @Router /person [get]
func GetPersons(c *gin.Context) {
	if len(listPersonGlobal) == 0 {
		Response := Response{
			Data:    listPersonGlobal,
			Message: "not found",
			Status:  400,
		}
		c.JSON(200, Response)
		return
	}

	Response := Response{
		Data:    listPersonGlobal,
		Message: "success",
		Status:  200,
	}
	c.JSON(200, Response)
}

// @BasePath /api/v1
// StorePerson godoc
// @Summary store person
// @Schemes
// @Description do store person
// @Tags person
// @Accept json
// @Produce json
// @Param   Body body   Person  true  "payload"
// @Success 200 {object} Response
// @Success 400 {object} Response
// @Router /person [post]
func StorePerson(c *gin.Context) {
	var request Person
	err := c.BindJSON(&request)
	if err != nil {
		Response := Response{
			Data:    request,
			Message: err.Error(),
			Status:  400,
		}
		c.JSON(400, Response)
		return
	}

	listPersonGlobal = append(listPersonGlobal, &request)

	Response := Response{
		Data:    request,
		Message: "success",
		Status:  200,
	}
	c.JSON(200, Response)
}

// @BasePath /api/v1
// UpdatePerson godoc
// @Summary update person
// @Schemes
// @Description do update person
// @Tags person
// @Accept json
// @Produce json
// @Param   id  query   string  true  "id"
// @Param   Body body   Person  true  "payload"
// @Success 200 {object} Response
// @Success 400 {object} Response
// @Router /person/:id [put]
func UpdatePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		request Person
		isIsset bool
	)
	err := c.BindJSON(&request)
	if err != nil {
		Response := Response{
			Data:    request,
			Message: err.Error(),
			Status:  400,
		}
		c.JSON(400, Response)
		return
	}

	for i, data := range listPersonGlobal {
		if data.ID == id {
			isIsset = true
			listPersonGlobal[i].ID = request.ID
			listPersonGlobal[i].FirstName = request.FirstName
			listPersonGlobal[i].LastName = request.LastName

			break
		}
	}

	if !isIsset {
		Response := Response{
			Data:    request,
			Message: "not found",
			Status:  200,
		}
		c.JSON(200, Response)
		return
	}

	Response := Response{
		Data:    request,
		Message: "success",
		Status:  200,
	}
	c.JSON(200, Response)
}

// @BasePath /api/v1
// DeletePerson godoc
// @Summary delete person
// @Schemes
// @Description do delete person
// @Tags person
// @Accept json
// @Produce json
// @Param   id  query   string  true  "id"
// @Param   Body body   Person  true  "payload"
// @Success 200 {object} Response
// @Success 400 {object} Response
// @Router /person/:id [delete]
func DeletePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		isIsset bool
	)

	for i, data := range listPersonGlobal {
		if data.ID == id {
			isIsset = true
			listPersonGlobal = RemoveIndex(listPersonGlobal, i)

			break
		}
	}

	if !isIsset {
		Response := Response{
			Data:    id,
			Message: "not found",
			Status:  200,
		}
		c.JSON(200, Response)
		return
	}

	Response := Response{
		Data:    id,
		Message: "success",
		Status:  200,
	}
	c.JSON(200, Response)
}

func RemoveIndex(listPerson []*Person, index int) []*Person {
	return append(listPerson[:index], listPerson[index+1:]...)
}
