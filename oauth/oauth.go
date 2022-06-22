package oauth

import "github.com/gin-gonic/gin"

func NewOauthRouter(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/oauth")
}
