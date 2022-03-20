package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Sprow/todo/internal/session"
	"github.com/Sprow/todo/internal/todo"
	"github.com/go-chi/chi/v5"
	"github.com/rs/xid"
	"net/http"
	"time"
)

type Handler struct {
	managerSession *session.ManagerSession
}

func NewHandler(managerSession *session.ManagerSession) *Handler {
	return &Handler{
		managerSession: managerSession,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	r.Post("/add_item", h.sessionMiddleware(h.addItem))
	r.Post("/update_item_status", h.sessionMiddleware(h.updateItemStatus))
	r.Post("/update_item_text", h.sessionMiddleware(h.updateItemText))
	r.Get("/", h.sessionMiddleware(h.getAllTodos))
}

func (h *Handler) addItem(w http.ResponseWriter, r *http.Request) {
	todoManager := h.managerSession.GetSession(h.getToken(r))

	fmt.Println("add item Header >>>", r.Header)
	d := json.NewDecoder(r.Body)
	var item todo.Item
	err := d.Decode(&item)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}

	todoManager.AddItem(item)
	fmt.Println("add item>>>", item)
}

type indexDTO struct {
	Index int `json:"index"`
}

func (h *Handler) updateItemStatus(w http.ResponseWriter, r *http.Request) {
	todoManager := h.managerSession.GetSession(h.getToken(r))

	d := json.NewDecoder(r.Body)
	var index indexDTO
	err := d.Decode(&index)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}
	todoManager.UpdateItemStatus(index.Index)
}

func (h *Handler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	todoManager := h.managerSession.GetSession(h.getToken(r))
	e := json.NewEncoder(w)
	data := todoManager.GetAllTodos()
	e.SetIndent("", "  ") // фомотирует джейсон которвый приходит
	err := e.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *Handler) removeItem(w http.ResponseWriter, r *http.Request) {
	todoManager := h.managerSession.GetSession(h.getToken(r))
	d := json.NewDecoder(r.Body)
	var index indexDTO
	err := d.Decode(&index)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}
	todoManager.RemoveItem(index.Index)
}

type updateTextDTO struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

func (h *Handler) updateItemText(w http.ResponseWriter, r *http.Request) {
	todoManager := h.managerSession.GetSession(h.getToken(r))
	d := json.NewDecoder(r.Body)
	var textToUpdate updateTextDTO
	err := d.Decode(&textToUpdate)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}
	todoManager.UpdateItemText(textToUpdate.Index, textToUpdate.Text)
}

func (h *Handler) getToken(r *http.Request) string {
	c, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return c.Value
}

func (h *Handler) sessionMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := h.getToken(r)

		if t == "" {
			cookie := http.Cookie{
				Name:     "token",
				Value:    xid.New().String(),             //new token
				Expires:  time.Now().Add(24 * time.Hour), //user lose his todolist after 24h
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			r.AddCookie(&cookie)
		}

		handler(w, r)
	}

}
