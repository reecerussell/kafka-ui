package ws

// Handler is an interface used to handle incoming messages.
type Handler interface {
	Handle(m *Message) error
}
