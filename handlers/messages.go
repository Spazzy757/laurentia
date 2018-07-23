package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Spazzy757/laurentia/messages"
	"strconv"
	"time"
)

type JSONString string

type DynamicMessage struct {
	Key string `json:"key"`
	Timestamp time.Time `json:"timestamp"`
	ID  string `json:"id"`
	Payload interface{} `json:"payload"`
}

func GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}

func GetMessagesHandler(c *gin.Context) {
	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {limit = 10}
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	if err != nil {page = 0}
	messageList, _ := messages.GetMessageList(limit, page)
	var messageJson []DynamicMessage
	for i := 0; i < len(messageList); i++ {
		message := messageList[i]
		m := DynamicMessage{message.Key, message.Timestamp,
		message.ID, message.Payload}
		messageJson = append(messageJson, m)
	}
	c.JSON(http.StatusOK, gin.H{
		"count": len(messageJson),
		"messages": messageJson,
	})
}

func GetSubscriberList(c *gin.Context) {
	event := c.Request.URL.Query().Get("event")
	lookUp := "pubsub.events." + event + ".subscribers"
	subscriberList := messages.GetSMembers(lookUp)
	c.JSON(http.StatusOK, gin.H{
		"subscribers": subscriberList,
	})
}

func GetAcknowledgedSubscribers(c *gin.Context)  {
	event := c.Request.URL.Query().Get("event")
	messageID := c.Query("messageID")
	lookUp := "pubsub.events.actions." + event + "." + messageID + ".received"
	acknowledgedList := messages.GetSMembers(lookUp)
	c.JSON(http.StatusOK, gin.H{
		"event": event,
		"eventID": messageID,
		"acknowledged": acknowledgedList,
	})
}


//var clientList = make(map[ClientConn]int)
//var clientListRWMutex sync.RWMutex
//var dataChannel = make(chan string)
//var mongoChannel = make(chan string)
//var ack = make(chan string)
//
//type ClientConn struct {
//	uuid      string
//	websocket *websocket.Conn
//	ip        net.Addr
//}
//
//func addClient(clientconnection ClientConn) {
//	clientListRWMutex.Lock()
//	clientList[clientconnection] = 0
//	clientListRWMutex.Unlock()
//	//sendMessage(clientconnection, []byte(clientconnection.uuid))
//}
//func removeClient(clientconnection ClientConn) {
//	clientListRWMutex.Lock()
//	delete(clientList, clientconnection)
//	clientListRWMutex.Unlock()
//}
//func sendMessage(clientconnection ClientConn, message []byte) {
//	clientconnection.websocket.WriteMessage(1, message)
//}

//
//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//func WSHandler(w http.ResponseWriter, r *http.Request, dataChannel chan string, mongoChannel chan string) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		fmt.Printf("Failed to set websocket upgrade: %v", err)
//		return
//	}
//	defer conn.Close()
//	clientUUID := uuid.New()
//	client := conn.RemoteAddr()
//	socketClient := ClientConn{clientUUID.String(), conn, client}
//	addClient(socketClient)
//
//	for {
//		msg := <- dataChannel
//		msg = strings.Replace(msg, "'", "\"", -1)
//		broadcastMessage(1, []byte(msg))
//		fmt.Println("Message Received")
//		mongoChannel <- msg
//	}
//}
//
//func broadcastMessage(messageType int, message []byte) {
//	for client := range clientList {
//		err := client.websocket.WriteMessage(messageType, message)
//		if err != nil {
//			log.Println("Failed to send message to client, " + client.ip.String())
//			removeClient(client)
//		}
//	}
//}