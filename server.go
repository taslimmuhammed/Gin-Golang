package main

import (
	"io"
	"net/http"
	"os"
	service "taslimmuhammed/golang-Gin/Service"
	controller "taslimmuhammed/golang-Gin/controller"
	"taslimmuhammed/golang-Gin/middlewares"

	gindump "github.com/tpkeeper/gin-dump"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	//server := gin.Default()
	//Creating route
	// server.GET("/test",func(ctx *gin.Context){
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Ok!",
	// 	})
	// })
	setupLogOutput()
	server := gin.New()

	//adding static html files
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
	//adding middlewares
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())
    
	//grouping
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/posts", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"})
			}
	
		})
	}
	server.POST("/login", func(ctx *gin.Context){
		token := loginController.Login(ctx)
	})
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos",videoController.ShowAll)
	}

	//runing server
	server.Run(":8080")

}


//steps in hosting heroku
//-create Procfile
//-Run cmd "go build -o bin/golang-gin-poc"