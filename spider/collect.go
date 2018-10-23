package spider

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"spidernovel/models"
	"strconv"
	"sync"
	"time"
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
	}
	ChapterContent struct {
		BookId  int
		Title   string
		Content string
		Url     string
		Sort    int
		Pre     int
		Next    int
	}
)

func GetBook(){
	fmt.Println("spider start")
	books, _ := models.GetBookList("status", 0)
	for _, book := range books{
		s := spiderBookInfos(book.Url)
		domainUrls, err := url.Parse(book.Url)
		if err != nil {
			log.Fatal(err)
		}
		domainUrl := domainUrls.Scheme+"://"+domainUrls.Host+"/"
		chapters := spiderChapterUrl(book.Url)
		lastchapterids := models.GetLastChapterIds(book.Id)
		limitChan := make(chan int, 10)
		cacheChan := make(chan *ChapterContent, 10)
		var wg sync.WaitGroup
		for _,chapter := range chapters{
			if len(chapter.URI) < 18 {//较短的网址过滤
				continue
			}
			url := domainUrl+chapter.URI
			reg := regexp.MustCompile(`[\d]+`)
			ids := reg.FindAllString(url, -1)
			curid, err := strconv.Atoi(ids[1])
			if err != nil {
				log.Fatal(err)
			}

			if _,ok := lastchapterids[curid]; ok {
				fmt.Println(curid )
				continue
			}
			limitChan <- 1
			wg.Add(1)
			go SpiderProcessor(book.Id,chapter.Name,curid,cacheChan, url, &wg)
			wg.Add(1)
			go WriterProcessor(book.Id,cacheChan, limitChan, &wg)
		}
		wg.Wait()
		models.UpdateBookInfo(book.Id,s)
	}
}

func SpiderProcessor(bookid int,title string,curid int,cacheChan chan *ChapterContent, url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	content := getContents(url)
	if content == ""{
		return nil
	}
	cacheChan <- &ChapterContent{
		BookId:bookid,
		Title:title,
		Content:content,
		Url:url,
		Sort:curid,
		Pre:curid -1,
		Next:curid +1,
	}
	return nil
}
func WriterProcessor(bookid int,cacheChan chan *ChapterContent, limitChan chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	bodyContent := <-cacheChan
	ch := models.Chapter{BookId:bookid, Title:bodyContent.Title, Content:bodyContent.Content,Volume:" ",Status:0,Sort:bodyContent.Sort,Url: bodyContent.Url,Pre:bodyContent.Pre, Next:bodyContent.Next,CreateTime:time.Now(),LastUpdateTime:time.Now()}
	models.ChapterAdd(&ch)
	<-limitChan
	return nil
}
