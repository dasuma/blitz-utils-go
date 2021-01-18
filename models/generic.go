package models

//GenericResponse model
type GenericResponse struct {
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Response  interface{} `json:"result"`
	IsSuccess bool        `json:"is_success"`
}
