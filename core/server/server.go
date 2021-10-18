package server

type Type int8

const (
	ServerType_Fasthttp Type = 1
)

type Server interface {
	NewServer() *Server
	StartServer() error
	StopServer() error
}

//
//func NewServer(
//	t Type,) (server *Server) {
//	switch t {
//	case ServerType_Fasthttp:
//		return
//	}
//}
