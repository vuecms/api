package controller

import (
	"api/spi"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type Controller struct {
	App *spi.Application
}


func (c *Controller) ListQuestionLibs(ctx *gin.Context) {
	paginator := c.App.QuestionSpi.ListQuestionLibsPaginate()
	ctx.JSON(200, paginator)
}


func (c *Controller) AddQuestionLib(ctx *gin.Context) {
	//b := binding.Default(ctx.Request.Method, ctx.ContentType())
	s := &AddQuestionParam{}
	//err := ctx.ShouldBind(s)
	err := ctx.ShouldBindJSON(s)
	if err != nil {
		ctx.JSON(400, NewValidatorError(err))
	} else {
		ctx.JSON(201, nil)
	}
}

func (c *Controller) UpdateQuestionLib(ctx *gin.Context) {

}

func NewController(app *spi.Application) *Controller {
	var ctrl = new(Controller)
	ctrl.App = app
	return ctrl
}

type CommonError struct {
	Errors []FieldError `json:"errors"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidatorError(err error) CommonError {

	res := CommonError{}
	res.Errors = make([]FieldError, 0)
	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		fe := FieldError{
			Field:   e.Field(),
			Message: fmt.Sprint(e),
		}
		res.Errors = append(res.Errors, fe)
	}
	return res
}
