package data_response

import "srttoken/models/data_error"

type ResponseBalance struct {
	Data int64 `json:"data"`
	Err *data_error.ErrorResponse `json:"error"`
}

