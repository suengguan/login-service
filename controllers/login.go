package controllers

import (
	"app-service/login-service/models"
	"app-service/login-service/service"
	"encoding/json"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Login
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description login
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router / [post]
func (this *LoginController) Login() {
	var err error
	var user model.User
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err == nil {
		var svc service.LoginService
		var result []byte
		var newUser *model.User
		newUser, err = svc.Login(&user)
		if err == nil {
			result, err = json.Marshal(newUser)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}
