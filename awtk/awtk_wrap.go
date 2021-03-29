package awtk

/*
#cgo CFLAGS: -DWITH_FS_RES=1 -DAWTK_GO=1
#cgo CFLAGS: -I../../awtk/src
#cgo CFLAGS: -I../../awtk/3rd
#cgo CFLAGS: -I../../awtk/src/ext_widgets
#cgo LDFLAGS: -L../../awtk/bin -lawtk
#include "./awtk_wrap.h"
*/
import "C"
import "fmt"
import "unsafe"
import gopointer "github.com/mattn/go-pointer"

type OnEventFunc func(ctx interface{}, e TEvent) TRet

type OnEventInfo struct {
	onEvent OnEventFunc
	ctx     interface{}
}

//export GoOnEventDestroy
func GoOnEventDestroy(ctx unsafe.Pointer) C.int {
	info := gopointer.Restore(ctx).(OnEventInfo)
	fmt.Println("GoOnEventDestroy", info)
	gopointer.Unref(ctx)

	return C.int(C.RET_OK)
}

//export GoOnEvent
func GoOnEvent(ctx unsafe.Pointer, e unsafe.Pointer) C.int {
	evt := TEvent{}
	evt.handle = e

	info := gopointer.Restore(ctx).(OnEventInfo)
	info.onEvent(info.ctx, evt)
	fmt.Println("GoOnEvent", info)

	return C.int(C.RET_OK)
}

func (this TWidget) On(eventType uint32, onEvent OnEventFunc, ctx interface{}) uint32 {
	c := OnEventInfo{onEvent: onEvent, ctx: ctx}
	p := gopointer.Save(c)

	id := C.wrap_widget_on((*C.widget_t)(this.handle), C.uint32_t(eventType), p)

	return uint32(id)
}
