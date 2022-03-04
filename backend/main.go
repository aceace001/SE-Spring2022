package main

import "main/app"

func main() {
	app.StartApplication()
	// sprint 2
	server := NewServer("127.0.0.1", 8081)
	server.Start()
}
