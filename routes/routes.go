package routes

import (
    "golang-crud-krs/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/krs", controllers.CreateKRS)
    r.GET("/krs", controllers.GetAllKRS)
    r.GET("/krs/:id", controllers.GetKRS)
    r.PUT("/krs/:id", controllers.UpdateKRS)
    r.DELETE("/krs/:id", controllers.DeleteKRS)
    return r
}
