package session

import (
	"github.com/Sprow/todo/internal/todo"
	"sync"
)

type ManagerSession struct {
	mu      sync.Mutex
	session map[string]*todo.Manager
}

func NewManagerSession() *ManagerSession { // ????? any input???
	return &ManagerSession{
		session: make(map[string]*todo.Manager),
	}
}

func (s *ManagerSession) GetSession(token string) *todo.Manager {
	sm, ok := s.session[token]
	if ok {
		return sm
	}

	newTodoListManager := todo.NewManager(todo.List{
		Title: "",
		Items: nil,
	})
	s.mu.Lock()
	defer s.mu.Unlock()
	s.session[token] = newTodoListManager
	return newTodoListManager
}
