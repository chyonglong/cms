package model

type UserCatogeryRel struct {
	Id         int64
	Userid     int64
	Catogeryid int64
	Isdel      int8 `orm:"default(1)"`
}
