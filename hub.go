package main

type Hub struct {
	clients map[*Client]bool

	broadcast  chan []ClientMessage
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []ClientMessage),
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}

}

func (h *Hub) isClientExist(c *Client) bool{
	if _, ok := h.clients[c]; ok {
        return true
	} 
	return false
}

func (h *Hub) run() {
	for {
		select {
	case client := <-h.register:
			h.clients[client] = true
		}
	case client := <-h.unregister:
		if _,ok := h.clients[client]; ok {
			delete(h.clients, client)
			close(client.send)
		}
	case message := <-h.broadcast:
		 if client, ok := h.clients[message.receiver]; ok {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		 }
		}
	}
}
