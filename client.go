package irc

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"syscall"
	"time"
)

// The Client struct
type Client struct {
	Nick     string
	RealName string
	UserName string
	Mode     int
	Conn     net.Conn
	Messages chan *Message
}

// Create a new client providing the client's NICK, username, and real name.
func NewClient(nick string, username string, realname string) *Client {
	client := new(Client)

	client.Nick = nick
	client.UserName = username
	client.RealName = realname
	client.Mode = 0

	client.Messages = make(chan *Message, 10)

	return client
}

// Connect connects to the IRC server and blocks until you
// are fully connected.
func (c *Client) Connect(server string) error {
	conn, err := net.Dial("tcp", server)

	if err != nil {
		return err
	}

	c.Conn = conn

	// Wait for first message -- give the server 1 second
	// after acknowledgment
	c.GetMessage()
	time.Sleep(time.Second)

	// Write the NICK and USER commands
	c.writeNick()
	c.writeUser()

	// Block until message 375 comes through
	for true {
		message, err := c.GetMessage()

		if err != nil {
			return err
		}

		if message.Command == "375" {
			break
		}

	}

	return nil
}

// Write a NICK message to the server
func (c *Client) writeNick() {
	message := NewMessage("NICK", []string{c.Nick})
	c.Write(message)
}

// Write a USER message to the server
func (c *Client) writeUser() {
	message := NewMessage("USER", []string{
		c.UserName,           // <user>
		strconv.Itoa(c.Mode), // <mode>
		"*",                  // <unused>
		c.RealName,           // <realname>
	})
	c.Write(message)
}

// Connects to the server, and blocks until connected.  Once connected, the
// client will start listening and pushing messages to the client.Messages
// channel
func (c *Client) ConnectAndListen(server string) error {
	err := c.Connect(server)

	if err != nil {
		return err
	}

	go c.Listen()
	return nil
}

// Get one message from the underlying socket, and convert it into a Message
// struct (see irc.ParseMessage)
func (c *Client) GetMessage() (*Message, error) {
	status, err := bufio.NewReader(c.Conn).ReadString('\n')

	if err != nil {
		return nil, err
	}

	fmt.Printf(status)

	message := ParseMessage(status)

	return message, nil
}

// Send a Channel Join message
func (c *Client) Join(channel string) {
	message := NewMessage("JOIN", []string{channel})
	c.Write(message)
}

// Write a message to the client
func (c *Client) Write(message *Message) {
	fmt.Fprintf(c.Conn, message.ToString())
}

// Listen to the underlying connection for messages.
func (c *Client) Listen() {
	for true {
		message, err := c.GetMessage()

		if err == syscall.EINVAL {
			break
		}

		if err != nil {
			panic(err)
		}

		//fmt.Printf("%v\n", message)
		c.Messages <- message
	}
}
