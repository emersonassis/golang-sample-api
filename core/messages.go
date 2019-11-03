package core

//SuccessMessage ...
type SuccessMessage struct {
	Message string `json:"message"`
}

//ErrMessage is return message default
type ErrMessage struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Erro    string `json:"erro"`
}

//ErrDetail ...
type ErrDetail struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Message  string `json:"message"`
}
