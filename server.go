package demo

type Server interface {
	Serve(port int) error
}

var A Server
