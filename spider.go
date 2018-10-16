package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"strings"
)

type (
	Chapter struct {
		Index int
		Name  string
		URI   string
	}
)
func getChapters(charterUrl string) []Chapter {
	doc, err := getHtmlDocument(charterUrl, "uft-8")
	CheckErr(err)
	chapters := []Chapter{}
	doc.Find("#list").First().Find("dd").Each(func(i int, s *goquery.Selection) {
		chapter := Chapter{}
		chapter.Index = i
		chapter.Name = strings.Replace(s.Text(),"?","",-1)
		chapter.URI, _ = s.Find("a").First().Attr("href")
		chapters = append(chapters, chapter)
	})
	return chapters
}

func getContents(contentUrl string) string {
	doc, err := getHtmlDocument(contentUrl, "uft-8")
	CheckErr(err)
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

func getHtmlDocument(url string, charset string) (*goquery.Document, error) {
	res, err := http.Get(url)
	CheckErr(err)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		Error.Fatalln("Http错误码:%d", res.StatusCode)
	}
	responseString, err := ioutil.ReadAll(res.Body)
	CheckErr(err)
	if charset == "gbk" {
		decoder := mahonia.NewDecoder("GB18030")
		decodedString := decoder.ConvertString(string(responseString))
		responseString = []byte(decodedString)
	}
	decodeReader := bytes.NewReader(responseString)
	doc, err := goquery.NewDocumentFromReader(decodeReader)
	return doc, err
}
