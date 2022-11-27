package api

import (
	"github.com/gin-gonic/gin"

	"github.com/vanbien2402/highest-prime-number/internal/api/handler"
	"github.com/vanbien2402/highest-prime-number/internal/api/middlware"
)

func initRouter(router *gin.Engine) {
	router.Use(
		middlware.Logger(),
		gin.Recovery(),
	)
	router.GET("/heath_check", handler.HeathCheck)
	router.GET("/prime_number", handler.FindHighestPrimeNumber)
}
