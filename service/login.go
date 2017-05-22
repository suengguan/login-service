package service

import (
	"fmt"
	//"time"

	"model"

	daoApi "api/dao_service"
	//systemApi "api/system_service"

	"github.com/astaxie/beego"
)

type LoginService struct {
}

func (this *LoginService) Login(user *model.User) (*model.User, error) {
	var err error
	var newUser *model.User

	// check login
	beego.Debug("->check login")
	err = this.checkLogin(user)
	if err != nil {
		beego.Debug("check login data failed")
		return nil, err
	}

	// get user by name
	beego.Debug("->get user by name")
	newUser, err = daoApi.UserDaoApi.GetByName(user.Name)
	if err != nil {
		beego.Debug("get user failed!")
		return nil, err
	}

	// check password
	beego.Debug("->check password")
	if user.EncryptedPassword != newUser.EncryptedPassword {
		err = fmt.Errorf("%s", "user password is incorrect")
		return nil, err
	}

	// get resource
	beego.Debug("->get resource")
	newUser.Resource, err = daoApi.ResourceDaoApi.GetByUserId(newUser.Id)
	if err != nil {
		beego.Debug("get resource failed")
		return nil, err
	}

	// get project
	beego.Debug("->get all projects")
	var projects []*model.Project
	projects, err = daoApi.BussinessDaoApi.GetAllProjects(newUser.Id)
	if err != nil {
		beego.Debug("get projects failed")
		return nil, err
	}
	for _, p := range projects {
		newUser.Projects = append(newUser.Projects, p)
	}

	beego.Debug("result:", *newUser)

	// todo send mail
	// var emails []*model.Email
	// var e1 model.Email
	// e1.To = newUser.Email
	// e1.Subject = "登录提醒"
	// currentTime := time.Now().Format("2006-01-02 15:04:05")
	// e1.Body = "【科思世纪官网登录提醒】欢迎您于 " + currentTime + " 登录PME系统。特此提醒。【科思世纪】"
	// emails = append(emails, &e1)

	//systemApi.ApiSendEmails(emails)

	return newUser, err
}

func (this *LoginService) checkLogin(user *model.User) error {
	var err error

	if user.Name == "" {
		err = fmt.Errorf("%s", "please input your name")
		return err
	}

	if user.EncryptedPassword == "" {
		err = fmt.Errorf("%s", "please input your password")
		return err
	}

	return err
}
