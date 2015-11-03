package polling

import (
	"socket.io/Godeps/_workspace/src/github.com/googollee/go-engine.io/transport"
)

var Creater = transport.Creater{
	Name:      "polling",
	Upgrading: false,
	Server:    NewServer,
	Client:    NewClient,
}
