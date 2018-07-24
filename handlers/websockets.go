package handlers

import (
	"net/http"
	"log"
	"sync"
	"net"
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
	"encoding/json"
	"strings"
	"time"
)

var clientList = make(map[ClientConn]int)
var clientListRWMutex sync.RWMutex
var MessageChannel = make(chan DynamicMessage)
type ClientConn struct {
	uuid      string
	websocket *websocket.Conn
	ip        net.Addr
}

func addClient(clientConnection ClientConn) {
	clientListRWMutex.Lock()
	clientList[clientConnection] = 0
	clientListRWMutex.Unlock()
}
func removeClient(clientConnection ClientConn) {
	clientListRWMutex.Lock()
	delete(clientList, clientConnection)
	clientListRWMutex.Unlock()
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MessageWSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %v", err)
		return
	}
	defer conn.Close()
	clientUUID := uuid.New()
	client := conn.RemoteAddr()
	socketClient := ClientConn{clientUUID.String(), conn, client}
	addClient(socketClient)
	for {
		msg :=  <-MessageChannel
		broadcastMessage(&msg)
	}
}

func broadcastMessage(message *DynamicMessage) {
	for client := range clientList {
		err := client.websocket.WriteJSON(message)
		// If there is a error communicating with client remove them from list
		if err != nil {removeClient(client)}
	}
}


func AddMessageToChannel(msg string) {
	var m *DynamicMessage
	if err := json.Unmarshal([]byte(formatPythonDict(msg)), &m); err != nil {log.Println("Invalid Message")}
	m.Timestamp = time.Now()
	MessageChannel <- *m
}

func formatPythonDict(message string) string {
	message = strings.Replace(message, `None`, `null`, -1)
	message = strings.Replace(message, `True`, `true`, -1)
	message = strings.Replace(message, `False`, `false`, -1)
	message = strings.Replace(message, `'`, `"`, -1)
	return message
}