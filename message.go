package irc

import (
	"strings"
)

type Message struct {
	Prefix     string
	Command    string
	Parameters []string
}

// Construct a new message with a command and parameters.
func NewMessage(command string, parameters []string) *Message {
	message := new(Message)

	message.Command = command
	message.Parameters = parameters

	return message
}

// Parse a string into a message.
func ParseMessage(message string) *Message {
	message = strings.Trim(message, "\r\n")

	msg := new(Message)
	msg.Prefix = ""

	// Look for 'source'
	if message[0] == ':' {
		tmp := strings.SplitN(message, " ", 2)
		msg.Prefix = tmp[0][1:]
		message = tmp[1]
	}

	// Look for trailing param
	if strings.Index(message, " :") != -1 {
		tmp := strings.SplitN(message, " :", 2)
		message = tmp[0]
		msg.Parameters = strings.Split(message, " ")
		msg.Parameters = append(msg.Parameters, tmp[1])
	} else {
		msg.Parameters = strings.Split(message, " ")
	}

	msg.Command = msg.Parameters[0]
	msg.Parameters = msg.Parameters[1:]

	return msg
}

// Convert a message to a raw string for sending to the server.
func (m *Message) ToString() string {
	params := m.Parameters
	var words []string

	if len(m.Prefix) > 0 {
		words = append(words, ":"+m.Prefix)
	}

	words = append(words, m.Command)

	for idx, param := range m.Parameters {
		if idx == len(params)-1 && strings.ContainsAny(param, " \t") {
			words = append(words, ":"+param)
		} else {
			words = append(words, param)
		}
	}

	return strings.Join(words, " ") + "\r\n"
}
