package main

import (
	"five/oauth"
	oauthutils "five/oauth/utils"
	oauthhandler "five/oauth/v1/delivery/http"
	"five/oauth/v1/repository"
	"five/oauth/v1/usecase"
	"five/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health-check", func(context *gin.Context) {
		context.JSON(http.StatusOK, utils.SuccessResponse("ok"))
	})

	oauthRouter := oauth.NewOauthRouter(r)
	config := utils.InitConfig()
	oauthConfig := oauthutils.InitConfig(config)
	client := &http.Client{}
	oauthRepository := repository.NewOauthRepository(oauthConfig, client)
	oauthUseCase := usecase.NewOauthUseCase(oauthRepository)
	oauthhandler.NewOauthHandler(oauthRouter, oauthUseCase)
	r.Run(fmt.Sprintf(":%s", config.Port))
}
