package usecase

import (
	"five/oauth/v1/domain"
	"github.com/gin-gonic/gin"
)

type oauthUseCase struct {
	oauthRepository domain.OauthRepository
}

func (o oauthUseCase) GetLinkUseCase() string {
	return o.oauthRepository.GetLinkRepository()
}

func (o oauthUseCase) GetTokenUseCase(c *gin.Context, code string) (string, error) {
	return o.oauthRepository.GetTokenRepository(c, code)
}

func (o oauthUseCase) GetInfoUseCase(token string) (domain.OauthUserInfo, error) {
	return o.oauthRepository.GetInfoRepository(token)
}

func NewOauthUseCase(oauthRepository domain.OauthRepository) domain.OauthUseCase {
	return oauthUseCase{
		oauthRepository: oauthRepository,
	}
}
