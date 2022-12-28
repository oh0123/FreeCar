// Code generated by hertz generator.

package Api

import (
	"context"
	"fmt"

	"github.com/CyanAsterisk/FreeCar/server/cmd/api/biz/errno"
	"github.com/CyanAsterisk/FreeCar/shared/middleware"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createcarMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _getcarMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _submitprofileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _clearprofileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _tripMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _updatetripMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _gettripMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _gettripsMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.Recovery(),
		middleware.JWTAuth(),
	}
}

func _authMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _profileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use recovery mw
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.BadRequest,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		middleware.JWTAuth(),
	}
}

func _createprofilephotoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _clearprofilephotoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _photoMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		// use recovery mw
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.BadRequest,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		middleware.JWTAuth(),
	}
}

func _completeprofilephotoMw() []app.HandlerFunc {
	// your code...
	return nil
}
