package exception

import (
	"context"
	"fmt"
	"github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"home/web/utils"
	"net/http"
)

type s struct {}

func Exception(ctx context.Context, req, err interface{}) error {
	// 两者的实际类型不一样，但是地址相同
	//cb := context.WithValue(context.Background(), s{}, 1)
	//fmt.Printf("---->%v\n", reflect.TypeOf(s{})) // ---->exception.s
	//fmt.Printf("---->%v \n",reflect.TypeOf(struct{}{})) // ---->struct {}
	//val := cb.Value(struct{}{}).(int)
	//
	//fmt.Println(val)
	fmt.Println("===捕获到全局的异常===")
	// FromGinContext
	c, _ := gin.FromGinContext(ctx)
	fmt.Println("c:", c)
	c.JSON(http.StatusOK, utils.OK(nil))
	return errors.InternalServer("RECOVERY", fmt.Sprintf("panic triggered: %v", err))
}
