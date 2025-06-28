package routes

import (
	_ "library-app/docs"
	"library-app/src/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
