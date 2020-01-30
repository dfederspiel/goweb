package main

import (
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/services"
)

/*
* What is a package?
* Wow, cool! So, then.. where are they how do I reference them?
* Ahhh.. ok, yeah.. I think I get it. So why not just keep everything in main?
* Sweet, what are good some good examples for how to organize all this stuff?
* Solid!
 */

func main() {
	g := gin.Default()
	services.RegisterRoutes(&g.RouterGroup)
	_ = g.Run()
}
