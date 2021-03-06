package chat

// DefaultHub ...
var DefaultHub = NewHub()

// Hub ...
type Hub struct {
	Join  chan *Conn
	Conns map[*Conn]bool
	Echo  chan string
}

// NewHub ...
func NewHub() *Hub {
	return &Hub{
		Join:  make(chan *Conn),
		Conns: make(map[*Conn]bool),
		Echo:  make(chan string),
	}
}

// Start ...
func (hub *Hub) Start() {
	for {
		select {
		case conn := <-hub.Join:
			DefaultHub.Conns[conn] = true
		case msg := <-hub.Echo:
			for conn := range hub.Conns {
				conn.Send <- msg
			}
		}
	}
}
