package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"spidernovel/models"
	"strconv"
	"strings"
	"sync"
	"net/url"
)

type BiqugeSpider struct {

}

func (self *BiqugeSpider)SpiderBook(bookurl string,bookid int)(error){
	s := spiderBookInfos(bookurl)
	domainUrls, err := url.Parse(bookurl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	domainUrl := domainUrls.Scheme+"://"+domainUrls.Host
	chapters := spiderChapterUrl(bookurl)
	lastchapterids := models.GetLastChapterIds(bookid)
	limitChan := make(chan int, 10)
	cacheChan := make(chan *ChapterContent, 10)
	var wg sync.WaitGroup
	for _,chapter := range chapters{
		url := domainUrl+chapter.URI
		if _,ok := lastchapterids[chapter.Sort]; ok {
			fmt.Println(chapter.Sort)
			continue
		}
		limitChan <- 1
		wg.Add(1)
		content := getContents(url)
		if content == ""{
			return nil
		}
		go SpiderProcessor(bookid,chapter.Name,content,chapter.Sort,cacheChan, url, &wg)
		wg.Add(1)
		go WriterProcessor(bookid,cacheChan, limitChan, &wg)
	}
	wg.Wait()
	models.UpdateBookInfo(bookid,s)
	return nil
}

/**
   采集书本基本介绍
 */

func spiderBookInfos(url string) (book []string)  {
	doc,_ := GetHtmlDocument(url, "uft-8")
	doc.Find("#info").First().Find("h1").Each(func(i int, s *goquery.Selection) {
		ss := make([]string,3)
		ss[0] = s.Text()
		ss[1] = strings.Replace(s.Parent().Find("p").First().Text(),"作  者：","",-1)
		ss[2] = strings.TrimSpace(doc.Find("#intro").First().Text())
		book = ss
	})
	return book
}

/**
	采集小说章节网址
 */
func spiderChapterUrl(url string)(chapters []Chapter){
	doc, _ := GetHtmlDocument(url, "uft-8")
	doc.Find("#list").First().Find("dd").Each(func(i int, s *goquery.Selection){
		chapter := Chapter{}
		chapter.Index = i
		chapter.Name = strings.Replace(s.Text(),"?","",-1)
		chapter.URI, _ = s.Find("a").First().Attr("href")
		if len(chapter.URI) > 18 {//较短的网址过滤
			reg := regexp.MustCompile(`[\d]+`)
			ids := reg.FindAllString(chapter.URI, -1)
			curid, _ := strconv.Atoi(ids[1])
			chapter.Sort = curid
			chapters = append(chapters, chapter)
		}
	})
	return chapters
}

/**
	采集章节内容
 */
func getContents(contentUrl string) string {
	doc, _ := GetHtmlDocument(contentUrl, "uft-8")
	title := doc.Find(".bookname > h1").Text()
	if len(title)== 0 {
		return ""
	}
	var contents []string
	doc.Find("#content").First().Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		html = strings.Replace(html,"\t","",-1)
		for _, value := range strings.Split(html, "<br/>　　<br/>") {
			content := strings.Replace(value,"&nbsp;"," ",4)+ "\r\n"
			content = strings.Replace(content,"<br/>","",-1)
			content = strings.Replace(content, "Ps:书友们，我是未苍，推荐一款免费小说App，支持小说下载、听书、零广告、多种阅读模式。请您关注微信公众号：dazhuzaiyuedu（长按三秒复制）书友们快关注起来吧！","",-1)
			content = strings.Replace(content,"<script>chaptererror();</script>", " ", -1)
			contents = append(contents, content)
		}
	})
	return title + "\r\n"+ strings.Join(contents, "")
}