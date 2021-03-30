package awtk

/*
#include "./awtk_wrap.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

type OnEventFunc func(ctx interface{}, e TEvent) TRet

type OnEventInfo struct {
	onEvent OnEventFunc
	ctx     interface{}
}

//export GoReleasePointer
func GoReleasePointer(ctx unsafe.Pointer) C.int {
	fmt.Println("GoReleasePointer", ctx)
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

func (this TWidget) AddTimer(onTimer OnTimerFunc, ctx interface{}, duration uint32) uint32 {
	c := OnTimerInfo{onTimer: onTimer, ctx: ctx}
	p := gopointer.Save(c)

	id := C.wrap_widget_add_timer((*C.widget_t)(this.handle), p, C.uint32_t(duration))

	return uint32(id)
}

func (this TWidget) AddIdle(onIdle OnIdleFunc, ctx interface{}) uint32 {
	c := OnIdleInfo{onIdle: onIdle, ctx: ctx}
	p := gopointer.Save(c)

	id := C.wrap_widget_add_idle((*C.widget_t)(this.handle), p)

	return uint32(id)
}

func (this TEmitter) On(eventType uint32, onEvent OnEventFunc, ctx interface{}) uint32 {
	c := OnEventInfo{onEvent: onEvent, ctx: ctx}
	p := gopointer.Save(c)

	id := C.wrap_emitter_on((*C.emitter_t)(this.handle), C.uint32_t(eventType), p)

	return uint32(id)
}

/////////////////////timer/////////////////////////////
type OnTimerFunc func(ctx interface{}) TRet

type OnTimerInfo struct {
	onTimer OnTimerFunc
	ctx     interface{}
}

//export GoOnTimer
func GoOnTimer(ctx unsafe.Pointer) C.int {
	info := gopointer.Restore(ctx).(OnTimerInfo)
	ret := info.onTimer(info.ctx)
	fmt.Println("GoOnTimer", info, ret)

	return C.int(ret)
}

func AddTimer(onTimer OnTimerFunc, ctx interface{}, duration uint32) uint32 {
	c := OnTimerInfo{onTimer: onTimer, ctx: ctx}
	p := gopointer.Save(c)

	id := C.wrap_add_timer(p, C.uint32_t(duration))

	return uint32(id)
}

/////////////////////idle/////////////////////////////
type OnIdleFunc func(ctx interface{}) TRet

type OnIdleInfo struct {
	onIdle OnIdleFunc
	ctx    interface{}
}

//export GoOnIdle
func GoOnIdle(ctx unsafe.Pointer) C.int {
	info := gopointer.Restore(ctx).(OnIdleInfo)
	ret := info.onIdle(info.ctx)
	fmt.Println("GoOnIdle", info, ret)

	return C.int(ret)
}

func AddIdle(onIdle OnIdleFunc, ctx interface{}) uint32 {
	c := OnIdleInfo{onIdle: onIdle, ctx: ctx}
	p := gopointer.Save(c)

	id := C.wrap_add_idle(p)

	return uint32(id)
}
