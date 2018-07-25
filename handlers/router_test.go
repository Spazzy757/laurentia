package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
)

type Token struct {
	Token string `json:"token"`
	Expire string `json:"expire"`
	Code int `json:"code"`
}


func getLoginPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "admin")
	params.Add("password", "admin")

	return params.Encode()
}

func TestAuthMiddleware(t *testing.T) {
	var tokenSet *Token
	t.Run(`Test Logging In`, func (t *testing.T) {
		r := SetupRouter()
		registrationPayload := getLoginPOSTPayload()
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", `/login`, strings.NewReader(registrationPayload))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.ServeHTTP(resp, req)
		err := json.NewDecoder(resp.Body).Decode(&tokenSet)
		if err != nil {t.Fatal(err)}
		if resp.Code != 200 {t.Fatal("Response was not 200")}
	})
	t.Run(`Test Refresh Token`, func (t *testing.T) {
		r := SetupRouter()
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", `/auth/refresh_token`, nil)
		req.Header.Add("Authorization", "Bearer " + tokenSet.Token)
		req.Header.Add("Content-Type", "application/json")
		r.ServeHTTP(resp, req)
		if resp.Code != 200 {t.Fatal("Response was not 200")}
	})
}
