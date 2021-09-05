package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	_ "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"home/web/controller"
	"home/web/internal/conf"
	"home/web/internal/service"
	"reflect"
	"time"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, captcha *service.CaptchaService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	//srv := http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)

	router := gin.Default()

	// 使用中间件
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))

	// 加载web
	router.Static("/home", "web/view")

	addRoute(router, captcha)

	httpSrv := http.NewServer(opts...)

	httpSrv.HandlePrefix("/", router)

	return httpSrv

}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}

		fmt.Println("====>",reflect.TypeOf(req))

		reply, err = handler(ctx, req)
		return
	}
}

//  addRoute 添加路由
func addRoute(router *gin.Engine, captchaService *service.CaptchaService) {
	userController := controller.NewUserController(captchaService)
	v1 := router.Group("/api/v1.0")
	{
		//  获取 session 接口
		v1.GET("/session", userController.GetSession)
		// 获取验证码图片服务
		v1.GET("/imagecode/:uuid", userController.GetImageCd)
	}
}






// ============================================

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request


		// 拦截
		//c.Abort()

		// 放行
		c.Next()


		// after request
		latency := time.Since(t)
		fmt.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		fmt.Println(status)


	}
}

//func ginServer() *http.Server{
//
//	router := gin.Default()
//
//	router.Use(Logger())
//	// 使用中间件
//	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
//
//
//	router.Static("/home", "web/view")
//
//
//	v1 := router.Group("/api/v1.0")
//	{
//		//  获取ssession 接口
//		v1.GET("/session", controller.GetSession)
//	}
//
//
//	httpSrv := http.NewServer(http.Address(":8001"))
//	httpSrv.HandlePrefix("/", router)
//
//	return httpSrv
//}