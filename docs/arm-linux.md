# 部署 AWTK-GO 到嵌入式 linux 系统

## 0. 获取源码

```
git clone https://github.com/zlgopen/awtk.git
git clone https://github.com/zlgopen/awtk-linux-fb.git
git clone https://github.com/zlgopen/awtk-go.git
```

## 1. 交叉编译 awtk-linux-fb

具体方法请参考：https://github.com/zlgopen/awtk-linux-fb

## 2. 交叉编译 awtk-go

> 进入 awtk-go 目录

* 准备

> 交叉编译 awtk-go 需要修改 env.sh。打开文件 env.sh，放开下面被注释的代码，并将 CC 改成实际的交叉编译器的名称。

```
#export GOOS=linux
#export GOARCH=arm
#export CC=arm-linux-gnueabihf-gcc
#export AWTK_BIN=${PWD}/../awtk-linux-fb/bin
```

* 让环境变量生效

```
source env.sh
```

* 编译

```
cd awtk;go build; cd -
```

* 生成发布的文件

> 可以修改 release.sh 文件即可，其内容如下：

```
source env.sh

rm -rf release release.tgz

# 更新资源
python3 scripts/update_res.py all

# 拷贝 AWTK 动态库
mkdir -p release/bin
cp -rvf res release/
cp -fv ${AWTK_BIN}/*awtk.* release/bin
rm -rf release/res/assets/default/inc
rm release/res/assets*.inc

# 编译应用程序，请改为您自己的应用程序
go build -o release/bin/app demos/button.go

# 生成压缩包
tar czfv release.tgz release

#请将 release.tgz 拷贝到目标开发板
#解压后进入 release 目录，运行。/bin/app 启动应用程序。
```

* 将 release.tgz 拷贝到目标开发板

* 在目标板上解压运行

```
tar xf release.tgz 
cd release 
./bin/app 
```
