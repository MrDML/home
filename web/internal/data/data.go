package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	v1 "home/api/captcha/service/v1"
	"home/web/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCaptchaClient ,NewCaptchaRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	cc v1.CaptchaClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, captchaClient v1.CaptchaClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{cc: captchaClient}, cleanup, nil
}



func NewCaptchaClient() v1.CaptchaClient {
	connect, errr := grpc.DialInsecure(context.Background() ,grpc.WithEndpoint("127.0.0.1:9001"))
	if errr != nil {
		panic(errr)
	}
	c := v1.NewCaptchaClient(connect)
	return c
}