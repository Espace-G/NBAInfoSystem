package models

type Dto struct {
	Status  int         `json:"status"` //status表示状态，0为失败，1为成功
	Message string      `json:"message"`
	Obj     interface{} `json:"obj"`
}
