package messages


import (
	"testing"
)

func TestPubSub(t *testing.T) {

	// Message Channel
	messageChan := make(chan string, 1)
	defer close(messageChan)

	// Channel to confirm message channel
	done := make(chan bool)
	defer close(done)

	t.Run("Post Message on Pub Sub and Return it on Channel", func(t *testing.T) {
		var value string
		// Expected response on Channel
		expected := `{"foo": "bar"}`

		helper := false
		counter := 0

		go PubSubListener(messageChan)
		// Function to help handle go routine
		go func() {
			value = <-messageChan // Block, waiting
			helper = true
		}()
		for helper {
			err := PubSubSendMessage(expected)
			if err != nil {t.Fatal(err)}
			//
			if counter > 30 {t.Fatal("Message Timeout")}
			counter++
		}
	})
	t.Run("Get SMEMBERS List", func(t *testing.T) {
		client := GetClient()
		lookUp := "pubsub.events.{event}.subscribers"
		client.SAdd(lookUp, "Hello")
		client.SAdd(lookUp, "World")
		subscriberList := GetSMembers(lookUp)
		if len(subscriberList) != 2 {t.Fail()}
	})
	t.Run("Check Acknowledgment", func(t *testing.T) {
		client := GetClient()
		messageID := "12345"
		event := "clean.up"
		subscriber := "order.service"
		subscriberLookUp := "pubsub.events.order." + event + ".subscribers"
		client.SAdd(subscriberLookUp, subscriber)
		client.SAdd(lookUp, "World")
	})

}