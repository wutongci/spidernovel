package models
import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type Book struct{
	Id int
	Name string
	Author string
	Intro string
	Image string
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
		book.Intro = books[1]
		book.Author = books[2]
		if _, err := o.Update(&book); err == nil {
			fmt.Println("更新书本成功！")
		}
	}
}
