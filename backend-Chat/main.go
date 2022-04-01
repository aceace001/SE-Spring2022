package main

import (
	"flag"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {

	pool := websocket.NewPool()
	go pool.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(pool, w, r)
	})
	http.ListenAndServe(*addr, nil)
}
