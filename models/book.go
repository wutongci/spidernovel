package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Book struct{
	Id int
	Name string
	Author string
	Intro string
	Image string
	From  int
	Status int
	Url string
	CreateTime time.Time
	LastUpdateTime time.Time
}

func init()  {
	orm.RegisterModel(new(Book))
}
func GetBookList(filters ...interface{})([]*Book, int64){
	books := make([]*Book, 0)
	query := orm.NewOrm().QueryTable("book")
	if len(filters) > 0{
		l := len(filters)
		for i := 0; i < l; i += 2{
			query = query.Filter(filters[i].(string), filters[i + 1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").All(&books)
	return books, total
}

func UpdateBookInfo(bookid int,books []string){
	o := orm.NewOrm()
	book := Book{Id: bookid}
	if o.Read(&book) == nil {
		book.Status = 1
		book.Name = books[0]
		book.Author = books[1]
		book.Intro = books[2]
		if _, err := o.Update(&book); err == nil {
			fmt.Println("更新书本成功！")
		}
	}
}
