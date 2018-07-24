package handlers
import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gorilla/websocket"
	"strings"
	"encoding/json"
	"sync"
)

func TestMessageWSHandler(t *testing.T) {
	var webSocketWaitGroup sync.WaitGroup
	webSocketWaitGroup.Add(10)
	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(MessageWSHandler))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	//var testMessageList []DynamicMessage
	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {t.Fatalf("%v", err)}
	defer ws.Close()

	var messageList []DynamicMessage
	// Send message to server, read response and check to see if it's what we expect.
	go func(w *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			var m *DynamicMessage
			_, p, err := ws.ReadMessage()
			if err != nil {t.Fatalf("%v", err)}
			if err := json.Unmarshal([]byte(p), &m); err != nil {t.Log(err)}
			messageList = append(messageList, *m)
			w.Done()
		}
	}(&webSocketWaitGroup)
	for i := 0; i < 10; i ++ {
		AddMessageToChannel(`{"key": "1234", "id": "12334", "payload": {}}`)
	}
	webSocketWaitGroup.Wait()
	if len(messageList) != 10 {
		t.Fatal("Web Sockets only recieved did not receieve all messages")
	}
}

