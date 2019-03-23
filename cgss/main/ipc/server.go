package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(ch chan string) {
		for {
			request := <-ch
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			ch <- string(b)
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created.")
	return session
}