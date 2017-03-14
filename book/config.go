package main

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

//var Root string = "/Users/mebius/Documents/story/"

func getRoot() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := filepath.Abs(file)

	pp := strings.Replace(dir, "\\", "/", -1)
	pp = pp + "_c/book.config"
	b, err := ioutil.ReadFile(pp)
	if err != nil {
		fmt.Println(pp, os.Args[0])
		panic(err)
	}
	return string(b)
}

//var barText string = " 目录:[m](fg-red) 跳转:[g](fg-red) 上一行/上翻页:[k](fg-red) 下一行/下翻页:[j](fg-red) 信息:[i](fg-red) 帮助:[h](fg-red) 确认:[enter](fg-red) 返回:[b](fg-red) 老板键:[o](fg-red) 退出:[q](fg-red) "
//var barText string = " 目录:[m](fg-red) 上一行/上翻页:[k](fg-red) 下一行/下翻页:[j](fg-red) 信息:[i](fg-red) 确认:[enter](fg-red) 老板键:[o](fg-red) 退出:[q](fg-red) "
var barText string = " 目录:[m](fg-red) 上一行/上翻页:[k](fg-red) 下一行/下翻页:[j](fg-red) 确认:[enter](fg-red) 退出:[q](fg-red) "

func getHomePath() string {
	user, _ := user.Current()
	return user.HomeDir + "/.book"
}

var config Config

//var cc Config

func ReadConfig() {
	_, err := toml.DecodeFile(getHomePath(), &config)
	//fmt.Println("...", err)
	if err != nil {
		//没有文件
		//fmt.Println(err)
		//panic(err)
		//config = new(Config)
		//config.Books = make([]Bookread, 0)
		config = Config{MD5: make([]string, 0), BOOKC: make([]int, 0)}
	}
	//fmt.Println(cc.BOOKC)
}

func WriteConfig() {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(config)
	if err != nil {
		//转化失败
		panic(err)
	}
	//fmt.Println("data:", buf.String())
	ioutil.WriteFile(getHomePath(), buf.Bytes(), 0777)
}

func setBookData(md5 string, _c int) {
	var rel bool = true
	for k, v := range config.MD5 {
		if v == md5 {
			config.BOOKC[k] = _c
			rel = false
			//fmt.Println(_c)
			//return
		}
	}
	if rel == true {
		//config.Books = append(config.Books, Bookread{MD5: md5, BOOKC: _c})
		config.MD5 = append(config.MD5, md5)
		config.BOOKC = append(config.BOOKC, _c)
	}
}
func getBookC(md5 string) int {
	//return 30
	for k, v := range config.MD5 {
		if v == md5 {

			return config.BOOKC[k]
		}
	}
	return -1
}

type Config struct {
	//Books []Bookread
	MD5   []string
	BOOKC []int
}

type Bookread struct {
	MD5   string
	BOOKC int
}
