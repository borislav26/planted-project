package main

import (
	"backend/api"
	"backend/builder"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	repositories := builder.BuildRepositories()
	services := builder.BuildServices(repositories)

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://frontend:3000"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	api.AddEndpoints(app, services)
	app.Run(":80")
}
