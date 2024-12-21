package entities

type Task struct {
	ID            int    `json:"id,omitempty"`
	Author_name   string `json:"author_name"`
	Assignee_name string `json:"Assignee_name"`
	Created       string `json:"created"`
	Resolved      string `json:"resolved"`
	Status        string `json:"status"`
}
