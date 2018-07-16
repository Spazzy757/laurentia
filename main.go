package main

import (
	"github.com/Spazzy757/laurentia/handlers"
)
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




func main() {
	//go PubSubListener("default", dataChannel)
	//go MongoStore(mongoChannel)
	handlers.GetMainEngine().Run("0.0.0.0:8000")
}

//func PubSubListener(channel string, pubSubChannel chan string) {
//	redisHost := os.Getenv("REDIS_HOST") + ":6379"
//	redisPassword := os.Getenv("REDIS_PASSWORD")
//	client := redis.NewClient(&redis.Options{
//		Addr:     redisHost,
//		Password: redisPassword,
//		DB:       1,
//	})
//	pubSub := client.Subscribe(channel)
//	defer pubSub.Close()
//	for {
//		msg, err := pubSub.ReceiveMessage()
//		if err != nil {
//			log.Println("Pub Sub Failed to Receive Message")
//		}
//		pubSubChannel <- msg.Payload
//		time.Sleep(1)
//	}
//}
//
//func MongoStore(dataChannel chan string) {
//	session, err := mgo.Dial("localhost")
//	if err != nil {
//		log.Fatal(err)
//	}
//	c := session.DB("channelDump").C("messages")
//	for {
//		msg := <-dataChannel
//		var v map[string]interface{}
//		data := strings.Replace(msg, "'", "\"", -1)
//		type Message struct {
//			Msg       string
//			Timestamp time.Time
//		}
//		message := &Message{Msg:data, Timestamp: time.Now()}
//		b, err := json.Marshal(message)
//		if err != nil {
//			panic(err)
//		}
//		if err := json.Unmarshal(b, &v); err != nil {
//			panic(err)
//		}
//		err = c.Insert(&v)
//		if err != nil {
//			panic(err)
//		}
//	}
//
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