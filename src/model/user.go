package model

import "time"

type User struct {
	Id         int64     `json:"id"`
	Accout     string    `json:"accout",orm:"unique"`
	Password   string    `json:"password"`
	Mail       string    `json:"mail"`
	Name       string    `json:"name"`
	Catogery   string    `json:"catogery"`
	Phone      string    `json:"phone"`
	Ischarge   string    `json:"ischarge"`
	Createtime time.Time `json:"createtime",orm:"auto_now_add;type(datetime)"`
	Updatetime time.Time `json:"updatetime",orm:"auto_now;type(datetime)"` // 更新时间
	Isdel      int8      `json:"isdel",orm:"default(1)"`
}
