package models

type Todo struct {
	ID      int64  `json:"id"`
	Text    string `json:"text"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}
