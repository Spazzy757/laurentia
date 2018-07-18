package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/Spazzy757/laurentia/messages"
)

func TestGetMessagesHandler(t *testing.T) {
	messages.SaveMessage(`{'foo'': 'bar'}`)
	messages.SaveMessage(`{'foo'': 'baz'}`)
	r := SetupRouter()
	url := `/v1/messages?limit=2`
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(resp, req)
	t.Log(resp.Body)
	if resp.Code != 200 {
		t.Fatal("Response was not 200")
	}
}

func TestGetSubScriberList(t *testing.T) {
	client := messages.GetClient()
	lookUp := "pubsub.events.order.created.subscribers"
	client.SAdd(lookUp, "Hello")
	client.SAdd(lookUp, "World")
	r := SetupRouter()
	url := `/v1/subscribers?event=order.created`
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(resp, req)
	t.Log(resp.Body)
	if resp.Code != 200 {
		t.Fatal("Response was not 200")
	}
}