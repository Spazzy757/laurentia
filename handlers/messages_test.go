package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"../messages"
)

func TestGetMessagesHandler(t *testing.T) {
	messages.SaveMessage(`{"foo": "bar"}`)
	r := GetMainEngine()
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	r.ServeHTTP(resp, req)
	t.Log(resp.Code)
}