rm -rf release

python3 scripts/update_res.py all

mkdir -p release/bin
cp -rvf res release/
cp -fv ${AWTK_BIN}/*awtk.* release/bin

#build you app

go build -o release/bin/button demos/button.go
