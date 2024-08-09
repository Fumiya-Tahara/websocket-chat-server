package controller

type Message struct {
	Type      string `json:"type"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
	Sender    string `json:"sender"`
}
