package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/Spazzy757/laurentia/messages"
)

func TestGetMessagesHandler(t *testing.T) {
	messages.SaveMessage(`{"foo": "bar"}`)
	messages.SaveMessage(`{"foo": "baz"}`)
	r := GetMainEngine()
	url := `/messages?limit=2`
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatal("Response was not 200")
	}
}

//func TestGetSubScriberList(t *testing.T) {
//	client := messages.getClient()
//	lookUp := "pubsub.events.{event}.subscribers"
//}