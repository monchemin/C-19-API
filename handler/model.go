package handler

type Response struct {
	ID    string      `json:"ID"`
	Error string      `json:"error"`
	Data  interface{} `json:"response"`
}
