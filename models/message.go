package models

//Message estructura de los mensajes de respuesta
type Message struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	DataString string      `json:"data_string,omitempty"`
	Message    string      `json:"message,omitempty"`
	NewToken   string      `json:"new_token,omitempty"`
}
