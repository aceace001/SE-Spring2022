package main

import "net"

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

//the method of user processes the message
func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}

//The method of monitoring the current User channel: It will be sent directly to the opposite client once there is a message.
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
