package app

import (
	"github.com/gin-gonic/gin"
)
var (
	router = gin.Default()
)
func StartApplication(){
	// Start the application
	mapUrls()
	router.Run(":8080")
}