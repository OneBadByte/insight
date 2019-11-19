package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"insight/src/go/routes"
)

func main() {
	fmt.Println("Starting Server")
	router := gin.Default()
	routes := routes.CreateRouteController()

	// Routes
	router.GET("/getAllPosts", routes.GetAllPosts)
	router.POST("/addPost", routes.AddPost)
	router.POST("/verifyUser", routes.VerifyUser)
	router.POST("/createUser", routes.CreateUser)
	router.Run(":3000")

}
