package data_response

import (
	"srttoken/models"
	"srttoken/models/data_error"
)

type ResponseBuyToken struct {
	Data *models.Transaction `json:"data"`
	Err *data_error.ErrorResponse `json:"error"`
}
