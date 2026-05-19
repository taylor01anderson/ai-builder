package main

import (
	"ai-builder/config"
	"ai-builder/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

    config.ConnectDB()

    r := gin.Default()

    r.Use(cors.Default())

    api := r.Group("/api")
    {
        routes.AuthRoutes(api)
        routes.AssistantRoutes(api)
        routes.ChatRoutes(api)
    }

    r.Run(":8080")
}