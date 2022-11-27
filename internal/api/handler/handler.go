package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vanbien2402/highest-prime-number/internal/pkg"
)

//InputParams input struct
type InputParams struct {
	InputNumber int `form:"inputNumber"`
}

//FindHighestPrimeNumber Find the highest prime number lower than input number
func FindHighestPrimeNumber(c *gin.Context) {
	var req InputParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please enter a positive integer. " + err.Error()})
		_ = c.Error(err)
		return
	}
	if req.InputNumber < 0 {
		err := errors.New("Please enter a positive integer.")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		_ = c.Error(err)
		return
	}
	resp := findHighestPrimeNumberBeforeN(req.InputNumber)
	c.JSON(http.StatusOK, &resp)
}

func findHighestPrimeNumberBeforeN(N int) string {
	highestPrime := pkg.HighestPrime(N)
	log.Printf("HighestPrime(%d) result: %d", N, highestPrime)
	if highestPrime == 0 {
		return fmt.Sprintf("There is not any prime number lower than %d", N)
	} else {
		return fmt.Sprintf("The highest prime number lower than %d is %d", N, highestPrime)
	}
}
