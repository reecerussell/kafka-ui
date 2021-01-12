package ws

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/reecerussell/kafka-ui/logging"
)

type Server struct {
	hub   *Hub
	mu    sync.Mutex
	hdlrs map[string]Handler
}

func New() *Server {
	hub := NewHub()

	return &Server{
		hub:   hub,
		mu:    sync.Mutex{},
		hdlrs: make(map[string]Handler),
	}
}

// Start starts the Server's underlying hub.
func (s *Server) Start() {
	s.hub.Run()
}

// Serve handles websocket requests from the peer.
func (s *Server) Serve(w http.ResponseWriter, r *http.Request) {
	conn, err := s.hub.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logging.Warning("An error occured while upgrading the connection: %v", err.Error())
		return
	}
	client := &Client{hub: s.hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	//go client.readPump()
}

// Handle adds a new handler for the given type.
func (s *Server) Handle(typ string, hdlr Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.hdlrs[typ] = hdlr
}

func (s *Server) Send(m *Message) {
	bytes, _ := json.Marshal(m)
	s.hub.broadcast <- bytes
}
