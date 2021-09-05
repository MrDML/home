package controller

import (
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"home/web/internal/service"
	"home/web/utils"
	"image/png"
	"net/http"
)

type UserController struct {
	captchaService *service.CaptchaService
}

func NewUserController(captchaService *service.CaptchaService) *UserController {
	return &UserController{
		captchaService: captchaService,
	}
}

// GetSession 获取用户session
func (c *UserController) GetSession(ctx *gin.Context) {
	// 初始化错误返回的map
	resp := make(map[string]string)

	resp["error"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSONP(http.StatusOK, resp)

}

// GetImageCd 获取验证码图片服务
func (c *UserController) GetImageCd(ctx *gin.Context) {


	// 获取图片的UUID
	uuid := ctx.Param("uuid")
	fmt.Println("uuid = ", uuid)

	 imgBytes, err := c.captchaService.GetCaptchaFromCaseAndData(ctx, uuid)

	if err != nil {
		panic(err)
	}

	 img := &captcha.Image{}

	 json.Unmarshal(imgBytes, img)

	//// 初始化对象
	//cap := captcha.New()
	//// 设置字体
	//cap.SetFont("./web/conf/comic.ttf")
	//// 创建验证码 4 个字符 captcha.NUM 字符模式数字类型
	//// 设置干扰强度
	//cap.SetDisturbance(captcha.MEDIUM)
	////  设置前景色 可以多个 随机替换文字颜色 默认黑色
	//cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	//// 设置背景色 可以多个 随机替换背景色 默认是白色
	//cap.SetBkgColor(color.RGBA{0, 128, 128, 128}, color.RGBA{255, 0, 0, 128})
	//
	//
	//// 生成字体
	//img, str := cap.Create(6, captcha.ALL)
	png.Encode(ctx.Writer, img)


	//fmt.Println("验证码:", str)
	//fmt.Println("验证码uuid:", uuid)

	//ctx.JSONP(http.StatusOK, gin.H{"error":utils.RECODE_SESSIONERR,"errmsg":utils.RecodeText(utils.RECODE_SESSIONERR) })

}