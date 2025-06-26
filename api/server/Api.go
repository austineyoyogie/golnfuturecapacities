package server

import (
	"github.com/gorilla/handlers"
	"golnfuturecapacities/api/routes"
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	mux := routes.RouteHandler()
	server := http.Server{
		Addr:           s.addr,
		Handler:        loadCors(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("\nStarting [SERVER] on Port %s", s.addr)
	return server.ListenAndServe()
}

func loadCors(r http.Handler) http.Handler {
	headers := handlers.AllowedHeaders([]string{"Origin", "Access-Control-Allow-Origin", "Content-Type",
		"Accept", "Jwt-Token", "Authorization", "Origin, Accept", "X-Requested-With", "X-CSRF-Token",
		"Access-Control-Request-Method", "Access-Control-Request-Headers", "Location", "Entity", "Accept", "Authorization"})
	exposes := handlers.ExposedHeaders([]string{"Origin", "Content-Type", "Accept", "Jwt-Token", "Authorization",
		"Access-Control-Allow-Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "true"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	credentials := handlers.AllowCredentials()
	return handlers.CORS(headers, exposes, methods, origins, credentials)(r)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("User ID Index!!!\n"))
	if err != nil {
		return
	}
}
