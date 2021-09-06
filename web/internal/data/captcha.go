package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "home/api/captcha/service/v1"
	"home/web/internal/biz"
)

type captchaRepo struct {
	data *Data
	log  *log.Helper
}

func NewCaptchaRepo(data *Data, logger log.Logger) biz.CaptchaRepo {
	return &captchaRepo{data: data, log: log.NewHelper(logger)}
}

func (r *captchaRepo)GetCaptcha(ctx context.Context, uuid string)  (img []byte, err error) {

	// 调用远程方法 call
	reply, err := r.data.cc.GetCaptcha(ctx, &v1.GetCaptchaReq{
		Uuid: uuid,
	})
	if err != nil {
		panic(err)
	}
    imgBytes := reply.GetImg()

	return imgBytes, nil
}

func (r *captchaRepo) GetImageCodeFromRdb(ctx context.Context, uuid string, imgCode *string) error {
	// 调用远程方法 call
	reply, err := r.data.cc.GetImageCodeFromRdb(ctx, &v1.GetImageCodeFromRdbReq{
		Uuid: uuid,
	})
	if err != nil {
		r.log.Fatal("call remote GetImageCodeFromRdb err = ",err)
	}
	// 给验证码赋值
	*imgCode = reply.GetImgCode()
	return err
}