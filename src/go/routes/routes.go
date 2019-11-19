package routes

import (
	"github.com/gin-gonic/gin"
	"insight/src/go/database"
	"log"
)

// RouteController controls the database for each route
type RouteController struct {
	dbConn database.DatabaseConnection
}

// CreateRouteController will create a database connection and return a RouteController
func CreateRouteController() RouteController {
	return RouteController{
		database.CreateDatabaseConnection(),
	}
}

func (rc RouteController) VerifyUser(c *gin.Context) {
	if rc.ValidateUser(c) {
		c.Status(200)
	} else {
		c.Status(403)
	}
}

func (rc RouteController) CreateUser(c *gin.Context) {
	var user Auth
	c.BindJSON(&user)
	if !rc.dbConn.CheckIfUserExists(user.Username) {
		rc.dbConn.CreateUser(user.Username, user.Password)
	}
	if rc.dbConn.VerifyPasswordByUsername(user.Username, user.Password) {
		c.Status(200)
	} else {
		c.Status(500)
	}
}

func (rc RouteController) GetAllPosts(c *gin.Context) {
	if rc.ValidateUser(c) {
		c.JSON(200, rc.dbConn.GetAllPosts())
	} else {
		c.String(403, "username and password not valid")
	}

}

func (rc RouteController) AddPost(c *gin.Context) {
	if rc.ValidateUser(c) {
		var post database.Post
		c.BindJSON(&post)
		log.Printf("recieved %v", post)
		err := rc.dbConn.AddPost(post)
		if err != nil {
			c.String(500, "Couldn't add post: %v", err)
			return
		}
		c.Status(200)
	} else {
		c.String(403, "username and password not valid")
	}
}
