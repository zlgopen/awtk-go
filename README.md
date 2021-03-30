# awtk-go

ZLG 开源 GUI 引擎 [awtk](https://github.com/zlgopen/awtk) 针对 [GO 语言](https://golang.org/) 的绑定。

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

```
cd awtk 
go build
cd -
```

* Linux/MacOS 平台将动态库拷贝到 bin 目录下

```
mkdir bin
cp ../awtk/bin/libawtk.* bin
```

* Windows 平台将动态库拷贝到当前目录下

```
mkdir bin
cp ../awtk/bin/awtk.dll .
```

## 运行

> 在 Linux 下需要设置 LD\_LIBRARY\_PATH 

```
export LD_LIBRARY_PATH=bin
```

```
go run demos/button.go
```

## 更新绑定（由本项目的维护人员完成）

```
./sync.sh
```

> 在非 bash 终端（如 Windows 平台的 cmd.exe)，需要根据 sync.sh 的内容手工执行相应的命令。

## 文档

* [AWTK JS API 文档](https://github.com/zlgopen/awtk-binding/tree/master/docs/js)

* [AWTK 脚本绑定原理](https://github.com/zlgopen/awtk/blob/master/docs/script_binding.md)

* [AWTK Go 语言绑定笔记](docs/awtk_go.md)

> 本文以 Linux/MacOS 为例，Windows 可能会微妙差异，请酌情处理。
