package routes

import (
    "github.com/gin-gonic/gin"
    "library-app/src/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    api := r.Group("/books")
    {
        api.GET("", controllers.GetBooks)
        api.GET("/:id", controllers.GetBook)
        api.POST("", controllers.CreateBook)
        api.PUT("/:id", controllers.UpdateBook)
        api.DELETE("/:id", controllers.DeleteBook)
    }
    return r
}