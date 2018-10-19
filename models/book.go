package models
import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Book struct{
	Id int
	Name string
	Author string
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
