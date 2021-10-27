package response

import (
	"fmt"
	"gin_example/common/errcode"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToListResponse(result interface{}, count int64) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"result": result,
		"count":  count,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code, "msg": err.Msg}
	fmt.Println(response)
	r.Ctx.JSON(err.HttpCode, response)
}
