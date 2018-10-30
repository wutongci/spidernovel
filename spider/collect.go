package spider

import (
	"fmt"
	"github.com/pkg/errors"
	"spidernovel/models"
)

type (
	Book struct {
		Name string
		Author string
		Intro  string
	}
	Chapter struct {
		Index int
		Name  string
		URI   string
		Sort  int
	}
	ChapterContent struct {
		BookId  int
		Title   string
		Content string
		Url     string
		Sort    int
	}
)

type Spider interface{
	SpiderBook(url string,bookid int) error
}

func NewSpider(from int) (Spider, error){
	switch from{
	case 1:
		return new(BiqugeSpider),nil
	default:
		return nil, errors.New("暂时不支持该站点")
	}
}

func GetBook(){
	fmt.Println("spider start")
	books, _ := models.GetBookList("status", 0)
	for _, book := range books{
		c,err:= NewSpider(book.From)
		if err != nil{
			continue
		}
		c.SpiderBook(book.Url,book.Id)
	}
	fmt.Println("完成下载更新所有书本")
}
