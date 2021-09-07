package service

import (
	"context"
)

// GetCaptchaFromCaseAndData 获取验证码图片二进制数据
func (s *CaptchaService) GetCaptchaFromCaseAndData(ctx context.Context, uuid string) (img []byte, err error) {
	return s.cc.GetCaptcha(ctx, uuid)
}

// GetImageCodeFromRedis 获取短信验证码字符串
func (s *CaptchaService) GetImageCodeFromRedis(ctx context.Context, uuid string, imgCode *string) error {
	err := s.cc.GetImageCodeFromRdb(ctx, uuid, imgCode)
	return err
}

func (s *CaptchaService) SendSmsCodeFromCaseInData(ctx context.Context, phone string) error {
	 err := s.cc.SendSmsCode(ctx, phone)
	return err
}



