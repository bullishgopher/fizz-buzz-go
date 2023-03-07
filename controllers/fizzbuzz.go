package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FizzBuzzController struct{}

type fizzbuzzBody struct {
	Count int `json:"count"`
}

func Fizzbuzz(n int) string {

	var name string = ""

	if n%3 == 0 {
		name = os.Getenv("Fizz")
	}

	if n%5 == 0 {
		name += os.Getenv("Buzz")
	}

	return name
}

// Status godoc
// @Summary Get a fizzbuzz message
// @Description Get a fizzbuzz message
// @Produce  json
// @Success 200 {object} message
// @Router /fizzbuzz/fizzbuzz [post]
func (h FizzBuzzController) FizzBuzz(c *gin.Context) {
	var newCount fizzbuzzBody

	if err := c.BindJSON(&newCount); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": Fizzbuzz(newCount.Count),
	})
}
