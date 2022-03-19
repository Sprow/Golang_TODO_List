package todo

// data transfer objects

type List struct {
	Title string `json:"title"`
	Items []Item `json:"items"`
}

type Item struct {
	Done bool   `json:"done"`
	Text string `json:"text"`
}

func (l List) Clone() List {
	res := l
	res.Items = make([]Item, len(l.Items))
	for i := range l.Items {
		res.Items[i] = l.Items[i]
	}
	return res
}
