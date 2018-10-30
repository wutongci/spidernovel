package spider

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
	"spidernovel/models"
)

func GetHtmlDocument(url string, charset string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalln("Http错误码:%d", res.StatusCode)
	}
	responseString, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if charset == "gbk" {
		decoder := mahonia.NewDecoder("GB18030")
		decodedString := decoder.ConvertString(string(responseString))
		responseString = []byte(decodedString)
	}
	decodeReader := bytes.NewReader(responseString)
	doc, err := goquery.NewDocumentFromReader(decodeReader)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return doc, err
}


func SpiderProcessor(bookid int,title string,content string,curid int,cacheChan chan *ChapterContent, url string, wg *sync.WaitGroup) error {
	defer wg.Done()
	cacheChan <- &ChapterContent{
		BookId:bookid,
		Title:title,
		Content:content,
		Url:url,
		Sort:curid,
	}
	return nil
}
func WriterProcessor(bookid int,cacheChan chan *ChapterContent, limitChan chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	bodyContent := <-cacheChan
	ch := models.Chapter{BookId:bookid, Title:bodyContent.Title, Content:bodyContent.Content,Volume:" ",Status:1,Sort:bodyContent.Sort,Url: bodyContent.Url,CreateTime:time.Now(),LastUpdateTime:time.Now()}
	models.ChapterAdd(&ch)
	<-limitChan
	return nil
}
