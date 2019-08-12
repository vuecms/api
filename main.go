package main

import (
	"api/controller"
	"api/spi"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	//_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	//"github.com/swaggo/files"
	//"github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	binding.Validator = new(spi.ValidatorV9)
	r := gin.Default()
	app, err := spi.DefaultInstance()
	if err != nil {
		log.Fatal(err)
	}
	c := controller.NewController(app)

	v1 := r.Group("/api/v1")
	{

		areas := v1.Group("/areas")
		{
			areas.GET("", c.ListAreas)
			areas.PUT("", c.CreateArea)
		}
		questionLibs := v1.Group("/question-libraries")
		{
			questionLibs.GET("", c.ListQuestionLibs)
			questionLibs.PUT("", c.AddQuestionLib)
			questionLibs.PATCH(":id", c.UpdateQuestionLib)
		}

	}
	//r.GET("/", func(context *gin.Context) {
	//	context.Redirect(http.StatusMovedPermanently,"/swagger/index.html")
	//})
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	addr := fmt.Sprintf("%s:%d", app.Config.Http.Host, app.Config.Http.Port)
	err = r.Run(addr)
	if err != nil {
		log.Fatal("Unable to start server")
	}
}
