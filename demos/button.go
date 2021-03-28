package main

import (
	"github.com/zlgopen/awtk-go/awtk"
)

func appInit() {
	win := awtk.TWindowCreateDefault()
	ok := awtk.TButtonCreate(win, 10, 10, 80, 30)
	ok.SetPropStr(awtk.WIDGET_PROP_TEXT, "Quit")
	ok.SetSelfLayout("default(x=c,y=m,w=80,h=40)")
	ok.SetPropStr("on:click", "quit()")
}

func main() {
	awtk.Init(320, 480, awtk.APP_DESKTOP, "demo", "res")
	awtk.InitAssets()
	appInit()

	awtk.Run()
}
