package awtk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWindow(t *testing.T) {
	Init(320, 480, APP_DESKTOP, "demo", "../res")
	win := TWindowCreateDefault()

	assert.NotNil(t, win)
	fmt.Println(win.SetText("Hello"), RET_OK)

	assert.Equal(t, win.SetText("Hello"), RET_OK)
	assert.Equal(t, win.GetText(), "Hello")

	assert.Equal(t, win.SetName("main"), RET_OK)
	assert.Equal(t, win.GetName(), "main")

	assert.Equal(t, win.GetEnable(), true)
	assert.Equal(t, win.GetVisible(), true)
	assert.Equal(t, win.GetSensitive(), true)
	assert.Equal(t, win.GetFocusable(), false)
	assert.Equal(t, win.GetFloating(), false)
	assert.Equal(t, win.GetAutoAdjustSize(), false)
	assert.Equal(t, win.GetFeedback(), false)

	assert.Equal(t, win.GetX(), 0)
	assert.Equal(t, win.GetY(), 0)
	assert.Equal(t, win.GetW(), 320)
	assert.Equal(t, win.GetH(), 480)

}

func TestButton(t *testing.T) {
	win := TWindowCreateDefault()
	assert.NotNil(t, win)
	ok := TButtonCreate(win, 10, 20, 80, 30)
	assert.NotNil(t, ok)
	assert.Equal(t, ok.GetX(), 10)
	assert.Equal(t, ok.GetY(), 20)
	assert.Equal(t, ok.GetW(), 80)
	assert.Equal(t, ok.GetH(), 30)

	assert.Equal(t, ok.Move(30, 40), RET_OK)
	assert.Equal(t, ok.GetX(), 30)
	assert.Equal(t, ok.GetY(), 40)

	assert.Equal(t, ok.Resize(50, 60), RET_OK)
	assert.Equal(t, ok.GetW(), 50)
	assert.Equal(t, ok.GetH(), 60)

	assert.Equal(t, ok.MoveResize(30, 40, 150, 30), RET_OK)
	assert.Equal(t, ok.GetX(), 30)
	assert.Equal(t, ok.GetY(), 40)
	assert.Equal(t, ok.GetW(), 150)
	assert.Equal(t, ok.GetH(), 30)

	btn := TButtonCast(ok)
	assert.Equal(t, btn.SetRepeat(1000), RET_OK)
	assert.Equal(t, btn.GetRepeat(), int32(1000))

	assert.Equal(t, btn.SetEnableLongPress(true), RET_OK)
	assert.Equal(t, btn.GetEnableLongPress(), true)

	assert.Equal(t, btn.SetLongPressTime(2000), RET_OK)
	assert.Equal(t, btn.GetLongPressTime(), uint32(2000))

	assert.NotEqual(t, ok.On(EVT_CLICK, func(ctx interface{}, e TEvent) TRet {
		fmt.Println("OK is clicked")
		return RET_OK
	}, win), TK_INVALID_ID)
}
