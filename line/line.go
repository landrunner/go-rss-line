package line

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

const push_api string = "https://api.line.me/v2/bot/message/broadcast"

/*
{
    "messages":[
        {
            "type":"text",
            "text":"Hello, world1"
        },
        {
            "type":"text",
            "text":"Hello, world2"
        }
    ]
}'
*/

// Message represents the messege to be sent to LINEbot.
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Body represents the full messege to be sent LINEbot.
// It is used for making JSON.
type Body struct {
	Messages []Message `json:"messages"`
}

// LineBot is a client to send message to LINEbot.
type LineBot struct {
	access_token string
}

// Create new LineBot Object
func New(access_token string) LineBot {
	return LineBot{access_token}
}

// Create UUID
func create_uuid() string {
	uuidObj, _ := uuid.NewUUID()
	return uuidObj.String()
}

// Send message to LINEbot.
func (l LineBot) SendMessage(message string) error {
	messages := []Message{{"text", message}}
	body := Body{messages}
	bodyjson, err := json.Marshal(body)
	fmt.Println(string(bodyjson))
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", push_api, bytes.NewBuffer(bodyjson))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+l.access_token)
	req.Header.Set("Content-Type", "application/json")
	uuid := create_uuid()
	req.Header.Set("X-Line-Retry-Key", uuid)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	bytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bytes))

	return nil
}
