# Go 语言绑定笔记

理论上 AWTK 本身脚本化是非常简单的事情，AWTK 已经支持 jerryscript、nodejs、quickjs、lua、python、java、minijvm 和 cpp 等语言的绑定。不过每种语言都有自己的特性，所以每次增加新的语言，也是对 AWTK 绑定机制的考验和完善。

这次实现 Go 语言绑定大概花了两天时间。一天时间用来写代码生成器，一天时间用来写 scriptable 为 custom 的函数和示例代码。Go 语言调用 C 语言的函数有些特殊的要求，所以对 AWTK 的部分 API 做来一些扩展或改进：

## 1. Go 不支持直接访问 C 语言中用位域定义的结构变量。

AWTK 中部分属性使用了位域。如：

```c
  /**
   * @property {bool_t} alt
   * @annotation ["readable", "scriptable"]
   * alt 键是否按下。
   */
  uint8_t alt : 1;
  /**
   * @property {bool_t} ctrl
   * @annotation ["readable", "scriptable"]
   * ctrl 键是否按下。
   */
  uint8_t ctrl : 1;
  /**
   * @property {bool_t} shift
   * @annotation ["readable", "scriptable"]
   * shift 键是否按下。
   */
  uint8_t shift : 1;
```

为此采用了两种方式去解决。

方案 1：如果结构的实例数量不大时，使用其它数据类型代替（如 bool_t)。如：

```c
  /**  
   * @property {bool_t} alt
   * @annotation ["readable", "scriptable"]
   * alt 键是否按下。
   */
  bool_t alt; 
  /**  
   * @property {bool_t} ctrl
   * @annotation ["readable", "scriptable"]
   * ctrl 键是否按下。
   */
  bool_t ctrl;
  /**  
   * @property {bool_t} shift
   * @annotation ["readable", "scriptable"]
   * shift 键是否按下。
   */
  bool_t shift;
```

方案 2：如果结构的实例数量大时，提供相应的 get 函数，代替直接访问成员。如：

```c
/**
 * @method widget_get_visible
 * 获取控件 visible 属性值。
 * @annotation ["scriptable"]
 * @param {widget_t*} widget 控件对象。
 *
 * @return {bool_t} 返回 visible。
 */
bool_t widget_get_visible(widget_t* widget);
```

## 2. Go 语言不支持访问 C 语言 union 的成员。

方案：提供相应的 get 函数，代替直接访问成员。如：

```c
/**
 * @class color_t
 * 颜色。
 * @annotation ["scriptable"]
 *
 */
typedef union _color_t {
  /**
   * @property {rgba_t} rgba
   * @annotation ["readable", "writable"]
   * 颜色的 RGBA 值。
   */
  rgba_t rgba;
  /**
   * @property {uint32_t} color
   * @annotation ["readable", "writable", "scriptable"]
   * 颜色的数值。
   */
  uint32_t color;
} color_t;

/**
 * @method color_get_color
 *
 * 获取颜色值。
 *
 * > 主要供脚本语言使用。
 *
 * @annotation ["scriptable"]
 * @param {color_t*} c color 对象。
 *
 * @return {uint32_t} 返回颜色值。
 *
 */
uint32_t color_get_color(color_t* c);
```

## 3. Go 不支持调用 C 语言中可变参数的函数。

AWTK 中的 widget_set_visible 由于历史原因弄成了可变参数。

方案：只好做些修改：

```c
/**
 * @method widget_set_visible
 * 设置控件的可见性。
 * @annotation ["scriptable"]
 * @param {widget_t*} widget 控件对象。
 * @param {bool_t} visible 是否可见。
 *
 * @return {ret_t} 返回 RET_OK 表示成功，否则表示失败。
 */
/*为了避免脚本绑定时冲突，去掉了 recursive 参数。为了兼容旧版本，只好改成可变参数。*/
#ifdef AWTK_GO
ret_t widget_set_visible(widget_t* widget, bool_t visible);
#else
ret_t widget_set_visible(widget_t* widget, bool_t visible, ...);
#endif/*AWTK_GO*/
```

## 4.type 在 Go 语言中是关键字。

AWTK 中有些结构的成员名为 type，直接访问会造成编译错误。

方案：提供相应的 get 函数，代替直接访问成员。

```c
/**
 * @method event_get_type
 * 获取 event 类型。 
 * @annotation ["scriptable"]
 * @param {event_t*} event event 对象。
 *
 * @return {int}  返回 event 类型。
 */
int event_get_type(event_t* event);
```

## 5. 不支持将 Go 的函数转换成 C 的函数指针。

方案：将函数和函数的上下文合并成一个对象，转递给 C 语言，由 C 语言的函数反过来调用 Go 的包装函数。

```go
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
```

```c
extern int GoOnTimer(void* ctx);

static ret_t timer_info_on_destroy(void *data) {
  timer_info_t *item = (timer_info_t *)data;
  GoReleasePointer(item->ctx);
  return RET_OK;
}

static ret_t wrap_on_timer(const timer_info_t* info) {
  return (ret_t)GoOnTimer(info->ctx);
}

static uint32_t wrap_add_timer(void* ctx, uint32_t duration) {
  uint32_t ret = timer_add(wrap_on_timer, ctx, duration);
  if(ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    timer_set_on_destroy(ret, timer_info_on_destroy, NULL); 
  }

  return ret;
}
```

## 6. 枚举类型的处理

最早枚举类型直接当作 int 类型处理。

```go
type TRet int

const (
	RET_OK              = C.RET_OK
	RET_OOM             = C.RET_OOM
	RET_FAIL            = C.RET_FAIL
	RET_NOT_IMPL        = C.RET_NOT_IMP
)
```

但是单元测试会失败：

```
assert.Equal(t, win.SetText("Hello"), RET_OK)
```

```go
    awtk_test.go:17: 
        	Error Trace:	awtk_test.go:17
        	Error:      	Not equal: 
        	            	expected: awtk.TRet(0)
        	            	actual  : int(0)
        	Test:       	TestSetGet
```

于是在常量前面增加一个类型。一切正常了：

```go
type TRet int

const (
	RET_OK             TRet = C.RET_OK
	RET_OOM            TRet = C.RET_OOM
	RET_FAIL           TRet = C.RET_FAIL
	RET_NOT_IMPL       TRet = C.RET_NOT_IMP
)
```

## 7. cgo 的坑

 * import "C" 之前不能有空行

 ```c
 /*
#include <awtk.h>
#include <tkc/rlog.h>
#include <conf_io/app_conf.h>
#include "../res/assets.inc"
*/
import "C"
```

## 总结

整个支持 Go 语言绑定过程还是很顺利也很有趣的。刚开始接触 Go 语言，初步感觉 Go 语言是很好用的，和 C 语言结合很方便。
