package main

import "os"

var BooklistData []string
var BookInfos []os.FileInfo

var BookListNum int = 0
var curBookName string = ""
var curBookIndex int = 0

var Scene int = MenuScene

const (
	MenuScene int = iota
	BookScene
	InfoScene
	BossScene
)

var bookmd5 string = ""
var bookc int = 0
