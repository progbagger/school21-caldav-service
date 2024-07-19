package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"s21calendar/api"
)

const authServerUrl string = "https://auth.sberclass.ru/auth/realms/EduPowerKeycloak/protocol/openid-connect/token"

type CredentialsStorage struct {
	Username string
	Password string
}

func (cs *CredentialsStorage) Authorization(ctx context.Context, operationName string) (result api.Authorization, err error) {
	resp, err := http.PostForm(authServerUrl, url.Values{
		"username":   {cs.Username},
		"password":   {cs.Password},
		"grant_type": {"password"},
		"client_id":  {"s21-open-api"},
	})
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	type GetTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	var getTokenResponse GetTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&getTokenResponse); err != nil {
		return result, err
	}

	result.SetAPIKey(fmt.Sprintf("%s %s", getTokenResponse.TokenType, getTokenResponse.AccessToken))
	return result, nil
}
