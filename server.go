package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//List of online users
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//message broadcast channel
	Message chan string
}

//Create a server interface
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

//The goroutine listens to the Message broadcast message channel, and it will be sent to all online users once there is a message.
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//Send messages to all online users.
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//broadcast message
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//...Currently connected business
	//fmt.Println("Connection established successfully")

	user := NewUser(conn, this)

	user.Online()

	//Accept messages sent by clients
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}

			//get user's message (remove '\n')
			msg := string(buf[:n-1])

			//The user performs message processing for msg
			user.DoMessage(msg)
		}
	}()

	//The current handler is blocked
	select {}
}

//The interface to start the server
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//Start a goroutine that listens for Messages
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}
}
