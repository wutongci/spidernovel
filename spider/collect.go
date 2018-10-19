package spider

import (
	"spidernovel/models"
	"fmt"
	"time"
	"net/url"
	"log"
	"sync"
	"strings"
	"regexp"
	"strconv"
)

type (
	Chapter struct {
		Index int
		Name  string
		URI   string
	}
)

func GetBook(){
	fmt.Println("spider start")
	books, _ := models.GetBookList("status", 0)
	for _, book := range books{
		domainUrls, err := url.Parse(book.Url)
		if err != nil {
			log.Fatal(err)
		}
		domainUrl := domainUrls.Scheme+"://"+domainUrls.Host+"/"
		chapters := spiderChapterUrl(book.Url)

		limitChan := make(chan int, 10)
		cacheChan := make(chan string, 10)
		var wg sync.WaitGroup
		for _,chapter := range chapters{
			if len(chapter.URI) < 18 {//较短的网址过滤
				continue
			}
			limitChan <- 1
			wg.Add(1)
			go SpiderProcessor(cacheChan, domainUrl+chapter.URI, &wg)
			wg.Add(1)
			go WriterProcessor(book.Id,chapter,domainUrl+chapter.URI,cacheChan, limitChan, &wg)
		}
		wg.Wait()
	}
}

func SpiderProcessor(cacheChan chan string, url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	content := getContents(url)
	if content == ""{
		return nil
	}
	cacheChan <-content
	return nil
}
func WriterProcessor(bookid int,chapter Chapter ,url string,cacheChan chan string, limitChan chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	bodyContent := <-cacheChan
	reg := regexp.MustCompile(`[\d]+`)
	ids := reg.FindAllString(url, -1)
	curid, err := strconv.Atoi(ids[1])
	if err != nil {
		log.Fatal(err)
	}
	sort := curid
	pre := curid -1
	next := curid +1
	contents := strings.Split(bodyContent,"===")
	ch := models.Chapter{BookId:bookid, Title:contents[0], Content:contents[1],Volume:" ",Status:0,Sort:sort,Url: url,Pre:pre, Next:next,CreateTime:time.Now(),LastUpdateTime:time.Now()}
	models.ChapterAdd(&ch)
	<-limitChan
	return nil
}
