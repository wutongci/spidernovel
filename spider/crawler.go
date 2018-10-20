package spider

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"log"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"strings"
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
	return doc, err
}

/**
	采集小说章节网址
 */
func spiderChapterUrl(url string)(chapters []Chapter){
	doc, err := GetHtmlDocument(url, "uft-8")
	if err != nil {
		log.Fatalln(err.Error())
	}
	doc.Find("#list").First().Find("dd").Each(func(i int, s *goquery.Selection) {
		chapter := Chapter{}
		chapter.Index = i
		chapter.Name = strings.Replace(s.Text(),"?","",-1)
		chapter.URI, _ = s.Find("a").First().Attr("href")
		chapters = append(chapters, chapter)
	})
	return chapters
}

/**
采集章节内容
 */
func getContents(contentUrl string) string {
	doc, err := GetHtmlDocument(contentUrl, "uft-8")
	if err != nil {
		log.Fatalln(err.Error())
	}
	title := doc.Find(".bookname > h1").Text()
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
	if len(title)== 0 {
		return ""
	}

	return title+"==="+title + "\r\n"+ strings.Join(contents, "")
}
