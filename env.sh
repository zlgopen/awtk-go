export AWTK_ROOT=/Users/jim/work/awtk-root/awtk

export PATH=${PATH}:${AWTK_ROOT}/bin
export LD_LIBRARY_PATH=${AWTK_ROOT}/bin

export CGO_CFLAGS="-DWITH_FS_RES=1 -DAWTK_GO=1 -DHAS_STDIO=1"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/src"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/3rd"
export CGO_CFLAGS="${CGO_CFLAGS} -I ${AWTK_ROOT}/src/ext_widgets"
export CGO_LDFLAGS="-L ${AWTK_ROOT}/bin -lawtk"
