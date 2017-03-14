package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getFiles() {
	var num int = 0
	BooklistData = make([]string, 0)
	BookInfos = make([]os.FileInfo, 0)
	filepath.Walk(getRoot(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
		}
		if info.IsDir() == false {
			num++
			BooklistData = append(BooklistData, "["+strconv.Itoa(num)+"] "+info.Name())
			BookInfos = append(BookInfos, info)
		}
		return err
	})

	BookListNum = len(BooklistData)
	if BookListNum > 0 {
		curBookName = BooklistData[0]
	}
}

//打开图书
// 1. 取图书信息
// 2. 读取文件流，整体读取
// 3. 转化string
// 4. 赋值给 BookScene
func openBook() {
	b, err := ioutil.ReadFile(getRoot() + BookInfos[curBookIndex].Name())
	if err != nil {
		panic(err)
	}
	str := string(b)
	bookmd5 = BytesToString(md5.Sum(b))
	bookc = getBookC(bookmd5)
	if bookc == -1 {
		bookc = 0
	}
	updateConfig()
	//setBookData(bookmd5, bookc)
	initBookScene(BookInfos[curBookIndex].Name(), str)
}

func updateConfig() {
	//ReadConfig()

	setBookData(bookmd5, bookc)
	WriteConfig()
}

func BytesToString(bs [16]byte) string {
	l := len(bs)
	buf := make([]string, 0, l)
	for i := 0; i < l; i++ {
		buf = appendString(buf, bs[i])
	}
	return strings.Join(buf, "")
}
func appendString(bs []string, b byte) []string {
	var a byte
	var s int
	for i := 0; i < 8; i++ {
		a = b
		b <<= 1
		b >>= 1
		switch a {
		case b:
			s += 0
		default:
			temp := 1
			for j := 0; j < 7-i; j++ {
				temp = temp * 2
			}
			s += temp
		}

		b <<= 1
	}

	return append(bs, strconv.Itoa(s))
}
