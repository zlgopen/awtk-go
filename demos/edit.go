package main

import (
	"fmt"

	"github.com/zlgopen/awtk-go/awtk"
)

func onCancel(ctx interface{}, e awtk.TEvent) awtk.TRet {
	fmt.Println("Cancel is clicked")
	return awtk.RET_OK
}

func appInit() {
	win := awtk.TWindowCreateDefault()

	entry:= awtk.TEditCreate(win, 10, 10, 80, 30)
	entry.SetSelfLayout("default(x=c,y=m,w=80%,h=30)")
	entry.SetPropStr(awtk.WIDGET_PROP_INPUT_TYPE, "text")

	quit := awtk.TButtonCreate(win, 10, 10, 80, 30)
	quit.SetPropStr(awtk.WIDGET_PROP_TEXT, "Quit")
	quit.SetSelfLayout("default(x=c,y=m:60,w=80,h=40)")
	quit.SetPropStr("on:click", "quit()")
}

func main() {
	awtk.Init(320, 480, awtk.APP_MOBILE, "demo", "res")
	awtk.TExtWidgetsInit()
	awtk.InitAssets()
	appInit()

	awtk.Run()
}
