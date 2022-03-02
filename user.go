package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

//Create a user API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//Start a goroutine that listens for messages on the current user channel
	go user.ListenMessage()

	return user
}

//User's login business
func (this *User) Online() {

	//When the user logged in, add the user to the onlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//Broadcast current user logged in message
	this.server.BroadCast(this, "is online")
}

//User's offline business
func (this *User) Offline() {

	//delete the user from the onlineMap when they are offline
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//Broadcast user who is offline in message
	this.server.BroadCast(this, "is offline")

}

//Send a message to the client corresponding to the current User
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//the method of user processes the message
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//lookup who are the currently online users

		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "is online...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//order format: rename:Michael
		newName := strings.Split(msg, "|")[1]

		//Check if the name already exists
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("The current username is already in use, and please choose another name\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("You have updated your username:" + this.Name + "\n")
		}

	     } else if len(msg) > 4 && msg[:3] == "to|" {
		//Message format: to|Michael|message information

		//1 Get the username of the other user
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("The message format is incorrect, please use the \"to|Michael|Hello\" format.\n")
			return
		}

		//2 Get the other party's User object according to the username
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("The username does not exist\n")
			return
		}

		//3 Get the message content and send the message content through the other party's User object.
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("There is no message content, and please resend.\n")
			return
		}
		remoteUser.SendMsg(this.Name + "says to you:" + content)

	} else {
		this.server.BroadCast(this, msg)
	}
}

//The method of monitoring the current User channel: It will be sent directly to the opposite client once there is a message.
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
