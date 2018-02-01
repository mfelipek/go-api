package server

import (
	"go-api/domain"
	"github.com/codegangsta/negroni"
	"gopkg.in/tylerb/graceful.v1"
	"net/http"
	"time"
)

// Request JSON body limit is set at 5MB (currently not enforced)
const BodyLimitBytes uint32 = 1048576 * 5

// Server type
type Server struct {
	negroni        *negroni.Negroni
	router         *Router
	gracefulServer *graceful.Server
	timeout        time.Duration
}

// Options for running the server
type Options struct {
	Timeout time.Duration
	ShutdownHandler func()
}

// NewServer Returns a new Server object
func NewServer() *Server {

	// set up server and middlewares
	n := negroni.Classic()

	s := &Server{n, nil, nil, 0}

	return s
}

func (s *Server) UseMiddleware(middleware domain.IMiddleware) *Server {
	// next convert it into negroni style handlerfunc
	s.negroni.Use(negroni.HandlerFunc(middleware.Handler))
	return s
}

func (s *Server) UseRouter(router *Router) *Server {
	// add router
	s.negroni.UseHandler(router)
	return s
}

func (s *Server) Run(address string, options Options) *Server {
	s.timeout = options.Timeout
	s.gracefulServer = &graceful.Server{
		Timeout:           options.Timeout,
		Server:            &http.Server{Addr: address, Handler: s.negroni},
		ShutdownInitiated: options.ShutdownHandler,
	}
	s.gracefulServer.ListenAndServe()
	return s
}

func (s *Server) Stop() {
	s.gracefulServer.Stop(s.timeout)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) *Server {
	s.negroni.ServeHTTP(w, r)
	return s
}