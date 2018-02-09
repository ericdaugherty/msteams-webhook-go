package teams

import (
	"encoding/json"
	"strings"
	"testing"
)

const requestString = `{
    "type": "message",
    "id": "1485983408511",
    "timestamp": "2017-02-01T21:10:07.437Z",
    "localTimestamp": "2017-02-01T14:10:07.437-07:00",
    "serviceUrl": "https://smba.trafficmanager.net/amer-client-ss.msg/",
    "channelId": "msteams", 
    "from": {
        "id": "29:1XJKJMvc5GBtc2JwZq0oj8tHZmzrQgFmB39ATiQWA85gQtHieVkHilBZ9XHoq9j7Zaqt7CZ-NJWi7me2kHTL3Bw",
        "name": "Tim Jones"
    },
    "conversation": {
        "id": "19:253b1f341670408fb6fe51050b6e5ceb@thread.skype;messageid=1485983194839"
    },
    "recipient": {
        "id": "null", 
        "name": "null"
    },
    "textFormat": "plain",
    "text": "<at>MyCustomBot</at> Hello <at>Larry Brown</at>",
    "attachments": [{
        "contentType": "text/html",
        "content": "<div><span contenteditable=\"false\" itemscope=\"\" itemtype=\"http://schema.skype.com/Mention\" itemid=\"0\">MyWebHook </span> Hello <span contenteditable=\"false\" itemscope=\"\" itemtype=\"http://schema.skype.com/Mention\" itemid=\"1\">Larry Brown </span></div>"
    }],
    "entities": [{
        "type": "mention",
        "mentioned": {
            "id": "28:c9e8c047-2a74-40a2-b28a-b162d5f5327c", 
            "name": "MyWebHook"
        },
        "text": "<at>MyWebHook</at>"
    }, {
        "type": "mention",
        "mentioned": {
            "id": "29:1jnFbZYs0qXMLH-O4S9-sDLNc3NVEIMWMnC-q0tVdEa-8BRosfojI35QdNoB-yW8iutWLJzHUm_mqEZSSU8si0Q",
            "name": "Larry Brown"
        },
        "text": "<at>Larry Brown</at>"
    }, {
        "type": "clientInfo",
        "locale": "en-US",
        "country": "US",
        "platform": "Windows"
    }],
    "channelData": {
        "teamsChannelId": "19:253b1f341670408fb6fe51050b6e5ceb@thread.skype",
        "teamsTeamId": "19:712c61d0ef384e5fa681ba90ca943398@thread.skype"
    } 
}`

func TestParseRequest(t *testing.T) {
	var r = &Request{}
	err := json.NewDecoder(strings.NewReader(requestString)).Decode(&r)
	if err != nil {
		t.Errorf("Error decoding JSON %s", err)
	}

	if r.Type != "message" {
		t.Errorf("Type should be message but was %s", r.Type)
	}
	if r.ServiceURL != "https://smba.trafficmanager.net/amer-client-ss.msg/" {
		t.Errorf("ServiceURL should be https://smba.trafficmanager.net/amer-client-ss.msg/ but was %s", r.ServiceURL)
	}
	if r.FromUser.Name != "Tim Jones" {
		t.Errorf("FromUser.Name should be Tim Jones but was %s", r.FromUser.Name)
	}
	if r.Conversation.ID != "19:253b1f341670408fb6fe51050b6e5ceb@thread.skype;messageid=1485983194839" {
		t.Errorf("Converation.ID should be 19:253b1f341670408fb6fe51050b6e5ceb@thread.skype;messageid=1485983194839 but was %s", r.Conversation.ID)
	}
	if r.Text != "<at>MyCustomBot</at> Hello <at>Larry Brown</at>" {
		t.Errorf("Text should be <at>MyCustomBot</at> Hello <at>Larry Brown</at> but was %s", r.Text)
	}
	if r.Attachments[0].ContentType != "text/html" {
		t.Errorf("Attachement[0].ContentType should be text/html but was %s", r.Attachments[0].ContentType)
	}
}
