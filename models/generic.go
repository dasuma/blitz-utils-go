package models

//GenericResponse model
type GenericResponse struct {
	Response interface{}
	Error    error
}

//GenericErrorResponse model
type GenericErrorResponse struct {
	Message  string      `json:"message"`
	Code     int         `json:"code"`
	Response interface{} `json:"result"`
}
