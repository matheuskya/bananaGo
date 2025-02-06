package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type fruit struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Flavor string `json:"flavor"`
	Color  string `json:"color"`
}

var fruits []fruit = []fruit{
	{
		Id:     0,
		Name:   "banana",
		Flavor: "sweet",
		Color:  "yelllow",
	},
	{
		Id:     1,
		Name:   "lemon",
		Flavor: "sour",
		Color:  "yellow",
	},
}

func getFruits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fruits)
}

func postAlbums(c *gin.Context) {
	var newFruit fruit

	if err := c.BindJSON(&newFruit); err != nil {
		return
	}

	fruits = append(fruits, newFruit)
	c.IndentedJSON(http.StatusCreated, newFruit)
}

func getFruitById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range fruits {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Fruit not found :("})

}

func main() {
	fmt.Println("Hello world")
	router := gin.Default()
	router.GET("/fruits", getFruits)
	router.Run("localhost:8080")
}
