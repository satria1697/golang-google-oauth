package http

import (
	"errors"
	"five/oauth/v1/domain"
	"five/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type OauthHandler struct {
	r            *gin.RouterGroup
	oauthUseCase domain.OauthUseCase
}

func NewOauthHandler(r *gin.RouterGroup, oauthUseCase domain.OauthUseCase) {
	handler := OauthHandler{
		r:            r,
		oauthUseCase: oauthUseCase,
	}
	r.GET("/link", handler.GetOauthLink)
	r.GET("/token", handler.GetOauthToken)
	r.GET("/info", handler.GetOauthInfo)
}

func (h OauthHandler) GetOauthToken(c *gin.Context) {
	var query domain.OauthTokenQuery
	err := c.Bind(&query)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	res, err := h.oauthUseCase.GetTokenUseCase(c, query.Code)
	if err != nil {
		e := err
		if strings.Contains(err.Error(), "cannot fetch token") {
			e = errors.New("token-fail")
		}
		c.JSON(http.StatusForbidden, utils.ErrorResponse(e))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(&gin.H{
		"token": res,
	}))
}

func (h OauthHandler) GetOauthInfo(c *gin.Context) {
	var request domain.OauthTokenUserInfo
	c.Bind(&request)
	res, err := h.oauthUseCase.GetInfoUseCase(request.Token)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	if res.Name == "" {
		err = errors.New("not-found")
		c.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(res))
}

func (h OauthHandler) GetOauthLink(c *gin.Context) {
	res := h.oauthUseCase.GetLinkUseCase()
	c.JSON(http.StatusOK, utils.SuccessResponse(&gin.H{
		"url": res,
	}))
}
