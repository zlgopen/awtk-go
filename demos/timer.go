package main

import (
	"fmt"
	"strconv"

	"github.com/zlgopen/awtk-go/awtk"
)

func onTimer(ctx interface{}) awtk.TRet {
	label := ctx.(awtk.TWidget)
	text := label.GetText()
	n, _ := strconv.Atoi(text)
	value := strconv.Itoa(n + 1)
	label.SetText(value)

	fmt.Println("onTimer", value)
	return awtk.RET_REPEAT
}

func appInit() {
	win := awtk.TWindowCreateDefault()

	label := awtk.TButtonCreate(win, 10, 10, 80, 30)
	label.SetPropStr(awtk.WIDGET_PROP_TEXT, "0")
	label.SetSelfLayout("default(x=c,y=m,w=80,h=40)")
	awtk.AddTimer(onTimer, label, 1000)

	quit := awtk.TButtonCreate(win, 10, 10, 80, 30)
	quit.SetPropStr(awtk.WIDGET_PROP_TEXT, "Quit")
	quit.SetSelfLayout("default(x=c,y=m:60,w=80,h=40)")
	quit.SetPropStr("on:click", "quit()")

	label.AddTimer(func(ctx interface{}) awtk.TRet {
		fmt.Println("Widget Timer Repeat")
		return awtk.RET_REPEAT
	}, label, 1000)

	label.AddTimer(func(ctx interface{}) awtk.TRet {
		fmt.Println("Widget Timer Once")
		return awtk.RET_REMOVE
	}, label, 1000)
}

func main() {
	awtk.Init(320, 480, awtk.APP_DESKTOP, "demo", "res")
	awtk.InitAssets()
	appInit()

	awtk.Run()
}
