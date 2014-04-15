package libirc

import (
	"errors"
	"net"
	"testing"
	"time"
)

type FakeConnection struct {
	written  []string
	readable []string
	errors   []error
}

func (c *FakeConnection) Write(p []byte) (int, error) {
	c.written = append(c.written, string(p[:len(p)]))
	return len(p), nil
}

func (c *FakeConnection) Read(p []byte) (int, error) {
	if len(c.errors) > 0 {
		err, remaining := c.errors[0], c.errors[1:]
		c.errors = remaining

		return 0, err
	}

	if len(c.readable) > 0 {
		message, remaining := c.readable[0], c.readable[1:]
		c.readable = remaining

		copy(p[:], message)
		return len(message), nil
	}

	return 0, nil
}

func (c *FakeConnection) Close() error {
	return nil
}

func (c *FakeConnection) LocalAddr() net.Addr {
	return nil
}

func (c *FakeConnection) RemoteAddr() net.Addr {
	return nil
}

func (c *FakeConnection) SetDeadline(t time.Time) error {
	return nil
}

func (c *FakeConnection) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *FakeConnection) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestWrite(t *testing.T) {
	client := NewClient("NickName", "UserName", "Full Name")

	conn := new(FakeConnection)
	client.Conn = conn

	client.Write(new(Message))

	if len(conn.written) != 1 {
		t.Errorf("Expected 1 message to be written, got %d", len(conn.written))
	}
}

func TestGetMessage(t *testing.T) {
	client := NewClient("NickName", "UserName", "Full Name")

	conn := new(FakeConnection)
	client.Conn = conn

	conn.readable = append(conn.readable, "NICK hello\r\n")

	message, err := client.GetMessage()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assertMessage(t, message, "", "NICK", []string{"hello"})
}

func TestGetMessageError(t *testing.T) {
	client := NewClient("NickName", "UserName", "Full Name")

	conn := new(FakeConnection)
	client.Conn = conn

	expected := errors.New("Something went wrong")
	conn.errors = append(conn.errors, expected)

	_, actual := client.GetMessage()

	if actual != expected {
		t.Errorf("Expected:\n%v\nGot:\n%v", expected, actual)
	}
}

func TestWriteNick(t *testing.T) {
	client := NewClient("NickName", "UserName", "Full Name")

	conn := new(FakeConnection)
	client.Conn = conn

	client.writeNick()
	message := conn.written[0]
	expected := "NICK NickName\r\n"

	if message != expected {
		t.Errorf("Expected, Actual\n%s\n%s", expected, message)
	}
}

func TestWriteUser(t *testing.T) {
	client := NewClient("NickName", "UserName", "Full Name")

	client.Mode = 8

	conn := new(FakeConnection)
	client.Conn = conn

	client.writeUser()
	message := conn.written[0]
	expected := "USER UserName 8 * :Full Name\r\n"

	if message != expected {
		t.Errorf("Expected, Actual\n%s\n%s", expected, message)
	}
}

func TestJoin(t *testing.T) {
	client := NewClient("NickName", "UserName", "Full Name")

	conn := new(FakeConnection)
	client.Conn = conn

	client.Join("#helloWorld")
	message := conn.written[0]
	expected := "JOIN #helloWorld\r\n"

	if message != expected {
		t.Errorf("Expected, Actual\n%s\n%s", expected, message)
	}
}
