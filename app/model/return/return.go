package _return

type Entity struct {
	Data interface{}      `json:"Data"`
	Code      int         `json:"Code"`
	Msg      string         `json:"Msg"`
}
const OkMsg string = "成功"
const OKCode  =0
const ErrMsg string = "错误"
const ErrCode  =100
func Ok(data interface{}) Entity {
	return Entity{Data:data,Code:OKCode,Msg: OkMsg}
}
func Err()Entity  {
	return Entity{Code:ErrCode,Msg: ErrMsg}
}
func OkNoData() Entity {
	return Entity{Code:OKCode,Msg: OkMsg}
}
func ErrAddMsg(msg string)Entity  {
	return Entity{Code:ErrCode,Msg: msg}
}