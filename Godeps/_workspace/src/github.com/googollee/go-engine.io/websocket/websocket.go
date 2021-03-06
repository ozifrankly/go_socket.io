package websocket

import (
	"socket.io/Godeps/_workspace/src/github.com/googollee/go-engine.io/transport"
)

var Creater = transport.Creater{
	Name:      "websocket",
	Upgrading: true,
	Server:    NewServer,
	Client:    NewClient,
}
