package model

type Task struct {
	Id      TaskId
	Title   string
	Content string
}

type TaskId int64
