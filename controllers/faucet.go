package controllers

import (
	"github.com/astaxie/beego"
	"srttoken/models/data_body"
	"fmt"
	"encoding/json"
	"srttoken/handle"
	"srttoken/models/data_response"
	"srttoken/models/data_error"
)

type FaucetController struct {
	beego.Controller
}

// URLMapping ...
func (f *FaucetController) URLMapping() {
	f.Mapping("buy", f.Buy)
	f.Mapping("balance", f.Balance)

}

// Post ...
// @Title Buy token
// @Description buy token by share link
// @Param	body		body 	data_body.BodyBuyToken	true		"body for User content"
// @Success 201 {object} string
// @Failure 403 body is empty
// @router /buy [post]
func (f *FaucetController) Buy() {
	// check link
	var body data_body.BodyBuyToken
	json.Unmarshal(f.Ctx.Input.RequestBody, &body)
	fmt.Println(body)
	response_buy := handle.RequestSendToken(body)
	f.Data["json"] = response_buy
	f.ServeJSON()
}

// Balance ...
// @Title Get balance address
// @Description get Balance address by id
// @Param	address		query 	string	true		"The key for staticblock"
// @Success 200 {object} string
// @Failure 403 address is empty
// @router /balance [get]
func (f *FaucetController) Balance() {
	address := f.GetString("address")
	balance := handle.GetBalance(address)
	balance_response := data_response.ResponseBalance{
		Data:balance,
		Err:data_error.ErrorSuccess(),
	}
	f.Data["json"] = balance_response
	f.ServeJSON()
}
