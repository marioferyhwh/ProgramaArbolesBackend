package models

//Message estructura de los mensajes de respuesta
type Message struct {
	Code     int    `json:"code"`
	Data     string `json:"data,omitempty"`
	Message  string `json:"message,omitempty"`
	NewToken string `json:"new_token,omitempty"`
}
