package model

type Response struct {
	Status int
}

type ResponseError struct {
	Error string `json:"error"`
}
