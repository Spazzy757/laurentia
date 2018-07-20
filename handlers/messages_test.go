package handlers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"gopkg.in/mgo.v2"
	"github.com/Spazzy757/laurentia/messages"
	"log"
	"os"
)
const MESSAGE = `{'key': 'health.check', 'id': '0396660a-e111-4cbb-9abc-f9db68574480', 'payload': {'workorder_type': 'install', 'id': 133446, 'backend_references': [{'id': 197116, 'backend_name': 'quickbase', 'backend_field_name': 'online order id', 'backend_id': '216536'}, {'id': 197078, 'backend_name': 'quickbase', 'backend_field_name': 'partial order id', 'backend_id': '11796'}], 'status_changes': [{'id': 26697, 'old_status': 'ORDER_UPDATED', 'new_status': 'IN_PROGRESS', 'date': '2018-07-18T11:58:35.087023+02:00', 'notes': 'Installation - Service'}, {'id': 26646, 'old_status': 'ORDER_CREATED', 'new_status': 'ORDER_UPDATED', 'date': '2018-07-18T11:26:35.438613+02:00', 'notes': None}, {'id': 26641, 'old_status': 'NEW', 'new_status': 'ORDER_CREATED', 'date': '2018-07-18T11:25:51.470154+02:00', 'notes': 'Order Created in Quickbase'}], 'first_name': 'Geraldine', 'last_name': 'Rajagopal', 'cell_number': '0824144955', 'email': 'geraldiner@relatirgroup.co.za', 'identity_number': '8208280007080', 'za_resident': True, 'type': 'Residential', 'business_name': None, 'vatnumber': None, 'businessregistrationnumber': None, 'alternativecontactname': None, 'alternatecontactnumber': None, 'alternatecontactemail': None, 'promocode': 'HD77', 'mduname': 'Centenary Park', 'mduunitno': '29', 'mdurid': '', 'street_number': 'N/A', 'street_name': '2', 'fibrehoodonorder': 'honeydew_manor', 'city': 'Johannesburg', 'installationstreetaddress': '29 Centenary Park, 2, Honeydew Manor', 'postal_code': '2170', 'lat': '-26.09656380000000000000', 'long': '27.90073299999994600000', 'province': 'Gauteng', 'country': 'South Africa', 'location_id': '320665', 'orderchannel': 'Online (Website)', 'getupdates': False, 'networkname': 'vumaerial', 'networktype': 'GPON', 'orderreference': 'VA-18718-225186', 'objectgroupid': 'default', 'objectgroupname': 'default', 'relatedpremiseid_ifapplicable': 'NA', 'serviceproviderselected': 'Home Connect', 'serviceproviderid': '64D629C9-11C9-4DE2-9BF4-B70D30B2ECBA', 'serviceselected': '20/2 FibreTrend Ignite', 'ispserviceid': '5AE1E5EA-24B7-4FC1-ABD3-B426EB7EA9DC', 'ordercomplete': True, 'datecreated': '2018-07-18T11:25:48.127309+02:00', 'datemodified': '2018-07-18T11:58:35.011766+02:00', 'recordid': '11796', 'recordowner': None, 'lastmodifiedby': None, 'processed': True, 'transactionref': '987b2d0eb88762093878f66670a99c75674de0995cf629b601a2262e93f31596', 'ordertype': 'order', 'order_already_placed': False, 'order_comments': None, 'url': None, 'location_reference': '2489880210', 'appointment_date': None, 'isp_order_reference': None, 'organization_id': None, 'user_id': None, 'estate_name': None, 'status': 'ORDER_UPDATED'}}`
const SECONDMESSAGE = `{'key': 'health.check', 'id': '0396870a-e111-4cbb-9abc-f9db68574480', 'payload': {'service': 'appointments'}}`
var host  = os.Getenv("MONGO_HOST")
const DB = "laurentia"
const COLLECTION = "messages"

func removeAllMessages() {
	s, err := mgo.Dial(host)
	c := s.DB(DB).C(COLLECTION)
	if err != nil { log.Fatal(`Error Connecting to Mongo DB`)}
	c.RemoveAll(nil)
}

func TestGetMessagesHandler(t *testing.T) {
	t.Run(`Test Geting Messages With A Limit of Two`, func (t *testing.T) {
		messages.SaveMessage(MESSAGE)
		messages.SaveMessage(SECONDMESSAGE)
		r := SetupRouter()
		url := `/v1/messages?limit=2`
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		r.ServeHTTP(resp, req)
		if resp.Code != 200 {t.Fatal("Response was not 200")}
	})
	removeAllMessages()
}

func TestGetSubScriberList(t *testing.T) {
	t.Run(`Get Subscribers of Messages`, func (t *testing.T) {
		client := messages.GetClient()
		lookUp := "pubsub.events.order.created.subscribers"
		client.SAdd(lookUp, "Hello")
		client.SAdd(lookUp, "World")
		r := SetupRouter()
		url := `/v1/subscribers?event=order.created`
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		r.ServeHTTP(resp, req)
		if resp.Code != 200 {t.Fatal("Response was not 200")}
	})
}

func TestGetAcknowledgedSubscribers(t *testing.T) {
	t.Run(`Test Getting Acknowledgment fo a Message`, func (t *testing.T) {
		client := messages.GetClient()
		lookUp := "pubsub.events.actions.order.12345.received"
		client.SAdd(lookUp, "subscriber-1.received")
		r := SetupRouter()
		url := `/v1/acknowledged?event=order&messageID=12345`
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		r.ServeHTTP(resp, req)
		if resp.Code != 200 {t.Fatal("Response was not 200")}
	})
}

func TestGetHealth(t *testing.T) {
	t.Run(`Test Get Health Endpoint`, func (t *testing.T) {
		r := SetupRouter()
		url := `/health`
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", url, nil)
		r.ServeHTTP(resp, req)
		if resp.Code != 200 {t.Fatal("Response was not 200")}
	})
}
