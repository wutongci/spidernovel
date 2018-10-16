package main

import (
	"log"
	"os"
	"sync"
	"strings"
)

func WriteWithFileWrite(name, content string) {
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Failed to open the file", err.Error())
		return
	}
	defer fileObj.Close()
	if _, err := fileObj.WriteString(content); err == nil {
		return
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
func WriterProcessor(bookid string,cacheChan chan string, limitChan chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	bodyContent := <-cacheChan
	contents := strings.Split(bodyContent,"===")
	filename := strings.Replace(contents[0],"?","",-1)
	if len(filename) == 0 {
		return nil
	}
	name := "output/"+bookid+"/" + filename + ".txt"
	WriteWithFileWrite(name, contents[1])
	<-limitChan
	return nil
}
