package data_error

type ErrorResponse struct {
	Code int32  	`json:"code"`
	Message string 	`json:"message"`
}

func ErrorSuccess() *ErrorResponse {
	return &ErrorResponse{200, "Successful."}
}

func ErrorParam(Err string) *ErrorResponse{
	return &ErrorResponse{403, Err}
}

func ErrorConnect() *ErrorResponse{
	return &ErrorResponse{500, "Connection error."}
}

func ErrorInsert() *ErrorResponse{
	return &ErrorResponse{300, "Insert data fail."}
}

func ErrorNotFoundId() *ErrorResponse{
	return &ErrorResponse{404, "Not found id."}
}

func ErrorQuery(Err string) *ErrorResponse {
	return &ErrorResponse{403, Err}
}

