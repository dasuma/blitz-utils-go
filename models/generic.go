package models

//GenericResponse model
type GenericResponse struct {
	Response interface{} `json:"result"`
	Error    error
	Message  string `json:"message"`
}

//GenericErrorResponse model
type GenericErrorResponse struct {
	Message  string      `json:"message"`
	Code     int         `json:"code"`
	Response interface{} `json:"result"`
}
