package validator

import (
	"gin_example/common/errcode"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) {
	var errs ValidErrors
	var err error
	switch c.Request.Method {
	case "POST":
		err = c.ShouldBindJSON(v)
	case "PUT":
		err = c.ShouldBindJSON(v)
	case "GET":
		err = c.ShouldBindQuery(v)
	}
	paramError := errcode.ParamError
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			panic(paramError)
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		paramError.Msg = errs.Error()
		panic(paramError)
	}
	return
}
