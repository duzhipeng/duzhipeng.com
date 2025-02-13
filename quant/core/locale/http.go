package locale

// 全局错误码

type ResponseField struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type StatusCode struct {
	// 2xx系列
	Ok                 ResponseField
	OKButAlreadyExists ResponseField

	// 4xx系列
	BadInput     ResponseField
	Unauthorized ResponseField
	NotFound     ResponseField
	Forbidden    ResponseField

	// 5xx系列
	ServerError ResponseField
}

var Result = StatusCode{
	// 2xx系列
	Ok:                 ResponseField{200000, "", nil},
	OKButAlreadyExists: ResponseField{200001, "该信息已存在，不能重复创建。", nil},

	// 4xx系列
	BadInput:     ResponseField{400000, "输入内容有误。请按要求输入信息。", nil},
	Unauthorized: ResponseField{401000, "密码等身份信息有误。请核实后重试。", nil},
	Forbidden:    ResponseField{403000, "您无权进行此操作。", nil},
	NotFound:     ResponseField{404000, "查无此信息。", nil},

	// 5xx系列
	ServerError: ResponseField{500000, "服务异常。请联系工作人员处理。", nil},
}
