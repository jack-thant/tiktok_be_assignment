package main

type Message struct {
	Sender    string  `json:"sender"`
	Message   string  `json:"message"`
	Timestamp float64 `json:"timestamp"`
}
