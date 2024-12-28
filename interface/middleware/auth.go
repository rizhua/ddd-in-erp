package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"rizhua.com/infrastructure/adapter"
)

// 授权中间件:解析token
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 验证token
		token := ctx.GetHeader("Access-Token")
		if token != "" {
			cache := adapter.NewCache()
			_, err := cache.Get("token:" + token)
			if err != nil {
				ctx.Abort()
				ctx.JSON(http.StatusOK, gin.H{"code": 3071, "desc": "非法 token"})
				return
			}
			// 刷新登录时间
			expire := cache.TTL("token:" + token)
			if expire.Hours() < 72 {
				cache.Expire("token:"+token, 720*time.Hour)
			}
		} else {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{"code": 3070, "desc": "请携带 token"})
			return
		}

		ctx.Next()
	}
}
