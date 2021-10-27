package errcode

var (
	NotFound    = NewError(404, "没有这个地址", 404000)
	ServerError = NewError(500, "服务内部错误", 5000000)
	ParamError  = NewError(400, "参数错误", 400001) // 参数错误 验证的时候可以替换
)

type Error struct {
	HttpCode int    `json:"-"`
	Msg      string `json:"msg"`
	Code     int    `json:"code"`
}

func NewError(httpCode int, msg string, code int) *Error {
	return &Error{HttpCode: httpCode, Msg: msg, Code: code}
}
