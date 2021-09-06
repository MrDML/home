package exception

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
)

func Exception(ctx context.Context, req, err interface{}) error {
	fmt.Println("===捕获到全局的异常===")
	return errors.InternalServer("RECOVERY", fmt.Sprintf("panic triggered: %v", err))
}
