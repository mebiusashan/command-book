package main

import (
	//"fmt"
	"github.com/gizak/termui"
	//"os"
	"strconv"
)

var title string = ""

var wen []rune
var wenl int = 0
var context string = ""

//var curCharPrev int = 0
var curCharStart int = 0
var curCharEnd int = 0

func initBookScene(_title string, _context string) {
	curCharStart = 0
	curCharEnd = getBookC(bookmd5)

	if Book == nil {
		Book = termui.NewPar("")
		Book.BorderFg = termui.ColorYellow
		updateBookScene()
	}
	title = _title
	//Book.BorderLabel = title
	context = _context
	wen = []rune(context)
	wenl = len(wen)

	NextPage()
	//updateBookContext()
}

func updateBookContext() {
	var b int = int(float32(curCharStart) / float32(wenl) * 100)

	//Book.Text = strconv.Itoa(b) + ":" + strconv.Itoa(curCharStart) + ":" + strconv.Itoa(curCharEnd)
	Book.Text = string(wen[curCharStart:curCharEnd])
	//var b float32 = float32(curCharStart) / float32(wenl)
	bookc = curCharStart

	//129399420713388218218150176814619124117183
	//updateConfig()
	//setBookData(bookmd5, 60) //curCharStart)
	/*err := os.Remove(getHomePath())
	if err != nil {
		panic(err)
	}*/
	//WriteConfig()strconv.Itoa(bookc) + ":" +
	Book.BorderLabel = "【" + title + "】 [" + strconv.Itoa(b) + "%](fg-red)"
	updateConfig()
}

func NextPage() {
	if curCharEnd < len(wen) {
		curCharStart = curCharEnd
		var linenum = termui.TermHeight() - 3 //可用行数
		var charnum = termui.TermWidth() - 2  //每行可用字符数
		var cchar int = charnum

		var i int = 0
		for i = 0; i < (len(wen) - curCharStart); i++ {
			if string(wen[curCharStart+i]) == "\n" {
				linenum--
				cchar = charnum
			} else if string(wen[curCharStart+i]) == "\t" {
				cchar -= 4
			} else {
				cchar -= 2
			}

			if cchar <= 0 {
				linenum--
				cchar += charnum
			}
			curCharEnd = i + curCharStart
			if linenum == 0 {
				break
			}
		}
		updateBookContext()
	}
}

func PrevPage() {
	if curCharStart > 0 {
		var linenum = termui.TermHeight() - 3 //可用行数
		var charnum = termui.TermWidth() - 2  //每行可用字符数
		var cchar int = charnum

		var i int = 0
		for i = 0; i <= curCharStart; i++ {
			if string(wen[curCharStart-i]) == "\n" {
				linenum--
				cchar = charnum
			} else if string(wen[curCharStart-i]) == "\t" {
				cchar -= 4
			} else {
				cchar -= 2
			}

			if cchar <= 0 {
				linenum--
				cchar += charnum
			}
			curCharEnd = curCharStart - i
			if linenum == 0 {
				break
			}
		}

		v := curCharEnd
		curCharEnd = curCharStart
		curCharStart = v

		updateBookContext()
	}
}

func updateBookScene() {
	if Book != nil {
		Book.Width = termui.TermWidth()
		Book.Height = termui.TermHeight() - 1
	}
}
