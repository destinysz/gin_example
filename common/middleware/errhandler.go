package middleware

import (
	"fmt"
	"gin_example/common/errcode"
	"gin_example/common/response"
	"gin_example/global"
	"reflect"

	"github.com/gin-gonic/gin"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				response := response.NewResponse(c)
				t := reflect.TypeOf(r)
				switch fmt.Sprint(t) {
				case "*errcode.Error":
					response.ToErrorResponse(r.(*errcode.Error))
				default:
					fmt.Println(r)
					global.Log.Error(fmt.Sprint(r))
					response.ToErrorResponse(errcode.ServerError)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
