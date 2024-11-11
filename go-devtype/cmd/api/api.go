package api

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/ktarafder/devtype-backend/service/user"
)

type APIServer struct {
    addr    string
    db      *sql.DB
    Handler http.Handler
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
    server := &APIServer{
        addr: addr,
        db:   db,
    }
    server.setupRouter()
    return server
}

func (s *APIServer) setupRouter() {
    router := mux.NewRouter()
    subrouter := router.PathPrefix("/api/v1").Subrouter()

    userStore := user.NewStore(s.db)
    userHandler := user.NewHandler(userStore)
    userHandler.RegisterRoutes(subrouter)

    s.Handler = router
}

// ServeHTTP implements http.Handler.
func (s *APIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.Handler.ServeHTTP(w, r)
}

func (s *APIServer) Run() error {
    log.Println("Listening on", s.addr)
    return http.ListenAndServe(s.addr, s)
}