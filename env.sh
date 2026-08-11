#Linux/MasOS
export AWTK_ROOT=${PWD}/../awtk
export AWTK_BIN=${PWD}/../awtk/bin

#for arm-linux-fb
#export GOOS=linux
#export GOARCH=arm
#export CC=arm-linux-gnueabihf-gcc
#export AWTK_BIN=${PWD}/../awtk-linux-fb/bin

#windows bash
#export AWTK_ROOT=d:/work/awtk-root/awtk

export PATH=${PATH}:${AWTK_ROOT}/bin
export LD_LIBRARY_PATH=${AWTK_BIN}:${LD_LIBRARY_PATH}
# macOS: libawtk.dylib 使用 @rpath，需同时设置 DYLD_LIBRARY_PATH 与 CGO rpath
export DYLD_LIBRARY_PATH=${AWTK_BIN}:${DYLD_LIBRARY_PATH}

export CGO_ENABLED=1
export CGO_CFLAGS="-DWITH_FS_RES=1 -DAWTK_GO=1 -DHAS_STDIO=1"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/src"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/3rd"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/src/ext_widgets"
export CGO_LDFLAGS="-L${AWTK_BIN} -lawtk"
# macOS/Linux: 写入 LC_RPATH，避免 dyld 报 no LC_RPATH's found
case "$(uname -s)" in
  Darwin)
    export CGO_LDFLAGS="${CGO_LDFLAGS} -Wl,-rpath,${AWTK_BIN}"
    ;;
  Linux)
    export CGO_LDFLAGS="${CGO_LDFLAGS} -Wl,-rpath,${AWTK_BIN}"
    ;;
esac


#for window mingw, uncomment this following line
#export CGO_CFLAGS="${CGO_CFLAGS} -DMINGW=1"


