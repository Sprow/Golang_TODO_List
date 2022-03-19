package todo

// todo list methods

type Manager struct {
	list List
}

func NewManager(list List) *Manager {
	return &Manager{
		list: list,
	}
}

func (m *Manager) AddItem(item Item) {
	m.list.Items = append(m.list.Items, item)
}

func (m *Manager) UpdateItemStatus(index int) {
	if index < len(m.list.Items) {
		m.list.Items[index].Done = !m.list.Items[index].Done
	}
}

func (m *Manager) RemoveItem(index int) {
	s := m.list.Items
	if index < len(s) {
		m.list.Items = append(s[:index], s[index+1:]...)
	}
}
func (m *Manager) UpdateItemText(index int, text string) {
	if index < len(m.list.Items) {
		m.list.Items[index].Text = text
	}
}

func (m Manager) GetAllTodos() List {
	return m.list.Clone()
}
