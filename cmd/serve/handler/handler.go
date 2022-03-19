package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Sprow/todo/internal/todo"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	todoManager *todo.Manager
}

func NewHandler(todoManager *todo.Manager) *Handler {
	return &Handler{
		todoManager: todoManager,
	}
}

func (h *Handler) Register(r *chi.Mux) {
	r.Post("/add_item", h.addItem)
	r.Post("/update_item_status", h.updateItemStatus)
	r.Post("/update_item_text", h.updateItemText)
	r.Get("/", h.getAllTodos)
}

func (h *Handler) addItem(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var item todo.Item
	err := d.Decode(&item)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}

	h.todoManager.AddItem(item)
	fmt.Println("add item>>>", item)
}

type indexDTO struct {
	Index int `json:"index"`
}

func (h *Handler) updateItemStatus(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var index indexDTO
	err := d.Decode(&index)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}
	h.todoManager.UpdateItemStatus(index.Index)
}

func (h *Handler) getAllTodos(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	data := h.todoManager.GetAllTodos()
	e.SetIndent("", "  ") // фомотирует джейсон которвый приходит
	err := e.Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}

func (h *Handler) removeItem(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var index indexDTO
	err := d.Decode(&index)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}
	h.todoManager.RemoveItem(index.Index)
}

type updateTextDTO struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

func (h *Handler) updateItemText(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var textToUpdate updateTextDTO
	err := d.Decode(&textToUpdate)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "не сомог декодировать", http.StatusBadRequest)
		return
	}
	h.todoManager.UpdateItemText(textToUpdate.Index, textToUpdate.Text)
}
