package controllers

import (
	"github.com/BitAssetManagement/cms/src/common"
	"github.com/BitAssetManagement/cms/src/model"
	"github.com/BitAssetManagement/cms/src/service"
	"time"

	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

/**
进入管理员列表页面
*/
func (this *UserController) List() {
	this.show("user/userList.html")
}

/**
获取分页展示数据
*/
func (this *UserController) Gridlist() {
	pageNum, _ := this.GetInt("page")
	rowsNum, _ := this.GetInt("rows")
	usermail := this.GetString("usermail")
	userphone := this.GetString("userphone")
	username := this.GetString("username")
	userid := this.GetString("userid")
	account := this.GetString("accout")
	p := common.NewPager(pageNum, rowsNum)
	count, user := service.UserService.Gridlist(p, userid, usermail, username, userphone, account)
	this.jsonResultPager(count, user)
}

/**
进入添加页面
*/
func (this *UserController) Toadduser() {
	this.show("user/addUser.html")
}

/**
添加管理员
*/
func (this *UserController) Adduser() {
	account := this.GetString("account")
	mail := this.GetString("mail")
	name := this.GetString("name")
	phone := this.GetString("phone")
	catogery := this.GetString("catogery")
	password := this.GetString("password")
	// groupIds := this.GetString("ids")

	//参数校验
	valid := validation.Validation{}
	valid.Required(account, "账号").Message("不能为空")
	valid.MaxSize(account, 20, "账号").Message("长度不能超过20个字符")
	valid.Required(mail, "邮箱").Message("不能为空")
	valid.MaxSize(mail, 50, "邮箱").Message("长度不能超过50个字符")
	valid.Email(mail, "邮箱").Message("格式错误")
	valid.Required(name, "姓名").Message("不能为空")
	valid.MaxSize(name, 20, "姓名").Message("长度不能超过20个字符")
	valid.Required(phone, "手机号码").Message("不能为空")
	valid.MaxSize(phone, 15, "手机号码").Message("长度不能超过15个字符")
	// valid.Required(department, "部门").Message("不能为空")
	// valid.MaxSize(department, 20, "部门").Message("长度不能超过20个字符")
	valid.Required(password, "密码").Message("不能为空")
	valid.MaxSize(password, 20, "密码").Message("长度不能超过20个字符")
	// valid.MinSize(groupIds, 1, "组信息").Message("请至少选择一个")
	valid.Required(catogery, "用户分类").Message("不能为空")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			this.jsonResult((err.Key + err.Message))
		}
	}

	password = common.EncodeMessageMd5(password)

	user := &model.User{
		Accout:     account,
		Name:       name,
		Catogery:   catogery,
		Mail:       mail,
		Phone:      phone,
		Password:   password,
		Createtime: time.Now(),
		Updatetime: time.Now(),
		Isdel:      1}

	// if err := service.UserService.AddUser(user, groupIds); err != nil {

	if err := service.UserService.AddUser(user); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
进入修改页面
*/
func (this *UserController) Tomodifyuser() {
	userId, _ := this.GetInt64("userId")
	// this.jsonResult(userId)
	user, _ := service.UserService.GetUserById(userId)
	// this.jsonResult(user)
	this.Data["user"] = user
	this.show("user/modifyUser.html")
}

/**
修改管理员
*/
func (this *UserController) Modifyyuser() {
	userId, _ := this.GetInt64("userId")
	account := this.GetString("account")
	mail := this.GetString("mail")
	name := this.GetString("name")
	phone := this.GetString("phone")
	catogery := this.GetString("catogery")
	password := this.GetString("password")
	// groupIds := this.GetString("groupids")

	//参数校验
	valid := validation.Validation{}
	valid.Required(account, "账号").Message("不能为空")
	valid.MaxSize(account, 20, "账号").Message("长度不能超过20个字符")
	valid.Required(mail, "邮箱").Message("不能为空")
	valid.MaxSize(mail, 50, "邮箱").Message("长度不能超过50个字符")
	valid.Email(mail, "邮箱").Message("格式错误")
	valid.Required(name, "姓名").Message("不能为空")
	valid.MaxSize(name, 20, "姓名").Message("长度不能超过20个字符")
	valid.Required(phone, "手机号码").Message("不能为空")
	valid.MaxSize(phone, 15, "手机号码").Message("长度不能超过15个字符")
	// valid.Required(department, "部门").Message("不能为空")
	// valid.MaxSize(department, 20, "部门").Message("长度不能超过20个字符")

	if len(password) > 0 {
		valid.Required(password, "密码").Message("不能为空")
		valid.MaxSize(password, 20, "密码").Message("长度不能超过20个字符")
	}

	// valid.MinSize(groupIds, 1, "组信息").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			this.jsonResult((err.Key + err.Message))
		}
	}

	if len(password) != 0 {
		password = common.EncodeMessageMd5(password)
	}

	user := &model.User{
		Id:         userId,
		Accout:     account,
		Name:       name,
		Catogery:   catogery,
		Mail:       mail,
		Phone:      phone,
		Password:   password,
		Createtime: time.Now(),
		Updatetime: time.Now(),
		Isdel:      1}

	// if err := service.UserService.ModifyUser(user, groupIds); err != nil {
	if err := service.UserService.ModifyUser(user); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
删除
*/
func (this *UserController) Delete() {
	userids := this.GetString("userids")
	if err := service.UserService.Delete(userids); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
获取管理员组列表数据
修改管理员的时候需要加载管理员组列表，并且设置已经选择的权限为选中状态
// */
// func (this *UserController) Gridgrouplist() {
// 	userId, _ := this.GetInt64("userId")
// 	groupName := this.GetString("groupName")
// 	pageNum, _ := this.GetInt("page")
// 	rowsNum, _ := this.GetInt("rows")
// 	p := common.NewPager(pageNum, rowsNum)

// 	count, userGroup := service.UserCatogeryService.Gridlist(groupName, p)
// 	checkedGroupId := service.UserService.GetAllCheckGroup(userId)

// 	userCheckGroup := make([]model.Usercatogerycheck, len(userGroup))

// 	for index, user := range userGroup {
// 		userCheck := model.Usercatogerycheck{
// 			Id:           user.Id,
// 			Catogeryname: user.Catogeryname,
// 			Des:          user.Des,
// 			Createtime:   user.Createtime,
// 			Updatetime:   user.Updatetime,
// 			Isdel:        user.Isdel,
// 			Check:        checkedGroupId[user.Id]}
// 		userCheckGroup[index] = userCheck
// 	}

// 	this.jsonResultPager(count, userCheckGroup)
// }
