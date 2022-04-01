package main

import (
	"flag"
	"net/http"

	"github.com/aceace001/SE-Spring2022/backend-Chat/pkg/websocket"
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
