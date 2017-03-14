package main

import (
	"github.com/gizak/termui"
)

var BookList *termui.List
var Bar *termui.Par
var Book *termui.Par

func initUI() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	initBookList()
	initBar()
	updateBookListContext()

	UIhandle()
	render()
	termui.Loop()
}

func render() {

	switch Scene {
	case BookScene:
		termui.Render(Book)
		termui.Render(Bar)
		break
	case MenuScene:
		termui.Render(BookList)
		termui.Render(Bar)
		break
	case InfoScene:

		termui.Render(Bar)
		break
	case BossScene:

		break
	}
}

func UIhandle() {
	//退出
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/<escape>", func(termui.Event) {
		termui.StopLoop()
	})

	//上下操作
	termui.Handle("/sys/kbd/k", func(termui.Event) {
		if Scene == MenuScene {
			BookListPrev()
			termui.Render(BookList)
		} else if Scene == BookScene {
			PrevPage()
			termui.Render(Book)
		}
	})
	termui.Handle("/sys/kbd/j", func(termui.Event) {
		if Scene == MenuScene {
			BookListNext()
			termui.Render(BookList)
		} else if Scene == BookScene {
			NextPage()
			termui.Render(Book)
		}
	})
	termui.Handle("/sys/kbd/<space>", func(termui.Event) {
		if Scene == BookScene {
			NextPage()
			termui.Render(Book)
		}
	})

	//打开菜单
	termui.Handle("/sys/kbd/m", func(termui.Event) {
		if Scene != MenuScene {
			Scene = MenuScene
			render()
		}
	})

	//信息界面
	termui.Handle("/sys/kbd/i", func(termui.Event) {
		if Scene != InfoScene {
			Scene = InfoScene
			render()
		}
	})

	//老板界面
	termui.Handle("/sys/kbd/i", func(termui.Event) {
		if Scene != BossScene {
			Scene = BossScene
			render()
		}
	})

	//确认键
	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		//只有才菜单界面才有用
		if Scene == MenuScene {
			//进入图书
			openBook()
			Scene = BookScene
			render()
		}
	})

	//界面缩放
	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		updateBookList()
		updateBar()
		updateBookScene()
		termui.Clear()
		render()
	})
}
