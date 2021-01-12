package ws

// Message represents the structure of a the data sent
// through a web socket.
type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
