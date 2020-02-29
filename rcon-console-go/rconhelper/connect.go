package rconhelper

import (
	"strings"

	"github.com/cpf/rcon"
)

/*
RconConnection is an abstractio for a connection
*/
type Connection struct {
	host   string
	port   int
	pass   string
	client *rcon.Client
}

/*
NewRconConnection creates a new object of type Connection
*/
func NewRconConnection(host string, port int, pass string) *Connection {
	c := new(Connection)
	c.host = host
	c.port = port
	c.pass = pass

	c.connect()
	return c
}

/*
Connect returns an authenticated connection
*/
func (c *Connection) connect() {

	c.client = rcon.NewClient(c.host /* Your servers IP address */, c.port /* Its port */, c.pass /* The RCON password for your server */)
}

/*
Execute executes a single command
*/
func (c *Connection) Execute(cmd []string) string {

	response, err := c.client.Execute(strings.Join(cmd[:], " "))
	if nil != err {
		// Failed to authorize your connection with the server.
		panic(err)
	}

	return response.Body
}
