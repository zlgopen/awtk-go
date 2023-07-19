source env.sh

rm -rf release release.tgz

# 更新资源
python3 scripts/update_res.py all

# 拷贝AWTK动态库
mkdir -p release/bin
cp -rvf res release/
cp -fv ${AWTK_BIN}/*awtk.* release/bin
rm -rf release/res/assets/default/inc
rm release/res/assets*.inc

# 编译应用程序，请改为您自己的应用程序
go build -o release/bin/app demos/button.go

# 生成压缩包
tar czfv release.tgz release

#请将release.tgz拷贝到目标开发板
#解压后进入release目录，运行./bin/app启动应用程序。

