package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"log"
	"spidernovel/components"
	"strconv"
	"time"
)

type Chapter struct{
	Id int
	BookId int
	Title string
	Content string
	Volume string
    Url string
	Sort int
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

func GetAllConents(bookid int) (lists []interface{})  {
	var res []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("SELECT content FROM chapter WHERE book_id =? ORDER BY sort ASC",bookid).Values(&res)

	if err == nil && num > 0 {
		for _, v := range res {
			lists = append(lists, v)
		}
		return lists
	}
	return  nil
}

func GetLastChapterIds(bookid int) (chapterids map[int]int){
	var maps []orm.Params
	res := make(map[int]int)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT sort FROM chapter WHERE book_id =?",bookid).Values(&maps)
	if err == nil && num > 0 {
		for _,vv := range maps {
			lastchapterid, err := strconv.Atoi(vv["sort"].(string))
			if err != nil {
				log.Fatal(err)
			}
			res[lastchapterid] = lastchapterid
		}
		return res
	}
	return  nil
}