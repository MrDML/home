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

	// 调用远程服务
	imgBytes, err := c.captchaService.GetCaptchaFromCaseAndData(ctx, uuid)

	if err != nil {
		panic(err)
	}
	img := &captcha.Image{}
	json.Unmarshal(imgBytes, img)

	// 回写客户端
	png.Encode(ctx.Writer, img)
}

// GetSmscd 获取短信验证码服务
func (c *UserController) GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	mobile := ctx.Param("mobile")
	if mobile == "" {
		ctx.JSONP(http.StatusOK, gin.H{"error": utils.RECODE_PARAMERR, "errmsg": utils.RecodeText(utils.RECODE_PARAMERR)})
	}
	// 短信验证码
	imgCode := ctx.Query("text")
	// 获取图片验证码
	uuid := ctx.Query("id")
	if imgCode == "" || uuid == "" {
		ctx.JSONP(http.StatusOK, gin.H{"error": utils.RECODE_PARAMERR, "errmsg": utils.RecodeText(utils.RECODE_PARAMERR)})
	}

	var imageCode string
	// 调用远程服务 查询redis服务中存的图片验证码code 使用 传入的 id 进行获取如果获取成功那么验证码没有失效，如果没去获取到要么失效了，要么不存在该值
	err := c.captchaService.GetImageCodeFromRdb(ctx, uuid, &imageCode)
	if err != nil {
		panic(err)
	}

	// 调用远程服务发送短信
	c.captchaService.SendSmsCode(ctx, imgCode)

	// 成功
	ctx.JSONP(http.StatusOK, gin.H{"error": utils.RECODE_OK, "errmsg": utils.RecodeText(utils.RECODE_OK)})

}
