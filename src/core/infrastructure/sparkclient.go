package infrastructure

import (
	"encoding/json"
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
