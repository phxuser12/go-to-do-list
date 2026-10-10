package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//go:embed templates/*.html
var templateFS embed.FS

type Server struct {
	store TodoStore
	tmpl  *template.Template
}

func NewServer(store TodoStore) *Server {
	return &Server{
		store: store,
		tmpl:  template.Must(template.ParseFS(templateFS, "templates/*.html")),
	}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.index)
	mux.HandleFunc("POST /todos", s.add)
	mux.HandleFunc("POST /todos/{id}/toggle", s.toggle)
	mux.HandleFunc("DELETE /todos/{id}", s.delete)
	return mux
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	s.render(w, "index.html")
}

func (s *Server) add(w http.ResponseWriter, r *http.Request) {
	if text := strings.TrimSpace(r.FormValue("text")); text != "" {
		s.store.Add(text)
	}
	s.render(w, "list")
}

func (s *Server) toggle(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.PathValue("id")); err == nil {
		s.store.Toggle(id)
	}
	s.render(w, "list")
}

func (s *Server) delete(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.Atoi(r.PathValue("id")); err == nil {
		s.store.Delete(id)
	}
	s.render(w, "list")
}

func (s *Server) render(w http.ResponseWriter, name string) {
	if err := s.tmpl.ExecuteTemplate(w, name, s.store.List()); err != nil {
		log.Println("render:", err)
	}
}

func main() {
	addr := ":8080"
	log.Println("listening on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, NewServer(NewMemoryStore()).Routes()))
}
