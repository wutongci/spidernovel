package main

import (
	"spidernovel/spider"
	"os"
	"fmt"
	"regexp"
	"bufio"
	"spidernovel/utils"
	"spidernovel/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func main()  {
	startSpider()
}
func startSpider()  {
	fmt.Println("1. 下载小说")
	fmt.Println("2. 合并小说")
	fmt.Println("3. 退出")
	command := utils.GetInput()
	switch command {
	case "1":
		spider.GetBook()
		fmt.Println("下载完成!")
	case "2":
		BookidCommand := utils.GetInput()
		reg := regexp.MustCompile(`^[0-9]*$`)
		str := reg.FindAllString(BookidCommand, -1)
		if len(str)==0 {
			fmt.Println("err num")
			os.Exit(0)
		}
		bookid,_ := strconv.Atoi(BookidCommand)
		mergeNovels(bookid)
		fmt.Println("合并完成!")
	case "3":
		fmt.Println("再见!")
		os.Exit(0)
	default:
		fmt.Println("错误的指令")
		fmt.Println()
	}
}

func mergeNovels(bookid int)  {
	utils.CreateFile("mergeoutput")
	bookstr := strconv.Itoa(bookid)
	outFileName := "./mergeoutput/merge_result_"+bookstr+".txt"
	outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)
	if openErr != nil {
		outFile, _ := os.Create(outFileName)
		defer outFile.Close()
	}
	bWriter := bufio.NewWriter(outFile)
	lists := models.GetAllConents(bookid)
	for _,values := range lists {
		for _,content := range values.(orm.Params){
			bWriter.WriteString(content.(string))
		}

	}
	bWriter.Flush()
}