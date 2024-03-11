package routes

import (
	"ChargPiles/api"
	"ChargPiles/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	user := r.Group("chargpiles/user")
	{
		user.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		user.POST("register", api.UserRegisterHandler())
		user.POST("login", api.UserLoginHandler())
		user.POST("verification_code", api.UserVerificationCodeHandler())

		authed := user.Group("/")
		authed.Use(middleware.AuthMiddleware())
		{
			authed.POST("update", api.UserUpdateHandler())
			authed.GET("show_info", api.ShowUserInfoHandler())
			authed.POST("avatar", api.UploadAvatarHandler())
		}
	}

	return r
}
