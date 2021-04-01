#Linux/MasOS
export AWTK_ROOT=${PWD}/../awtk

#windows bash
#export AWTK_ROOT=d:/work/awtk-root/awtk

export PATH=${PATH}:${AWTK_ROOT}/bin
export LD_LIBRARY_PATH=${AWTK_ROOT}/bin

export CGO_CFLAGS="-DWITH_FS_RES=1 -DAWTK_GO=1 -DHAS_STDIO=1"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/src"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/3rd"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/src/ext_widgets"
export CGO_LDFLAGS="-L ${AWTK_ROOT}/bin -lawtk"

#for window mingw, uncomment this following line
#export CGO_CFLAGS="${CGO_CFLAGS} -DMINGW=1"
