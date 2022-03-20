package todo

import "sync"

// todo list methods

type Manager struct {
	list List
	// mutex на случай если todolist будет общим для нескольких человек
	mutex sync.Mutex
}

func NewManager(list List) *Manager {
	return &Manager{
		list: list,
	}
}

func (m *Manager) AddItem(item Item) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.list.Items = append(m.list.Items, item)
}

func (m *Manager) UpdateItemStatus(index int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if index < len(m.list.Items) {
		m.list.Items[index].Done = !m.list.Items[index].Done
	}
}

func (m *Manager) RemoveItem(index int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	s := m.list.Items
	if index < len(s) {
		m.list.Items = append(s[:index], s[index+1:]...)
	}
}
func (m *Manager) UpdateItemText(index int, text string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if index < len(m.list.Items) {
		m.list.Items[index].Text = text
	}
}

func (m *Manager) GetAllTodos() List {
	return m.list.Clone()
}
