package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vanbien2402/highest-prime-number/internal/pkg"
)

//InputParams input struct
type InputParams struct {
	InputNumber int `form:"inputNumber" binding:"required"`
}

// FindHighestPrimeNumber Find the highest prime number
// @Summary		 FindHighestPrimeNumber
// @Description  Get input number and return the highest prime number
// @Accept       json
// @Produce      json
// @param inputNumber query int true "positive number"
// @Success      200  {object}  string "The highest prime number lower than 10 is 7"
// @Failure      400  {object}  string "Please enter a positive integer."
// @Failure      404  {object}  string "There is not any prime number lower than 2"
// @Failure      500  {object}  string
// @Router       /prime_number [get]
func FindHighestPrimeNumber(c *gin.Context) {
	var req InputParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a positive integer. " + err.Error()})
		_ = c.Error(err)
		return
	}
	if req.InputNumber < 0 {
		err := fmt.Errorf("Please enter a positive integer.")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		_ = c.Error(err)
		return
	}
	resp, err := findHighestPrimeNumberBeforeN(req.InputNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		_ = c.Error(err)
		return
	}
	c.JSON(http.StatusOK, &resp)
}

func findHighestPrimeNumberBeforeN(N int) (string, error) {
	highestPrime := pkg.HighestPrime(N)
	log.Printf("HighestPrime(%d) result: %d", N, highestPrime)
	if highestPrime == 0 {
		return "", fmt.Errorf("There is not any prime number lower than %d", N)
	} else {
		return fmt.Sprintf("The highest prime number lower than %d is %d", N, highestPrime), nil
	}
}

// HeathCheck Heath Check
// @Summary		 Heath Check
// @Description  Heath Check
// @Accept       json
// @Produce      json
// @Success      200  {object}  string "connected!"
// @Router       /heath_check [get]
func HeathCheck(c *gin.Context) {
	c.String(http.StatusOK, "connected!")
}
