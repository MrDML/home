package service

import "context"

func (s *CaptchaService)GetCaptchaFromCaseAndData(ctx context.Context, uuid string)  (img []byte, err error) {
	return  s.uc.GetCaptcha(ctx, uuid)
}

