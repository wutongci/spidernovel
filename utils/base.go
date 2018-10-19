package utils

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"bufio"
	"strings"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)

func init() {
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stdout,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func CheckErr(err error) {
	if err != nil {
		Error.Fatalln(err.Error())
	}
}

func CreateFile(filePath string)  error  {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		if err ==nil {
			fmt.Printf("mkdir success!\n")
			return nil
		}
		fmt.Printf("mkdir failed![%v]\n", err)
		os.Exit(0)
		return err
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, "\r")
	return input
}