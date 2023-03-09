package controllers

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FizzBuzzController struct{}

type fizzbuzzBody struct {
	Count int `json:"count"`
}

var ErrEnv = errors.New("cannot find environment variable")

func Fizzbuzz(n int) (string, error) {

	var name string = ""
	fizz := os.Getenv("Fizz")
	buzz := os.Getenv("Buzz")
	if fizz == "" || buzz == "" {
		return "", ErrEnv
	}

	if n%3 == 0 {
		name = fizz
	}

	if n%5 == 0 {
		name += buzz
	}

	return name, nil
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Input correct number",
		})
		return
	}

	res, resErr := Fizzbuzz(newCount.Count)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}

// Status godoc
// @Summary Get a fizzbuzz message
// @Description Get a fizzbuzz message
// @Produce  json
// @Success 200 {object} message
// @Router /fizzbuzz/messages [get]
func (h FizzBuzzController) GetMessageByCount(c *gin.Context) {
	count := c.Query("count")

	countVal, countErr := strconv.Atoi(count)
	if countErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Input correct number",
		})
		return
	}

	res, resErr := Fizzbuzz(countVal)
	if resErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": res,
	})
}
