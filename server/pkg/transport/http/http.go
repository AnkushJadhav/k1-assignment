package http

import (
	"fmt"
	"net"
	"strconv"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"

	"github.com/gin-gonic/gin"
)

// Server represents a http server
type Server struct {
	engine   *gin.Engine
	bindIP   net.IP
	bindPort uint
	store    persistance.Client
}

// New creates a Server to run on ip and port
func New(ip, port string) (*Server, error) {
	parsedIP := validateIP(ip)
	if parsedIP == nil {
		return nil, fmt.Errorf("Invalid bind IP")
	}
	parsedPort, err := validatePort(port)
	if err != nil {
		return nil, fmt.Errorf("Invalid bind port")
	}

	return &Server{
		engine:   gin.Default(),
		bindIP:   parsedIP,
		bindPort: uint(parsedPort),
	}, nil
}

// SetStore adds a persistant store client to the server
func (s *Server) SetStore(db persistance.Client) {
	s.store = db
}

// AttachInternalRoutes attaches the internal - authenticated routes to the server
func (s *Server) AttachInternalRoutes(authenticator gin.HandlerFunc) {
	internalGrp := s.engine.Group("/api", authenticator)
	setInternalRoutes(internalGrp, s.store)
}

// AttachExternalRoutes attaches the external - unauthenticated routes to the server
func (s *Server) AttachExternalRoutes() {

}

// validateIP validates the IP
func validateIP(ip string) net.IP {
	return net.ParseIP(ip)
}

// validatePort validates the port
func validatePort(port string) (int, error) {
	return strconv.Atoi(port)

}
