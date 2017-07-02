package controllers

import (
	"github.com/BitAssetManagement/cms/src/common"
	"github.com/BitAssetManagement/cms/src/model"
	"github.com/BitAssetManagement/cms/src/service"
	"time"

	"github.com/astaxie/beego/validation"
)

type UserCatogeryController struct {
	BaseController
}

/**
进入管理员组管理页面
*/
func (this *UserCatogeryController) List() {
	this.show("usercatogery/usercatogeryList.html")
}

/**
获取管理员组列表数据
*/
func (this *UserCatogeryController) Gridlist() {
	catogeryname := this.GetString("catogeryname")
	pageNum, _ := this.GetInt("page")
	rowsNum, _ := this.GetInt("rows")
	p := common.NewPager(pageNum, rowsNum)

	count, usercatogery := service.UserCatogeryService.Gridlist(catogeryname, p)
	this.jsonResultPager(count, usercatogery)
}

/**
获取管理员组列表数据
*/
func (this *UserCatogeryController) ListAll() {
	usercatogery := service.UserCatogeryService.ListAll()
	this.jsonResult(usercatogery)
}

/**
进入添加页面
*/
func (this *UserCatogeryController) Toadd() {
	this.show("usercatogery/addUsercatogery.html")
}

/**
添加管理员组
*/
func (this *UserCatogeryController) Addusercatogery() {
	// ids := this.GetString("ids")
	catogeryname := this.GetString("catogeryname")
	describe := this.GetString("describe")

	//参数校验
	valid := validation.Validation{}
	valid.Required(catogeryname, "管理员组名称").Message("不能为空")
	valid.MaxSize(catogeryname, 20, "管理员组名称").Message("长度不能超过20个字符")
	valid.Required(describe, "描述信息").Message("不能为空")
	valid.MaxSize(describe, 50, "描述信息").Message("长度不能超过50个字符")
	// valid.MinSize(ids, 1, "权限").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			this.jsonResult((err.Key + err.Message))
		}
	}

	usercatogery := &model.Usercatogery{
		Catogeryname: catogeryname,
		Des:          describe,
		Createtime:   time.Now(),
		Updatetime:   time.Now(),
		Isdel:        1}
	if err := service.UserCatogeryService.AddUserCatogery(usercatogery); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
进入修改管理员组页面
*/
func (this *UserCatogeryController) Tomodify() {
	id, _ := this.GetInt64("usercatogeryid")
	usercatogery := service.UserCatogeryService.GetUserCatogeryById(id)
	this.Data["usercatogery"] = usercatogery
	this.show("usercatogery/modifyUsercatogery.html")
}

/**
修改管理员组
*/
func (this *UserCatogeryController) Modifyusercatogery() {
	// ids := this.GetString("ids")
	catogeryname := this.GetString("catogeryname")
	describe := this.GetString("describe")
	id, _ := this.GetInt64("id")

	//参数校验
	valid := validation.Validation{}
	valid.Required(catogeryname, "管理员组名称").Message("不能为空")
	valid.MaxSize(catogeryname, 20, "管理员组名称").Message("长度不能超过20个字符")
	valid.Required(describe, "描述信息").Message("不能为空")
	valid.MaxSize(describe, 50, "描述信息").Message("长度不能超过50个字符")
	// valid.MinSize(ids, 1, "权限").Message("请至少选择一个")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			this.jsonResult((err.Key + err.Message))
		}
	}

	usercatogery := &model.Usercatogery{
		Id:           id,
		Catogeryname: catogeryname,
		Des:          describe,
		Createtime:   time.Now(),
		Updatetime:   time.Now(),
		Isdel:        1}
	if err := service.UserCatogeryService.Modifyusercatogery(usercatogery); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
删除管理员组
*/
func (this *UserCatogeryController) Delete() {
	ids := this.GetString("ids")
	if err := service.UserCatogeryService.Delete(ids); err != nil {
		this.jsonResult(err.Error())
	}
	this.jsonResult(SUCCESS)
}

/**
加载权限树(用于添加管理员组的时候选择权限)
*/
func (this *UserCatogeryController) Loadtreewithoutroot() {
	//查询树结构不加载root节点
	roles := service.RoleService.Listtree(false)
	//展开一级目录
	for i, role := range roles {
		if role.Pid == 0 {
			roles[i].Open = true
		}
	}
	this.jsonResult(roles)
}

/**
加载权限树(用于修改管理员组的时候选择权限-添加时选择的权限在修改的时候需要选中)
*/
func (this *UserCatogeryController) Loadtreechecked() {
	admgroupuserid, _ := this.GetInt64("admgroupuserid")
	roleIdMap := service.UserCatogeryService.GetAllRoleByGroupId(admgroupuserid)
	//查询树结构不加载root节点
	roles := service.RoleService.Listtree(false)
	if roleIdMap == nil {
		//展开一级目录
		for i, role := range roles {
			if role.Pid == 0 {
				roles[i].Open = true
			}
		}
	} else {
		for i, role := range roles {
			if role.Pid == 0 {
				roles[i].Open = true
			}
			if _, ok := roleIdMap[role.Id]; ok {
				roles[i].Checked = true
			}
		}
	}
	this.jsonResult(roles)
}
