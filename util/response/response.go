package response

type Response struct {
	Status int
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func MakeResponse(status int) Response {
	return Response{Status: status}
}

func MakeErrorResponse(error string) ErrorResponse {
	return ErrorResponse{Error: error}
}
