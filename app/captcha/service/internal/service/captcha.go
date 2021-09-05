package service

import (
	"context"
	v1 "home/api/captcha/service/v1"
)



// GetCaptcha rpc方法
func (s *CaptchaService) GetCaptcha(ctx context.Context, req *v1.GetCaptchaReq) (reply *v1.GetCaptchaReply, err error) {

	// 业务逻辑处理
	imgBytes, err :=  s.uc.GetCaptcha(ctx,req.Uuid)

	reply = &v1.GetCaptchaReply {
		Img: imgBytes,
	}

	return reply,err
}