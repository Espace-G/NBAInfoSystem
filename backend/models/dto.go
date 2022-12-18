package models

type Dto struct {
	//status表示状态，0为失败，1为成功
	Status int `json:"status"`
	//message返回处理信息
	Message string `json:"message"`
	//obj返回查询结果等数据
	Obj interface{} `json:"obj"`
}
