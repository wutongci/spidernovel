package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"spidernovel/components"
)

type Chapter struct{
	Id int
	BookId int
	Title string
	Content string
	Volume string
    Url string
	Sort int
	Pre int
	Next int
	Status int
	CreateTime time.Time
	LastUpdateTime time.Time
}

func init()  {
	orm.RegisterModel(new(Chapter))
}

func ChapterAdd(chapter *Chapter)(int64, error){
	return orm.NewOrm().Insert(chapter)
}


func GetAllList() (lists []interface{}) {
	var res []orm.Params
	if(components.Cache.IsExist("list")) {
		var imapGet []interface{}
		valueGet := components.Cache.Get("list")
		bytes, _ := valueGet.([]byte)
		json.Unmarshal(bytes,&imapGet)
		return imapGet
	}
	o := orm.NewOrm()
	num, err := o.Raw("SELECT id,title FROM chapter").Values(&res)

	if err == nil && num > 0 {
		for _, v := range res {
			lists = append(lists, v)
		}
		cachelist,_ := json.Marshal(lists)
		components.Cache.Put("list",cachelist,3*time.Minute)
		return lists
	}
	return  nil
}