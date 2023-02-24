package middleware

import (
	"gateway/models"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取头信息，是否包含token
		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请登录",
			})
			c.Abort()
			return
		}

		// 解析token，判断是否正常
		j := public.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == public.TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
			return
		}

		// 3. 查询用户是否存在
		// 根据ID查询用户信息
		admin := models.Admin{}
		user, err := admin.Find(c, public.GormDB, &models.Admin{ID: int(claims.ID)})
		if err != nil {
			public.ResponseError(c, public.MiddleUserNotExist, errors.New("用户不存在"))
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Set("user", user)
		c.Next()
	}
}
