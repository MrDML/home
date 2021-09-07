package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Captcha struct {

}

type CaptchaRepo interface {
	// GetCaptcha 获取验证码
	GetCaptcha(ctx context.Context, uuid string) ([]byte, error)

	// GetImageCodeFromRdb 获取图片验证码code
	GetImageCodeFromRdb(ctx context.Context, uuid string) (string, error)
}

type CaptchaUseCase struct {
	repo CaptchaRepo
	log  *log.Helper
}

func NewCaptchaUseCase(repo CaptchaRepo, logger log.Logger) *CaptchaUseCase{
	return  &CaptchaUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (us *CaptchaUseCase) GetCaptcha(ctx context.Context, uuid string)  (img []byte, err error) {
	return us.repo.GetCaptcha(ctx, uuid)
}

func (us *CaptchaUseCase) GetImageCodeFromRdb(ctx context.Context, uuid string) (string ,error) {
	return us.repo.GetImageCodeFromRdb(ctx, uuid)
}