package models

import (
	"net/http"
)

type Order struct {
	Id           uint64 `json:"id"`
	UserFullName string `json:"userFullName"`
	Isbn         uint64 `json:"isbn"`
	Date         string `json:"order_date"`
}

type OrderList struct {
	Orders []Order `json:"orders"`
}

// todo: Bind (эти ошибки не возвращаются обратно юзеру)
// тут нужна валидация полей
func (i *Order) Bind(r *http.Request) error {
	return nil
}

func (*OrderList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Order) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Result struct {
	Id        uint64 `json:"id"`
	FullName  string `json:"fullName"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	OrderDate string `json:"order_date"`
}

type ResultList struct {
	Orders []Order `json:"orders"`
}

func (i *Result) Bind(r *http.Request) error {
	return nil
}

func (*ResultList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Result) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
