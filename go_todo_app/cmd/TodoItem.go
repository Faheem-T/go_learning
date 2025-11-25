package main

type Status string

const (
	StatusPending  Status = "pending"
	StatusActive   Status = "active"
	StatusComplete Status = "complete"
)

type TodoItem struct {
	ID     uint8
	Task   string
	Status Status
}
