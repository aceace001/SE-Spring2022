// package main

// import (
// 	"fmt"
// 	"net"
// )

// type Client struct {
// 	ServerIp 	string 
// 	ServerPort 	int
// 	Name 		string 
// 	conn		net.Conn 
// }

// func NewClient(serverIp string, serverPort int) *Client {
// 	// create client object
// 	client := &Client{
// 		ServerIp:	serverIp,
// 		ServerPort: serverPort,
// 	}

// 	// connect client 
// 	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
// 	if err != nil {
// 		fmt.Println("net.Dial error: ", err)

// 		return nil 
// 	}

// 	client.conn = conn 

// 	// return 
// 	return client 
// }
 
// func main() {
// 	client := NewClient("127.0.0.1", 8888) 
// 	if client == nil {
// 		fmt.Println(">>>>>connection failed...")

// 		return 
// 	}

// 	fmt.Println(">>>>>>connection succeed...")

// 	// start client 
// 	select {}
// }