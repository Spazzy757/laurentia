package messages

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"os"
	"encoding/json"
	"log"
	//"errors"
	"errors"
)

type Message struct {
	MongoID bson.ObjectId `bson:"_id,omitempty" `
	Key string `bson:"key"`
	ID string `bson:"id"`
	Timestamp time.Time `bson:"timestamp"`
	Payload  interface{} `json:"payload"`
}

var host  = os.Getenv("MONGO_HOST")
const DB = "laurentia"
const COLLECTION = "messages"

func getCollection () *mgo.Collection {
	s, err := mgo.Dial(host)
	c := s.DB(DB).C(COLLECTION)
	if err != nil { log.Fatal(`Error Connecting to Mongo DB`)}
	return c
}

func SaveMessage (message string) (bool, error){
	var m *Message
	if err := json.Unmarshal([]byte(formatPythonDict(message)), &m); err != nil {
		return false, err
	}
	log.Println(`*************************************************`)
	if m.ID == "" {
		return false, errors.New(`invalid input`)
	}
	c := getCollection()
	err := c.Insert(&m)
	if err != nil {return false, err}
	log.Println(m)
	log.Println(`*************************************************`)
	return true, nil
}

func GetMessageList (limit int, page int) ([]Message, error) {
	var results []Message
	c := getCollection()
	err := c.Find(bson.M{}).Skip(limit * page).Sort("-timestamp").Limit(limit).All(&results)
	if err != nil {return nil, err}
	return results, nil
}

func GetMessageByID (id string) (Message, error) {
	var results Message
	c := getCollection()
	err := c.Find(bson.M{"id": id}).One(&results)
	if err != nil {return Message{}, err}
	return results, nil
}

func DeleteMessage (id string) error{
	if id != "" {
		c := getCollection()
		err := c.Remove(bson.M{"id": id})
		if err != nil {return err}
		return nil
	}
	return nil
}
