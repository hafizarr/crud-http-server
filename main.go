package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
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

var listPersonGlobal *[]Person

func main() {
	var listPerson []Person
	listPerson = append(listPerson, Person{
		ID:        1,
		FirstName: "Hafiz",
		LastName:  "Arrahman",
	})
	listPersonGlobal = &listPerson

	r := gin.Default()
	r.GET("/", GetPersons)
	r.POST("/", StorePerson)
	r.PUT("/:id", UpdatePerson)
	r.DELETE("/:id", DeletePerson)

	r.Run()
}

func GetPersons(c *gin.Context) {
	if len(*listPersonGlobal) == 0 {
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

	var listPerson []Person
	listPerson = *listPersonGlobal

	listPerson = append(listPerson, request)
	listPersonGlobal = &listPerson

	Response := Response{
		Data:    request,
		Message: "success",
		Status:  200,
	}
	c.JSON(200, Response)
}

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

	var listPerson []Person
	listPerson = *listPersonGlobal

	for i, data := range listPerson {
		if data.ID == id {
			isIsset = true
			listPerson[i].ID = request.ID
			listPerson[i].FirstName = request.FirstName
			listPerson[i].LastName = request.LastName

			listPersonGlobal = &listPerson

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

func DeletePerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var (
		isIsset bool
	)

	var listPerson []Person
	listPerson = *listPersonGlobal

	for i, data := range listPerson {
		if data.ID == id {
			isIsset = true
			listPerson = RemoveIndex(listPerson, i)

			listPersonGlobal = &listPerson

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

func RemoveIndex(listPerson []Person, index int) []Person {
	return append(listPerson[:index], listPerson[index+1:]...)
}
