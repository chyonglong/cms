package service

import (
	"github.com/BitAssetManagement/cms/src/common"
	"github.com/BitAssetManagement/cms/src/model"
	"strconv"
	// "strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type userCatogeryService struct{}

/**
查询管理员组分页列表
*/
func (this *userCatogeryService) Gridlist(catogeryName string, pager *common.Pager) (count int, userCatogery []model.Usercatogery) {
	coutsql := "select count(1) from t_usercatogery t "
	condition := genUserCatogeryCondition(catogeryName)
	if err := o.Raw(coutsql + condition).QueryRow(&count); err != nil || count < 1 {
		//如果查询出错或者查询结果为空返回默认空值
		return
	}

	listsql := "SELECT id,catogeryname,des,createtime,updatetime,isdel from t_usercatogery t "
	if num, err := o.Raw(listsql+condition+common.LIMIT, pager.GetBegin(), pager.GetLen()).QueryRows(&userCatogery); err != nil || num < 1 {
		//如果查询出错返回默认空值
		return
	}
	return
}

/**
查询管理员组分页列表
*/
func (this *userCatogeryService) ListAll() (userCatogery []model.Usercatogery) {

	listsql := "SELECT id,catogeryname,des,createtime,updatetime,isdel from t_usercatogery t where isdel = 1"
	if num, err := o.Raw(listsql).QueryRows(&userCatogery); err != nil || num < 1 {
		//如果查询出错返回默认空值
		return
	}
	return
}

func genUserCatogeryCondition(catogeryName string) (condition string) {
	condition = " where t.isdel = 1 "
	if catogeryName != "" {
		condition += " and t.catogeryname = " + catogeryName
	}
	return
}

/**
添加管理员组
*/
func (this *userCatogeryService) AddUserCatogery(usercatogery *model.Usercatogery) error {
	id, err := o.Insert(usercatogery)
	if err != nil || id < 1 {
		return &common.BizError{"添加失败"}
	}
	// flag := false
	// idArray := strings.Split(ids, ",")
	// for _, roleId := range idArray {
	// 	beego.Debug("给ID为", id, "的管理员组添加", roleId, "权限")
	// 	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	// 	if err != nil {
	// 		beego.Warn(roleId, "不是数字")
	// 		flag = true
	// 		continue
	// 	}
	// 	groupRoleRel := &model.GroupRoleRel{
	// 		Groupid: id,
	// 		Roleid:  roleIdInt,
	// 		Isdel:   1}
	// 	if _, err := o.Insert(groupRoleRel); err != nil {
	// 		beego.Warn("给ID为", id, "的管理员组添加", groupRoleRel.Roleid, "权限失败")
	// 		flag = true
	// 		continue
	// 	}
	// }
	// if flag {
	// 	return &common.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
	// }
	return nil
}

/**
修改管理员组
*/
func (this *userCatogeryService) Modifyusercatogery(usercatogery *model.Usercatogery) error {
	//修改基础信息
	if _, err := o.Update(usercatogery); err != nil {
		beego.Warn("update usercatogery db error.", err.Error())
		return &common.BizError{"修改失败"}
	}

	id := usercatogery.Id
	//删除当前组关联的所有权限
	delsql := "update t_group_role_rel t set t.isdel = 0 where t.groupid = ? and t.isdel =1"
	if _, err := o.Raw(delsql, id).Exec(); err != nil {
		beego.Warn("del group's role fail.", err.Error())
		return &common.BizError{"修改失败"}
	}

	//重新添加权限
	// flag := false
	// idArray := strings.Split(ids, ",")
	// for _, roleId := range idArray {
	// 	beego.Debug("给ID为", id, "的管理员组添加", roleId, "权限")
	// 	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	// 	if err != nil {
	// 		beego.Warn(roleId, "不是数字")
	// 		flag = true
	// 		continue
	// 	}
	// 	groupRoleRel := &model.GroupRoleRel{
	// 		Groupid: id,
	// 		Roleid:  roleIdInt,
	// 		Isdel:   1}
	// 	if _, err := o.Insert(groupRoleRel); err != nil {
	// 		beego.Warn("给ID为", id, "的管理员组添加", groupRoleRel.Roleid, "权限失败", err.Error())
	// 		flag = true
	// 		continue
	// 	}
	// }
	// if flag {
	// 	return &common.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
	// }
	return nil
}

/**
删除管理员组
*/
func (this *userCatogeryService) Delete(ids string) error {
	delsql := "update t_usercatogery t set t.isdel = 0 where t.id in (" + ids + ")"
	if _, err := o.Raw(delsql).Exec(); err != nil {
		beego.Warn("delete fail id:", ids, err.Error())
		return &common.BizError{"删除失败"}
	}

	//删除当前组关联的所有权限
	delrolesql := "update t_group_role_rel t set t.isdel = 0 where t.groupid in (" + ids + ") and t.isdel =1"
	if _, err := o.Raw(delrolesql).Exec(); err != nil {
		beego.Warn("del group's role fail.", err.Error())
		return &common.BizError{"删除失败"}
	}
	return nil
}

/**
根据ID获取管理员组信息
*/
func (this *userCatogeryService) GetUserCatogeryById(id int64) model.Usercatogery {
	usercatogery := model.Usercatogery{Id: id}
	if err := o.Read(&usercatogery); err != nil {
		return model.Usercatogery{}
	}
	return usercatogery
}

/**
根据管理员组ID获取所有的权限列表
*/
func (this *userCatogeryService) GetAllRoleByGroupId(id int64) map[int64]bool {
	var list orm.ParamsList
	num, err := o.Raw("SELECT roleid from t_group_role_rel t where t.groupid = ? and t.isdel =1", id).ValuesFlat(&list)
	if err != nil || num < 1 {
		return nil
	}
	roleIdMap := make(map[int64]bool, len(list))
	for i := 0; i < len(list); i++ {
		idStr := list[i].(string)
		id, _ := strconv.ParseInt(idStr, 10, 64)
		roleIdMap[id] = true
	}
	return roleIdMap
}
