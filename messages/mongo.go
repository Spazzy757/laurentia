package messages

import (
	"gopkg.in/mgo.v2"
	"time"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type Message struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time `bson:"timestamp"`
	Message string `bson:"result"`
}

var host  = os.Getenv("MONGO_HOST")
const DB = "laurentia"
const COLLECTION = "messages"


func SaveMessage(message string) (bool, error){
	m := &Message{Timestamp: time.Now(), Message: message}
	session, err := mgo.Dial(host)
	if err != nil { return false, err}
	c := session.DB(DB).C(COLLECTION)
	err = c.Insert(&m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetMessageList (limit int, page int) ([]Message, error) {
	var results []Message
	session, err := mgo.Dial(host)
	c := session.DB(DB).C(COLLECTION)
	err = c.Find(bson.M{}).Skip(limit * page).Sort("-timestamp").Limit(limit).All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func DeleteMessage (objectID bson.ObjectId) (error) {
	session, err := mgo.Dial(host)
	if err != nil {return err}
	c := session.DB(DB).C(COLLECTION)
	err = c.RemoveId(objectID)
	if err != nil {return err}
	return nil
}
