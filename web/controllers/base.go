package controllers

type ApiJson struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

func ApiResource(status bool, objects interface{}, msg string) (apijson *ApiJson) {
	apijson = &ApiJson{Status: status, Data: objects, Msg: msg}
	return
}
