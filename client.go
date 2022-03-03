package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	// create client object
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	// connect client
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)

		return nil
	}

	client.conn = conn

	// return
	return client
}

// deal with server's response
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)

	// for {
	// 	buf := make()
	// 	client.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}
func (client *Client) menu() bool {
	var flag int

	fmt.Println("1. public chat")
	fmt.Println("2. private chat")
	fmt.Println("3. update profile")
	fmt.Println("0. log out")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>Please input valid integer number(0-3): ")

		return false
	}
}

// select active users
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

// create private chat function
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	client.SelectUsers()
	fmt.Println(">>>>please select a user to chat! ('exit' quit chat)")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println("please start chat! ('exit' quit chat)")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// message is not NULL
			if len(chatMsg) != 0 {
				sendMsg := chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}			
		}
	}
}
func (client *Client) PublicChat() {
	// prompt user input msg
	var chatMsg string

	fmt.Println(">>>>Please start chat, 'exit' quit chat! ")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// send to server

		// message is not NULL
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>Please type, 'exit' to quit.")
		fmt.Scanln(&chatMsg)
	}
}
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>Please input your username: ")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// deal with different options
		switch client.flag {
		case 1:
			// public chat mode
			client.PublicChat()
			break
		case 2:
			// private chat mode
			fmt.Println("2. private chat mode...")
			break
		case 3:
			// update profile
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "set IP address(default 127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "set port(default 8888)")
}

func main() {
	// parse comment
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>connection failed...")

		return
	}

	go client.DealResponse()

	fmt.Println(">>>>>>connection succeed...")

	// start client
	client.Run()
}
