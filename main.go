package main

import (
	"sync"
	"bufio"
	"os"
	"strings"
	"fmt"
	"io"
	"path/filepath"
	"regexp"
)

func main()  {
	fmt.Println("1. 下载小说")
	fmt.Println("2. 合并小说")
	fmt.Println("3. 退出")
	command := getInput()
	switch command {
	case "1":
		BookidCommand := getInput()
		reg := regexp.MustCompile(`^[0-9]*$`)
		str :=reg.FindAllString(BookidCommand, -1)
		if len(str)==0 {
			fmt.Println("err num")
			os.Exit(0)
		}
		getNovel(BookidCommand)
		fmt.Println("下载完成!")
	case "2":
		BookidCommand := getInput()
		reg := regexp.MustCompile(`^[0-9]*$`)
		str :=reg.FindAllString(BookidCommand, -1)
		if len(str)==0 {
			fmt.Println("err num")
			os.Exit(0)
		}
		mergeNovel(BookidCommand)
		fmt.Println("合并完成!")
	case "3":
		fmt.Println("再见!")
		os.Exit(0)
	default:
		fmt.Println("错误的指令")
		fmt.Println()
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, "\r")
	return input
}

/**
  只支持qu.la小说站下载，例如 https://www.qu.la/book/3987/
 */
func getNovel(bookid string){
	url := "https://www.qu.la/book/"+bookid+"/"
	createFile("./output/"+bookid)
	charpts := getChapters(url)
	limitChan := make(chan int, 10)
	cacheChan := make(chan string, 10)
	var wg sync.WaitGroup
	for _, charpt := range charpts {
		limitChan <- 1
		wg.Add(1)
		go SpiderProcessor(cacheChan, url+charpt.URI, &wg)
		wg.Add(1)
		go WriterProcessor(bookid,cacheChan, limitChan, &wg)
	}
	wg.Wait()
}

func mergeNovel(bookid string)  {
	rootPath := "output/"+bookid
	outFileName := "./mergeoutput/merge_result_"+bookid+".txt"
	outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)
	if openErr != nil {
		outFile, _ := os.Create(outFileName)
		defer outFile.Close()
	}
	bWriter := bufio.NewWriter(outFile)
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		fmt.Println("正在合并:", path)
		if strings.HasSuffix(path, ".txt") {
			fp, fpOpenErr := os.Open(path)
			if fpOpenErr != nil {
				fmt.Printf("该文件不能打开 %v", fpOpenErr)
				return fpOpenErr
			}
			bReader := bufio.NewReader(fp)
			for {
				buffer := make([]byte, 1024)
				readCount, readErr := bReader.Read(buffer)
				if readErr == io.EOF {
					break
				} else {
					bWriter.Write(buffer[:readCount])
				}
			}
		}
		return err
	})
	bWriter.Flush()
}