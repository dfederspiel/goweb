package main

/*
* What is a package?
* Wow, cool! So, then.. where are they how do I reference them?
* Ahhh.. ok, yeah.. I think I get it. So why not just keep everything in main?
* Sweet, what are good some good examples for how to organize all this stuff?
* Solid!
 */
import (
	"fmt"
	"rsi.com/go-training/services"

	"github.com/gin-gonic/gin"
)

func main() {

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
