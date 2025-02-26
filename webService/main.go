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
		Id:     "0",
		Name:   "banana",
		Flavor: "sweet",
		Color:  "yelllow",
	},
	{
		Id:     "1",
		Name:   "lemon",
		Flavor: "sour",
		Color:  "yellow",
	},
}

func getFruits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fruits)
}

func postFruits(c *gin.Context) {
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

func updateFruitById(c *gin.Context) {
	var updatedFruit fruit
	id := c.Param("id")
	for i, a := range fruits {
		if a.Id == id {
			// if err := c.BindJSON(&a); err != nil {
			// 	return
			// }
			if updatedFruit.Name != "" {
				fruits[i].Name = updatedFruit.Name
				if err := c.BindJSON(&a); err != nil {
					return
				}
				c.IndentedJSON(http.StatusOK, a)
			}
			if updatedFruit.Color != "" {
				fruits[i].Color = updatedFruit.Color
			}
			if updatedFruit.Flavor != "" {
				fruits[i].Flavor = updatedFruit.Flavor
			}
		}
	}
	c.IndentedJSON(http.StatusOK, updatedFruit)
}

func teste(c *gin.Context) {
	var frutaTeste fruit
	// if err := c.BindJSON(&frutaTeste); err != nil {
	// 	return
	// }
	c.BindJSON(&frutaTeste)
	fmt.Println(frutaTeste)
	c.IndentedJSON(http.StatusOK, frutaTeste)
}

func main() {
	fmt.Println("Hello world")

	router := gin.Default()
	router.GET("/fruits", getFruits)
	router.POST("/fruits", postFruits)
	router.PUT("/fruits/:id", updateFruitById)
	router.GET("/fruits/:id", getFruitById)
	router.GET("/teste", teste)
	router.POST("/teste", teste)

	router.Run("localhost:8080")
}
