package types

type ResponseData struct {
	Code    uint
	Message string
	Data    interface{}
}

func CreateResponseData() *ResponseData {
	resp := &ResponseData{
		Code:    0,
		Message: "",
	}
	return resp
}

func (rd *ResponseData) ToMap() map[string]interface{} {
	m := map[string]interface{}{
		"code":    rd.Code,
		"message": rd.Message,
		"data":    rd.Data,
	}
	return m
}
