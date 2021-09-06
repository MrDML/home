package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "home/api/captcha/service/v1"
	"home/web/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCaptchaService)

type CaptchaService struct {
	v1.UnimplementedCaptchaServer

	cc  *biz.CaptchaUseCase
	log *log.Helper
}



func NewCaptchaService(cc *biz.CaptchaUseCase, logger log.Logger) *CaptchaService {
	return &CaptchaService{cc: cc, log: log.NewHelper(logger)}
}



