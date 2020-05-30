package models

//Message estructura de los mensajes de respuesta
type Message struct {
	Code    uint16 `json:"code"`
	Data    string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
