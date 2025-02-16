package routes

import (
    "golang-crud-users/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/users", controllers.CreateUser)
    r.GET("/users", controllers.GetUsers)
    r.GET("/users/:id", controllers.GetUser)
    r.PUT("/users/:id", controllers.UpdateUser)
    r.DELETE("/users/:id", controllers.DeleteUser)
    return r
}
