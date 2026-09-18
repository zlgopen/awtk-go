cat awtk_pre.txt >awtk.go
cat ../../awtk-binding/tools/code_gen/go/output/*.go >>awtk.go
# cgo: C 结构体字段 type 是 Go 关键字，访问时需写成 _type
sed -i.bak 's/\.type);$/._type);/' awtk.go && rm -f awtk.go.bak
# cgo: char[N] 字段不能直接传给 C.GoString，需取 &field[0]
sed -i.bak 's/C\.GoString((\*C\.scroll_bar_t)(unsafe\.Pointer(this\.handle))\.wheel_modifier_key);/C.GoString(\&(*C.scroll_bar_t)(unsafe.Pointer(this.handle)).wheel_modifier_key[0]);/' awtk.go && rm -f awtk.go.bak
# C 签名是 int32_t，IDL 仍写 uint32_t
sed -i.bak 's/combo_box_set_selected_index((\*C.widget_t)(this.handle), (C.uint32_t)(index))/combo_box_set_selected_index((*C.widget_t)(this.handle), (C.int32_t)(index))/' awtk.go && rm -f awtk.go.bak
# C 签名是 float_t x/y，IDL 仍写 xy_t
sed -i.bak 's/vgcanvas_fill_text_by_glyphs(\(.*\), (C.xy_t)(x), (C.xy_t)(y), (C.float_t)(max_width))/vgcanvas_fill_text_by_glyphs(\1, (C.float_t)(x), (C.float_t)(y), (C.float_t)(max_width))/' awtk.go && rm -f awtk.go.bak
