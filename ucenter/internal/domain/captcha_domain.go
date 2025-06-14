package domain

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/svc"
)

type CaptchaDomain struct {
	ServiceContext *svc.ServiceContext
}

func NewCaptchaDomain(ctx *svc.ServiceContext) *CaptchaDomain {
	return &CaptchaDomain{
		ServiceContext: ctx,
	}
}

type vaptchaReq struct {
	Id        string `json:"id"`
	Secretkey string `json:"secretkey"`
	Scene     int    `json:"scene"`
	Token     string `json:"token"`
	Ip        string `json:"ip"`
}
type vaptchaRsp struct {
	Success int    `json:"success"`
	Score   int    `json:"score"`
	Msg     string `json:"msg"`
}

func (d *CaptchaDomain) Verify(
	server string,
	vid string,
	key string,
	token string,
	scene int,
	ip string) bool {
	//发送一个post请求
	resp, err := tools.Post(server, &vaptchaReq{
		Id:        vid,
		Secretkey: key,
		Token:     token,
		Scene:     scene,
		Ip:        ip,
	})
	if err != nil {
		logx.Error(err)
		return false
	}
	result := &vaptchaRsp{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		logx.Error(err)
		return false
	}
	return result.Success == 1
}
