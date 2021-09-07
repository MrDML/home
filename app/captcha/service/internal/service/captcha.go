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


func (s *CaptchaService) GetImageCodeFromRdb(ctx context.Context, req *v1.GetImageCodeFromRdbReq) (reply * v1.GetImageCodeFromRdbReply, err error) {

	code, err := s.uc.GetImageCodeFromRdb(ctx, req.Uuid)
	reply = &v1.GetImageCodeFromRdbReply{
		ImgCode: code,
	}
	return reply, err
}


// SendSmsCode 发送短信验证码
func (s *CaptchaService) SendSmsCode(ctx context.Context, req *v1.SendSmsCodeReq) (*v1.SendSmsCodeReply, error) {
	 err := s.uc.SendSmsCode(ctx, req.GetPhone())
	 return nil, err
}
// GetSmsCode 获取短信验证码
func (s *CaptchaService) GetSmsCode(ctx context.Context, req *v1.GetSmsCodeReq) (*v1.GetSmsCodeReply, error) {

	 smsCode, err := s.uc.GetSmsCode(ctx, req.GetPhone())
	if err != nil {
		return nil, err
	}
	 reply := &v1.GetSmsCodeReply{
	 	Code: smsCode,
	 }
	 return reply, nil

}