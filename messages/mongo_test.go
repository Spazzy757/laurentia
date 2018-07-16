package messages

import (
	"testing"
	"time"
)

func TestMessages(t *testing.T) {
	t.Run("Test Saving Message To Mongo", func(t *testing.T) {
		//Test Creating Messages
		_, err := SaveMessage(`{"foo": "bar"}`)
		if err != nil {t.Fail()}
	})
	t.Run("Test Getting Results with a limit of two", func(t *testing.T) {
		// Test Messages Passed to Function
		messages := []string{`{"foo": "baz"}`, `{"foo": "maya"}`, `{"foo": "binte"}`}
		// Create Messages before getting Messages
		for i := 0; i < len(messages); i++ {
			_, err := SaveMessage(messages[i])
			if err != nil {t.Fail()}
			time.Sleep(1)
		}
		resp, err := GetMessageList(2, 0)
		//Check for errors
		if err != nil {t.Fatal(err)}
		//Check the amount returned is correct (the limit)
		if len(resp) != 2 {t.Fatal("Pagination Failure")}
		//Remove all messages in Mongo
		for i := 0; i < len(resp); i++ {DeleteMessage(resp[i].ID)}
	})
}