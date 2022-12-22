package models

type Todo struct {
	Text    string `json:"text"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}
