package main

import "time"

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	// TODO: answer here
	todos.items = append(todos.items, item)
}

func (todos *Todos) GetAll() []Item {
	var arr []Item

	for _, items := range todos.items {
		arr = append(arr, items)
	}
	return arr // TODO: replace this
}

func (todos *Todos) GetUpcoming() []Item {
	var arr []Item
	return append(arr, todos.items[1]) // TODO: replace this
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}
