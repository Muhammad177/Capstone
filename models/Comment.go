package models

type Comment struct {
	Comment  string `json:"comment" form:"comment"`
	ThreadID int    `json:"thread_id" form:"thread_id"`
	Threads  Thread `json:"thread"`
}
