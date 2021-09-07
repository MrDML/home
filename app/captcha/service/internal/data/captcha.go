package data

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"github.com/go-kratos/kratos/v2/log"
	"home/app/captcha/service/internal/biz"
	"image/color"
	"time"
)

type captchaRepo struct {
	data *Data
	log *log.Helper
}


func NewCaptchaRepo(data *Data, logger log.Logger) biz.CaptchaRepo {
	return &captchaRepo{data: data, log: log.NewHelper(logger)}
}


func (r *captchaRepo) GetCaptcha(ctx context.Context, uuid string)  (img []byte, err error) {

	// 业务逻辑处理

	// 初始化对象
	cap := captcha.New()
	// 设置字体
	cap.SetFont("./app/captcha/service/conf/comic.ttf")
	// 创建验证码 4 个字符 captcha.NUM 字符模式数字类型
	// 设置干扰强度
	cap.SetDisturbance(captcha.MEDIUM)
	//  设置前景色 可以多个 随机替换文字颜色 默认黑色
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	// 设置背景色 可以多个 随机替换背景色 默认是白色
	cap.SetBkgColor(color.RGBA{0, 128, 128, 128}, color.RGBA{255, 0, 0, 128})

	image, code := cap.Create(6, captcha.ALL)

	if code != "" {
		// 将code存入redis
		var key = "imageCode:" + uuid
		set, err := r.data.rdb.SetEX(ctx, key, code, 60 * time.Second).Result()
		if err != nil {
			panic(err)
		}
		r.log.Infof("======>save image code :%s", set)
	}

	// 序列化
	imgBytes, errMarshal :=  json.Marshal(image)

	return imgBytes, errMarshal
}


func (r*captchaRepo) GetImageCodeFromRdb(ctx context.Context, uuid string) (string, error) {
	imageCode, err := r.data.rdb.Get(ctx, uuid).Result()
	return imageCode, err
}


