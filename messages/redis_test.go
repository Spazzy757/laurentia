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
		message := `{'key': 'testing.ack', 'id': '1234567890', 'payload': {'service': 'appointments'}}`
		messageID := "1234567890"
		event := "testing.ack"
		subscriber := "order.service"
		subscriberTwo := "order.fulfilment"
		subscriberLookUp := "pubsub.events.order." + event + ".subscribers"
		ackLookUp := "pubsub.events.actions." + event + "." + messageID + ".received"
		client.SAdd(subscriberLookUp, subscriber)
		client.SAdd(subscriberLookUp, subscriberTwo)
		client.SAdd(ackLookUp, subscriber)
		VerifyMessageAndNotify(message)
	})

}