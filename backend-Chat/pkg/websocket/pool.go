package websocket

type Pool struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Run() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}

		case client := <-pool.Unregister:

			delete(pool.Clients, client)

			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}

		case message := <-pool.Broadcast:

			for client := range pool.Clients {
				client.Conn.WriteJSON(message)

			}
		}
	}
}
