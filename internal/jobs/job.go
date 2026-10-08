package jobs

type Job struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Command string `json:"command"`
}