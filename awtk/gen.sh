cat awtk_pre.txt >awtk.go
cat ../../awtk-binding/tools/code_gen/go/output/*.go >>awtk.go
# cgo: C 结构体字段 type 是 Go 关键字，访问时需写成 _type
sed -i.bak 's/\.type);$/._type);/' awtk.go && rm -f awtk.go.bak
# cgo: char[N] 字段不能直接传给 C.GoString，需取 &field[0]
sed -i.bak 's/C\.GoString((\*C\.scroll_bar_t)(unsafe\.Pointer(this\.handle))\.wheel_modifier_key);/C.GoString(\&(*C.scroll_bar_t)(unsafe.Pointer(this.handle)).wheel_modifier_key[0]);/' awtk.go && rm -f awtk.go.bak
