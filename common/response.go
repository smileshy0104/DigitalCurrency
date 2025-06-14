package common

// BizCode 是业务状态码的整数类型
type BizCode int

// SuccessCode 表示操作成功的业务状态码
const SuccessCode BizCode = 0

// Result 是 API 返回结果的结构体，包含业务状态码、消息和数据
type Result struct {
	Code    BizCode `json:"code"`    // 业务状态码
	Message string  `json:"message"` // 状态描述信息
	Data    any     `json:"data"`    // 返回的数据
}

// NewResult 创建并返回一个新的 Result 对象
func NewResult() *Result {
	return &Result{}
}

// Success 将 Result 设置为成功状态，并填充数据
func (r *Result) Success(data any) {
	r.Code = SuccessCode
	r.Message = "success"
	r.Data = data
}

// Fail 将 Result 设置为失败状态，并设置错误码和错误信息
func (r *Result) Fail(code BizCode, msg string) {
	r.Code = code
	r.Message = msg
}

// Deal 根据 error 是否为 nil 来决定将 Result 设为成功或失败状态
// 如果 err 不为 nil，则设为失败并使用错误信息；否则设为成功并填充数据
func (r *Result) Deal(data any, err error) *Result {
	if err != nil {
		r.Fail(-999, err.Error())
	} else {
		r.Success(data)
	}
	return r
}
