package http

import (
	"fmt"
	"net"
	"strconv"

	"github.com/AnkushJadhav/k1-assignment/server/pkg/persistance"
	"github.com/AnkushJadhav/k1-assignment/server/pkg/transport/http/jwt"

	"github.com/gin-gonic/gin"
)

// Server represents a http server
type Server struct {
	engine   *gin.Engine
	bindIP   net.IP
	bindPort int
	store    persistance.Client
	issuer   *jwt.Issuer
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
		bindPort: parsedPort,
	}, nil
}

// SetAuth adds a JWT auth isuer and validator
func (s *Server) SetAuth(secret string) {
	iss := jwt.NewIssuer(secret)
	s.issuer = iss
	return
}

// SetStore adds a persistant store client to the server
func (s *Server) SetStore(db persistance.Client) {
	s.store = db
	return
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.engine.Run(s.bindIP.String() + ":" + strconv.Itoa(s.bindPort))
}

// AttachRoutes attaches the routes to the server
func (s *Server) AttachRoutes() {
	authGrp := s.engine.Group("/auth")
	s.setAuthRoutes(authGrp)

	apiGrp := s.engine.Group("/api")
	s.setAPIRoutes(apiGrp)
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
