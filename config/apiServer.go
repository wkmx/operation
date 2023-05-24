package config

import (
	"log"
	"net"
	"net/http"
)

type ApiServer struct {
	Server *http.Server
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello wkmx!"))
}

func InitApiServer() (err error) {
	var (
		mux      *http.ServeMux
		server   *http.Server
		listener net.Listener
	)
	mux = http.NewServeMux()
	mux.HandleFunc("/", handlerIndex)

	if listener, err = net.Listen("tcp", ":8080"); err != nil {
		return
	}

	server = &http.Server{
		Handler: mux,
	}

	log.Println("server start, port: 8080 ...")
	err = server.Serve(listener)
	if err != nil {
		return
	}

	return
}
