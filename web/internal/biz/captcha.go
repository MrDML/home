package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Captcha struct {
	str string
}

type CaptchaRepo interface {
	GetCaptcha(ctx context.Context, uuid string) (img []byte, err error)

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