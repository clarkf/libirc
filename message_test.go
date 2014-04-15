package libirc

import "testing"

func assertMessage(t *testing.T, m *Message, prefix string, command string, params []string) {
	if m.Prefix != prefix {
		t.Errorf("Expected prefix %s, got %s", prefix, m.Prefix)
	}

	if m.Command != command {
		t.Errorf("Expected command %s, got %s", command, m.Command)
	}

	for idx, param := range params {
		actual := m.Parameters[idx]

		if actual != param {
			t.Errorf("Expected parameter at position %d  to be %s, got %s", idx, param, actual)
		}
	}
}

func TestParseSimpleMessage(t *testing.T) {
	message := ParseMessage("NICK gobot")
	assertMessage(t, message, "", "NICK", []string{"gobot"})

	message = ParseMessage("PRIVMSG #gobot :Hello World!")
	assertMessage(t, message, "", "PRIVMSG", []string{"#gobot", "Hello World!"})
}

func TestParseComplexMessage(t *testing.T) {
	message := ParseMessage(":irc.localhost 232 :Welcome to the network")
	assertMessage(t, message, "irc.localhost", "232", []string{"Welcome to the network"})

	message = ParseMessage(":anotheruser!user@host PRIVMSG #channel :This is my message body")
	assertMessage(t, message, "anotheruser!user@host", "PRIVMSG", []string{"#channel", "This is my message body"})
}

func TestNewMessage(t *testing.T) {
	message := NewMessage("NICK", []string{"gobot"})
	assertMessage(t, message, "", "NICK", []string{"gobot"})
}

func TestToString(t *testing.T) {
	message := NewMessage("PRIVMSG", []string{"#gobot", "Hello World"})
	expected := "PRIVMSG #gobot :Hello World\r\n"
	actual := message.ToString()

	if actual != expected {
		t.Errorf("Expected message to equal\n\"%s\"\nInstead got\n\"%s\"", expected, actual)
	}
}

func TestPrefixedToString(t *testing.T) {
	message := NewMessage("PRIVMSG", []string{"#gobot", "Hello World"})
	message.Prefix = "gobot!gobot@localhost"
	expected := ":gobot!gobot@localhost PRIVMSG #gobot :Hello World\r\n"
	actual := message.ToString()

	if actual != expected {
		t.Errorf("Expected message to equal\n\"%s\"\nInstead got\n\"%s\"", expected, actual)
	}
}
