package jsonapi

import (
	"encoding/json"
)

type Error struct {
	ID     string `json:"id,omitempty"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type Errors struct {
	Errors []*Error `json:"errors"`
}

func (e *Error) Set() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Errors) GetErrors() ([]byte, error) {
	return json.Marshal(e)
}

func NewError(status int, title string, detail string) *Error {
	e := new(Error)
	e.Status = status
	e.Title = title
	e.Detail = detail

	return e
}

func NewErrors(errors ...*Error) *Errors {
	e := new(Errors)
	e.Errors = errors

	return e
}
