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

//
// GetSession 获取用户session
//  @Description:
//  @receiver c
//  @param ctx
//
func (c *UserController) GetSession(ctx *gin.Context) {


      index :=  []int{1,2}

      i := index[3]

      fmt.Println("======> i", i)


	ctx.JSON(http.StatusOK, utils.OK(""))
}


//
// GetImageCd 获取验证码图片服务
//  @Description:
//  @receiver c
//  @param ctx
//
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

//
// GetSmscd
//  @Description: 获取短信验证码服务
//  @receiver c
//  @param ctx
//
func (c *UserController) GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	mobile := ctx.Param("mobile")
	if mobile == "" {
		ctx.JSON(http.StatusOK, utils.Fail(utils.RECODE_PARAMERR))

	}
	// 短信验证码
	imgCode := ctx.Query("text")
	// 获取图片验证码
	uuid := ctx.Query("id")
	if imgCode == "" || uuid == "" {
		ctx.JSON(http.StatusOK, utils.Fail(utils.RECODE_PARAMERR))
	}

	var imageCode string
	// 调用远程服务 查询redis服务中存的图片验证码code 使用 传入的 id 进行获取如果获取成功那么验证码没有失效，如果没去获取到要么失效了，要么不存在该值
	err := c.captchaService.GetImageCodeFromRedis(ctx, uuid, &imageCode)
	if err != nil {
		panic(err)
	}

	// 调用远程服务发送短信
	c.captchaService.SendSmsCode(ctx, imgCode)

	// 成功
	ctx.JSON(http.StatusOK, utils.OK(nil))

}
