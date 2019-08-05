package controller

import (
	"exam/spi"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type Controller struct {
	App *spi.Application
}

// ListQuestionLibs godoc
// @Summary List question libs
// @Description get question libs
// @Tags question-libraries
// @Accept  json
// @Produce  json
// @Success 200 {object} protocol.QuestionLibraryPaginator
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /question-libraries [get]
func (c *Controller) ListQuestionLibs(ctx *gin.Context) {
	paginator := c.App.QuestionSpi.ListQuestionLibsPaginate()
	ctx.JSON(200, paginator)
}

// AddQuestionLib godoc
// @Summary Add a question lib
// @Description add by json question lib
// @Tags question-libraries
// @Accept  json
// @Produce  json
// @Param body questionLib body controller.AddQuestionParam true "Add QuestionLibrary"
// @Success 200 {object} model.QuestionLibrary
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /question-libraries [put]
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

// UpdateQuestionLib godoc
// @Summary Update a question lib
// @Description Update by json question lib
// @Tags question-libraries
// @Accept  json
// @Produce  json
// @Param  id path int true "QuestionLib ID"
// @Param  questionLib body model.QuestionLibrary true "Update account"
// @Success 200 {object} model.QuestionLibrary
// @Failure 400 {object} protocol.Error
// @Failure 404 {object} protocol.Error
// @Failure 500 {object} protocol.Error
// @Router /question-libraries/{id} [patch]
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
