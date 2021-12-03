package main

import (
	"fmt"
	"github.com/zlgopen/awtk-go/awtk"
)

func appInit() {
	win := awtk.TWindowOpen("basic")
	bar1 := win.Lookup("bar1", true)
	bar2 := win.Lookup("bar2", true)

	incButton := win.Lookup("inc_value", true)
	incButton.On(awtk.EVT_CLICK, func(ctx interface{}, e awtk.TEvent) awtk.TRet {
		fmt.Println("Inc is clicked")
		bar1.AddValueInt(10)
		bar2.AddValueInt(10)
		return awtk.RET_OK
	}, win)
	decButton := win.Lookup("dec_value", true)
	decButton.On(awtk.EVT_CLICK, func(ctx interface{}, e awtk.TEvent) awtk.TRet {
		fmt.Println("Dec is clicked")
		bar1.AddValueInt(-10)
		bar2.AddValueInt(-10)
		return awtk.RET_OK
	}, win)
}

func main() {
	awtk.Init(320, 480, awtk.APP_DESKTOP, "demo", "res")
	awtk.TExtWidgetsInit()
	awtk.InitAssets()
	appInit()

	awtk.Run()
}
