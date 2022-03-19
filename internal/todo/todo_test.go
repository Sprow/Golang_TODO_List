package todo_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/Sprow/todo/internal/todo"
)

func TestAddItem(t *testing.T) {
	initList := todo.List{
		Title: "zzz",
		Items: []todo.Item{
			{
				Done: false,
				Text: "learn python",
			},
		},
	}

	type testStruct struct {
		name     string
		add      todo.Item
		expected todo.List
		initList todo.List
	}
	tests := []testStruct{
		{"first", todo.Item{Done: false, Text: "new item 1"},
			todo.List{
				Title: "zzz",
				Items: []todo.Item{
					{
						Done: false,
						Text: "learn python",
					},
					{
						Done: false,
						Text: "new item 1",
					},
				},
			}, initList},
		{"second", todo.Item{Done: true, Text: "math"},
			todo.List{
				Title: "zzz",
				Items: []todo.Item{
					{
						Done: false,
						Text: "learn python",
					},
					{
						Done: true,
						Text: "math",
					},
				},
			}, initList},
		{"third", todo.Item{Done: false, Text: "biology"},
			todo.List{
				Title: "zzz",
				Items: []todo.Item{
					{
						Done: false,
						Text: "learn python",
					},
					{
						Done: false,
						Text: "biology",
					},
				},
			}, initList},
	}

	for _, test := range tests {
		// t.Run даём имена тестам (name string in testStruct)
		t.Run(test.name, func(t *testing.T) {
			m := todo.NewManager(initList)
			m.AddItem(test.add)
			newList := m.GetAllTodos()
			require.Equal(t, test.expected, newList)
		})
	}
}

func TestUpdateItemStatus(t *testing.T) {
	type testStruct struct {
		init      todo.List
		itemIndex int
		expected  todo.List
	}

	tests := []testStruct{{
		init: todo.List{
			Title: "TODO List",
			Items: []todo.Item{
				{
					Done: true,
					Text: "learn Python",
				},
				{
					Done: true,
					Text: "learn Go",
				},
				{
					Done: false,
					Text: "learn japanese",
				},
			},
		},
		itemIndex: 2,
		expected: todo.List{
			Title: "TODO List",
			Items: []todo.Item{
				{
					Done: true,
					Text: "learn Python",
				},
				{
					Done: true,
					Text: "learn Go",
				},
				{
					Done: true,
					Text: "learn japanese",
				},
			},
		},
	},
		{
			init: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
			itemIndex: 0,
			expected: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: true,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
		},
		{
			init: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
			itemIndex: 5,
			expected: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
		},
	}
	for _, test := range tests {
		m := todo.NewManager(test.init)
		m.UpdateItemStatus(test.itemIndex)
		newList := m.GetAllTodos()
		require.Equal(t, test.expected, newList)
	}
}

func TestRemoveItem(t *testing.T) {
	type testStruct struct {
		init      todo.List
		itemIndex int
		expected  todo.List
	}

	tests := []testStruct{{
		init: todo.List{
			Title: "TODO List",
			Items: []todo.Item{
				{
					Done: true,
					Text: "learn Python",
				},
				{
					Done: true,
					Text: "learn Go",
				},
				{
					Done: false,
					Text: "learn japanese",
				},
			},
		},
		itemIndex: 2,
		expected: todo.List{
			Title: "TODO List",
			Items: []todo.Item{
				{
					Done: true,
					Text: "learn Python",
				},
				{
					Done: true,
					Text: "learn Go",
				},
			},
		},
	},
		{
			init: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
			itemIndex: 0,
			expected: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
		},
		{
			init: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
			itemIndex: 10,
			expected: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
		},
	}
	for _, test := range tests {
		m := todo.NewManager(test.init)
		m.RemoveItem(test.itemIndex)
		newList := m.GetAllTodos()
		require.Equal(t, test.expected, newList)
	}
}

func TestUpdateItemText(t *testing.T) {
	type testStruct struct {
		init      todo.List
		itemIndex int
		newString string
		expected  todo.List
	}

	tests := []testStruct{{
		init: todo.List{
			Title: "TODO List",
			Items: []todo.Item{
				{
					Done: true,
					Text: "learn Python",
				},
				{
					Done: true,
					Text: "learn Go",
				},
				{
					Done: false,
					Text: "learn japanese",
				},
			},
		},
		itemIndex: 2,
		newString: "learn economy",
		expected: todo.List{
			Title: "TODO List",
			Items: []todo.Item{
				{
					Done: true,
					Text: "learn Python",
				},
				{
					Done: true,
					Text: "learn Go",
				},
				{
					Done: false,
					Text: "learn economy",
				},
			},
		},
	},
		{
			init: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
			itemIndex: 0,
			newString: "learn JS",
			expected: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "learn JS",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
		},
		{
			init: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
			itemIndex: 77,
			newString: "learn JS",
			expected: todo.List{
				Title: "TODO List",
				Items: []todo.Item{
					{
						Done: false,
						Text: "football rules",
					},
					{
						Done: true,
						Text: "learn Go",
					},
					{
						Done: false,
						Text: "learn japanese",
					},
				},
			},
		},
	}
	for _, test := range tests {
		m := todo.NewManager(test.init)
		m.UpdateItemText(test.itemIndex, test.newString)
		newList := m.GetAllTodos()
		require.Equal(t, test.expected, newList)
	}
}

func TestGetAllTodos(t *testing.T) {
	init := todo.List{
		Title: "TODO List",
		Items: []todo.Item{
			{
				Done: false,
				Text: "football rules",
			},
			{
				Done: true,
				Text: "learn Go",
			},
			{
				Done: false,
				Text: "learn japanese",
			},
		},
	}
	expected := todo.List{
		Title: "TODO List",
		Items: []todo.Item{
			{
				Done: false,
				Text: "football rules",
			},
			{
				Done: true,
				Text: "learn Go",
			},
			{
				Done: false,
				Text: "learn japanese",
			},
		},
	}

	m := todo.NewManager(init)

	initClone := m.GetAllTodos()
	m.UpdateItemText(0, "learn f")
	m.AddItem(todo.Item{Done: false, Text: "biology"})
	m.UpdateItemStatus(3)
	require.Equal(t, expected, initClone)
}
