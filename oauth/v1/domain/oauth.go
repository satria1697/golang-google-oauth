package domain

import "github.com/gin-gonic/gin"

type OauthTokenQuery struct {
	Code string `form:"code"`
}

type OauthTokenUserInfo struct {
	Token string `form:"token"`
}

type OauthUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

type OauthUseCase interface {
	GetLinkUseCase() string
	GetTokenUseCase(c *gin.Context, code string) (string, error)
	GetInfoUseCase(token string) (OauthUserInfo, error)
}

type OauthRepository interface {
	GetLinkRepository() string
	GetTokenRepository(c *gin.Context, code string) (string, error)
	GetInfoRepository(token string) (OauthUserInfo, error)
}
