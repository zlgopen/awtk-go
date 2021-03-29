package main

import (
	"fmt"
	"strconv"

	"github.com/zlgopen/awtk-go/awtk"
)

func onIdle(ctx interface{}) awtk.TRet {
	label := ctx.(awtk.TWidget)
	text := label.GetText()
	n, _ := strconv.Atoi(text)
	value := strconv.Itoa(n + 1)

	label.SetText(value)

	fmt.Println("onIdle", value)
	return awtk.RET_REPEAT
}

func appInit() {
	win := awtk.TWindowCreateDefault()

	label := awtk.TButtonCreate(win, 10, 10, 80, 30)
	label.SetPropStr(awtk.WIDGET_PROP_TEXT, "0")
	label.SetSelfLayout("default(x=c,y=m,w=80,h=40)")
	awtk.AddIdle(onIdle, label)

	quit := awtk.TButtonCreate(win, 10, 10, 80, 30)
	quit.SetPropStr(awtk.WIDGET_PROP_TEXT, "Quit")
	quit.SetSelfLayout("default(x=c,y=m:60,w=80,h=40)")
	quit.SetPropStr("on:click", "quit()")
}

func main() {
	awtk.Init(320, 480, awtk.APP_DESKTOP, "demo", "res")
	awtk.InitAssets()
	appInit()

	awtk.Run()
}
