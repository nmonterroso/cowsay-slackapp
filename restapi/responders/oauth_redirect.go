package responders

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/nmonterroso/cowsay-slackapp/models"
	"github.com/nmonterroso/cowsay-slackapp/restapi/operations"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"os"
)

const (
	clientIdEnv     = "CLIENT_ID"
	clientSecretEnv = "CLIENT_SECRET"
	redirectUriEnv  = "REDIRECT_URI"
)

var (
	missingClientCredentials = errors.New("missing_oauth_client_credentials")
	missingRedirectUri       = errors.New("missing_redirect_uri")
	missingCode              = errors.New("missing_oauth_code")
	clientId, clientSecret   = slackClientCredentials()
	redirectUri              = slackRedirectUri()
)

type SlackAccessTokenResponse struct {
	Ok          bool   `json:"ok"`
	Error       string `json:"error"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
}

func OauthRedirectResponder(params operations.OauthRedirectParams) middleware.Responder {
	if params.Error != "" {
		return &operations.OauthRedirectNoContent{}
	} else if params.Code == "" {
		return &operations.OauthRedirectDefault{
			&models.Error{
				Code:    http.StatusInternalServerError,
				Message: missingCode.Error(),
			},
		}
	}

	_, err := slackAccessToken(params.Code)

	if err != nil {
		return &operations.OauthRedirectDefault{
			&models.Error{
				Code:    http.StatusInternalServerError,
				Message: missingCode.Error(),
			},
		}
	}

	return &operations.OauthRedirectNoContent{}
}

func slackClientCredentials() (string, string) {
	id := os.Getenv(clientIdEnv)
	secret := os.Getenv(clientSecretEnv)

	if id == "" || secret == "" {
		panic(missingClientCredentials)
	}

	return id, secret
}

func slackRedirectUri() string {
	uri := os.Getenv(redirectUriEnv)

	if uri == "" {
		panic(missingRedirectUri)
	}

	return uri
}

func slackAccessToken(code string) (*SlackAccessTokenResponse, error) {
	resp, body, errs := gorequest.New().Get(slackOauthAccessEndpoint(code)).EndBytes()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("slack responded with %d", resp.StatusCode))
	}

	var response SlackAccessTokenResponse
	err := json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	if !response.Ok || response.Error != "" {
		return nil, response.Error
	}

	return response, nil
}

func slackOauthAccessEndpoint(code string) string {
	return fmt.Sprintf(
		"https://slack.com/api/oauth.access?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		clientId,
		clientSecret,
		code,
		redirectUri)
}
