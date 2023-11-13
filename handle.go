package myokxapi

import (
	"fmt"
)

type OkxErrorRes struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type OkxTimeRes struct {
	InTime  string `json:"inTime"`  //REST网关接收请求时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123返回的时间是请求验证后的时间
	OutTime string `json:"outTime"` //REST网关发送响应时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123
}
type OkxRestRes[T any] struct {
	OkxErrorRes   //错误信息
	OkxTimeRes    //时间戳
	Data        T `json:"data"` //请求结果
}

func handlerCommonRest[T any](data []byte) (*OkxRestRes[T], error) {
	res := &OkxRestRes[T]{}
	// log.Warn(string(data))
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Error("rest返回值获取失败", err)
	}
	return res, err
}
func (err *OkxErrorRes) handlerError() error {
	if err.Code != "0" && err.Code != "200" {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.Code, err.Msg)
	} else {
		return nil
	}

}
