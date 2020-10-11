package main

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Time   int64  `json:"time"`
}
