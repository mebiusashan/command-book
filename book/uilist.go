package main

import (
	//"fmt"
	"github.com/gizak/termui"
	"strconv"
)

var cTopIndex int = 0

func initBookList() {
	BookList = termui.NewList()
	BookList.ItemFgColor = termui.ColorYellow
	BookList.BorderLabel = "我的图书 共[" + strconv.Itoa(BookListNum) + "](fg-red)本"
	BookList.Height = termui.TermHeight() - 1
	BookList.Width = termui.TermWidth()
	BookList.Y = 0

	BooklistData[curBookIndex] = "[" + BooklistData[curBookIndex] + "](fg-white,bg-green)"
}

func BookListNext() {
	if (BookListNum - curBookIndex) != 1 {
		BooklistData[curBookIndex] = curBookName
		curBookIndex++
		curBookName = BooklistData[curBookIndex]
		BooklistData[curBookIndex] = "[" + BooklistData[curBookIndex] + "](fg-white,bg-green)"
	}
	updateBookListContext()
}

func BookListPrev() {
	if curBookIndex != 0 {
		BooklistData[curBookIndex] = curBookName
		curBookIndex--
		curBookName = BooklistData[curBookIndex]
		BooklistData[curBookIndex] = "[" + BooklistData[curBookIndex] + "](fg-white,bg-green)"
	}
	updateBookListContext()
}

func updateBookListContext() {
	if (curBookIndex - cTopIndex) >= (BookList.Height - 2) {
		cTopIndex++
	} else if (curBookIndex - cTopIndex) < 0 {
		cTopIndex--
	}

	BookList.Items = BooklistData[cTopIndex:]
}

func updateBookList() {
	BookList.Width = termui.TermWidth()
	BookList.Height = termui.TermHeight() - 1
	if BookListNum <= BookList.Height {
		cTopIndex = 0
	}
}
