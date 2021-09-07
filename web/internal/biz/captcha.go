package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Captcha struct {
	str string
}

type CaptchaRepo interface {

	// GetCaptcha 获取验证码
	GetCaptcha(ctx context.Context, uuid string) (img []byte, err error)

	// GetImageCodeFromRdb 从redis获取验证码
	GetImageCodeFromRdb(ctx context.Context, uuid string, imgCode *string) error

	// SendSmsCode 发送短信验证码
	SendSmsCode(ctx context.Context, phone string) error

}

type CaptchaUseCase struct {
	repo CaptchaRepo
	log *log.Helper
}

func NewCaptchaUseCase(repo CaptchaRepo, logger log.Logger) *CaptchaUseCase{
	return &CaptchaUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (us *CaptchaUseCase) GetCaptcha(ctx context.Context, uuid string) (img []byte, err error) {
	return us.repo.GetCaptcha(ctx ,uuid)
}

func (us *CaptchaUseCase) GetImageCodeFromRdb(ctx context.Context, uuid string, imgCode *string) error {
	return us.repo.GetImageCodeFromRdb(ctx, uuid, imgCode)
}

func (us *CaptchaUseCase) SendSmsCode(ctx context.Context, phone string) error  {
	return us.repo.SendSmsCode(ctx, phone)
}