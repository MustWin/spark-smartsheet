package infrastructure

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

const sparkBase = "https://api.ciscospark.com/v1/"

// SparkClient is a Spark specific REST client
type SparkClient struct {
	*RESTClient
}

// NewSparkClient returns an initialized SparkClient
func NewSparkClient(apiKey string) *SparkClient {
	return &SparkClient{RESTClient: NewRESTClient(apiKey)}
}

// Room represents a Spark Room as returned from the API
type Room struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	IsLocked     bool      `json:"isLocked"`
	LastActivity time.Time `json:"lastActivity"`
	CreatorID    string    `json:"creatorId"`
	Created      time.Time `json:"created"`
}

// RoomResponse wraps the slice of Room that is returned
// in certain API requests
type RoomResponse struct {
	Items []Room `json:"items"`
}

// Hook represents a web hook registration in Spark
type Hook struct {
	Name      string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource  string `json:"resource"`
	Event     string `json:"event"`
	Filter    string `json:"filter"`
	Secret    string `json:"secret"`
}

// HookRequest encapsulates the request to register a web hook
type HookRequest struct {
	Hook
}

// HookResponse encapsulates the response to registering a web hook
type HookResponse struct {
	Hook
	ID      string    `json:"id"`
	Created time.Time `json:"created"`
}

// HooksResponse describes the data coming back from a request to list
// the web hooks
type HooksResponse struct {
	Items []HookResponse `json:"items"`
}

// Message represents a Spark message
type Message struct {
	RoomID        string   `json:"roomId"`
	ToPersonID    string   `json:"toPersonId"`
	ToPersonEmail string   `json:"toPersonEmail"`
	Text          string   `json:"text"`
	Markdown      string   `json:"markdown"`
	Files         []string `json:"files"`
}

// MessageResponse describes the full Message including ID
type MessageResponse struct {
	Message
	ID          string    `json:"id"`
	RoomType    string    `json:"roomType"`
	PersonID    string    `json:"personId"`
	PersonEmail string    `json:"personEmail"`
	Created     time.Time `json:"created"`
	Mentioned   []string  `json:"mentionedPeople"`
}

// MessagesResponse describes a response including a list of Messages
type MessagesResponse struct {
	Items []MessageResponse `json:"items"`
}

// ListRooms returns a slice of Room representing all Rooms
// visible to the logged in user
func (c *SparkClient) ListRooms() ([]Room, error) {
	body, err := c.GetResource(sparkBase + "rooms")
	if err != nil {
		return nil, err
	}
	rooms := RoomResponse{}
	err = json.Unmarshal(body, &rooms)
	return rooms.Items, err
}

// ListHooks returns a slice of HookResponse describing all
// registered web hooks
func (c *SparkClient) ListHooks() ([]HookResponse, error) {
	body, err := c.GetResource(sparkBase + "webhooks")
	if err != nil {
		return nil, err
	}
	hooks := HooksResponse{}
	err = json.Unmarshal(body, &hooks)
	return hooks.Items, err
}

// ListMessages returns a slice of Message representing messages
func (c *SparkClient) ListMessages(roomName string) ([]MessageResponse, error) {
	if roomName == "" {
		return nil, fmt.Errorf("no room specified")
	}
	id := ""
	rooms, err := c.ListRooms()
	if err != nil {
		return nil, err
	}
	for _, room := range rooms {
		if room.Title == roomName {
			id = room.ID
			break
		}
	}
	if id == "" {
		return nil, fmt.Errorf("room %q not found", roomName)
	}

	body, err := c.GetResource(sparkBase + "messages?roomId=" + id)
	if err != nil {
		return nil, err
	}
	messages := MessagesResponse{}
	err = json.Unmarshal(body, &messages)
	sort.Sort(byCreated(messages.Items))
	return messages.Items, err
}

// CreateHook creates a new webhook, filtered by an optional room name
func (c *SparkClient) CreateHook(name string, callbackURL string, roomName string) (HookResponse, error) {
	req := Hook{Name: name, TargetURL: callbackURL, Resource: "messages", Event: "created", Filter: "roomType=group", Secret: "s3cr37"}
	res := HookResponse{}

	if roomName != "" {
		req.Filter = ""
		rooms, err := c.ListRooms()
		if err != nil {
			return res, err
		}
		for _, room := range rooms {
			if room.Title == roomName {
				req.Filter = "roomId=" + room.ID
				break
			}
		}
		if req.Filter == "" {
			return res, fmt.Errorf("room named %q not found", roomName)
		}
	}
	body, err := c.PostResource(sparkBase+"webhooks", req)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(body, &res)
	return res, err
}

// DeleteHook removes a registered hook
func (c *SparkClient) DeleteHook(name string) error {
	hooks, err := c.ListHooks()
	if err != nil {
		return err
	}
	id := ""
	for _, hook := range hooks {
		if hook.Name == name {
			id = hook.ID
			break
		}
	}
	if id == "" {
		return fmt.Errorf("no matching webhook found")
	}
	_, err = c.DeleteResource(sparkBase + "webhooks/" + id)
	return err
}

type byCreated []MessageResponse

func (b byCreated) Len() int           { return len(b) }
func (b byCreated) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byCreated) Less(i, j int) bool { return b[i].Created.Before(b[j].Created) }
