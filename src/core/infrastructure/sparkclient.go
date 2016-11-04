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
