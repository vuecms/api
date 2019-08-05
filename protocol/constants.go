package protocol

var success = &Success{
	Code:    1000,
	Message: "Success",
}

type ResultCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	ResultCode
	Code    int    `json:"code" example:"10000"`
	Message string `json:"message" example:"Success"`
}

type Error struct {
	ResultCode
	Code    int    `json:"code" example:"10001"`
	Message string `json:"message" example:"Some Error"`
}
