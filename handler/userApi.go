package handler

import (
	"context"
	"encoding/json"
	"errors"
	serviceUser "github.com/ashoreDove/parasite-user/proto/user"

	userApi "github.com/ashoreDove/parasite-userApi/proto/userApi"
	log "github.com/micro/go-micro/v2/logger"
)

type UserApi struct {
	UserModuleService serviceUser.UserService
}

// UserApi.Call 通过API向外暴露为/userApi/call，接收http请求
// 即：/userApi/call请求会调用go.micro.api.userApi 服务的UserApi.Call方法
func (u *UserApi) Register(ctx context.Context, req *userApi.Request, resp *userApi.Response) error {
	log.Info("收到对 /userApi/register 访问请求")
	var post map[string]interface{}
	err := json.Unmarshal([]byte(req.Body), &post)
	if err != nil {
		log.Error(err)
		return err
	}
	params := post["params"]
	param := params.(map[string]interface{})
	if _, ok := param["account"]; !ok {
		resp.StatusCode = 500
		return errors.New("参数account异常")
	}
	if _, ok := param["password"]; !ok {
		resp.StatusCode = 500
		return errors.New("参数password异常")
	}
	if _, ok := param["code"]; !ok {
		resp.StatusCode = 500
		return errors.New("验证码为空")
	}
	if _, ok := param["nickname"]; !ok {
		resp.StatusCode = 500
		return errors.New("参数nickname异常")
	}
	log.Info(req)
	registerResp, err := u.UserModuleService.Register(context.TODO(), &serviceUser.RegisterRequest{
		Account:  param["account"].(string),
		Password: param["password"].(string),
		Nickname: param["nickname"].(string),
		Code:     param["code"].(string),
	})
	if err != nil {
		resp.StatusCode = 400
		return err
	}
	//转类型
	b, err := json.Marshal(registerResp)
	if err != nil {
		return err
	}
	resp.StatusCode = 200
	resp.Body = string(b)
	return nil
}

func (u *UserApi) Login(ctx context.Context, req *userApi.Request, resp *userApi.Response) error {
	log.Info("收到对 /userApi/login 访问请求")
	var post map[string]interface{}
	err := json.Unmarshal([]byte(req.Body), &post)
	if err != nil {
		log.Error(err)
		return err
	}
	params := post["params"]
	param := params.(map[string]interface{})
	if _, ok := param["account"]; !ok {
		resp.StatusCode = 500
		log.Info("参数account异常")
		return errors.New("参数account异常")
	}
	if _, ok := param["password"]; !ok {
		resp.StatusCode = 500
		log.Info("参数password异常")
		return errors.New("参数password异常")
	}
	if _, ok := param["code"]; !ok {
		resp.StatusCode = 500
		log.Info("验证码为空")
		return errors.New("验证码为空")
	}
	log.Info(req)
	loginResp, err := u.UserModuleService.Login(context.TODO(), &serviceUser.LoginRequest{
		Account:  param["account"].(string),
		Password: param["password"].(string),
		Code:     param["code"].(string),
	})
	if err != nil {
		resp.StatusCode = 400
		log.Info("err:", err)
		return err
	}
	//转类型
	b, err := json.Marshal(loginResp)
	if err != nil {
		return err
	}
	resp.StatusCode = 200
	resp.Body = string(b)
	return nil
}

func (u *UserApi) SendMessage(ctx context.Context, req *userApi.Request, resp *userApi.Response) error {
	log.Info("收到对 /userApi/sendMessage 访问请求")
	return nil
}
