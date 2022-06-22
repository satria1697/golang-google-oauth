package repository

import (
	"encoding/json"
	"five/oauth/v1/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

type oauthRepository struct {
	config *oauth2.Config
	client *http.Client
}

func (o oauthRepository) GetLinkRepository() string {

	url := o.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url
}

func (o oauthRepository) GetTokenRepository(c *gin.Context, code string) (string, error) {
	response, err := o.config.Exchange(c, code)
	if err != nil {
		return "", err
	}
	return response.AccessToken, nil
}

func (o oauthRepository) GetInfoRepository(token string) (domain.OauthUserInfo, error) {
	var userInfo domain.OauthUserInfo
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return userInfo, err
	}
	res, err := o.client.Do(req)
	if err != nil {
		return userInfo, err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return userInfo, err
	}
	json.Unmarshal(bytes, &userInfo)
	return userInfo, nil
}

func NewOauthRepository(config *oauth2.Config, client *http.Client) domain.OauthRepository {
	return &oauthRepository{
		config: config,
		client: client,
	}
}
