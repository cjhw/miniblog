package miniblog

import (
	"github.com/cjhw/miniblog/internal/miniblog/controller/v1/user"
	"github.com/cjhw/miniblog/internal/miniblog/store"
	"github.com/cjhw/miniblog/internal/pkg/core"
	"github.com/cjhw/miniblog/internal/pkg/errno"
	"github.com/cjhw/miniblog/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
		}
	}

	return nil
}
