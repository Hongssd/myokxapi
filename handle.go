package myokxapi

import (
	"fmt"
)

type OkxErrorRes struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type OkxRestRes[T any] struct {
	OkxErrorRes   //错误信息
	Result      T `json:"data"` //请求结果
}

func handlerCommonRest[T any](data []byte) (*OkxRestRes[T], error) {
	res := &OkxRestRes[T]{}
	log.Warn(string(data))
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
