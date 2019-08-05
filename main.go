package main

import (
	"exam/controller"
	_ "exam/docs"
	"exam/spi"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/go-sql-driver/mysql"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"net/http"
)

// @title Exam API
// @version 1.0
// @description This is an api server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

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
	r.GET("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently,"/swagger/index.html")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	addr := fmt.Sprintf("%s:%d", app.Config.Http.Host, app.Config.Http.Port)
	err = r.Run(addr)
	if err != nil {
		log.Fatal("Unable to start server")
	}
}
