package model

import "time"

type Usercatogery struct {
	Id           int64     `json:"id"`
	Catogeryname string    `json:"catogeryname"`
	Des          string    `json:"describe"`
	Createtime   time.Time `json:"createtime",orm:"auto_now_add;type(datetime)"`
	Updatetime   time.Time `json:"updatetime",orm:"auto_now;type(datetime)"`
	Isdel        int8      `json:"isdel",orm:"default(1)"`
}

type Usercatogerycheck struct {
	Id           int64     `json:"id"`
	Catogeryname string    `json:"catogeryname"`
	Des          string    `json:"describe"`
	Createtime   time.Time `json:"createtime",orm:"auto_now_add;type(datetime)"`
	Updatetime   time.Time `json:"updatetime",orm:"auto_now;type(datetime)"`
	Isdel        int8      `json:"isdel",orm:"default(1)"`
	Check        bool      `json:"check"`
}
