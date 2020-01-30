package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"rsi.com/go-training/services"
)

func main() {
	g := gin.Default()
	g.Use(static.Serve("/", static.LocalFile("./build", true)))

	api := g.Group("/api")
	{
		services.RegisterRoutes(api)
	}

	_ = g.Run()
}
