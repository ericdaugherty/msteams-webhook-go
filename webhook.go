package teams

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

var auth bool
var keyBytes []byte
var webhook WebHook

// WebHook represnts the interface needed to handle Microsoft Teams WebHook Requests.
type WebHook interface {
	OnMessage(Request) (Response, error)
}

// Request data representing an inbound WebHook request from Microsoft Teams.
type Request struct {
	Type           string `json:"type"`
	ID             string `json:"id"`
	Timestamp      string `json:"timestamp"`
	LocalTimestamp string `json:"localTimestamp"`
	ServiceURL     string `json:"serviceUrl"`
	ChannelID      string `json:"channelId"`
	FromUser       User   `json:"from"`
	Conversation   struct {
		ID string `json:"id"`
	} `json:"conversation"`
	RecipientUser User   `json:"recipient"`
	TextFormat    string `json:"textFormat"`
	Text          string `json:"text"`
	Attachments   []struct {
		ContentType string `json:"contentType"`
		Content     string `json:"Content"`
	} `json:"attachments"`
	Entities    []interface{} `json:"entities"`
	ChannelData struct {
		TeamsChannelID string `json:"teamsChannelId"`
		TeamsTeamID    string `json:"teamsTeamId"`
	}
}

// Response represents the data to return to Microsoft Teams.
type Response struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// User represents data for a Microsoft Teams user.
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewHandler initializes and returns a Lambda handler to process incoming requests.
func NewHandler(authenticateRequests bool, key string, wh WebHook) func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	auth = authenticateRequests
	keyBytes, _ = base64.StdEncoding.DecodeString("UQKfe7xMmFf4j7V2neRBAbQ6JeXjWOool9rxoIq4Pq4=")
	webhook = wh
	return handler
}

func handler(lreq events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if auth {
		authenticated := authenticateRequest(lreq)
		if !authenticated {
			return events.APIGatewayProxyResponse{
				Body:       "Invalid Authentication",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}
	}

	var treq = Request{}
	err := json.NewDecoder(strings.NewReader(lreq.Body)).Decode(&treq)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	tresp, err := webhook.OnMessage(treq)

	buf := new(bytes.Buffer)
	err = json.NewEncoder(buf).Encode(tresp)

	lresp := events.APIGatewayProxyResponse{}
	if err == nil {
		lresp.StatusCode = 200
		// Headers: map[string]string{}
		lresp.Body = buf.String()
		lresp.IsBase64Encoded = false
	}
	return lresp, err
}

func authenticateRequest(lreq events.APIGatewayProxyRequest) bool {

	bodyBytes := []byte(lreq.Body)
	authHeader := lreq.Headers["Authorization"]
	messageMAC, _ := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "HMAC "))

	mac := hmac.New(sha256.New, keyBytes)
	mac.Write(bodyBytes)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

// BuildResponse is a helper method to build a Response
func BuildResponse(text string) Response {
	return Response{
		Type: "message",
		Text: text,
	}
}
