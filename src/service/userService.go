package service

import (
	"github.com/BitAssetManagement/cms/src/common"
	"github.com/BitAssetManagement/cms/src/model"
	"strings"
	"time"

	// "strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type userService struct{}

/**
分页查询管理员列表*
*/
func (this *userService) Gridlist(pager *common.Pager, userid, usermail, username, userphone, accout string) (count int, users []model.User) {
	countsql := "select count(1) from t_user t "
	condition := genUserCondition(userid, usermail, username, userphone, accout)
	if err := o.Raw(countsql + condition).QueryRow(&count); err != nil || count < 1 {
		beego.Debug("select user count err or result is null.")
		return
	}

	listsql := "select id,accout,mail,name,phone,catogery,password,createtime,updatetime,isdel from t_user t "
	if _, err := o.Raw(listsql+condition+common.LIMIT, pager.GetBegin(), pager.GetLen()).QueryRows(&users); err != nil {
		beego.Warn("select userList from db error.")
		return
	}
	return
}

/**
按照参数拼接sql查询条件
*/
func genUserCondition(userid, usermail, username, userphone, accout string) (condition string) {
	condition = " where t.isdel = 1 "
	if !strings.EqualFold(userid, "") {
		condition += " and t.id = " + userid + "'"
	}
	if !strings.EqualFold(usermail, "") {
		condition += " and t.mail = '" + usermail + "'"
	}
	if !strings.EqualFold(username, "") {
		condition += " and t.name =  '" + username + "'"
	}
	if !strings.EqualFold(userphone, "") {
		condition += " and t.phone =  '" + userphone + "'"
	}
	if !strings.EqualFold(accout, "") {
		condition += " and t.accout =  '" + accout + "'"
	}
	beego.Debug("condition is : ", condition)
	return
}

/**
添加管理员
// */

// func (this *userService) AddUser(user *model.User, groupIds string) error {
// 	flag := false
// 	if userId, err := o.Insert(user); err != nil {
// 		beego.Warn("insert user fail, user:", user, err.Error())
// 		return &common.BizError{"添加失败,账号已经存在"}
// 	} else {
// 		idArray := strings.Split(groupIds, ",")
// 		for _, gid := range idArray {
// 			gidint, err := strconv.ParseInt(gid, 10, 64)
// 			if err != nil {
// 				beego.Debug("id 转换成数字异常，id：", gid)
// 				flag = true
// 			}
// 			rel := model.UserCatogeryRel{
// 				Userid:     userId,
// 				Catogeryid: gidint,
// 				Isdel:      1}
// 			if _, err := o.Insert(&rel); err != nil {
// 				flag = true
// 			}
// 		}
// 	}
// 	if flag {
// 		return &common.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
// 	}
// 	return nil
// }

func (this *userService) AddUser(user *model.User) error {
	if _, err := o.Insert(user); err != nil {
		beego.Warn("insert user fail, user:", user, err.Error())
		return &common.BizError{"添加失败,账号已经存在"}
	}
	return nil
}

// func (this *userService) AddUser(user *model.User) error {
// 	flag := false
// 	if userId, err := o.Insert(user); err != nil {
// 		beego.Warn("insert user fail, user:", user, err.Error())
// 		return &common.BizError{"添加失败,账号已经存在"}
// 	}
// 	if flag {
// 		return &common.BizError{"出现异常，部分权限添加失败，请补充添加权限。"}
// 	}
// 	return nil
// }

/**
修改管理员
*/
// func (this *userService) ModifyUser(user *model.User, groupIds string) error {
// 	flag := false
// 	updateSql := "UPDATE user SET "

// 	set := updateUserSet(user)
// 	condition := " where id = ? "

// 	// if _, err := o.Raw(updateSql, user.Accout, user.Mail, user.Name, user.Phone, user.Department, time.Now(), user.Id).Exec(); err != nil {
// 	id := user.Id
// 	if _, err := o.Raw(updateSql+set+condition, id).Exec(); err != nil {
// 		beego.Warn("update user fail, user:", user, err.Error())
// 		return &common.BizError{"修改失败"}
// 	} else {
// 		//逻辑删除所有用户和组关联关系UserGroupRel
// 		delRelSql := "update user_catogery_rel set isdel = 0 where userid = ?"
// 		if _, err := o.Raw(delRelSql, user.Id).Exec(); err != nil {
// 			return &common.BizError{"修改失败"}
// 		}

// 		idArray := strings.Split(groupIds, ",")
// 		//重新添加关联关系
// 		for _, gid := range idArray {
// 			gidint, err := strconv.ParseInt(gid, 10, 64)
// 			if err != nil {
// 				beego.Debug("id 转换成数字异常，id：", gid)
// 				flag = true
// 			}
// 			rel := model.UserCatogeryRel{
// 				Userid:     user.Id,
// 				Catogeryid: gidint,
// 				Isdel:      1}
// 			if _, err := o.Insert(&rel); err != nil {
// 				beego.Warn("添加组关系失败", rel, err.Error())
// 				flag = true
// 			}
// 		}
// 	}
// 	if flag {
// 		return &common.BizError{"出现异常，部分权限修改失败，请补充添加权限。"}
// 	}

// 	return nil
// }

func (this *userService) ModifyUser(user *model.User) error {
	updateSql := "UPDATE t_user SET "

	set := updateUserSet(user)
	condition := " where id = ? "

	// if _, err := o.Raw(updateSql, user.Accout, user.Mail, user.Name, user.Phone, user.Department, time.Now(), user.Id).Exec(); err != nil {
	id := user.Id
	if _, err := o.Raw(updateSql+set+condition, id).Exec(); err != nil {
		beego.Warn("update user fail, user:", user, err.Error())
		return &common.BizError{"修改失败"}
	}

	return nil
}

func updateUserSet(user *model.User) string {
	set := ""
	if !strings.EqualFold(user.Password, "") {
		set += " password = '" + user.Password + "',"
	}
	if !strings.EqualFold(user.Accout, "") {
		set += " accout = '" + user.Accout + "',"
	}
	if !strings.EqualFold(user.Mail, "") {
		set += " mail = '" + user.Mail + "',"
	}
	if !strings.EqualFold(user.Name, "") {
		set += " name = '" + user.Name + "',"
	}
	if !strings.EqualFold(user.Phone, "") {
		set += " phone = '" + user.Phone + "',"
	}
	if !strings.EqualFold(user.Catogery, "") {
		set += " Catogery = '" + user.Catogery + "',"
	}
	set += " updatetime = '" + time.Now().Format("2006-01-02 15:04:05") + "'"

	return set
}

/**
删除管理员基本信息
*/
func (this *userService) Delete(userids string) error {
	delUserSql := "update t_user set isdel = 0 where id in (" + userids + ")"
	if _, err := o.Raw(delUserSql).Exec(); err != nil {
		return &common.BizError{"删除管理员基本信息失败"}
	}
	// delRelSql := "update user_group_rel set isdel = 0 where userid in (" + userids + ")"
	// if _, err := o.Raw(delRelSql).Exec(); err != nil {
	// 	return &common.BizError{"删除管理员和组关系失败"}
	// }
	return nil
}

/**
登陆鉴权
*/
// func (this *userService) Authentication(accout, encodePwd string) (user *model.User, err error) {
// 	selectSql := "select id,password from user t where t.accout = '" + accout + "' and isdel =1"
// 	if err := o.Raw(selectSql).QueryRow(&user); err != nil {
// 		if err == orm.ErrNoRows {
// 			return nil, &common.BizError{"账号不存在"}
// 		}
// 		return nil, &common.BizError{"登陆失败，请稍后重试"}
// 	}
// 	if !strings.EqualFold(encodePwd, user.Password) {
// 		return nil, &common.BizError{"密码错误"}
// 	}
// 	return user, nil
// }

/**
根据ID查询管理员
*/
func (this *userService) GetUserById(id int64) (user *model.User, err error) {
	user = &model.User{Id: id}
	if err := o.Read(user); err != nil {
		if err == orm.ErrNoRows {
			err = &common.BizError{"账号不存在"}
			return nil, err
		}
		err = &common.BizError{"系统错误"}
		return nil, err
	}
	return user, nil
}

// func (this *userService) GetAllCheckGroup(id int64) map[int64]bool {
// 	var list orm.ParamsList
// 	num, err := o.Raw("SELECT catogery from user_catogery_rel t where isdel=1 and t.userid = ?", id).ValuesFlat(&list)
// 	if err != nil || num < 1 {
// 		return nil
// 	}
// 	roleIdMap := make(map[int64]bool, len(list))
// 	for i := 0; i < len(list); i++ {
// 		idStr := list[i].(string)
// 		id, _ := strconv.ParseInt(idStr, 10, 64)
// 		roleIdMap[id] = true
// 	}
// 	return roleIdMap
// }
