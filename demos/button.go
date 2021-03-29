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

	ok := awtk.TButtonCreate(win, 10, 10, 80, 30)
	ok.SetPropStr(awtk.WIDGET_PROP_TEXT, "OK")
	ok.SetSelfLayout("default(x=c,y=m,w=80,h=40)")
	ok.On(awtk.EVT_CLICK, func(ctx interface{}, e awtk.TEvent) awtk.TRet {
		fmt.Println("OK is clicked")
		return awtk.RET_OK
	}, win)

	cancel := awtk.TButtonCreate(win, 10, 10, 80, 30)
	cancel.SetPropStr(awtk.WIDGET_PROP_TEXT, "Cancel")
	cancel.SetSelfLayout("default(x=c,y=m:-60,w=80,h=40)")
	cancel.On(awtk.EVT_CLICK, onCancel, win)

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
