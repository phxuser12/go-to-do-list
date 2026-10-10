package main

import "sync"

type Todo struct {
	ID   int
	Text string
	Done bool
}

// TodoStore is the seam for swapping in SQLite later.
type TodoStore interface {
	List() []Todo
	Add(text string) Todo
	Toggle(id int)
	Delete(id int)
}

type MemoryStore struct {
	mu     sync.Mutex
	todos  []Todo
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{nextID: 1}
}

func (s *MemoryStore) List() []Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Todo(nil), s.todos...)
}

func (s *MemoryStore) Add(text string) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	t := Todo{ID: s.nextID, Text: text}
	s.nextID++
	s.todos = append(s.todos, t)
	return t
}

func (s *MemoryStore) Toggle(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos[i].Done = !s.todos[i].Done
		}
	}
}

func (s *MemoryStore) Delete(id int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.todos {
		if s.todos[i].ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return
		}
	}
}
