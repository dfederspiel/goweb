package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"rsi.com/go-training/services"
)

/*StartServer sets up and runs our web server
 */
func StartServer() {
	r := gin.Default()
	r.Routes()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	r.POST("/ping", func(c *gin.Context) {
		var a services.Animal
		c.BindJSON(&a)
		fmt.Println(a)
	})
	r.PUT("/ping/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
	})
	r.DELETE("/ping/:id", func(c *gin.Context) {
		fmt.Println(c.Param("id"))
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
