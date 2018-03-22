package controllers

import (

	"github.com/astaxie/beego"
)

//  WalletController operations for Wallet
type WalletController struct {
	beego.Controller
}

// URLMapping ...
func (c *WalletController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Wallet
// @Param	body		body 	models.Wallet	true		"body for Wallet content"
// @Success 201 {int} models.Wallet
// @Failure 403 body is empty
// @router / [post]
func (c *WalletController) Post() {
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Wallet by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Wallet
// @Failure 403 :id is empty
// @router /:id [get]
func (c *WalletController) GetOne() {


	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Wallet
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Wallet
// @Failure 403
// @router / [get]
func (c *WalletController) GetAll() {
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Wallet
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Wallet	true		"body for Wallet content"
// @Success 200 {object} models.Wallet
// @Failure 403 :id is not int
// @router /:id [put]
func (c *WalletController) Put() {
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Wallet
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *WalletController) Delete() {
	c.ServeJSON()
}
