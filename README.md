# awtk-go

ZLG 开源 GUI 引擎 [AWTK](https://github.com/zlgopen/awtk) 针对 [GO 语言](https://golang.org/) 的绑定。

> 目前处于实验阶段。

## 准备

1. 获取 awtk 并编译

```
git clone https://github.com/zlgopen/awtk.git
cd awtk; scons; cd -
```

> AWTK 的编译环境请参考 AWTK 的文档。

2. 获取 awtk-go 并编译

```
git clone https://github.com/zlgopen/awtk-go.git
```

* 生成资源

```
python ./scripts/update_res.py all
```

> 或者通过 designer 生成资源

* 编译 PC 版本

> 修改环境变量 (MacOS/Linux/Windows(bash) 下修改 env.sh)，将 AWTK_ROOT 指向 awtk 的根目录，并使其生效。

```
source env.sh
```

> 安装依赖的库

```
go mod download github.com/mattn/go-pointer
```

```
cd awtk 
go build
cd -
```

* MacOS 平台将动态库拷贝到 bin 目录下

```
mkdir bin
cp ../awtk/bin/libawtk.* bin
```

## 运行

```
go run demos/button.go
```

## 更新绑定（由本项目的维护人员完成）

```
./sync.sh
```

> 在非 bash 终端（如 Windows 平台的 cmd.exe)，需要根据 sync.sh 的内容手工执行相应的命令。

## 代码示例

```go
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
	awtk.TExtWidgetsInit()
	awtk.InitAssets()
	appInit()

	awtk.Run()
}
```

## 注意事项

* 应用程序类型

> 初始化时指定为桌面类型类型时，使用PC原生输入法，不会弹出AWTK自身软键盘。

```go
awtk.Init(320, 480, awtk.APP_DESKTOP, "demo", "res")
```

> 如果希望使用AWTK内置输入法，请初始化为MOBILE类型。

```go
awtk.Init(320, 480, awtk.APP_MOBILE, "demo", "res")
```

## 文档

* [AWTK JS API 文档](https://github.com/zlgopen/awtk-binding/tree/master/docs/js)

* [AWTK 脚本绑定原理](https://github.com/zlgopen/awtk/blob/master/docs/script_binding.md)

* [AWTK Go 语言绑定笔记](docs/awtk_go.md)

* [部署 AWTK-GO 到嵌入式 linux 系统](docs/arm-linux.md)

> 本文以 Linux/MacOS 为例，Windows 可能会微妙差异，请酌情处理。
