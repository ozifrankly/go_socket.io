package main

import (
	"log"
	"net/http"
	"socket.io/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"socket.io/Godeps/_workspace/src/github.com/googollee/go-socket.io"
	"socket.io/Godeps/_workspace/src/github.com/unrolled/render"
)

func main() {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			log.Println("new message", msg)

			so.Emit("chat message", msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	r := render.New(render.Options{Extensions: []string{".html"}})
	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})

	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir("./assets")))
	n.UseHandler(mux)
	n.Run(":3000")
}
