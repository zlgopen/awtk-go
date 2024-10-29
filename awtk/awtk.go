package awtk

/*
#include <awtk.h>
#include <tkc/rlog.h>
#include <conf_io/app_conf.h>
#include "../res/assets.inc"
*/
import "C"
import "unsafe"

func InitAssets() TRet {
  return TRet(C.assets_init())
}

type TAlignH int
const (
  ALIGN_H_NONE TAlignH = C.ALIGN_H_NONE
  ALIGN_H_CENTER TAlignH = C.ALIGN_H_CENTER
  ALIGN_H_LEFT TAlignH = C.ALIGN_H_LEFT
  ALIGN_H_RIGHT TAlignH = C.ALIGN_H_RIGHT
)
type TAlignV int
const (
  ALIGN_V_NONE TAlignV = C.ALIGN_V_NONE
  ALIGN_V_MIDDLE TAlignV = C.ALIGN_V_MIDDLE
  ALIGN_V_TOP TAlignV = C.ALIGN_V_TOP
  ALIGN_V_BOTTOM TAlignV = C.ALIGN_V_BOTTOM
)
type TAppBar struct {
  TWidget
}

func TAppBarCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TAppBar{}
  retObj.handle = unsafe.Pointer(C.app_bar_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TAppBarCast(widget TWidget) TAppBar {
  retObj := TAppBar{}
  retObj.handle = unsafe.Pointer(C.app_bar_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func TAppConfSave() TRet {
  return TRet(C.app_conf_save());
}

func TAppConfReload() TRet {
  return TRet(C.app_conf_reload());
}

func TAppConfDeinit() TRet {
  return TRet(C.app_conf_deinit());
}

func TAppConfExist(key string) bool {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return (bool)(C.app_conf_exist(akey));
}

func TAppConfSetInt(key string, v int32) TRet {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return TRet(C.app_conf_set_int(akey, (C.int32_t)(v)));
}

func TAppConfSetInt64(key string, v int64) TRet {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return TRet(C.app_conf_set_int64(akey, (C.int64_t)(v)));
}

func TAppConfSetBool(key string, v bool) TRet {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return TRet(C.app_conf_set_bool(akey, (C.bool_t)(v)));
}

func TAppConfSetDouble(key string, v float64) TRet {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return TRet(C.app_conf_set_double(akey, (C.double)(v)));
}

func TAppConfSetStr(key string, v string) TRet {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  av := C.CString(v)
  defer C.free(unsafe.Pointer(av))
  return TRet(C.app_conf_set_str(akey, av));
}

func TAppConfGetInt(key string, defval int32) int32 {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return (int32)(C.app_conf_get_int(akey, (C.int32_t)(defval)));
}

func TAppConfGetInt64(key string, defval int64) int64 {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return (int64)(C.app_conf_get_int64(akey, (C.int64_t)(defval)));
}

func TAppConfGetBool(key string, defval bool) bool {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return (bool)(C.app_conf_get_bool(akey, (C.bool_t)(defval)));
}

func TAppConfGetDouble(key string, defval float64) float64 {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return (float64)(C.app_conf_get_double(akey, (C.double)(defval)));
}

func TAppConfGetStr(key string, defval string) string {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  adefval := C.CString(defval)
  defer C.free(unsafe.Pointer(adefval))
  return C.GoString(C.app_conf_get_str(akey, adefval));
}

func TAppConfRemove(key string) TRet {
  akey := C.CString(key)
  defer C.free(unsafe.Pointer(akey))
  return TRet(C.app_conf_remove(akey));
}

type TAppType int
const (
  APP_MOBILE TAppType = C.APP_MOBILE
  APP_SIMULATOR TAppType = C.APP_SIMULATOR
  APP_DESKTOP TAppType = C.APP_DESKTOP
  APP_CONSOLE TAppType = C.APP_CONSOLE
)
type TAssetInfo struct {
  handle unsafe.Pointer
}

func (this TAssetInfo) GetType() uint16 {
  return (uint16)(C.asset_info_get_type((*C.asset_info_t)(this.handle)));
}

func (this TAssetInfo) GetName() string {
  return C.GoString(C.asset_info_get_name((*C.asset_info_t)(this.handle)));
}

func (this TAssetInfo) IsInRom() bool {
  return (bool)(C.asset_info_is_in_rom((*C.asset_info_t)(this.handle)));
}

func (this TAssetInfo) SetIsInRom(is_in_rom bool) TRet {
  return TRet(C.asset_info_set_is_in_rom((*C.asset_info_t)(this.handle), (C.bool_t)(is_in_rom)));
}

func (this TAssetInfo) GetSubtype() uint8 {
  return (uint8)((*C.asset_info_t)(unsafe.Pointer(this.handle)).subtype);
}

func (this TAssetInfo) GetFlags() uint8 {
  return (uint8)((*C.asset_info_t)(unsafe.Pointer(this.handle)).flags);
}

func (this TAssetInfo) GetSize() uint32 {
  return (uint32)((*C.asset_info_t)(unsafe.Pointer(this.handle)).size);
}

func (this TAssetInfo) GetRefcount() uint32 {
  return (uint32)((*C.asset_info_t)(unsafe.Pointer(this.handle)).refcount);
}

type TAssetType int
const (
  ASSET_TYPE_NONE TAssetType = C.ASSET_TYPE_NONE
  ASSET_TYPE_FONT TAssetType = C.ASSET_TYPE_FONT
  ASSET_TYPE_IMAGE TAssetType = C.ASSET_TYPE_IMAGE
  ASSET_TYPE_STYLE TAssetType = C.ASSET_TYPE_STYLE
  ASSET_TYPE_UI TAssetType = C.ASSET_TYPE_UI
  ASSET_TYPE_XML TAssetType = C.ASSET_TYPE_XML
  ASSET_TYPE_STRINGS TAssetType = C.ASSET_TYPE_STRINGS
  ASSET_TYPE_SCRIPT TAssetType = C.ASSET_TYPE_SCRIPT
  ASSET_TYPE_FLOW TAssetType = C.ASSET_TYPE_FLOW
  ASSET_TYPE_DATA TAssetType = C.ASSET_TYPE_DATA
)
type TAssetsManager struct {
  TEmitter
}

func TAssetsManagerInstance() TAssetsManager {
  retObj := TAssetsManager{}
  retObj.handle = unsafe.Pointer(C.assets_manager())
  return retObj
}

func (this TAssetsManager) SetTheme(theme string) TRet {
  atheme := C.CString(theme)
  defer C.free(unsafe.Pointer(atheme))
  return TRet(C.assets_manager_set_theme((*C.assets_manager_t)(this.handle), atheme));
}

func (this TAssetsManager) Ref(typex TAssetType, name string) TAssetInfo {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TAssetInfo{}
  retObj.handle = unsafe.Pointer(C.assets_manager_ref((*C.assets_manager_t)(this.handle), (C.asset_type_t)(typex), aname))
  return retObj
}

func (this TAssetsManager) RefEx(typex TAssetType, subtype uint16, name string) TAssetInfo {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TAssetInfo{}
  retObj.handle = unsafe.Pointer(C.assets_manager_ref_ex((*C.assets_manager_t)(this.handle), (C.asset_type_t)(typex), (C.uint16_t)(subtype), aname))
  return retObj
}

func (this TAssetsManager) Unref(info TAssetInfo) TRet {
  return TRet(C.assets_manager_unref((*C.assets_manager_t)(this.handle), (*C.asset_info_t)(info.handle)));
}

type TBidiType int
const (
  BIDI_TYPE_AUTO TBidiType = C.BIDI_TYPE_AUTO
  BIDI_TYPE_LTR TBidiType = C.BIDI_TYPE_LTR
  BIDI_TYPE_RTL TBidiType = C.BIDI_TYPE_RTL
  BIDI_TYPE_LRO TBidiType = C.BIDI_TYPE_LRO
  BIDI_TYPE_RLO TBidiType = C.BIDI_TYPE_RLO
  BIDI_TYPE_WLTR TBidiType = C.BIDI_TYPE_WLTR
  BIDI_TYPE_WRTL TBidiType = C.BIDI_TYPE_WRTL
)
type TBitmap struct {
  handle unsafe.Pointer
}

func TBitmapCreate() TBitmap {
  retObj := TBitmap{}
  retObj.handle = unsafe.Pointer(C.bitmap_create())
  return retObj
}

func TBitmapCreateEx(w uint32, h uint32, line_length uint32, format TBitmapFormat) TBitmap {
  retObj := TBitmap{}
  retObj.handle = unsafe.Pointer(C.bitmap_create_ex((C.uint32_t)(w), (C.uint32_t)(h), (C.uint32_t)(line_length), (C.bitmap_format_t)(format)))
  return retObj
}

func (this TBitmap) GetBpp() uint32 {
  return (uint32)(C.bitmap_get_bpp((*C.bitmap_t)(this.handle)));
}

func (this TBitmap) Destroy() TRet {
  return TRet(C.bitmap_destroy_with_self((*C.bitmap_t)(this.handle)));
}

func TBitmapGetBppOfFormat(format TBitmapFormat) uint32 {
  return (uint32)(C.bitmap_get_bpp_of_format((C.bitmap_format_t)(format)));
}

func (this TBitmap) GetW() int {
  return (int)((*C.bitmap_t)(unsafe.Pointer(this.handle)).w);
}

func (this TBitmap) GetH() int {
  return (int)((*C.bitmap_t)(unsafe.Pointer(this.handle)).h);
}

func (this TBitmap) GetLineLength() uint32 {
  return (uint32)((*C.bitmap_t)(unsafe.Pointer(this.handle)).line_length);
}

func (this TBitmap) GetFlags() uint16 {
  return (uint16)((*C.bitmap_t)(unsafe.Pointer(this.handle)).flags);
}

func (this TBitmap) GetFormat() uint16 {
  return (uint16)((*C.bitmap_t)(unsafe.Pointer(this.handle)).format);
}

func (this TBitmap) GetName() string {
  return C.GoString((*C.bitmap_t)(unsafe.Pointer(this.handle)).name);
}

type TBitmapFlag int
const (
  BITMAP_FLAG_NONE TBitmapFlag = C.BITMAP_FLAG_NONE
  BITMAP_FLAG_OPAQUE TBitmapFlag = C.BITMAP_FLAG_OPAQUE
  BITMAP_FLAG_IMMUTABLE TBitmapFlag = C.BITMAP_FLAG_IMMUTABLE
  BITMAP_FLAG_TEXTURE TBitmapFlag = C.BITMAP_FLAG_TEXTURE
  BITMAP_FLAG_CHANGED TBitmapFlag = C.BITMAP_FLAG_CHANGED
  BITMAP_FLAG_PREMULTI_ALPHA TBitmapFlag = C.BITMAP_FLAG_PREMULTI_ALPHA
  BITMAP_FLAG_LCD_ORIENTATION TBitmapFlag = C.BITMAP_FLAG_LCD_ORIENTATION
  BITMAP_FLAG_GPU_FBO_TEXTURE TBitmapFlag = C.BITMAP_FLAG_GPU_FBO_TEXTURE
)
type TBitmapFormat int
const (
  BITMAP_FMT_NONE TBitmapFormat = C.BITMAP_FMT_NONE
  BITMAP_FMT_RGBA8888 TBitmapFormat = C.BITMAP_FMT_RGBA8888
  BITMAP_FMT_ABGR8888 TBitmapFormat = C.BITMAP_FMT_ABGR8888
  BITMAP_FMT_BGRA8888 TBitmapFormat = C.BITMAP_FMT_BGRA8888
  BITMAP_FMT_ARGB8888 TBitmapFormat = C.BITMAP_FMT_ARGB8888
  BITMAP_FMT_RGB565 TBitmapFormat = C.BITMAP_FMT_RGB565
  BITMAP_FMT_BGR565 TBitmapFormat = C.BITMAP_FMT_BGR565
  BITMAP_FMT_RGB888 TBitmapFormat = C.BITMAP_FMT_RGB888
  BITMAP_FMT_BGR888 TBitmapFormat = C.BITMAP_FMT_BGR888
  BITMAP_FMT_GRAY TBitmapFormat = C.BITMAP_FMT_GRAY
  BITMAP_FMT_MONO TBitmapFormat = C.BITMAP_FMT_MONO
)
type TButton struct {
  TWidget
}

func TButtonCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TButton{}
  retObj.handle = unsafe.Pointer(C.button_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TButtonCast(widget TWidget) TButton {
  retObj := TButton{}
  retObj.handle = unsafe.Pointer(C.button_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TButton) SetRepeat(repeat int32) TRet {
  return TRet(C.button_set_repeat((*C.widget_t)(this.handle), (C.int32_t)(repeat)));
}

func (this TButton) SetLongPressTime(long_press_time uint32) TRet {
  return TRet(C.button_set_long_press_time((*C.widget_t)(this.handle), (C.uint32_t)(long_press_time)));
}

func (this TButton) SetEnableLongPress(enable_long_press bool) TRet {
  return TRet(C.button_set_enable_long_press((*C.widget_t)(this.handle), (C.bool_t)(enable_long_press)));
}

func (this TButton) SetEnablePreview(enable_preview bool) TRet {
  return TRet(C.button_set_enable_preview((*C.widget_t)(this.handle), (C.bool_t)(enable_preview)));
}

func (this TButton) GetRepeat() int32 {
  return (int32)((*C.button_t)(unsafe.Pointer(this.handle)).repeat);
}

func (this TButton) GetEnableLongPress() bool {
  return (bool)((*C.button_t)(unsafe.Pointer(this.handle)).enable_long_press);
}

func (this TButton) GetEnablePreview() bool {
  return (bool)((*C.button_t)(unsafe.Pointer(this.handle)).enable_preview);
}

func (this TButton) GetLongPressTime() uint32 {
  return (uint32)((*C.button_t)(unsafe.Pointer(this.handle)).long_press_time);
}

func (this TButton) GetPressed() bool {
  return (bool)((*C.button_t)(unsafe.Pointer(this.handle)).pressed);
}

type TButtonGroup struct {
  TWidget
}

func TButtonGroupCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TButtonGroup{}
  retObj.handle = unsafe.Pointer(C.button_group_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TButtonGroupCast(widget TWidget) TButtonGroup {
  retObj := TButtonGroup{}
  retObj.handle = unsafe.Pointer(C.button_group_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TCalibrationWin struct {
  TWindowBase
}

func TCalibrationWinCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TCalibrationWin{}
  retObj.handle = unsafe.Pointer(C.calibration_win_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TCalibrationWinCast(widget TWidget) TCalibrationWin {
  retObj := TCalibrationWin{}
  retObj.handle = unsafe.Pointer(C.calibration_win_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TCandidates struct {
  TWidget
}

func TCandidatesCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TCandidates{}
  retObj.handle = unsafe.Pointer(C.candidates_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TCandidatesCast(widget TWidget) TCandidates {
  retObj := TCandidates{}
  retObj.handle = unsafe.Pointer(C.candidates_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TCandidates) SetPre(pre bool) TRet {
  return TRet(C.candidates_set_pre((*C.widget_t)(this.handle), (C.bool_t)(pre)));
}

func (this TCandidates) SetSelectByNum(select_by_num bool) TRet {
  return TRet(C.candidates_set_select_by_num((*C.widget_t)(this.handle), (C.bool_t)(select_by_num)));
}

func (this TCandidates) SetAutoHide(auto_hide bool) TRet {
  return TRet(C.candidates_set_auto_hide((*C.widget_t)(this.handle), (C.bool_t)(auto_hide)));
}

func (this TCandidates) SetButtonStyle(button_style string) TRet {
  abutton_style := C.CString(button_style)
  defer C.free(unsafe.Pointer(abutton_style))
  return TRet(C.candidates_set_button_style((*C.widget_t)(this.handle), abutton_style));
}

func (this TCandidates) GetPre() bool {
  return (bool)((*C.candidates_t)(unsafe.Pointer(this.handle)).pre);
}

func (this TCandidates) GetSelectByNum() bool {
  return (bool)((*C.candidates_t)(unsafe.Pointer(this.handle)).select_by_num);
}

func (this TCandidates) GetAutoHide() bool {
  return (bool)((*C.candidates_t)(unsafe.Pointer(this.handle)).auto_hide);
}

func (this TCandidates) GetButtonStyle() string {
  return C.GoString((*C.candidates_t)(unsafe.Pointer(this.handle)).button_style);
}

func (this TCandidates) GetEnablePreview() bool {
  return (bool)((*C.candidates_t)(unsafe.Pointer(this.handle)).enable_preview);
}

type TCanvas struct {
  handle unsafe.Pointer
}

func (this TCanvas) GetWidth() int {
  return (int)(C.canvas_get_width((*C.canvas_t)(this.handle)));
}

func (this TCanvas) GetHeight() int {
  return (int)(C.canvas_get_height((*C.canvas_t)(this.handle)));
}

func (this TCanvas) GetClipRect(r TRect) TRet {
  return TRet(C.canvas_get_clip_rect((*C.canvas_t)(this.handle), (*C.rect_t)(r.handle)));
}

func (this TCanvas) SetClipRect(r TRect) TRet {
  return TRet(C.canvas_set_clip_rect((*C.canvas_t)(this.handle), (*C.rect_t)(r.handle)));
}

func (this TCanvas) SetClipRectEx(r TRect, translate bool) TRet {
  return TRet(C.canvas_set_clip_rect_ex((*C.canvas_t)(this.handle), (*C.rect_t)(r.handle), (C.bool_t)(translate)));
}

func (this TCanvas) SetFillColor(color string) TRet {
  acolor := C.CString(color)
  defer C.free(unsafe.Pointer(acolor))
  return TRet(C.canvas_set_fill_color_str((*C.canvas_t)(this.handle), acolor));
}

func (this TCanvas) SetTextColor(color string) TRet {
  acolor := C.CString(color)
  defer C.free(unsafe.Pointer(acolor))
  return TRet(C.canvas_set_text_color_str((*C.canvas_t)(this.handle), acolor));
}

func (this TCanvas) SetStrokeColor(color string) TRet {
  acolor := C.CString(color)
  defer C.free(unsafe.Pointer(acolor))
  return TRet(C.canvas_set_stroke_color_str((*C.canvas_t)(this.handle), acolor));
}

func (this TCanvas) SetGlobalAlpha(alpha uint8) TRet {
  return TRet(C.canvas_set_global_alpha((*C.canvas_t)(this.handle), (C.uint8_t)(alpha)));
}

func (this TCanvas) Translate(dx int, dy int) TRet {
  return TRet(C.canvas_translate((*C.canvas_t)(this.handle), (C.xy_t)(dx), (C.xy_t)(dy)));
}

func (this TCanvas) Untranslate(dx int, dy int) TRet {
  return TRet(C.canvas_untranslate((*C.canvas_t)(this.handle), (C.xy_t)(dx), (C.xy_t)(dy)));
}

func (this TCanvas) DrawVline(x int, y int, h int) TRet {
  return TRet(C.canvas_draw_vline((*C.canvas_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(h)));
}

func (this TCanvas) DrawHline(x int, y int, w int) TRet {
  return TRet(C.canvas_draw_hline((*C.canvas_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w)));
}

func (this TCanvas) FillRect(x int, y int, w int, h int) TRet {
  return TRet(C.canvas_fill_rect((*C.canvas_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)));
}

func (this TCanvas) ClearRect(x int, y int, w int, h int) TRet {
  return TRet(C.canvas_clear_rect((*C.canvas_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)));
}

func (this TCanvas) StrokeRect(x int, y int, w int, h int) TRet {
  return TRet(C.canvas_stroke_rect((*C.canvas_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)));
}

func (this TCanvas) SetFont(name string, size int) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.canvas_set_font((*C.canvas_t)(this.handle), aname, (C.font_size_t)(size)));
}

func (this TCanvas) ResetFont() TRet {
  return TRet(C.canvas_reset_font((*C.canvas_t)(this.handle)));
}

func (this TCanvas) MeasureText(str string) float64 {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  return (float64)(C.canvas_measure_utf8((*C.canvas_t)(this.handle), astr));
}

func (this TCanvas) DrawText(str string, x int, y int) TRet {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  return TRet(C.canvas_draw_utf8((*C.canvas_t)(this.handle), astr, (C.xy_t)(x), (C.xy_t)(y)));
}

func (this TCanvas) DrawTextInRect(str string, r TRect) TRet {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  return TRet(C.canvas_draw_utf8_in_rect((*C.canvas_t)(this.handle), astr, (*C.rect_t)(r.handle)));
}

func (this TCanvas) DrawIcon(img TBitmap, cx int, cy int) TRet {
  return TRet(C.canvas_draw_icon((*C.canvas_t)(this.handle), (*C.bitmap_t)(img.handle), (C.xy_t)(cx), (C.xy_t)(cy)));
}

func (this TCanvas) DrawImage(img TBitmap, src TRect, dst TRect) TRet {
  return TRet(C.canvas_draw_image((*C.canvas_t)(this.handle), (*C.bitmap_t)(img.handle), (*C.rect_t)(src.handle), (*C.rect_t)(dst.handle)));
}

func (this TCanvas) DrawImageEx(img TBitmap, draw_type TImageDrawType, dst TRect) TRet {
  return TRet(C.canvas_draw_image_ex((*C.canvas_t)(this.handle), (*C.bitmap_t)(img.handle), (C.image_draw_type_t)(draw_type), (*C.rect_t)(dst.handle)));
}

func (this TCanvas) DrawImageEx2(img TBitmap, draw_type TImageDrawType, src TRect, dst TRect) TRet {
  return TRet(C.canvas_draw_image_ex2((*C.canvas_t)(this.handle), (*C.bitmap_t)(img.handle), (C.image_draw_type_t)(draw_type), (*C.rect_t)(src.handle), (*C.rect_t)(dst.handle)));
}

func (this TCanvas) GetVgcanvas() TVgcanvas {
  retObj := TVgcanvas{}
  retObj.handle = unsafe.Pointer(C.canvas_get_vgcanvas((*C.canvas_t)(this.handle)))
  return retObj
}

func TCanvasCast(c TCanvas) TCanvas {
  retObj := TCanvas{}
  retObj.handle = unsafe.Pointer(C.canvas_cast((*C.canvas_t)(c.handle)))
  return retObj
}

func (this TCanvas) Reset() TRet {
  return TRet(C.canvas_reset((*C.canvas_t)(this.handle)));
}

func (this TCanvas) ResetCache() TRet {
  return TRet(C.canvas_reset_cache((*C.canvas_t)(this.handle)));
}

func (this TCanvas) GetOx() int {
  return (int)((*C.canvas_t)(unsafe.Pointer(this.handle)).ox);
}

func (this TCanvas) GetOy() int {
  return (int)((*C.canvas_t)(unsafe.Pointer(this.handle)).oy);
}

func (this TCanvas) GetFontName() string {
  return C.GoString((*C.canvas_t)(unsafe.Pointer(this.handle)).font_name);
}

func (this TCanvas) GetFontSize() int {
  return (int)((*C.canvas_t)(unsafe.Pointer(this.handle)).font_size);
}

func (this TCanvas) GetGlobalAlpha() uint8 {
  return (uint8)((*C.canvas_t)(unsafe.Pointer(this.handle)).global_alpha);
}

type TCanvasOffline struct {
  handle unsafe.Pointer
}

type TCanvasWidget struct {
  TWidget
}

func TCanvasWidgetCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TCanvasWidget{}
  retObj.handle = unsafe.Pointer(C.canvas_widget_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TCanvasWidgetCast(widget TWidget) TCanvasWidget {
  retObj := TCanvasWidget{}
  retObj.handle = unsafe.Pointer(C.canvas_widget_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TCheckButton struct {
  TWidget
}

func TCheckButtonCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TCheckButton{}
  retObj.handle = unsafe.Pointer(C.check_button_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TCheckButtonCreateRadio(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TCheckButton{}
  retObj.handle = unsafe.Pointer(C.check_button_create_radio((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TCheckButton) SetValue(value bool) TRet {
  return TRet(C.check_button_set_value((*C.widget_t)(this.handle), (C.bool_t)(value)));
}

func (this TCheckButton) SetIndeterminate(indeterminate bool) TRet {
  return TRet(C.check_button_set_indeterminate((*C.widget_t)(this.handle), (C.bool_t)(indeterminate)));
}

func (this TCheckButton) GetIndeterminate() bool {
  return (bool)(C.check_button_get_indeterminate((*C.widget_t)(this.handle)));
}

func TCheckButtonCast(widget TWidget) TCheckButton {
  retObj := TCheckButton{}
  retObj.handle = unsafe.Pointer(C.check_button_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func TCheckButtonCreateEx(parent TWidget, x int, y int, w int, h int, typex string, radio bool) TWidget {
  atypex := C.CString(typex)
  defer C.free(unsafe.Pointer(atypex))
  retObj := TCheckButton{}
  retObj.handle = unsafe.Pointer(C.check_button_create_ex((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h), atypex, (C.bool_t)(radio)))
  return retObj.TWidget
}

type TClipBoard struct {
  handle unsafe.Pointer
}

func TClipBoardSetText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.clip_board_set_text(atext));
}

func TClipBoardGetText() string {
  return C.GoString(C.clip_board_get_text());
}

type TClipBoardDataType int
const (
  CLIP_BOARD_DATA_TYPE_NONE TClipBoardDataType = C.CLIP_BOARD_DATA_TYPE_NONE
  CLIP_BOARD_DATA_TYPE_TEXT TClipBoardDataType = C.CLIP_BOARD_DATA_TYPE_TEXT
)
type TClipView struct {
  TWidget
}

func TClipViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TClipView{}
  retObj.handle = unsafe.Pointer(C.clip_view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TClipViewCast(widget TWidget) TClipView {
  retObj := TClipView{}
  retObj.handle = unsafe.Pointer(C.clip_view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TCmdExecEvent struct {
  TEvent
}

func TCmdExecEventCast(event TEvent) TCmdExecEvent {
  retObj := TCmdExecEvent{}
  retObj.handle = unsafe.Pointer(C.cmd_exec_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TCmdExecEvent) GetName() string {
  return C.GoString((*C.cmd_exec_event_t)(unsafe.Pointer(this.handle)).name);
}

func (this TCmdExecEvent) GetArgs() string {
  return C.GoString((*C.cmd_exec_event_t)(unsafe.Pointer(this.handle)).args);
}

func (this TCmdExecEvent) GetResult() TRet {
  return TRet((*C.cmd_exec_event_t)(unsafe.Pointer(this.handle)).result);
}

func (this TCmdExecEvent) GetCanExec() bool {
  return (bool)((*C.cmd_exec_event_t)(unsafe.Pointer(this.handle)).can_exec);
}

type TColor struct {
  handle unsafe.Pointer
}

func TColorCreate(r uint8, g uint8, b uint8, a uint8) TColor {
  retObj := TColor{}
  retObj.handle = unsafe.Pointer(C.color_create((C.uint8_t)(r), (C.uint8_t)(g), (C.uint8_t)(b), (C.uint8_t)(a)))
  return retObj
}

func (this TColor) FromStr(str string) TColor {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  retObj := TColor{}
  retObj.handle = unsafe.Pointer(C.color_from_str((*C.color_t)(this.handle), astr))
  return retObj
}

func (this TColor) R() uint8 {
  return (uint8)(C.color_r((*C.color_t)(this.handle)));
}

func (this TColor) G() uint8 {
  return (uint8)(C.color_g((*C.color_t)(this.handle)));
}

func (this TColor) B() uint8 {
  return (uint8)(C.color_b((*C.color_t)(this.handle)));
}

func (this TColor) A() uint8 {
  return (uint8)(C.color_a((*C.color_t)(this.handle)));
}

func (this TColor) GetColor() uint32 {
  return (uint32)(C.color_get_color((*C.color_t)(this.handle)));
}

func TColorCast(color TColor) TColor {
  retObj := TColor{}
  retObj.handle = unsafe.Pointer(C.color_cast((*C.color_t)(color.handle)))
  return retObj
}

func (this TColor) Destroy() TRet {
  return TRet(C.color_destroy((*C.color_t)(this.handle)));
}

type TColorComponent struct {
  TWidget
}

func TColorComponentCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TColorComponent{}
  retObj.handle = unsafe.Pointer(C.color_component_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TColorComponentCast(widget TWidget) TColorComponent {
  retObj := TColorComponent{}
  retObj.handle = unsafe.Pointer(C.color_component_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TColorPicker struct {
  TWidget
}

func TColorPickerCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TColorPicker{}
  retObj.handle = unsafe.Pointer(C.color_picker_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TColorPicker) SetColor(color string) TRet {
  acolor := C.CString(color)
  defer C.free(unsafe.Pointer(acolor))
  return TRet(C.color_picker_set_color((*C.widget_t)(this.handle), acolor));
}

func TColorPickerCast(widget TWidget) TColorPicker {
  retObj := TColorPicker{}
  retObj.handle = unsafe.Pointer(C.color_picker_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TColorTile struct {
  TWidget
}

func TColorTileCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TColorTile{}
  retObj.handle = unsafe.Pointer(C.color_tile_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TColorTileCast(widget TWidget) TColorTile {
  retObj := TColorTile{}
  retObj.handle = unsafe.Pointer(C.color_tile_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TColorTile) SetBgColor(color string) TRet {
  acolor := C.CString(color)
  defer C.free(unsafe.Pointer(acolor))
  return TRet(C.color_tile_set_bg_color((*C.widget_t)(this.handle), acolor));
}

func (this TColorTile) GetBgColor() string {
  return C.GoString(C.color_tile_get_bg_color((*C.widget_t)(this.handle)));
}

func (this TColorTile) GetBorderColor() string {
  return C.GoString(C.color_tile_get_border_color((*C.widget_t)(this.handle)));
}

type TColumn struct {
  TWidget
}

func TColumnCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TColumn{}
  retObj.handle = unsafe.Pointer(C.column_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TColumnCast(widget TWidget) TColumn {
  retObj := TColumn{}
  retObj.handle = unsafe.Pointer(C.column_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TComboBox struct {
  TEdit
}

func TComboBoxCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TComboBox{}
  retObj.handle = unsafe.Pointer(C.combo_box_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TComboBoxCast(widget TWidget) TComboBox {
  retObj := TComboBox{}
  retObj.handle = unsafe.Pointer(C.combo_box_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TComboBox) SetOpenWindow(open_window string) TRet {
  aopen_window := C.CString(open_window)
  defer C.free(unsafe.Pointer(aopen_window))
  return TRet(C.combo_box_set_open_window((*C.widget_t)(this.handle), aopen_window));
}

func (this TComboBox) SetThemeOfPopup(theme_of_popup string) TRet {
  atheme_of_popup := C.CString(theme_of_popup)
  defer C.free(unsafe.Pointer(atheme_of_popup))
  return TRet(C.combo_box_set_theme_of_popup((*C.widget_t)(this.handle), atheme_of_popup));
}

func (this TComboBox) ResetOptions() TRet {
  return TRet(C.combo_box_reset_options((*C.widget_t)(this.handle)));
}

func (this TComboBox) CountOptions() int32 {
  return (int32)(C.combo_box_count_options((*C.widget_t)(this.handle)));
}

func (this TComboBox) SetSelectedIndex(index uint32) TRet {
  return TRet(C.combo_box_set_selected_index((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TComboBox) SetSelectedIndexByText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.combo_box_set_selected_index_by_text((*C.widget_t)(this.handle), atext));
}

func (this TComboBox) SetLocalizeOptions(localize_options bool) TRet {
  return TRet(C.combo_box_set_localize_options((*C.widget_t)(this.handle), (C.bool_t)(localize_options)));
}

func (this TComboBox) SetValue(value int32) TRet {
  return TRet(C.combo_box_set_value((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TComboBox) SetItemHeight(item_height uint32) TRet {
  return TRet(C.combo_box_set_item_height((*C.widget_t)(this.handle), (C.uint32_t)(item_height)));
}

func (this TComboBox) AppendOption(value int32, text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.combo_box_append_option((*C.widget_t)(this.handle), (C.int32_t)(value), atext));
}

func (this TComboBox) RemoveOption(value int32) TRet {
  return TRet(C.combo_box_remove_option((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TComboBox) RemoveOptionByIndex(index uint32) TRet {
  return TRet(C.combo_box_remove_option_by_index((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TComboBox) SetOptions(options string) TRet {
  aoptions := C.CString(options)
  defer C.free(unsafe.Pointer(aoptions))
  return TRet(C.combo_box_set_options((*C.widget_t)(this.handle), aoptions));
}

func (this TComboBox) GetValueInt() int32 {
  return (int32)(C.combo_box_get_value((*C.widget_t)(this.handle)));
}

func (this TComboBox) HasOptionText(text string) bool {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return (bool)(C.combo_box_has_option_text((*C.widget_t)(this.handle), atext));
}

func (this TComboBox) GetTextValue() string {
  return C.GoString(C.combo_box_get_text((*C.widget_t)(this.handle)));
}

func (this TComboBox) GetTextOfSelected() string {
  return C.GoString(C.combo_box_get_text_of_selected((*C.widget_t)(this.handle)));
}

func (this TComboBox) GetOpenWindow() string {
  return C.GoString((*C.combo_box_t)(unsafe.Pointer(this.handle)).open_window);
}

func (this TComboBox) GetThemeOfPopup() string {
  return C.GoString((*C.combo_box_t)(unsafe.Pointer(this.handle)).theme_of_popup);
}

func (this TComboBox) GetSelectedIndex() int32 {
  return (int32)((*C.combo_box_t)(unsafe.Pointer(this.handle)).selected_index);
}

func (this TComboBox) GetLocalizeOptions() bool {
  return (bool)((*C.combo_box_t)(unsafe.Pointer(this.handle)).localize_options);
}

func (this TComboBox) GetOptions() string {
  return C.GoString((*C.combo_box_t)(unsafe.Pointer(this.handle)).options);
}

func (this TComboBox) GetItemHeight() int32 {
  return (int32)((*C.combo_box_t)(unsafe.Pointer(this.handle)).item_height);
}

type TComboBoxEx struct {
  TComboBox
}

func TComboBoxExCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TComboBoxEx{}
  retObj.handle = unsafe.Pointer(C.combo_box_ex_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

type TComboBoxItem struct {
  TWidget
}

func TComboBoxItemCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TComboBoxItem{}
  retObj.handle = unsafe.Pointer(C.combo_box_item_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TComboBoxItemCast(widget TWidget) TComboBoxItem {
  retObj := TComboBoxItem{}
  retObj.handle = unsafe.Pointer(C.combo_box_item_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TComboBoxItem) SetChecked(checked bool) TRet {
  return TRet(C.combo_box_item_set_checked((*C.widget_t)(this.handle), (C.bool_t)(checked)));
}

func (this TComboBoxItem) SetValue(value int32) TRet {
  return TRet(C.combo_box_item_set_value((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TComboBoxItem) GetChecked() bool {
  return (bool)((*C.combo_box_item_t)(unsafe.Pointer(this.handle)).checked);
}

type TDateTime struct {
  handle unsafe.Pointer
}

func TDateTimeCreate() TDateTime {
  retObj := TDateTime{}
  retObj.handle = unsafe.Pointer(C.date_time_create())
  return retObj
}

func (this TDateTime) SetYear(year uint32) TRet {
  return TRet(C.date_time_set_year((*C.date_time_t)(this.handle), (C.uint32_t)(year)));
}

func (this TDateTime) SetMonth(month uint32) TRet {
  return TRet(C.date_time_set_month((*C.date_time_t)(this.handle), (C.uint32_t)(month)));
}

func (this TDateTime) SetDay(day uint32) TRet {
  return TRet(C.date_time_set_day((*C.date_time_t)(this.handle), (C.uint32_t)(day)));
}

func (this TDateTime) SetHour(hour uint32) TRet {
  return TRet(C.date_time_set_hour((*C.date_time_t)(this.handle), (C.uint32_t)(hour)));
}

func (this TDateTime) SetMinute(minute uint32) TRet {
  return TRet(C.date_time_set_minute((*C.date_time_t)(this.handle), (C.uint32_t)(minute)));
}

func (this TDateTime) SetSecond(second uint32) TRet {
  return TRet(C.date_time_set_second((*C.date_time_t)(this.handle), (C.uint32_t)(second)));
}

func (this TDateTime) Set() TRet {
  return TRet(C.date_time_set((*C.date_time_t)(this.handle)));
}

func (this TDateTime) FromTime(time int64) TRet {
  return TRet(C.date_time_from_time((*C.date_time_t)(this.handle), (C.int64_t)(time)));
}

func (this TDateTime) ToTime() int64 {
  return (int64)(C.date_time_to_time((*C.date_time_t)(this.handle)));
}

func (this TDateTime) AddDelta(delta int64) TRet {
  return TRet(C.date_time_add_delta((*C.date_time_t)(this.handle), (C.int64_t)(delta)));
}

func TDateTimeIsLeap(year uint32) bool {
  return (bool)(C.date_time_is_leap((C.uint32_t)(year)));
}

func TDateTimeGetDays(year uint32, month uint32) int32 {
  return (int32)(C.date_time_get_days((C.uint32_t)(year), (C.uint32_t)(month)));
}

func TDateTimeGetWday(year uint32, month uint32, day uint32) int32 {
  return (int32)(C.date_time_get_wday((C.uint32_t)(year), (C.uint32_t)(month), (C.uint32_t)(day)));
}

func TDateTimeGetMonthName(month uint32) string {
  return C.GoString(C.date_time_get_month_name((C.uint32_t)(month)));
}

func TDateTimeGetWdayName(wday uint32) string {
  return C.GoString(C.date_time_get_wday_name((C.uint32_t)(wday)));
}

func (this TDateTime) Destroy() TRet {
  return TRet(C.date_time_destroy((*C.date_time_t)(this.handle)));
}

func (this TDateTime) GetSecond() int32 {
  return (int32)((*C.date_time_t)(unsafe.Pointer(this.handle)).second);
}

func (this TDateTime) GetMinute() int32 {
  return (int32)((*C.date_time_t)(unsafe.Pointer(this.handle)).minute);
}

func (this TDateTime) GetHour() int32 {
  return (int32)((*C.date_time_t)(unsafe.Pointer(this.handle)).hour);
}

func (this TDateTime) GetDay() int32 {
  return (int32)((*C.date_time_t)(unsafe.Pointer(this.handle)).day);
}

func (this TDateTime) GetMonth() int32 {
  return (int32)((*C.date_time_t)(unsafe.Pointer(this.handle)).month);
}

func (this TDateTime) GetYear() int32 {
  return (int32)((*C.date_time_t)(unsafe.Pointer(this.handle)).year);
}

type TDialog struct {
  TWindowBase
}

func TDialogCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDialog{}
  retObj.handle = unsafe.Pointer(C.dialog_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDialogCreateSimple(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDialog{}
  retObj.handle = unsafe.Pointer(C.dialog_create_simple((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDialogCast(widget TWidget) TDialog {
  retObj := TDialog{}
  retObj.handle = unsafe.Pointer(C.dialog_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TDialog) GetTitle() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.dialog_get_title((*C.widget_t)(this.handle)))
  return retObj
}

func (this TDialog) GetClient() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.dialog_get_client((*C.widget_t)(this.handle)))
  return retObj
}

func TDialogOpen(name string) TWidget {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TDialog{}
  retObj.handle = unsafe.Pointer(C.dialog_open(aname))
  return retObj.TWidget
}

func (this TDialog) SetTitle(title string) TRet {
  atitle := C.CString(title)
  defer C.free(unsafe.Pointer(atitle))
  return TRet(C.dialog_set_title((*C.widget_t)(this.handle), atitle));
}

func (this TDialog) Modal() TDialogQuitCode {
  return TDialogQuitCode(C.dialog_modal((*C.widget_t)(this.handle)));
}

func (this TDialog) Quit(code uint32) TRet {
  return TRet(C.dialog_quit((*C.widget_t)(this.handle), (C.uint32_t)(code)));
}

func (this TDialog) IsQuited() bool {
  return (bool)(C.dialog_is_quited((*C.widget_t)(this.handle)));
}

func (this TDialog) IsModal() bool {
  return (bool)(C.dialog_is_modal((*C.widget_t)(this.handle)));
}

func TDialogToast(text string, duration uint32) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.dialog_toast(atext, (C.uint32_t)(duration)));
}

func TDialogInfo(title string, text string) TRet {
  atitle := C.CString(title)
  defer C.free(unsafe.Pointer(atitle))
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.dialog_info(atitle, atext));
}

func TDialogWarn(title string, text string) TRet {
  atitle := C.CString(title)
  defer C.free(unsafe.Pointer(atitle))
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.dialog_warn(atitle, atext));
}

func TDialogConfirm(title string, text string) TRet {
  atitle := C.CString(title)
  defer C.free(unsafe.Pointer(atitle))
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.dialog_confirm(atitle, atext));
}

func (this TDialog) GetHighlight() string {
  return C.GoString((*C.dialog_t)(unsafe.Pointer(this.handle)).highlight);
}

type TDialogClient struct {
  TWidget
}

func TDialogClientCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDialogClient{}
  retObj.handle = unsafe.Pointer(C.dialog_client_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDialogClientCast(widget TWidget) TDialogClient {
  retObj := TDialogClient{}
  retObj.handle = unsafe.Pointer(C.dialog_client_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TDialogQuitCode int
const (
  DIALOG_QUIT_NONE TDialogQuitCode = C.DIALOG_QUIT_NONE
  DIALOG_QUIT_OK TDialogQuitCode = C.DIALOG_QUIT_OK
  DIALOG_QUIT_YES TDialogQuitCode = C.DIALOG_QUIT_YES
  DIALOG_QUIT_CANCEL TDialogQuitCode = C.DIALOG_QUIT_CANCEL
  DIALOG_QUIT_NO TDialogQuitCode = C.DIALOG_QUIT_NO
  DIALOG_QUIT_OTHER TDialogQuitCode = C.DIALOG_QUIT_OTHER
)
type TDialogTitle struct {
  TWidget
}

func TDialogTitleCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDialogTitle{}
  retObj.handle = unsafe.Pointer(C.dialog_title_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDialogTitleCast(widget TWidget) TDialogTitle {
  retObj := TDialogTitle{}
  retObj.handle = unsafe.Pointer(C.dialog_title_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TDigitClock struct {
  TWidget
}

func TDigitClockCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDigitClock{}
  retObj.handle = unsafe.Pointer(C.digit_clock_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDigitClockCast(widget TWidget) TDigitClock {
  retObj := TDigitClock{}
  retObj.handle = unsafe.Pointer(C.digit_clock_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TDigitClock) SetFormat(format string) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.digit_clock_set_format((*C.widget_t)(this.handle), aformat));
}

func (this TDigitClock) GetFormat() string {
  return C.GoString((*C.digit_clock_t)(unsafe.Pointer(this.handle)).format);
}

type TDoneEvent struct {
  TEvent
}

func TDoneEventCast(event TEvent) TDoneEvent {
  retObj := TDoneEvent{}
  retObj.handle = unsafe.Pointer(C.done_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TDoneEvent) GetResult() TRet {
  return TRet((*C.done_event_t)(unsafe.Pointer(this.handle)).result);
}

type TDraggable struct {
  TWidget
}

func TDraggableCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDraggable{}
  retObj.handle = unsafe.Pointer(C.draggable_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDraggableCast(widget TWidget) TDraggable {
  retObj := TDraggable{}
  retObj.handle = unsafe.Pointer(C.draggable_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TDraggable) SetTop(top int32) TRet {
  return TRet(C.draggable_set_top((*C.widget_t)(this.handle), (C.int32_t)(top)));
}

func (this TDraggable) SetBottom(bottom int32) TRet {
  return TRet(C.draggable_set_bottom((*C.widget_t)(this.handle), (C.int32_t)(bottom)));
}

func (this TDraggable) SetLeft(left int32) TRet {
  return TRet(C.draggable_set_left((*C.widget_t)(this.handle), (C.int32_t)(left)));
}

func (this TDraggable) SetRight(right int32) TRet {
  return TRet(C.draggable_set_right((*C.widget_t)(this.handle), (C.int32_t)(right)));
}

func (this TDraggable) SetVerticalOnly(vertical_only bool) TRet {
  return TRet(C.draggable_set_vertical_only((*C.widget_t)(this.handle), (C.bool_t)(vertical_only)));
}

func (this TDraggable) SetHorizontalOnly(horizontal_only bool) TRet {
  return TRet(C.draggable_set_horizontal_only((*C.widget_t)(this.handle), (C.bool_t)(horizontal_only)));
}

func (this TDraggable) SetAllowOutOfScreen(allow_out_of_screen bool) TRet {
  return TRet(C.draggable_set_allow_out_of_screen((*C.widget_t)(this.handle), (C.bool_t)(allow_out_of_screen)));
}

func (this TDraggable) SetDragWindow(drag_window bool) TRet {
  return TRet(C.draggable_set_drag_window((*C.widget_t)(this.handle), (C.bool_t)(drag_window)));
}

func (this TDraggable) SetDragNativeWindow(drag_native_window bool) TRet {
  return TRet(C.draggable_set_drag_native_window((*C.widget_t)(this.handle), (C.bool_t)(drag_native_window)));
}

func (this TDraggable) SetDragParent(drag_parent uint32) TRet {
  return TRet(C.draggable_set_drag_parent((*C.widget_t)(this.handle), (C.uint32_t)(drag_parent)));
}

func (this TDraggable) GetTop() int32 {
  return (int32)((*C.draggable_t)(unsafe.Pointer(this.handle)).top);
}

func (this TDraggable) GetBottom() int32 {
  return (int32)((*C.draggable_t)(unsafe.Pointer(this.handle)).bottom);
}

func (this TDraggable) GetLeft() int32 {
  return (int32)((*C.draggable_t)(unsafe.Pointer(this.handle)).left);
}

func (this TDraggable) GetRight() int32 {
  return (int32)((*C.draggable_t)(unsafe.Pointer(this.handle)).right);
}

func (this TDraggable) GetAllowOutOfScreen() bool {
  return (bool)((*C.draggable_t)(unsafe.Pointer(this.handle)).allow_out_of_screen);
}

func (this TDraggable) GetVerticalOnly() bool {
  return (bool)((*C.draggable_t)(unsafe.Pointer(this.handle)).vertical_only);
}

func (this TDraggable) GetHorizontalOnly() bool {
  return (bool)((*C.draggable_t)(unsafe.Pointer(this.handle)).horizontal_only);
}

func (this TDraggable) GetDragWindow() bool {
  return (bool)((*C.draggable_t)(unsafe.Pointer(this.handle)).drag_window);
}

func (this TDraggable) GetDragNativeWindow() bool {
  return (bool)((*C.draggable_t)(unsafe.Pointer(this.handle)).drag_native_window);
}

func (this TDraggable) GetDragParent() uint32 {
  return (uint32)((*C.draggable_t)(unsafe.Pointer(this.handle)).drag_parent);
}

type TDragger struct {
  TWidget
}

func TDraggerCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TDragger{}
  retObj.handle = unsafe.Pointer(C.dragger_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TDraggerCast(widget TWidget) TDragger {
  retObj := TDragger{}
  retObj.handle = unsafe.Pointer(C.dragger_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TDragger) SetRange(x_min int, y_min int, x_max int, y_max int) TRet {
  return TRet(C.dragger_set_range((*C.widget_t)(this.handle), (C.xy_t)(x_min), (C.xy_t)(y_min), (C.xy_t)(x_max), (C.xy_t)(y_max)));
}

func (this TDragger) GetXMin() int {
  return (int)((*C.dragger_t)(unsafe.Pointer(this.handle)).x_min);
}

func (this TDragger) GetYMin() int {
  return (int)((*C.dragger_t)(unsafe.Pointer(this.handle)).y_min);
}

func (this TDragger) GetXMax() int {
  return (int)((*C.dragger_t)(unsafe.Pointer(this.handle)).x_max);
}

func (this TDragger) GetYMax() int {
  return (int)((*C.dragger_t)(unsafe.Pointer(this.handle)).y_max);
}

type TDropFileEvent struct {
  TEvent
}

func TDropFileEventCast(event TEvent) TDropFileEvent {
  retObj := TDropFileEvent{}
  retObj.handle = unsafe.Pointer(C.drop_file_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TDropFileEvent) GetFilename() string {
  return C.GoString((*C.drop_file_event_t)(unsafe.Pointer(this.handle)).filename);
}

type TEasingType int
const (
  EASING_LINEAR TEasingType = C.EASING_LINEAR
  EASING_QUADRATIC_IN TEasingType = C.EASING_QUADRATIC_IN
  EASING_QUADRATIC_OUT TEasingType = C.EASING_QUADRATIC_OUT
  EASING_QUADRATIC_INOUT TEasingType = C.EASING_QUADRATIC_INOUT
  EASING_CUBIC_IN TEasingType = C.EASING_CUBIC_IN
  EASING_CUBIC_OUT TEasingType = C.EASING_CUBIC_OUT
  EASING_SIN_IN TEasingType = C.EASING_SIN_IN
  EASING_SIN_OUT TEasingType = C.EASING_SIN_OUT
  EASING_SIN_INOUT TEasingType = C.EASING_SIN_INOUT
  EASING_POW_IN TEasingType = C.EASING_POW_IN
  EASING_POW_OUT TEasingType = C.EASING_POW_OUT
  EASING_POW_INOUT TEasingType = C.EASING_POW_INOUT
  EASING_CIRCULAR_IN TEasingType = C.EASING_CIRCULAR_IN
  EASING_CIRCULAR_OUT TEasingType = C.EASING_CIRCULAR_OUT
  EASING_CIRCULAR_INOUT TEasingType = C.EASING_CIRCULAR_INOUT
  EASING_ELASTIC_IN TEasingType = C.EASING_ELASTIC_IN
  EASING_ELASTIC_OUT TEasingType = C.EASING_ELASTIC_OUT
  EASING_ELASTIC_INOUT TEasingType = C.EASING_ELASTIC_INOUT
  EASING_BACK_IN TEasingType = C.EASING_BACK_IN
  EASING_BACK_OUT TEasingType = C.EASING_BACK_OUT
  EASING_BACK_INOUT TEasingType = C.EASING_BACK_INOUT
  EASING_BOUNCE_IN TEasingType = C.EASING_BOUNCE_IN
  EASING_BOUNCE_OUT TEasingType = C.EASING_BOUNCE_OUT
  EASING_BOUNCE_INOUT TEasingType = C.EASING_BOUNCE_INOUT
)
type TEdit struct {
  TWidget
}

func TEditCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TEdit{}
  retObj.handle = unsafe.Pointer(C.edit_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TEditCast(widget TWidget) TEdit {
  retObj := TEdit{}
  retObj.handle = unsafe.Pointer(C.edit_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TEdit) GetInt() int32 {
  return (int32)(C.edit_get_int((*C.widget_t)(this.handle)));
}

func (this TEdit) GetDouble() float64 {
  return (float64)(C.edit_get_double((*C.widget_t)(this.handle)));
}

func (this TEdit) SetInt(value int32) TRet {
  return TRet(C.edit_set_int((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TEdit) SetDouble(value float64) TRet {
  return TRet(C.edit_set_double((*C.widget_t)(this.handle), (C.double)(value)));
}

func (this TEdit) SetDoubleEx(format string, value float64) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.edit_set_double_ex((*C.widget_t)(this.handle), aformat, (C.double)(value)));
}

func (this TEdit) SetTextLimit(min uint32, max uint32) TRet {
  return TRet(C.edit_set_text_limit((*C.widget_t)(this.handle), (C.uint32_t)(min), (C.uint32_t)(max)));
}

func (this TEdit) SetIntLimit(min int32, max int32, step uint32) TRet {
  return TRet(C.edit_set_int_limit((*C.widget_t)(this.handle), (C.int32_t)(min), (C.int32_t)(max), (C.uint32_t)(step)));
}

func (this TEdit) SetFloatLimit(min float64, max float64, step float64) TRet {
  return TRet(C.edit_set_float_limit((*C.widget_t)(this.handle), (C.double)(min), (C.double)(max), (C.double)(step)));
}

func (this TEdit) SetReadonly(readonly bool) TRet {
  return TRet(C.edit_set_readonly((*C.widget_t)(this.handle), (C.bool_t)(readonly)));
}

func (this TEdit) SetCancelable(cancelable bool) TRet {
  return TRet(C.edit_set_cancelable((*C.widget_t)(this.handle), (C.bool_t)(cancelable)));
}

func (this TEdit) SetAutoFix(auto_fix bool) TRet {
  return TRet(C.edit_set_auto_fix((*C.widget_t)(this.handle), (C.bool_t)(auto_fix)));
}

func (this TEdit) SetSelectNoneWhenFocused(select_none_when_focused bool) TRet {
  return TRet(C.edit_set_select_none_when_focused((*C.widget_t)(this.handle), (C.bool_t)(select_none_when_focused)));
}

func (this TEdit) SetOpenImWhenFocused(open_im_when_focused bool) TRet {
  return TRet(C.edit_set_open_im_when_focused((*C.widget_t)(this.handle), (C.bool_t)(open_im_when_focused)));
}

func (this TEdit) SetCloseImWhenBlured(close_im_when_blured bool) TRet {
  return TRet(C.edit_set_close_im_when_blured((*C.widget_t)(this.handle), (C.bool_t)(close_im_when_blured)));
}

func (this TEdit) SetInputType(typex TInputType) TRet {
  return TRet(C.edit_set_input_type((*C.widget_t)(this.handle), (C.input_type_t)(typex)));
}

func (this TEdit) SetActionText(action_text string) TRet {
  aaction_text := C.CString(action_text)
  defer C.free(unsafe.Pointer(aaction_text))
  return TRet(C.edit_set_action_text((*C.widget_t)(this.handle), aaction_text));
}

func (this TEdit) SetTips(tips string) TRet {
  atips := C.CString(tips)
  defer C.free(unsafe.Pointer(atips))
  return TRet(C.edit_set_tips((*C.widget_t)(this.handle), atips));
}

func (this TEdit) SetTrTips(tr_tips string) TRet {
  atr_tips := C.CString(tr_tips)
  defer C.free(unsafe.Pointer(atr_tips))
  return TRet(C.edit_set_tr_tips((*C.widget_t)(this.handle), atr_tips));
}

func (this TEdit) SetKeyboard(keyboard string) TRet {
  akeyboard := C.CString(keyboard)
  defer C.free(unsafe.Pointer(akeyboard))
  return TRet(C.edit_set_keyboard((*C.widget_t)(this.handle), akeyboard));
}

func (this TEdit) SetPasswordVisible(password_visible bool) TRet {
  return TRet(C.edit_set_password_visible((*C.widget_t)(this.handle), (C.bool_t)(password_visible)));
}

func (this TEdit) SetFocus(focus bool) TRet {
  return TRet(C.edit_set_focus((*C.widget_t)(this.handle), (C.bool_t)(focus)));
}

func (this TEdit) SetCursor(cursor uint32) TRet {
  return TRet(C.edit_set_cursor((*C.widget_t)(this.handle), (C.uint32_t)(cursor)));
}

func (this TEdit) GetCursor() uint32 {
  return (uint32)(C.edit_get_cursor((*C.widget_t)(this.handle)));
}

func (this TEdit) SetSelect(start uint32, end uint32) TRet {
  return TRet(C.edit_set_select((*C.widget_t)(this.handle), (C.uint32_t)(start), (C.uint32_t)(end)));
}

func (this TEdit) GetSelectedText() string {
  return C.GoString(C.edit_get_selected_text((*C.widget_t)(this.handle)));
}

func (this TEdit) SetFocusNextWhenEnter(focus_next_when_enter bool) TRet {
  return TRet(C.edit_set_focus_next_when_enter((*C.widget_t)(this.handle), (C.bool_t)(focus_next_when_enter)));
}

func (this TEdit) GetTips() string {
  return C.GoString((*C.edit_t)(unsafe.Pointer(this.handle)).tips);
}

func (this TEdit) GetTrTips() string {
  return C.GoString((*C.edit_t)(unsafe.Pointer(this.handle)).tr_tips);
}

func (this TEdit) GetActionText() string {
  return C.GoString((*C.edit_t)(unsafe.Pointer(this.handle)).action_text);
}

func (this TEdit) GetValidator() string {
  return C.GoString((*C.edit_t)(unsafe.Pointer(this.handle)).validator);
}

func (this TEdit) GetKeyboard() string {
  return C.GoString((*C.edit_t)(unsafe.Pointer(this.handle)).keyboard);
}

func (this TEdit) GetMin() float64 {
  return (float64)((*C.edit_t)(unsafe.Pointer(this.handle)).min);
}

func (this TEdit) GetMax() float64 {
  return (float64)((*C.edit_t)(unsafe.Pointer(this.handle)).max);
}

func (this TEdit) GetStep() float64 {
  return (float64)((*C.edit_t)(unsafe.Pointer(this.handle)).step);
}

func (this TEdit) GetInputType() TInputType {
  return TInputType((*C.edit_t)(unsafe.Pointer(this.handle)).input_type);
}

func (this TEdit) GetReadonly() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).readonly);
}

func (this TEdit) GetPasswordVisible() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).password_visible);
}

func (this TEdit) GetAutoFix() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).auto_fix);
}

func (this TEdit) GetSelectNoneWhenFocused() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).select_none_when_focused);
}

func (this TEdit) GetOpenImWhenFocused() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).open_im_when_focused);
}

func (this TEdit) GetCloseImWhenBlured() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).close_im_when_blured);
}

func (this TEdit) GetCancelable() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).cancelable);
}

func (this TEdit) GetFocusNextWhenEnter() bool {
  return (bool)((*C.edit_t)(unsafe.Pointer(this.handle)).focus_next_when_enter);
}

type TEmitter struct {
  handle unsafe.Pointer
}

func TEmitterCreate() TEmitter {
  retObj := TEmitter{}
  retObj.handle = unsafe.Pointer(C.emitter_create())
  return retObj
}

func (this TEmitter) Dispatch(e TEvent) TRet {
  return TRet(C.emitter_dispatch((*C.emitter_t)(this.handle), (*C.event_t)(e.handle)));
}

func (this TEmitter) DispatchSimpleEvent(typex uint32) TRet {
  return TRet(C.emitter_dispatch_simple_event((*C.emitter_t)(this.handle), (C.uint32_t)(typex)));
}

func (this TEmitter) Off(id uint32) TRet {
  return TRet(C.emitter_off((*C.emitter_t)(this.handle), (C.uint32_t)(id)));
}

func (this TEmitter) Enable() TRet {
  return TRet(C.emitter_enable((*C.emitter_t)(this.handle)));
}

func (this TEmitter) Disable() TRet {
  return TRet(C.emitter_disable((*C.emitter_t)(this.handle)));
}

func (this TEmitter) Destroy() TRet {
  return TRet(C.emitter_destroy((*C.emitter_t)(this.handle)));
}

func TEmitterCast(emitter TEmitter) TEmitter {
  retObj := TEmitter{}
  retObj.handle = unsafe.Pointer(C.emitter_cast((*C.emitter_t)(emitter.handle)))
  return retObj
}

type TErrorEvent struct {
  TEvent
}

func TErrorEventCast(event TEvent) TErrorEvent {
  retObj := TErrorEvent{}
  retObj.handle = unsafe.Pointer(C.error_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TErrorEvent) GetCode() int32 {
  return (int32)((*C.error_event_t)(unsafe.Pointer(this.handle)).code);
}

func (this TErrorEvent) GetMessage() string {
  return C.GoString((*C.error_event_t)(unsafe.Pointer(this.handle)).message);
}

type TEvent struct {
  handle unsafe.Pointer
}

func TEventFromName(name string) int32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int32)(C.event_from_name(aname));
}

func TEventRegisterCustomName(event_type int32, name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.event_register_custom_name((C.int32_t)(event_type), aname));
}

func TEventUnregisterCustomName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.event_unregister_custom_name(aname));
}

func TEventCast(event TEvent) TEvent {
  retObj := TEvent{}
  retObj.handle = unsafe.Pointer(C.event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TEvent) GetType() uint32 {
  return (uint32)(C.event_get_type((*C.event_t)(this.handle)));
}

func TEventCreate(typex uint32) TEvent {
  retObj := TEvent{}
  retObj.handle = unsafe.Pointer(C.event_create((C.uint32_t)(typex)))
  return retObj
}

func (this TEvent) Destroy() TRet {
  return TRet(C.event_destroy((*C.event_t)(this.handle)));
}

func (this TEvent) GetSize() uint32 {
  return (uint32)((*C.event_t)(unsafe.Pointer(this.handle)).size);
}

func (this TEvent) GetTime() int64 {
  return (int64)((*C.event_t)(unsafe.Pointer(this.handle)).time);
}

func (this TEvent) GetTarget() unsafe.Pointer {
  return (unsafe.Pointer)((*C.event_t)(unsafe.Pointer(this.handle)).target);
}

type TEventType int
const (
  EVT_POINTER_DOWN TEventType = C.EVT_POINTER_DOWN
  EVT_POINTER_DOWN_BEFORE_CHILDREN TEventType = C.EVT_POINTER_DOWN_BEFORE_CHILDREN
  EVT_POINTER_MOVE TEventType = C.EVT_POINTER_MOVE
  EVT_POINTER_MOVE_BEFORE_CHILDREN TEventType = C.EVT_POINTER_MOVE_BEFORE_CHILDREN
  EVT_POINTER_UP TEventType = C.EVT_POINTER_UP
  EVT_POINTER_UP_BEFORE_CHILDREN TEventType = C.EVT_POINTER_UP_BEFORE_CHILDREN
  EVT_WHEEL TEventType = C.EVT_WHEEL
  EVT_WHEEL_BEFORE_CHILDREN TEventType = C.EVT_WHEEL_BEFORE_CHILDREN
  EVT_POINTER_DOWN_ABORT TEventType = C.EVT_POINTER_DOWN_ABORT
  EVT_CONTEXT_MENU TEventType = C.EVT_CONTEXT_MENU
  EVT_POINTER_ENTER TEventType = C.EVT_POINTER_ENTER
  EVT_POINTER_LEAVE TEventType = C.EVT_POINTER_LEAVE
  EVT_LONG_PRESS TEventType = C.EVT_LONG_PRESS
  EVT_CLICK TEventType = C.EVT_CLICK
  EVT_DOUBLE_CLICK TEventType = C.EVT_DOUBLE_CLICK
  EVT_FOCUS TEventType = C.EVT_FOCUS
  EVT_BLUR TEventType = C.EVT_BLUR
  EVT_KEY_DOWN TEventType = C.EVT_KEY_DOWN
  EVT_KEY_LONG_PRESS TEventType = C.EVT_KEY_LONG_PRESS
  EVT_KEY_DOWN_BEFORE_CHILDREN TEventType = C.EVT_KEY_DOWN_BEFORE_CHILDREN
  EVT_KEY_REPEAT TEventType = C.EVT_KEY_REPEAT
  EVT_KEY_UP TEventType = C.EVT_KEY_UP
  EVT_KEY_UP_BEFORE_CHILDREN TEventType = C.EVT_KEY_UP_BEFORE_CHILDREN
  EVT_WILL_MOVE TEventType = C.EVT_WILL_MOVE
  EVT_MOVE TEventType = C.EVT_MOVE
  EVT_WILL_RESIZE TEventType = C.EVT_WILL_RESIZE
  EVT_RESIZE TEventType = C.EVT_RESIZE
  EVT_WILL_MOVE_RESIZE TEventType = C.EVT_WILL_MOVE_RESIZE
  EVT_MOVE_RESIZE TEventType = C.EVT_MOVE_RESIZE
  EVT_PAINT TEventType = C.EVT_PAINT
  EVT_BEFORE_PAINT TEventType = C.EVT_BEFORE_PAINT
  EVT_AFTER_PAINT TEventType = C.EVT_AFTER_PAINT
  EVT_PAINT_DONE TEventType = C.EVT_PAINT_DONE
  EVT_LOCALE_CHANGED TEventType = C.EVT_LOCALE_CHANGED
  EVT_ANIM_START TEventType = C.EVT_ANIM_START
  EVT_ANIM_STOP TEventType = C.EVT_ANIM_STOP
  EVT_ANIM_PAUSE TEventType = C.EVT_ANIM_PAUSE
  EVT_ANIM_ONCE TEventType = C.EVT_ANIM_ONCE
  EVT_ANIM_END TEventType = C.EVT_ANIM_END
  EVT_WINDOW_LOAD TEventType = C.EVT_WINDOW_LOAD
  EVT_WIDGET_LOAD TEventType = C.EVT_WIDGET_LOAD
  EVT_WINDOW_WILL_OPEN TEventType = C.EVT_WINDOW_WILL_OPEN
  EVT_WINDOW_OPEN TEventType = C.EVT_WINDOW_OPEN
  EVT_WINDOW_TO_BACKGROUND TEventType = C.EVT_WINDOW_TO_BACKGROUND
  EVT_WINDOW_TO_FOREGROUND TEventType = C.EVT_WINDOW_TO_FOREGROUND
  EVT_WINDOW_CLOSE TEventType = C.EVT_WINDOW_CLOSE
  EVT_REQUEST_CLOSE_WINDOW TEventType = C.EVT_REQUEST_CLOSE_WINDOW
  EVT_TOP_WINDOW_CHANGED TEventType = C.EVT_TOP_WINDOW_CHANGED
  EVT_IM_START TEventType = C.EVT_IM_START
  EVT_IM_STOP TEventType = C.EVT_IM_STOP
  EVT_IM_COMMIT TEventType = C.EVT_IM_COMMIT
  EVT_IM_CLEAR TEventType = C.EVT_IM_CLEAR
  EVT_IM_CANCEL TEventType = C.EVT_IM_CANCEL
  EVT_IM_PREEDIT TEventType = C.EVT_IM_PREEDIT
  EVT_IM_PREEDIT_CONFIRM TEventType = C.EVT_IM_PREEDIT_CONFIRM
  EVT_IM_PREEDIT_ABORT TEventType = C.EVT_IM_PREEDIT_ABORT
  EVT_IM_SHOW_CANDIDATES TEventType = C.EVT_IM_SHOW_CANDIDATES
  EVT_IM_SHOW_PRE_CANDIDATES TEventType = C.EVT_IM_SHOW_PRE_CANDIDATES
  EVT_IM_LANG_CHANGED TEventType = C.EVT_IM_LANG_CHANGED
  EVT_IM_ACTION TEventType = C.EVT_IM_ACTION
  EVT_IM_ACTION_INFO TEventType = C.EVT_IM_ACTION_INFO
  EVT_DRAG_START TEventType = C.EVT_DRAG_START
  EVT_DRAG TEventType = C.EVT_DRAG
  EVT_DRAG_END TEventType = C.EVT_DRAG_END
  EVT_RESET TEventType = C.EVT_RESET
  EVT_SCREEN_SAVER TEventType = C.EVT_SCREEN_SAVER
  EVT_LOW_MEMORY TEventType = C.EVT_LOW_MEMORY
  EVT_OUT_OF_MEMORY TEventType = C.EVT_OUT_OF_MEMORY
  EVT_ORIENTATION_WILL_CHANGED TEventType = C.EVT_ORIENTATION_WILL_CHANGED
  EVT_ORIENTATION_CHANGED TEventType = C.EVT_ORIENTATION_CHANGED
  EVT_WIDGET_CREATED TEventType = C.EVT_WIDGET_CREATED
  EVT_REQUEST_QUIT_APP TEventType = C.EVT_REQUEST_QUIT_APP
  EVT_THEME_WILL_CHANGE TEventType = C.EVT_THEME_WILL_CHANGE
  EVT_THEME_CHANGED TEventType = C.EVT_THEME_CHANGED
  EVT_WIDGET_WILL_UPDATE_STYLE TEventType = C.EVT_WIDGET_WILL_UPDATE_STYLE
  EVT_WIDGET_UPDATE_STYLE TEventType = C.EVT_WIDGET_UPDATE_STYLE
  EVT_WIDGET_ADD_CHILD TEventType = C.EVT_WIDGET_ADD_CHILD
  EVT_WIDGET_REMOVE_CHILD TEventType = C.EVT_WIDGET_REMOVE_CHILD
  EVT_SCROLL_START TEventType = C.EVT_SCROLL_START
  EVT_SCROLL TEventType = C.EVT_SCROLL
  EVT_SCROLL_END TEventType = C.EVT_SCROLL_END
  EVT_MULTI_GESTURE TEventType = C.EVT_MULTI_GESTURE
  EVT_PAGE_CHANGED TEventType = C.EVT_PAGE_CHANGED
  EVT_PAGE_CHANGING TEventType = C.EVT_PAGE_CHANGING
  EVT_ASSET_MANAGER_LOAD_ASSET TEventType = C.EVT_ASSET_MANAGER_LOAD_ASSET
  EVT_ASSET_MANAGER_UNLOAD_ASSET TEventType = C.EVT_ASSET_MANAGER_UNLOAD_ASSET
  EVT_ASSET_MANAGER_CLEAR_CACHE TEventType = C.EVT_ASSET_MANAGER_CLEAR_CACHE
  EVT_TIMER TEventType = C.EVT_TIMER
  EVT_DATA TEventType = C.EVT_DATA
  EVT_CONNECT TEventType = C.EVT_CONNECT
  EVT_MODEL_CHANGE TEventType = C.EVT_MODEL_CHANGE
  EVT_SYSTEM TEventType = C.EVT_SYSTEM
  EVT_DROP_FILE TEventType = C.EVT_DROP_FILE
  EVT_LOCALE_INFOS_LOAD_INFO TEventType = C.EVT_LOCALE_INFOS_LOAD_INFO
  EVT_LOCALE_INFOS_UNLOAD_INFO TEventType = C.EVT_LOCALE_INFOS_UNLOAD_INFO
  EVT_ACTIVATED TEventType = C.EVT_ACTIVATED
  EVT_UNACTIVATED TEventType = C.EVT_UNACTIVATED
  EVT_UI_LOAD TEventType = C.EVT_UI_LOAD
  EVT_REQ_START TEventType = C.EVT_REQ_START
  EVT_USER_START TEventType = C.EVT_USER_START
  EVT_NONE TEventType = C.EVT_NONE
  EVT_PROP_WILL_CHANGE TEventType = C.EVT_PROP_WILL_CHANGE
  EVT_PROP_CHANGED TEventType = C.EVT_PROP_CHANGED
  EVT_CMD_WILL_EXEC TEventType = C.EVT_CMD_WILL_EXEC
  EVT_CMD_EXECED TEventType = C.EVT_CMD_EXECED
  EVT_CMD_CAN_EXEC TEventType = C.EVT_CMD_CAN_EXEC
  EVT_ITEMS_WILL_CHANGE TEventType = C.EVT_ITEMS_WILL_CHANGE
  EVT_ITEMS_CHANGED TEventType = C.EVT_ITEMS_CHANGED
  EVT_PROPS_CHANGED TEventType = C.EVT_PROPS_CHANGED
  EVT_PROGRESS TEventType = C.EVT_PROGRESS
  EVT_DONE TEventType = C.EVT_DONE
  EVT_ERROR TEventType = C.EVT_ERROR
  EVT_DESTROY TEventType = C.EVT_DESTROY
  EVT_VALUE_WILL_CHANGE TEventType = C.EVT_VALUE_WILL_CHANGE
  EVT_VALUE_CHANGED TEventType = C.EVT_VALUE_CHANGED
  EVT_VALUE_CHANGING TEventType = C.EVT_VALUE_CHANGING
  EVT_LOG_MESSAGE TEventType = C.EVT_LOG_MESSAGE
)
func TExtWidgetsInit() TRet {
  return TRet(C.tk_ext_widgets_init());
}

type TFileBrowserView struct {
  TWidget
}

func TFileBrowserViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TFileBrowserView{}
  retObj.handle = unsafe.Pointer(C.file_browser_view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TFileBrowserViewCast(widget TWidget) TFileBrowserView {
  retObj := TFileBrowserView{}
  retObj.handle = unsafe.Pointer(C.file_browser_view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TFileBrowserView) SetInitDir(init_dir string) TRet {
  ainit_dir := C.CString(init_dir)
  defer C.free(unsafe.Pointer(ainit_dir))
  return TRet(C.file_browser_view_set_init_dir((*C.widget_t)(this.handle), ainit_dir));
}

func (this TFileBrowserView) SetTopDir(top_dir string) TRet {
  atop_dir := C.CString(top_dir)
  defer C.free(unsafe.Pointer(atop_dir))
  return TRet(C.file_browser_view_set_top_dir((*C.widget_t)(this.handle), atop_dir));
}

func (this TFileBrowserView) SetFilter(filter string) TRet {
  afilter := C.CString(filter)
  defer C.free(unsafe.Pointer(afilter))
  return TRet(C.file_browser_view_set_filter((*C.widget_t)(this.handle), afilter));
}

func (this TFileBrowserView) Reload() TRet {
  return TRet(C.file_browser_view_reload((*C.widget_t)(this.handle)));
}

func (this TFileBrowserView) SetIgnoreHiddenFiles(ignore_hidden_files bool) TRet {
  return TRet(C.file_browser_view_set_ignore_hidden_files((*C.widget_t)(this.handle), (C.bool_t)(ignore_hidden_files)));
}

func (this TFileBrowserView) SetSortAscending(sort_ascending bool) TRet {
  return TRet(C.file_browser_view_set_sort_ascending((*C.widget_t)(this.handle), (C.bool_t)(sort_ascending)));
}

func (this TFileBrowserView) SetShowCheckButton(show_check_button bool) TRet {
  return TRet(C.file_browser_view_set_show_check_button((*C.widget_t)(this.handle), (C.bool_t)(show_check_button)));
}

func (this TFileBrowserView) SetSortBy(sort_by string) TRet {
  asort_by := C.CString(sort_by)
  defer C.free(unsafe.Pointer(asort_by))
  return TRet(C.file_browser_view_set_sort_by((*C.widget_t)(this.handle), asort_by));
}

func (this TFileBrowserView) SetOddItemStyle(odd_item_style string) TRet {
  aodd_item_style := C.CString(odd_item_style)
  defer C.free(unsafe.Pointer(aodd_item_style))
  return TRet(C.file_browser_view_set_odd_item_style((*C.widget_t)(this.handle), aodd_item_style));
}

func (this TFileBrowserView) SetEvenItemStyle(even_item_style string) TRet {
  aeven_item_style := C.CString(even_item_style)
  defer C.free(unsafe.Pointer(aeven_item_style))
  return TRet(C.file_browser_view_set_even_item_style((*C.widget_t)(this.handle), aeven_item_style));
}

func (this TFileBrowserView) GetCwd() string {
  return C.GoString(C.file_browser_view_get_cwd((*C.widget_t)(this.handle)));
}

func (this TFileBrowserView) CreateDir(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.file_browser_view_create_dir((*C.widget_t)(this.handle), aname));
}

func (this TFileBrowserView) CreateFile(name string, data string, size uint32) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  adata := C.CString(data)
  defer C.free(unsafe.Pointer(adata))
  return TRet(C.file_browser_view_create_file((*C.widget_t)(this.handle), aname, adata, (C.uint32_t)(size)));
}

func (this TFileBrowserView) GetInitDir() string {
  return C.GoString((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).init_dir);
}

func (this TFileBrowserView) GetTopDir() string {
  return C.GoString((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).top_dir);
}

func (this TFileBrowserView) GetFilter() string {
  return C.GoString((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).filter);
}

func (this TFileBrowserView) GetIgnoreHiddenFiles() bool {
  return (bool)((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).ignore_hidden_files);
}

func (this TFileBrowserView) GetSortAscending() bool {
  return (bool)((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).sort_ascending);
}

func (this TFileBrowserView) GetShowCheckButton() bool {
  return (bool)((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).show_check_button);
}

func (this TFileBrowserView) GetSortBy() string {
  return C.GoString((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).sort_by);
}

func (this TFileBrowserView) GetOddItemStyle() string {
  return C.GoString((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).odd_item_style);
}

func (this TFileBrowserView) GetEvenItemStyle() string {
  return C.GoString((*C.file_browser_view_t)(unsafe.Pointer(this.handle)).even_item_style);
}

type TFileChooser struct {
  TEmitter
}

func TFileChooserCreate() TFileChooser {
  retObj := TFileChooser{}
  retObj.handle = unsafe.Pointer(C.file_chooser_create())
  return retObj
}

func (this TFileChooser) SetInitDir(init_dir string) TRet {
  ainit_dir := C.CString(init_dir)
  defer C.free(unsafe.Pointer(ainit_dir))
  return TRet(C.file_chooser_set_init_dir((*C.file_chooser_t)(this.handle), ainit_dir));
}

func (this TFileChooser) SetTopDir(top_dir string) TRet {
  atop_dir := C.CString(top_dir)
  defer C.free(unsafe.Pointer(atop_dir))
  return TRet(C.file_chooser_set_top_dir((*C.file_chooser_t)(this.handle), atop_dir));
}

func (this TFileChooser) SetFilter(filter string) TRet {
  afilter := C.CString(filter)
  defer C.free(unsafe.Pointer(afilter))
  return TRet(C.file_chooser_set_filter((*C.file_chooser_t)(this.handle), afilter));
}

func TFileChooserCast(chooser TFileChooser) TFileChooser {
  retObj := TFileChooser{}
  retObj.handle = unsafe.Pointer(C.file_chooser_cast((*C.file_chooser_t)(chooser.handle)))
  return retObj
}

func (this TFileChooser) ChooseFileForSave() TRet {
  return TRet(C.file_chooser_choose_file_for_save((*C.file_chooser_t)(this.handle)));
}

func (this TFileChooser) ChooseFileForOpen() TRet {
  return TRet(C.file_chooser_choose_file_for_open((*C.file_chooser_t)(this.handle)));
}

func (this TFileChooser) ChooseFolder() TRet {
  return TRet(C.file_chooser_choose_folder((*C.file_chooser_t)(this.handle)));
}

func (this TFileChooser) GetDir() string {
  return C.GoString(C.file_chooser_get_dir((*C.file_chooser_t)(this.handle)));
}

func (this TFileChooser) GetFilename() string {
  return C.GoString(C.file_chooser_get_filename((*C.file_chooser_t)(this.handle)));
}

func (this TFileChooser) IsAborted() bool {
  return (bool)(C.file_chooser_is_aborted((*C.file_chooser_t)(this.handle)));
}

type TFontManager struct {
  TEmitter
}

func (this TFontManager) SetStandardFontSize(is_standard bool) TRet {
  return TRet(C.font_manager_set_standard_font_size((*C.font_manager_t)(this.handle), (C.bool_t)(is_standard)));
}

func (this TFontManager) GetStandardFontSize() bool {
  return (bool)(C.font_manager_get_standard_font_size((*C.font_manager_t)(this.handle)));
}

func (this TFontManager) UnloadFont(name string, size int) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.font_manager_unload_font((*C.font_manager_t)(this.handle), aname, (C.font_size_t)(size)));
}

func (this TFontManager) ShrinkCache(cache_size uint32) TRet {
  return TRet(C.font_manager_shrink_cache((*C.font_manager_t)(this.handle), (C.uint32_t)(cache_size)));
}

func (this TFontManager) UnloadAll() TRet {
  return TRet(C.font_manager_unload_all((*C.font_manager_t)(this.handle)));
}

type TGauge struct {
  TWidget
}

func TGaugeCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TGauge{}
  retObj.handle = unsafe.Pointer(C.gauge_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TGaugeCast(widget TWidget) TGauge {
  retObj := TGauge{}
  retObj.handle = unsafe.Pointer(C.gauge_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TGauge) SetImage(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.gauge_set_image((*C.widget_t)(this.handle), aname));
}

func (this TGauge) SetDrawType(draw_type TImageDrawType) TRet {
  return TRet(C.gauge_set_draw_type((*C.widget_t)(this.handle), (C.image_draw_type_t)(draw_type)));
}

func (this TGauge) GetImage() string {
  return C.GoString((*C.gauge_t)(unsafe.Pointer(this.handle)).image);
}

func (this TGauge) GetDrawType() TImageDrawType {
  return TImageDrawType((*C.gauge_t)(unsafe.Pointer(this.handle)).draw_type);
}

type TGaugePointer struct {
  TWidget
}

func TGaugePointerCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TGaugePointer{}
  retObj.handle = unsafe.Pointer(C.gauge_pointer_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TGaugePointerCast(widget TWidget) TGaugePointer {
  retObj := TGaugePointer{}
  retObj.handle = unsafe.Pointer(C.gauge_pointer_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TGaugePointer) SetAngle(angle float64) TRet {
  return TRet(C.gauge_pointer_set_angle((*C.widget_t)(this.handle), (C.float_t)(angle)));
}

func (this TGaugePointer) SetImage(image string) TRet {
  aimage := C.CString(image)
  defer C.free(unsafe.Pointer(aimage))
  return TRet(C.gauge_pointer_set_image((*C.widget_t)(this.handle), aimage));
}

func (this TGaugePointer) SetAnchor(anchor_x string, anchor_y string) TRet {
  aanchor_x := C.CString(anchor_x)
  defer C.free(unsafe.Pointer(aanchor_x))
  aanchor_y := C.CString(anchor_y)
  defer C.free(unsafe.Pointer(aanchor_y))
  return TRet(C.gauge_pointer_set_anchor((*C.widget_t)(this.handle), aanchor_x, aanchor_y));
}

func (this TGaugePointer) GetAngle() float64 {
  return (float64)((*C.gauge_pointer_t)(unsafe.Pointer(this.handle)).angle);
}

func (this TGaugePointer) GetImage() string {
  return C.GoString((*C.gauge_pointer_t)(unsafe.Pointer(this.handle)).image);
}

func (this TGaugePointer) GetAnchorX() string {
  return C.GoString((*C.gauge_pointer_t)(unsafe.Pointer(this.handle)).anchor_x);
}

func (this TGaugePointer) GetAnchorY() string {
  return C.GoString((*C.gauge_pointer_t)(unsafe.Pointer(this.handle)).anchor_y);
}

type TGifImage struct {
  TImageBase
}

func TGifImageCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TGifImage{}
  retObj.handle = unsafe.Pointer(C.gif_image_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TGifImage) Play() TRet {
  return TRet(C.gif_image_play((*C.widget_t)(this.handle)));
}

func (this TGifImage) Stop() TRet {
  return TRet(C.gif_image_stop((*C.widget_t)(this.handle)));
}

func (this TGifImage) Pause() TRet {
  return TRet(C.gif_image_pause((*C.widget_t)(this.handle)));
}

func (this TGifImage) SetLoop(loop uint32) TRet {
  return TRet(C.gif_image_set_loop((*C.widget_t)(this.handle), (C.uint32_t)(loop)));
}

func TGifImageCast(widget TWidget) TGifImage {
  retObj := TGifImage{}
  retObj.handle = unsafe.Pointer(C.gif_image_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TGifImage) GetLoop() uint32 {
  return (uint32)((*C.gif_image_t)(unsafe.Pointer(this.handle)).loop);
}

func PreInit() TRet {
  return TRet(C.tk_pre_init());
}

func Init(w int, h int, app_type TAppType, app_name string, app_root string) TRet {
  aapp_name := C.CString(app_name)
  defer C.free(unsafe.Pointer(aapp_name))
  aapp_root := C.CString(app_root)
  defer C.free(unsafe.Pointer(aapp_root))
  return TRet(C.tk_init((C.wh_t)(w), (C.wh_t)(h), (C.app_type_t)(app_type), aapp_name, aapp_root));
}

func Run() TRet {
  return TRet(C.tk_run());
}

func Quit() TRet {
  return TRet(C.tk_quit());
}

func QuitEx(delay uint32) TRet {
  return TRet(C.tk_quit_ex((C.uint32_t)(delay)));
}

func GetPointerX() int32 {
  return (int32)(C.tk_get_pointer_x());
}

func GetPointerY() int32 {
  return (int32)(C.tk_get_pointer_y());
}

func IsPointerPressed() bool {
  return (bool)(C.tk_is_pointer_pressed());
}

type TGlyphFormat int
const (
  GLYPH_FMT_ALPHA TGlyphFormat = C.GLYPH_FMT_ALPHA
  GLYPH_FMT_MONO TGlyphFormat = C.GLYPH_FMT_MONO
  GLYPH_FMT_RGBA TGlyphFormat = C.GLYPH_FMT_RGBA
  GLYPH_FMT_ALPHA2 TGlyphFormat = C.GLYPH_FMT_ALPHA2
  GLYPH_FMT_ALPHA4 TGlyphFormat = C.GLYPH_FMT_ALPHA4
)
type TGrid struct {
  TWidget
}

func TGridCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TGrid{}
  retObj.handle = unsafe.Pointer(C.grid_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TGridCast(widget TWidget) TGrid {
  retObj := TGrid{}
  retObj.handle = unsafe.Pointer(C.grid_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TGrid) SetRows(rows uint32) TRet {
  return TRet(C.grid_set_rows((*C.widget_t)(this.handle), (C.uint32_t)(rows)));
}

func (this TGrid) SetColumnsDefinition(columns_definition string) TRet {
  acolumns_definition := C.CString(columns_definition)
  defer C.free(unsafe.Pointer(acolumns_definition))
  return TRet(C.grid_set_columns_definition((*C.widget_t)(this.handle), acolumns_definition));
}

func (this TGrid) SetShowGrid(show_grid bool) TRet {
  return TRet(C.grid_set_show_grid((*C.widget_t)(this.handle), (C.bool_t)(show_grid)));
}

func (this TGrid) GetRows() uint32 {
  return (uint32)((*C.grid_t)(unsafe.Pointer(this.handle)).rows);
}

func (this TGrid) GetColumnsDefinition() string {
  return C.GoString((*C.grid_t)(unsafe.Pointer(this.handle)).columns_definition);
}

func (this TGrid) GetShowGrid() bool {
  return (bool)((*C.grid_t)(unsafe.Pointer(this.handle)).show_grid);
}

type TGridItem struct {
  TWidget
}

func TGridItemCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TGridItem{}
  retObj.handle = unsafe.Pointer(C.grid_item_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TGridItemCast(widget TWidget) TGridItem {
  retObj := TGridItem{}
  retObj.handle = unsafe.Pointer(C.grid_item_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TGroupBox struct {
  TWidget
}

func TGroupBoxCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TGroupBox{}
  retObj.handle = unsafe.Pointer(C.group_box_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TGroupBox) SetValue(value uint32) TRet {
  return TRet(C.group_box_set_value((*C.widget_t)(this.handle), (C.uint32_t)(value)));
}

func TGroupBoxCast(widget TWidget) TGroupBox {
  retObj := TGroupBox{}
  retObj.handle = unsafe.Pointer(C.group_box_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type THscrollLabel struct {
  TWidget
}

func THscrollLabelCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := THscrollLabel{}
  retObj.handle = unsafe.Pointer(C.hscroll_label_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this THscrollLabel) SetLull(lull int32) TRet {
  return TRet(C.hscroll_label_set_lull((*C.widget_t)(this.handle), (C.int32_t)(lull)));
}

func (this THscrollLabel) SetDuration(duration int32) TRet {
  return TRet(C.hscroll_label_set_duration((*C.widget_t)(this.handle), (C.int32_t)(duration)));
}

func (this THscrollLabel) SetSpeed(speed float64) TRet {
  return TRet(C.hscroll_label_set_speed((*C.widget_t)(this.handle), (C.float_t)(speed)));
}

func (this THscrollLabel) SetOnlyFocus(only_focus bool) TRet {
  return TRet(C.hscroll_label_set_only_focus((*C.widget_t)(this.handle), (C.bool_t)(only_focus)));
}

func (this THscrollLabel) SetOnlyParentFocus(only_parent_focus bool) TRet {
  return TRet(C.hscroll_label_set_only_parent_focus((*C.widget_t)(this.handle), (C.bool_t)(only_parent_focus)));
}

func (this THscrollLabel) SetLoop(loop bool) TRet {
  return TRet(C.hscroll_label_set_loop((*C.widget_t)(this.handle), (C.bool_t)(loop)));
}

func (this THscrollLabel) SetYoyo(yoyo bool) TRet {
  return TRet(C.hscroll_label_set_yoyo((*C.widget_t)(this.handle), (C.bool_t)(yoyo)));
}

func (this THscrollLabel) SetEllipses(ellipses bool) TRet {
  return TRet(C.hscroll_label_set_ellipses((*C.widget_t)(this.handle), (C.bool_t)(ellipses)));
}

func (this THscrollLabel) SetStopAtBegin(stop_at_begin bool) TRet {
  return TRet(C.hscroll_label_set_stop_at_begin((*C.widget_t)(this.handle), (C.bool_t)(stop_at_begin)));
}

func (this THscrollLabel) SetDelay(delay uint32) TRet {
  return TRet(C.hscroll_label_set_delay((*C.widget_t)(this.handle), (C.uint32_t)(delay)));
}

func (this THscrollLabel) SetLoopIntervalDistance(loop_interval_distance int32) TRet {
  return TRet(C.hscroll_label_set_loop_interval_distance((*C.widget_t)(this.handle), (C.int32_t)(loop_interval_distance)));
}

func (this THscrollLabel) SetXoffset(xoffset int32) TRet {
  return TRet(C.hscroll_label_set_xoffset((*C.widget_t)(this.handle), (C.int32_t)(xoffset)));
}

func (this THscrollLabel) Start() TRet {
  return TRet(C.hscroll_label_start((*C.widget_t)(this.handle)));
}

func (this THscrollLabel) Stop() TRet {
  return TRet(C.hscroll_label_stop((*C.widget_t)(this.handle)));
}

func THscrollLabelCast(widget TWidget) THscrollLabel {
  retObj := THscrollLabel{}
  retObj.handle = unsafe.Pointer(C.hscroll_label_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this THscrollLabel) GetOnlyFocus() bool {
  return (bool)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).only_focus);
}

func (this THscrollLabel) GetOnlyParentFocus() bool {
  return (bool)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).only_parent_focus);
}

func (this THscrollLabel) GetLoop() bool {
  return (bool)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).loop);
}

func (this THscrollLabel) GetYoyo() bool {
  return (bool)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).yoyo);
}

func (this THscrollLabel) GetEllipses() bool {
  return (bool)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).ellipses);
}

func (this THscrollLabel) GetLull() int32 {
  return (int32)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).lull);
}

func (this THscrollLabel) GetDuration() int32 {
  return (int32)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).duration);
}

func (this THscrollLabel) GetDelay() uint32 {
  return (uint32)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).delay);
}

func (this THscrollLabel) GetSpeed() float64 {
  return (float64)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).speed);
}

func (this THscrollLabel) GetXoffset() int32 {
  return (int32)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).xoffset);
}

func (this THscrollLabel) GetTextW() int32 {
  return (int32)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).text_w);
}

func (this THscrollLabel) GetStopAtBegin() bool {
  return (bool)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).stop_at_begin);
}

func (this THscrollLabel) GetLoopIntervalDistance() int32 {
  return (int32)((*C.hscroll_label_t)(unsafe.Pointer(this.handle)).loop_interval_distance);
}

func TIdleRemove(idle_id uint32) TRet {
  return TRet(C.idle_remove((C.uint32_t)(idle_id)));
}

func TIdleRemoveAllByCtx(ctx unsafe.Pointer) TRet {
  return TRet(C.idle_remove_all_by_ctx((unsafe.Pointer)(ctx)));
}

type TIdleInfo struct {
  TObject
}

func TIdleInfoCast(idle TIdleInfo) TIdleInfo {
  retObj := TIdleInfo{}
  retObj.handle = unsafe.Pointer(C.idle_info_cast((*C.idle_info_t)(idle.handle)))
  return retObj
}

func (this TIdleInfo) GetCtx() unsafe.Pointer {
  return (unsafe.Pointer)((*C.idle_info_t)(unsafe.Pointer(this.handle)).ctx);
}

func (this TIdleInfo) GetExtraCtx() unsafe.Pointer {
  return (unsafe.Pointer)((*C.idle_info_t)(unsafe.Pointer(this.handle)).extra_ctx);
}

func (this TIdleInfo) GetId() uint32 {
  return (uint32)((*C.idle_info_t)(unsafe.Pointer(this.handle)).id);
}

type TIdleManager struct {
  handle unsafe.Pointer
}

type TImage struct {
  TImageBase
}

func TImageCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TImage{}
  retObj.handle = unsafe.Pointer(C.image_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TImageIconCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TImage{}
  retObj.handle = unsafe.Pointer(C.icon_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TImage) SetDrawType(draw_type TImageDrawType) TRet {
  return TRet(C.image_set_draw_type((*C.widget_t)(this.handle), (C.image_draw_type_t)(draw_type)));
}

func TImageCast(widget TWidget) TImage {
  retObj := TImage{}
  retObj.handle = unsafe.Pointer(C.image_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TImage) GetDrawType() TImageDrawType {
  return TImageDrawType((*C.image_t)(unsafe.Pointer(this.handle)).draw_type);
}

type TImageAnimation struct {
  TWidget
}

func TImageAnimationCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TImageAnimation{}
  retObj.handle = unsafe.Pointer(C.image_animation_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TImageAnimation) SetLoop(loop bool) TRet {
  return TRet(C.image_animation_set_loop((*C.widget_t)(this.handle), (C.bool_t)(loop)));
}

func (this TImageAnimation) SetImage(image string) TRet {
  aimage := C.CString(image)
  defer C.free(unsafe.Pointer(aimage))
  return TRet(C.image_animation_set_image((*C.widget_t)(this.handle), aimage));
}

func (this TImageAnimation) SetInterval(interval uint32) TRet {
  return TRet(C.image_animation_set_interval((*C.widget_t)(this.handle), (C.uint32_t)(interval)));
}

func (this TImageAnimation) SetDelay(delay uint32) TRet {
  return TRet(C.image_animation_set_delay((*C.widget_t)(this.handle), (C.uint32_t)(delay)));
}

func (this TImageAnimation) SetAutoPlay(auto_play bool) TRet {
  return TRet(C.image_animation_set_auto_play((*C.widget_t)(this.handle), (C.bool_t)(auto_play)));
}

func (this TImageAnimation) SetSequence(sequence string) TRet {
  asequence := C.CString(sequence)
  defer C.free(unsafe.Pointer(asequence))
  return TRet(C.image_animation_set_sequence((*C.widget_t)(this.handle), asequence));
}

func (this TImageAnimation) SetRangeSequence(start_index uint32, end_index uint32) TRet {
  return TRet(C.image_animation_set_range_sequence((*C.widget_t)(this.handle), (C.uint32_t)(start_index), (C.uint32_t)(end_index)));
}

func (this TImageAnimation) Play() TRet {
  return TRet(C.image_animation_play((*C.widget_t)(this.handle)));
}

func (this TImageAnimation) Stop() TRet {
  return TRet(C.image_animation_stop((*C.widget_t)(this.handle)));
}

func (this TImageAnimation) Pause() TRet {
  return TRet(C.image_animation_pause((*C.widget_t)(this.handle)));
}

func (this TImageAnimation) Next() TRet {
  return TRet(C.image_animation_next((*C.widget_t)(this.handle)));
}

func (this TImageAnimation) SetFormat(format string) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.image_animation_set_format((*C.widget_t)(this.handle), aformat));
}

func (this TImageAnimation) SetUnloadAfterPaint(unload_after_paint bool) TRet {
  return TRet(C.image_animation_set_unload_after_paint((*C.widget_t)(this.handle), (C.bool_t)(unload_after_paint)));
}

func (this TImageAnimation) SetReverse(reverse bool) TRet {
  return TRet(C.image_animation_set_reverse((*C.widget_t)(this.handle), (C.bool_t)(reverse)));
}

func (this TImageAnimation) SetShowWhenDone(show_when_done bool) TRet {
  return TRet(C.image_animation_set_show_when_done((*C.widget_t)(this.handle), (C.bool_t)(show_when_done)));
}

func TImageAnimationCast(widget TWidget) TImageAnimation {
  retObj := TImageAnimation{}
  retObj.handle = unsafe.Pointer(C.image_animation_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TImageAnimation) IsPlaying() bool {
  return (bool)(C.image_animation_is_playing((*C.widget_t)(this.handle)));
}

func (this TImageAnimation) GetImage() string {
  return C.GoString((*C.image_animation_t)(unsafe.Pointer(this.handle)).image);
}

func (this TImageAnimation) GetSequence() string {
  return C.GoString((*C.image_animation_t)(unsafe.Pointer(this.handle)).sequence);
}

func (this TImageAnimation) GetStartIndex() uint32 {
  return (uint32)((*C.image_animation_t)(unsafe.Pointer(this.handle)).start_index);
}

func (this TImageAnimation) GetEndIndex() uint32 {
  return (uint32)((*C.image_animation_t)(unsafe.Pointer(this.handle)).end_index);
}

func (this TImageAnimation) GetReverse() bool {
  return (bool)((*C.image_animation_t)(unsafe.Pointer(this.handle)).reverse);
}

func (this TImageAnimation) GetLoop() bool {
  return (bool)((*C.image_animation_t)(unsafe.Pointer(this.handle)).loop);
}

func (this TImageAnimation) GetAutoPlay() bool {
  return (bool)((*C.image_animation_t)(unsafe.Pointer(this.handle)).auto_play);
}

func (this TImageAnimation) GetUnloadAfterPaint() bool {
  return (bool)((*C.image_animation_t)(unsafe.Pointer(this.handle)).unload_after_paint);
}

func (this TImageAnimation) GetFormat() string {
  return C.GoString((*C.image_animation_t)(unsafe.Pointer(this.handle)).format);
}

func (this TImageAnimation) GetInterval() uint32 {
  return (uint32)((*C.image_animation_t)(unsafe.Pointer(this.handle)).interval);
}

func (this TImageAnimation) GetDelay() uint32 {
  return (uint32)((*C.image_animation_t)(unsafe.Pointer(this.handle)).delay);
}

func (this TImageAnimation) GetShowWhenDone() bool {
  return (bool)((*C.image_animation_t)(unsafe.Pointer(this.handle)).show_when_done);
}

type TImageBase struct {
  TWidget
}

func (this TImageBase) SetImage(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.image_base_set_image((*C.widget_t)(this.handle), aname));
}

func (this TImageBase) SetRotation(rotation float64) TRet {
  return TRet(C.image_base_set_rotation((*C.widget_t)(this.handle), (C.float_t)(rotation)));
}

func (this TImageBase) SetScale(scale_x float64, scale_y float64) TRet {
  return TRet(C.image_base_set_scale((*C.widget_t)(this.handle), (C.float_t)(scale_x), (C.float_t)(scale_y)));
}

func (this TImageBase) SetAnchor(anchor_x float64, anchor_y float64) TRet {
  return TRet(C.image_base_set_anchor((*C.widget_t)(this.handle), (C.float_t)(anchor_x), (C.float_t)(anchor_y)));
}

func (this TImageBase) SetSelected(selected bool) TRet {
  return TRet(C.image_base_set_selected((*C.widget_t)(this.handle), (C.bool_t)(selected)));
}

func (this TImageBase) SetSelectable(selectable bool) TRet {
  return TRet(C.image_base_set_selectable((*C.widget_t)(this.handle), (C.bool_t)(selectable)));
}

func (this TImageBase) SetClickable(clickable bool) TRet {
  return TRet(C.image_base_set_clickable((*C.widget_t)(this.handle), (C.bool_t)(clickable)));
}

func TImageBaseCast(widget TWidget) TImageBase {
  retObj := TImageBase{}
  retObj.handle = unsafe.Pointer(C.image_base_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TImageBase) GetImage() string {
  return C.GoString((*C.image_base_t)(unsafe.Pointer(this.handle)).image);
}

func (this TImageBase) GetAnchorX() float64 {
  return (float64)((*C.image_base_t)(unsafe.Pointer(this.handle)).anchor_x);
}

func (this TImageBase) GetAnchorY() float64 {
  return (float64)((*C.image_base_t)(unsafe.Pointer(this.handle)).anchor_y);
}

func (this TImageBase) GetScaleX() float64 {
  return (float64)((*C.image_base_t)(unsafe.Pointer(this.handle)).scale_x);
}

func (this TImageBase) GetScaleY() float64 {
  return (float64)((*C.image_base_t)(unsafe.Pointer(this.handle)).scale_y);
}

func (this TImageBase) GetRotation() float64 {
  return (float64)((*C.image_base_t)(unsafe.Pointer(this.handle)).rotation);
}

func (this TImageBase) GetClickable() bool {
  return (bool)((*C.image_base_t)(unsafe.Pointer(this.handle)).clickable);
}

func (this TImageBase) GetSelectable() bool {
  return (bool)((*C.image_base_t)(unsafe.Pointer(this.handle)).selectable);
}

func (this TImageBase) GetSelected() bool {
  return (bool)((*C.image_base_t)(unsafe.Pointer(this.handle)).selected);
}

type TImageDrawType int
const (
  IMAGE_DRAW_DEFAULT TImageDrawType = C.IMAGE_DRAW_DEFAULT
  IMAGE_DRAW_CENTER TImageDrawType = C.IMAGE_DRAW_CENTER
  IMAGE_DRAW_ICON TImageDrawType = C.IMAGE_DRAW_ICON
  IMAGE_DRAW_SCALE TImageDrawType = C.IMAGE_DRAW_SCALE
  IMAGE_DRAW_SCALE_AUTO TImageDrawType = C.IMAGE_DRAW_SCALE_AUTO
  IMAGE_DRAW_SCALE_DOWN TImageDrawType = C.IMAGE_DRAW_SCALE_DOWN
  IMAGE_DRAW_SCALE_W TImageDrawType = C.IMAGE_DRAW_SCALE_W
  IMAGE_DRAW_SCALE_H TImageDrawType = C.IMAGE_DRAW_SCALE_H
  IMAGE_DRAW_FILL TImageDrawType = C.IMAGE_DRAW_FILL
  IMAGE_DRAW_REPEAT TImageDrawType = C.IMAGE_DRAW_REPEAT
  IMAGE_DRAW_REPEAT_X TImageDrawType = C.IMAGE_DRAW_REPEAT_X
  IMAGE_DRAW_REPEAT_Y TImageDrawType = C.IMAGE_DRAW_REPEAT_Y
  IMAGE_DRAW_REPEAT_Y_INVERSE TImageDrawType = C.IMAGE_DRAW_REPEAT_Y_INVERSE
  IMAGE_DRAW_PATCH9 TImageDrawType = C.IMAGE_DRAW_PATCH9
  IMAGE_DRAW_PATCH3_X TImageDrawType = C.IMAGE_DRAW_PATCH3_X
  IMAGE_DRAW_PATCH3_Y TImageDrawType = C.IMAGE_DRAW_PATCH3_Y
  IMAGE_DRAW_PATCH3_X_SCALE_Y TImageDrawType = C.IMAGE_DRAW_PATCH3_X_SCALE_Y
  IMAGE_DRAW_PATCH3_Y_SCALE_X TImageDrawType = C.IMAGE_DRAW_PATCH3_Y_SCALE_X
  IMAGE_DRAW_REPEAT9 TImageDrawType = C.IMAGE_DRAW_REPEAT9
  IMAGE_DRAW_REPEAT3_X TImageDrawType = C.IMAGE_DRAW_REPEAT3_X
  IMAGE_DRAW_REPEAT3_Y TImageDrawType = C.IMAGE_DRAW_REPEAT3_Y
)
type TImageManager struct {
  handle unsafe.Pointer
}

func TImageManagerInstance() TImageManager {
  retObj := TImageManager{}
  retObj.handle = unsafe.Pointer(C.image_manager())
  return retObj
}

func (this TImageManager) GetBitmap(name string, image TBitmap) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.image_manager_get_bitmap((*C.image_manager_t)(this.handle), aname, (*C.bitmap_t)(image.handle)));
}

func (this TImageManager) Preload(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.image_manager_preload((*C.image_manager_t)(this.handle), aname));
}

type TImageValue struct {
  TWidget
}

func TImageValueCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TImageValue{}
  retObj.handle = unsafe.Pointer(C.image_value_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TImageValue) SetImage(image string) TRet {
  aimage := C.CString(image)
  defer C.free(unsafe.Pointer(aimage))
  return TRet(C.image_value_set_image((*C.widget_t)(this.handle), aimage));
}

func (this TImageValue) SetFormat(format string) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.image_value_set_format((*C.widget_t)(this.handle), aformat));
}

func (this TImageValue) SetClickAddDelta(click_add_delta float64) TRet {
  return TRet(C.image_value_set_click_add_delta((*C.widget_t)(this.handle), (C.double)(click_add_delta)));
}

func (this TImageValue) SetValue(value float64) TRet {
  return TRet(C.image_value_set_value((*C.widget_t)(this.handle), (C.double)(value)));
}

func (this TImageValue) SetMin(min float64) TRet {
  return TRet(C.image_value_set_min((*C.widget_t)(this.handle), (C.double)(min)));
}

func (this TImageValue) SetMax(max float64) TRet {
  return TRet(C.image_value_set_max((*C.widget_t)(this.handle), (C.double)(max)));
}

func TImageValueCast(widget TWidget) TImageValue {
  retObj := TImageValue{}
  retObj.handle = unsafe.Pointer(C.image_value_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TImageValue) GetImage() string {
  return C.GoString((*C.image_value_t)(unsafe.Pointer(this.handle)).image);
}

func (this TImageValue) GetFormat() string {
  return C.GoString((*C.image_value_t)(unsafe.Pointer(this.handle)).format);
}

func (this TImageValue) GetClickAddDelta() float64 {
  return (float64)((*C.image_value_t)(unsafe.Pointer(this.handle)).click_add_delta);
}

func (this TImageValue) GetMin() float64 {
  return (float64)((*C.image_value_t)(unsafe.Pointer(this.handle)).min);
}

func (this TImageValue) GetMax() float64 {
  return (float64)((*C.image_value_t)(unsafe.Pointer(this.handle)).max);
}

type TIndicatorDefaultPaint int
const (
  INDICATOR_DEFAULT_PAINT_AUTO TIndicatorDefaultPaint = C.INDICATOR_DEFAULT_PAINT_AUTO
  INDICATOR_DEFAULT_PAINT_FILL_DOT TIndicatorDefaultPaint = C.INDICATOR_DEFAULT_PAINT_FILL_DOT
  INDICATOR_DEFAULT_PAINT_STROKE_DOT TIndicatorDefaultPaint = C.INDICATOR_DEFAULT_PAINT_STROKE_DOT
  INDICATOR_DEFAULT_PAINT_FILL_RECT TIndicatorDefaultPaint = C.INDICATOR_DEFAULT_PAINT_FILL_RECT
  INDICATOR_DEFAULT_PAINT_STROKE_RECT TIndicatorDefaultPaint = C.INDICATOR_DEFAULT_PAINT_STROKE_RECT
)
type TInputMethod struct {
  handle unsafe.Pointer
}

func (this TInputMethod) CommitText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.input_method_commit_text((*C.input_method_t)(this.handle), atext));
}

func (this TInputMethod) SetLang(lang string) TRet {
  alang := C.CString(lang)
  defer C.free(unsafe.Pointer(alang))
  return TRet(C.input_method_set_lang((*C.input_method_t)(this.handle), alang));
}

func (this TInputMethod) GetLang() string {
  return C.GoString(C.input_method_get_lang((*C.input_method_t)(this.handle)));
}

func (this TInputMethod) DispatchKey(key uint32) TRet {
  return TRet(C.input_method_dispatch_key((*C.input_method_t)(this.handle), (C.uint32_t)(key)));
}

func (this TInputMethod) DispatchKeys(keys string) TRet {
  akeys := C.CString(keys)
  defer C.free(unsafe.Pointer(akeys))
  return TRet(C.input_method_dispatch_keys((*C.input_method_t)(this.handle), akeys));
}

func (this TInputMethod) DispatchPreedit() TRet {
  return TRet(C.input_method_dispatch_preedit((*C.input_method_t)(this.handle)));
}

func (this TInputMethod) DispatchPreeditConfirm() TRet {
  return TRet(C.input_method_dispatch_preedit_confirm((*C.input_method_t)(this.handle)));
}

func (this TInputMethod) DispatchPreeditAbort() TRet {
  return TRet(C.input_method_dispatch_preedit_abort((*C.input_method_t)(this.handle)));
}

func TInputMethodInstance() TInputMethod {
  retObj := TInputMethod{}
  retObj.handle = unsafe.Pointer(C.input_method())
  return retObj
}

type TInputType int
const (
  INPUT_TEXT TInputType = C.INPUT_TEXT
  INPUT_INT TInputType = C.INPUT_INT
  INPUT_UINT TInputType = C.INPUT_UINT
  INPUT_HEX TInputType = C.INPUT_HEX
  INPUT_FLOAT TInputType = C.INPUT_FLOAT
  INPUT_UFLOAT TInputType = C.INPUT_UFLOAT
  INPUT_EMAIL TInputType = C.INPUT_EMAIL
  INPUT_PASSWORD TInputType = C.INPUT_PASSWORD
  INPUT_PHONE TInputType = C.INPUT_PHONE
  INPUT_IPV4 TInputType = C.INPUT_IPV4
  INPUT_DATE TInputType = C.INPUT_DATE
  INPUT_TIME TInputType = C.INPUT_TIME
  INPUT_TIME_FULL TInputType = C.INPUT_TIME_FULL
  INPUT_CUSTOM TInputType = C.INPUT_CUSTOM
  INPUT_CUSTOM_PASSWORD TInputType = C.INPUT_CUSTOM_PASSWORD
  INPUT_ASCII TInputType = C.INPUT_ASCII
)
type TKeyCode int
const (
  TK_KEY_RETURN TKeyCode = C.TK_KEY_RETURN
  TK_KEY_ESCAPE TKeyCode = C.TK_KEY_ESCAPE
  TK_KEY_BACKSPACE TKeyCode = C.TK_KEY_BACKSPACE
  TK_KEY_TAB TKeyCode = C.TK_KEY_TAB
  TK_KEY_SPACE TKeyCode = C.TK_KEY_SPACE
  TK_KEY_EXCLAIM TKeyCode = C.TK_KEY_EXCLAIM
  TK_KEY_QUOTEDBL TKeyCode = C.TK_KEY_QUOTEDBL
  TK_KEY_HASH TKeyCode = C.TK_KEY_HASH
  TK_KEY_PERCENT TKeyCode = C.TK_KEY_PERCENT
  TK_KEY_DOLLAR TKeyCode = C.TK_KEY_DOLLAR
  TK_KEY_AMPERSAND TKeyCode = C.TK_KEY_AMPERSAND
  TK_KEY_QUOTE TKeyCode = C.TK_KEY_QUOTE
  TK_KEY_LEFTPAREN TKeyCode = C.TK_KEY_LEFTPAREN
  TK_KEY_RIGHTPAREN TKeyCode = C.TK_KEY_RIGHTPAREN
  TK_KEY_ASTERISK TKeyCode = C.TK_KEY_ASTERISK
  TK_KEY_PLUS TKeyCode = C.TK_KEY_PLUS
  TK_KEY_COMMA TKeyCode = C.TK_KEY_COMMA
  TK_KEY_MINUS TKeyCode = C.TK_KEY_MINUS
  TK_KEY_PERIOD TKeyCode = C.TK_KEY_PERIOD
  TK_KEY_SLASH TKeyCode = C.TK_KEY_SLASH
  TK_KEY_0 TKeyCode = C.TK_KEY_0
  TK_KEY_1 TKeyCode = C.TK_KEY_1
  TK_KEY_2 TKeyCode = C.TK_KEY_2
  TK_KEY_3 TKeyCode = C.TK_KEY_3
  TK_KEY_4 TKeyCode = C.TK_KEY_4
  TK_KEY_5 TKeyCode = C.TK_KEY_5
  TK_KEY_6 TKeyCode = C.TK_KEY_6
  TK_KEY_7 TKeyCode = C.TK_KEY_7
  TK_KEY_8 TKeyCode = C.TK_KEY_8
  TK_KEY_9 TKeyCode = C.TK_KEY_9
  TK_KEY_COLON TKeyCode = C.TK_KEY_COLON
  TK_KEY_SEMICOLON TKeyCode = C.TK_KEY_SEMICOLON
  TK_KEY_LESS TKeyCode = C.TK_KEY_LESS
  TK_KEY_EQUAL TKeyCode = C.TK_KEY_EQUAL
  TK_KEY_GREATER TKeyCode = C.TK_KEY_GREATER
  TK_KEY_QUESTION TKeyCode = C.TK_KEY_QUESTION
  TK_KEY_AT TKeyCode = C.TK_KEY_AT
  TK_KEY_LEFTBRACKET TKeyCode = C.TK_KEY_LEFTBRACKET
  TK_KEY_BACKSLASH TKeyCode = C.TK_KEY_BACKSLASH
  TK_KEY_RIGHTBRACKET TKeyCode = C.TK_KEY_RIGHTBRACKET
  TK_KEY_CARET TKeyCode = C.TK_KEY_CARET
  TK_KEY_UNDERSCORE TKeyCode = C.TK_KEY_UNDERSCORE
  TK_KEY_BACKQUOTE TKeyCode = C.TK_KEY_BACKQUOTE
  TK_KEY_a TKeyCode = C.TK_KEY_a
  TK_KEY_b TKeyCode = C.TK_KEY_b
  TK_KEY_c TKeyCode = C.TK_KEY_c
  TK_KEY_d TKeyCode = C.TK_KEY_d
  TK_KEY_e TKeyCode = C.TK_KEY_e
  TK_KEY_f TKeyCode = C.TK_KEY_f
  TK_KEY_g TKeyCode = C.TK_KEY_g
  TK_KEY_h TKeyCode = C.TK_KEY_h
  TK_KEY_i TKeyCode = C.TK_KEY_i
  TK_KEY_j TKeyCode = C.TK_KEY_j
  TK_KEY_k TKeyCode = C.TK_KEY_k
  TK_KEY_l TKeyCode = C.TK_KEY_l
  TK_KEY_m TKeyCode = C.TK_KEY_m
  TK_KEY_n TKeyCode = C.TK_KEY_n
  TK_KEY_o TKeyCode = C.TK_KEY_o
  TK_KEY_p TKeyCode = C.TK_KEY_p
  TK_KEY_q TKeyCode = C.TK_KEY_q
  TK_KEY_r TKeyCode = C.TK_KEY_r
  TK_KEY_s TKeyCode = C.TK_KEY_s
  TK_KEY_t TKeyCode = C.TK_KEY_t
  TK_KEY_u TKeyCode = C.TK_KEY_u
  TK_KEY_v TKeyCode = C.TK_KEY_v
  TK_KEY_w TKeyCode = C.TK_KEY_w
  TK_KEY_x TKeyCode = C.TK_KEY_x
  TK_KEY_y TKeyCode = C.TK_KEY_y
  TK_KEY_z TKeyCode = C.TK_KEY_z
  TK_KEY_A TKeyCode = C.TK_KEY_A
  TK_KEY_B TKeyCode = C.TK_KEY_B
  TK_KEY_C TKeyCode = C.TK_KEY_C
  TK_KEY_D TKeyCode = C.TK_KEY_D
  TK_KEY_E TKeyCode = C.TK_KEY_E
  TK_KEY_F TKeyCode = C.TK_KEY_F
  TK_KEY_G TKeyCode = C.TK_KEY_G
  TK_KEY_H TKeyCode = C.TK_KEY_H
  TK_KEY_I TKeyCode = C.TK_KEY_I
  TK_KEY_J TKeyCode = C.TK_KEY_J
  TK_KEY_K TKeyCode = C.TK_KEY_K
  TK_KEY_L TKeyCode = C.TK_KEY_L
  TK_KEY_M TKeyCode = C.TK_KEY_M
  TK_KEY_N TKeyCode = C.TK_KEY_N
  TK_KEY_O TKeyCode = C.TK_KEY_O
  TK_KEY_P TKeyCode = C.TK_KEY_P
  TK_KEY_Q TKeyCode = C.TK_KEY_Q
  TK_KEY_R TKeyCode = C.TK_KEY_R
  TK_KEY_S TKeyCode = C.TK_KEY_S
  TK_KEY_T TKeyCode = C.TK_KEY_T
  TK_KEY_U TKeyCode = C.TK_KEY_U
  TK_KEY_V TKeyCode = C.TK_KEY_V
  TK_KEY_W TKeyCode = C.TK_KEY_W
  TK_KEY_X TKeyCode = C.TK_KEY_X
  TK_KEY_Y TKeyCode = C.TK_KEY_Y
  TK_KEY_Z TKeyCode = C.TK_KEY_Z
  TK_KEY_DOT TKeyCode = C.TK_KEY_DOT
  TK_KEY_DELETE TKeyCode = C.TK_KEY_DELETE
  TK_KEY_LEFTBRACE TKeyCode = C.TK_KEY_LEFTBRACE
  TK_KEY_RIGHTBRACE TKeyCode = C.TK_KEY_RIGHTBRACE
  TK_KEY_LSHIFT TKeyCode = C.TK_KEY_LSHIFT
  TK_KEY_RSHIFT TKeyCode = C.TK_KEY_RSHIFT
  TK_KEY_LCTRL TKeyCode = C.TK_KEY_LCTRL
  TK_KEY_RCTRL TKeyCode = C.TK_KEY_RCTRL
  TK_KEY_LALT TKeyCode = C.TK_KEY_LALT
  TK_KEY_RALT TKeyCode = C.TK_KEY_RALT
  TK_KEY_CAPSLOCK TKeyCode = C.TK_KEY_CAPSLOCK
  TK_KEY_HOME TKeyCode = C.TK_KEY_HOME
  TK_KEY_END TKeyCode = C.TK_KEY_END
  TK_KEY_INSERT TKeyCode = C.TK_KEY_INSERT
  TK_KEY_UP TKeyCode = C.TK_KEY_UP
  TK_KEY_DOWN TKeyCode = C.TK_KEY_DOWN
  TK_KEY_LEFT TKeyCode = C.TK_KEY_LEFT
  TK_KEY_RIGHT TKeyCode = C.TK_KEY_RIGHT
  TK_KEY_PAGEUP TKeyCode = C.TK_KEY_PAGEUP
  TK_KEY_PAGEDOWN TKeyCode = C.TK_KEY_PAGEDOWN
  TK_KEY_F1 TKeyCode = C.TK_KEY_F1
  TK_KEY_F2 TKeyCode = C.TK_KEY_F2
  TK_KEY_F3 TKeyCode = C.TK_KEY_F3
  TK_KEY_F4 TKeyCode = C.TK_KEY_F4
  TK_KEY_F5 TKeyCode = C.TK_KEY_F5
  TK_KEY_F6 TKeyCode = C.TK_KEY_F6
  TK_KEY_F7 TKeyCode = C.TK_KEY_F7
  TK_KEY_F8 TKeyCode = C.TK_KEY_F8
  TK_KEY_F9 TKeyCode = C.TK_KEY_F9
  TK_KEY_F10 TKeyCode = C.TK_KEY_F10
  TK_KEY_F11 TKeyCode = C.TK_KEY_F11
  TK_KEY_F12 TKeyCode = C.TK_KEY_F12
  TK_KEY_MENU TKeyCode = C.TK_KEY_MENU
  TK_KEY_COMMAND TKeyCode = C.TK_KEY_COMMAND
  TK_KEY_BACK TKeyCode = C.TK_KEY_BACK
  TK_KEY_CANCEL TKeyCode = C.TK_KEY_CANCEL
  TK_KEY_KP_DIVIDE TKeyCode = C.TK_KEY_KP_DIVIDE
  TK_KEY_KP_MULTIPLY TKeyCode = C.TK_KEY_KP_MULTIPLY
  TK_KEY_KP_MINUS TKeyCode = C.TK_KEY_KP_MINUS
  TK_KEY_KP_PLUS TKeyCode = C.TK_KEY_KP_PLUS
  TK_KEY_KP_ENTER TKeyCode = C.TK_KEY_KP_ENTER
  TK_KEY_KP_1 TKeyCode = C.TK_KEY_KP_1
  TK_KEY_KP_2 TKeyCode = C.TK_KEY_KP_2
  TK_KEY_KP_3 TKeyCode = C.TK_KEY_KP_3
  TK_KEY_KP_4 TKeyCode = C.TK_KEY_KP_4
  TK_KEY_KP_5 TKeyCode = C.TK_KEY_KP_5
  TK_KEY_KP_6 TKeyCode = C.TK_KEY_KP_6
  TK_KEY_KP_7 TKeyCode = C.TK_KEY_KP_7
  TK_KEY_KP_8 TKeyCode = C.TK_KEY_KP_8
  TK_KEY_KP_9 TKeyCode = C.TK_KEY_KP_9
  TK_KEY_KP_0 TKeyCode = C.TK_KEY_KP_0
  TK_KEY_KP_PERIOD TKeyCode = C.TK_KEY_KP_PERIOD
  TK_KEY_NUMLOCKCLEAR TKeyCode = C.TK_KEY_NUMLOCKCLEAR
  TK_KEY_WHEEL TKeyCode = C.TK_KEY_WHEEL
)
type TKeyEvent struct {
  TEvent
}

func TKeyEventCast(event TEvent) TKeyEvent {
  retObj := TKeyEvent{}
  retObj.handle = unsafe.Pointer(C.key_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TKeyEvent) GetKey() uint32 {
  return (uint32)((*C.key_event_t)(unsafe.Pointer(this.handle)).key);
}

func (this TKeyEvent) GetAlt() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).alt);
}

func (this TKeyEvent) GetLalt() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).lalt);
}

func (this TKeyEvent) GetRalt() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).ralt);
}

func (this TKeyEvent) GetCtrl() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).ctrl);
}

func (this TKeyEvent) GetLctrl() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).lctrl);
}

func (this TKeyEvent) GetRctrl() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).rctrl);
}

func (this TKeyEvent) GetShift() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).shift);
}

func (this TKeyEvent) GetLshift() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).lshift);
}

func (this TKeyEvent) GetRshift() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).rshift);
}

func (this TKeyEvent) GetCmd() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).cmd);
}

func (this TKeyEvent) GetMenu() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).menu);
}

func (this TKeyEvent) GetCapslock() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).capslock);
}

func (this TKeyEvent) GetNumlock() bool {
  return (bool)((*C.key_event_t)(unsafe.Pointer(this.handle)).numlock);
}

type TKeyboard struct {
  TWindowBase
}

func TKeyboardCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TKeyboard{}
  retObj.handle = unsafe.Pointer(C.keyboard_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TKeyboardCast(widget TWidget) TKeyboard {
  retObj := TKeyboard{}
  retObj.handle = unsafe.Pointer(C.keyboard_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TLabel struct {
  TWidget
}

func TLabelCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TLabel{}
  retObj.handle = unsafe.Pointer(C.label_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TLabel) SetLength(length int32) TRet {
  return TRet(C.label_set_length((*C.widget_t)(this.handle), (C.int32_t)(length)));
}

func (this TLabel) SetMaxW(max_w int32) TRet {
  return TRet(C.label_set_max_w((*C.widget_t)(this.handle), (C.int32_t)(max_w)));
}

func (this TLabel) SetLineWrap(line_wrap bool) TRet {
  return TRet(C.label_set_line_wrap((*C.widget_t)(this.handle), (C.bool_t)(line_wrap)));
}

func (this TLabel) SetWordWrap(word_wrap bool) TRet {
  return TRet(C.label_set_word_wrap((*C.widget_t)(this.handle), (C.bool_t)(word_wrap)));
}

func (this TLabel) SetEllipses(ellipses bool) TRet {
  return TRet(C.label_set_ellipses((*C.widget_t)(this.handle), (C.bool_t)(ellipses)));
}

func (this TLabel) ResizeToContent(min_w uint32, max_w uint32, min_h uint32, max_h uint32) TRet {
  return TRet(C.label_resize_to_content((*C.widget_t)(this.handle), (C.uint32_t)(min_w), (C.uint32_t)(max_w), (C.uint32_t)(min_h), (C.uint32_t)(max_h)));
}

func TLabelCast(widget TWidget) TLabel {
  retObj := TLabel{}
  retObj.handle = unsafe.Pointer(C.label_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TLabel) GetLength() int32 {
  return (int32)((*C.label_t)(unsafe.Pointer(this.handle)).length);
}

func (this TLabel) GetLineWrap() bool {
  return (bool)((*C.label_t)(unsafe.Pointer(this.handle)).line_wrap);
}

func (this TLabel) GetWordWrap() bool {
  return (bool)((*C.label_t)(unsafe.Pointer(this.handle)).word_wrap);
}

func (this TLabel) GetEllipses() bool {
  return (bool)((*C.label_t)(unsafe.Pointer(this.handle)).ellipses);
}

func (this TLabel) GetMaxW() int32 {
  return (int32)((*C.label_t)(unsafe.Pointer(this.handle)).max_w);
}

type TLangIndicator struct {
  TWidget
}

func TLangIndicatorCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TLangIndicator{}
  retObj.handle = unsafe.Pointer(C.lang_indicator_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TLangIndicator) SetImage(image string) TRet {
  aimage := C.CString(image)
  defer C.free(unsafe.Pointer(aimage))
  return TRet(C.lang_indicator_set_image((*C.widget_t)(this.handle), aimage));
}

func TLangIndicatorCast(widget TWidget) TLangIndicator {
  retObj := TLangIndicator{}
  retObj.handle = unsafe.Pointer(C.lang_indicator_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TLangIndicator) GetImage() string {
  return C.GoString((*C.lang_indicator_t)(unsafe.Pointer(this.handle)).image);
}

type TLineNumber struct {
  TWidget
}

func TLineNumberCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TLineNumber{}
  retObj.handle = unsafe.Pointer(C.line_number_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TLineNumber) SetTopMargin(top_margin int32) TRet {
  return TRet(C.line_number_set_top_margin((*C.widget_t)(this.handle), (C.int32_t)(top_margin)));
}

func (this TLineNumber) SetBottomMargin(bottom_margin int32) TRet {
  return TRet(C.line_number_set_bottom_margin((*C.widget_t)(this.handle), (C.int32_t)(bottom_margin)));
}

func (this TLineNumber) SetLineHeight(line_height int32) TRet {
  return TRet(C.line_number_set_line_height((*C.widget_t)(this.handle), (C.int32_t)(line_height)));
}

func (this TLineNumber) SetYoffset(yoffset int32) TRet {
  return TRet(C.line_number_set_yoffset((*C.widget_t)(this.handle), (C.int32_t)(yoffset)));
}

func TLineNumberCast(widget TWidget) TLineNumber {
  retObj := TLineNumber{}
  retObj.handle = unsafe.Pointer(C.line_number_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TLineNumber) AddHighlightLine(line int32) TRet {
  return TRet(C.line_number_add_highlight_line((*C.widget_t)(this.handle), (C.int32_t)(line)));
}

func (this TLineNumber) SetActiveLine(line int32) TRet {
  return TRet(C.line_number_set_active_line((*C.widget_t)(this.handle), (C.int32_t)(line)));
}

func (this TLineNumber) ClearHighlight() TRet {
  return TRet(C.line_number_clear_highlight((*C.widget_t)(this.handle)));
}

func (this TLineNumber) IsHighlightLine(line int32) bool {
  return (bool)(C.line_number_is_highlight_line((*C.widget_t)(this.handle), (C.int32_t)(line)));
}

type TListItem struct {
  TWidget
}

func TListItemCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TListItem{}
  retObj.handle = unsafe.Pointer(C.list_item_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TListItemCast(widget TWidget) TListItem {
  retObj := TListItem{}
  retObj.handle = unsafe.Pointer(C.list_item_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TListItemSeperator struct {
  TCheckButton
}

func TListItemSeperatorCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TListItemSeperator{}
  retObj.handle = unsafe.Pointer(C.list_item_seperator_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TListItemSeperatorCast(widget TWidget) TListItemSeperator {
  retObj := TListItemSeperator{}
  retObj.handle = unsafe.Pointer(C.list_item_seperator_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TListView struct {
  TWidget
}

func TListViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TListView{}
  retObj.handle = unsafe.Pointer(C.list_view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TListView) SetItemHeight(item_height int32) TRet {
  return TRet(C.list_view_set_item_height((*C.widget_t)(this.handle), (C.int32_t)(item_height)));
}

func (this TListView) SetDefaultItemHeight(default_item_height int32) TRet {
  return TRet(C.list_view_set_default_item_height((*C.widget_t)(this.handle), (C.int32_t)(default_item_height)));
}

func (this TListView) SetAutoHideScrollBar(auto_hide_scroll_bar bool) TRet {
  return TRet(C.list_view_set_auto_hide_scroll_bar((*C.widget_t)(this.handle), (C.bool_t)(auto_hide_scroll_bar)));
}

func (this TListView) SetFloatingScrollBar(floating_scroll_bar bool) TRet {
  return TRet(C.list_view_set_floating_scroll_bar((*C.widget_t)(this.handle), (C.bool_t)(floating_scroll_bar)));
}

func TListViewCast(widget TWidget) TListView {
  retObj := TListView{}
  retObj.handle = unsafe.Pointer(C.list_view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TListView) Reinit() TRet {
  return TRet(C.list_view_reinit((*C.widget_t)(this.handle)));
}

func (this TListView) GetItemHeight() int32 {
  return (int32)((*C.list_view_t)(unsafe.Pointer(this.handle)).item_height);
}

func (this TListView) GetDefaultItemHeight() int32 {
  return (int32)((*C.list_view_t)(unsafe.Pointer(this.handle)).default_item_height);
}

func (this TListView) GetAutoHideScrollBar() bool {
  return (bool)((*C.list_view_t)(unsafe.Pointer(this.handle)).auto_hide_scroll_bar);
}

func (this TListView) GetFloatingScrollBar() bool {
  return (bool)((*C.list_view_t)(unsafe.Pointer(this.handle)).floating_scroll_bar);
}

type TListViewH struct {
  TWidget
}

func TListViewHCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TListViewH{}
  retObj.handle = unsafe.Pointer(C.list_view_h_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TListViewH) SetItemWidth(item_width int32) TRet {
  return TRet(C.list_view_h_set_item_width((*C.widget_t)(this.handle), (C.int32_t)(item_width)));
}

func (this TListViewH) SetSpacing(spacing int32) TRet {
  return TRet(C.list_view_h_set_spacing((*C.widget_t)(this.handle), (C.int32_t)(spacing)));
}

func TListViewHCast(widget TWidget) TListViewH {
  retObj := TListViewH{}
  retObj.handle = unsafe.Pointer(C.list_view_h_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TListViewH) GetItemWidth() int32 {
  return (int32)((*C.list_view_h_t)(unsafe.Pointer(this.handle)).item_width);
}

func (this TListViewH) GetSpacing() int32 {
  return (int32)((*C.list_view_h_t)(unsafe.Pointer(this.handle)).spacing);
}

type TLocaleInfo struct {
  handle unsafe.Pointer
}

func TLocaleInfoInstance() TLocaleInfo {
  retObj := TLocaleInfo{}
  retObj.handle = unsafe.Pointer(C.locale_info())
  return retObj
}

func (this TLocaleInfo) Tr(text string) string {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return C.GoString(C.locale_info_tr((*C.locale_info_t)(this.handle), atext));
}

func (this TLocaleInfo) Change(language string, country string) TRet {
  alanguage := C.CString(language)
  defer C.free(unsafe.Pointer(alanguage))
  acountry := C.CString(country)
  defer C.free(unsafe.Pointer(acountry))
  return TRet(C.locale_info_change((*C.locale_info_t)(this.handle), alanguage, acountry));
}

func (this TLocaleInfo) Off(id uint32) TRet {
  return TRet(C.locale_info_off((*C.locale_info_t)(this.handle), (C.uint32_t)(id)));
}

func TLocaleInfosRef(name string) TLocaleInfo {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TLocaleInfo{}
  retObj.handle = unsafe.Pointer(C.locale_infos_ref(aname))
  return retObj
}

func TLocaleInfosUnref(locale_info TLocaleInfo) TRet {
  return TRet(C.locale_infos_unref((*C.locale_info_t)(locale_info.handle)));
}

func TLocaleInfosChange(language string, country string) TRet {
  alanguage := C.CString(language)
  defer C.free(unsafe.Pointer(alanguage))
  acountry := C.CString(country)
  defer C.free(unsafe.Pointer(acountry))
  return TRet(C.locale_infos_change(alanguage, acountry));
}

func TLocaleInfosOff(id uint32) TRet {
  return TRet(C.locale_infos_off((C.uint32_t)(id)));
}

func TLocaleInfosReloadAll() TRet {
  return TRet(C.locale_infos_reload_all());
}

type TLogMessageEvent struct {
  TEvent
}

func TLogMessageEventCast(event TEvent) TLogMessageEvent {
  retObj := TLogMessageEvent{}
  retObj.handle = unsafe.Pointer(C.log_message_event_cast((*C.event_t)(event.handle)))
  return retObj
}

type TMIME_TYPE string
const (
  MIME_TYPE_APPLICATION_ENVOY string = C.MIME_TYPE_APPLICATION_ENVOY
  MIME_TYPE_APPLICATION_FRACTALS string = C.MIME_TYPE_APPLICATION_FRACTALS
  MIME_TYPE_APPLICATION_FUTURESPLASH string = C.MIME_TYPE_APPLICATION_FUTURESPLASH
  MIME_TYPE_APPLICATION_HTA string = C.MIME_TYPE_APPLICATION_HTA
  MIME_TYPE_APPLICATION_JSON string = C.MIME_TYPE_APPLICATION_JSON
  MIME_TYPE_APPLICATION_UBJSON string = C.MIME_TYPE_APPLICATION_UBJSON
  MIME_TYPE_APPLICATION_MAC_BINHEX40 string = C.MIME_TYPE_APPLICATION_MAC_BINHEX40
  MIME_TYPE_APPLICATION_MSWORD string = C.MIME_TYPE_APPLICATION_MSWORD
  MIME_TYPE_APPLICATION_OCTET_STREAM string = C.MIME_TYPE_APPLICATION_OCTET_STREAM
  MIME_TYPE_APPLICATION_ODA string = C.MIME_TYPE_APPLICATION_ODA
  MIME_TYPE_APPLICATION_OLESCRIPT string = C.MIME_TYPE_APPLICATION_OLESCRIPT
  MIME_TYPE_APPLICATION_PDF string = C.MIME_TYPE_APPLICATION_PDF
  MIME_TYPE_APPLICATION_PICS_RULES string = C.MIME_TYPE_APPLICATION_PICS_RULES
  MIME_TYPE_APPLICATION_PKCS10 string = C.MIME_TYPE_APPLICATION_PKCS10
  MIME_TYPE_APPLICATION_PKIX_CRL string = C.MIME_TYPE_APPLICATION_PKIX_CRL
  MIME_TYPE_APPLICATION_POSTSCRIPT string = C.MIME_TYPE_APPLICATION_POSTSCRIPT
  MIME_TYPE_APPLICATION_RTF string = C.MIME_TYPE_APPLICATION_RTF
  MIME_TYPE_APPLICATION_VND_MS_EXCEL string = C.MIME_TYPE_APPLICATION_VND_MS_EXCEL
  MIME_TYPE_APPLICATION_VND_MS_OUTLOOK string = C.MIME_TYPE_APPLICATION_VND_MS_OUTLOOK
  MIME_TYPE_APPLICATION_VND_MS_PKICERTSTORE string = C.MIME_TYPE_APPLICATION_VND_MS_PKICERTSTORE
  MIME_TYPE_APPLICATION_VND_MS_PKISECCAT string = C.MIME_TYPE_APPLICATION_VND_MS_PKISECCAT
  MIME_TYPE_APPLICATION_VND_MS_PKISTL string = C.MIME_TYPE_APPLICATION_VND_MS_PKISTL
  MIME_TYPE_APPLICATION_VND_MS_POWERPOINT string = C.MIME_TYPE_APPLICATION_VND_MS_POWERPOINT
  MIME_TYPE_APPLICATION_VND_MS_PROJECT string = C.MIME_TYPE_APPLICATION_VND_MS_PROJECT
  MIME_TYPE_APPLICATION_VND_MS_WORKS string = C.MIME_TYPE_APPLICATION_VND_MS_WORKS
  MIME_TYPE_APPLICATION_WINHLP string = C.MIME_TYPE_APPLICATION_WINHLP
  MIME_TYPE_APPLICATION_X_BCPIO string = C.MIME_TYPE_APPLICATION_X_BCPIO
  MIME_TYPE_APPLICATION_X_CDF string = C.MIME_TYPE_APPLICATION_X_CDF
  MIME_TYPE_APPLICATION_X_COMPRESS string = C.MIME_TYPE_APPLICATION_X_COMPRESS
  MIME_TYPE_APPLICATION_X_COMPRESSED string = C.MIME_TYPE_APPLICATION_X_COMPRESSED
  MIME_TYPE_APPLICATION_X_CPIO string = C.MIME_TYPE_APPLICATION_X_CPIO
  MIME_TYPE_APPLICATION_X_CSH string = C.MIME_TYPE_APPLICATION_X_CSH
  MIME_TYPE_APPLICATION_X_DIRECTOR string = C.MIME_TYPE_APPLICATION_X_DIRECTOR
  MIME_TYPE_APPLICATION_X_DVI string = C.MIME_TYPE_APPLICATION_X_DVI
  MIME_TYPE_APPLICATION_X_GTAR string = C.MIME_TYPE_APPLICATION_X_GTAR
  MIME_TYPE_APPLICATION_X_GZIP string = C.MIME_TYPE_APPLICATION_X_GZIP
  MIME_TYPE_APPLICATION_X_HDF string = C.MIME_TYPE_APPLICATION_X_HDF
  MIME_TYPE_APPLICATION_X_IPHONE string = C.MIME_TYPE_APPLICATION_X_IPHONE
  MIME_TYPE_APPLICATION_X_JAVASCRIPT string = C.MIME_TYPE_APPLICATION_X_JAVASCRIPT
  MIME_TYPE_APPLICATION_X_LATEX string = C.MIME_TYPE_APPLICATION_X_LATEX
  MIME_TYPE_APPLICATION_X_MSACCESS string = C.MIME_TYPE_APPLICATION_X_MSACCESS
  MIME_TYPE_APPLICATION_X_MSCARDFILE string = C.MIME_TYPE_APPLICATION_X_MSCARDFILE
  MIME_TYPE_APPLICATION_X_MSCLIP string = C.MIME_TYPE_APPLICATION_X_MSCLIP
  MIME_TYPE_APPLICATION_X_MSDOWNLOAD string = C.MIME_TYPE_APPLICATION_X_MSDOWNLOAD
  MIME_TYPE_APPLICATION_X_MSMEDIAVIEW string = C.MIME_TYPE_APPLICATION_X_MSMEDIAVIEW
  MIME_TYPE_APPLICATION_X_MSMETAFILE string = C.MIME_TYPE_APPLICATION_X_MSMETAFILE
  MIME_TYPE_APPLICATION_X_MSMONEY string = C.MIME_TYPE_APPLICATION_X_MSMONEY
  MIME_TYPE_APPLICATION_X_MSPUBLISHER string = C.MIME_TYPE_APPLICATION_X_MSPUBLISHER
  MIME_TYPE_APPLICATION_X_MSSCHEDULE string = C.MIME_TYPE_APPLICATION_X_MSSCHEDULE
  MIME_TYPE_APPLICATION_X_MSTERMINAL string = C.MIME_TYPE_APPLICATION_X_MSTERMINAL
  MIME_TYPE_APPLICATION_X_MSWRITE string = C.MIME_TYPE_APPLICATION_X_MSWRITE
  MIME_TYPE_APPLICATION_X_NETCDF string = C.MIME_TYPE_APPLICATION_X_NETCDF
  MIME_TYPE_APPLICATION_X_PERFMON string = C.MIME_TYPE_APPLICATION_X_PERFMON
  MIME_TYPE_APPLICATION_X_PKCS12 string = C.MIME_TYPE_APPLICATION_X_PKCS12
  MIME_TYPE_APPLICATION_X_SH string = C.MIME_TYPE_APPLICATION_X_SH
  MIME_TYPE_APPLICATION_X_SHAR string = C.MIME_TYPE_APPLICATION_X_SHAR
  MIME_TYPE_APPLICATION_X_SHOCKWAVE_FLASH string = C.MIME_TYPE_APPLICATION_X_SHOCKWAVE_FLASH
  MIME_TYPE_APPLICATION_X_STUFFIT string = C.MIME_TYPE_APPLICATION_X_STUFFIT
  MIME_TYPE_APPLICATION_X_SV4CPIO string = C.MIME_TYPE_APPLICATION_X_SV4CPIO
  MIME_TYPE_APPLICATION_X_SV4CRC string = C.MIME_TYPE_APPLICATION_X_SV4CRC
  MIME_TYPE_APPLICATION_X_TAR string = C.MIME_TYPE_APPLICATION_X_TAR
  MIME_TYPE_APPLICATION_X_TCL string = C.MIME_TYPE_APPLICATION_X_TCL
  MIME_TYPE_APPLICATION_X_TEX string = C.MIME_TYPE_APPLICATION_X_TEX
  MIME_TYPE_APPLICATION_X_TEXINFO string = C.MIME_TYPE_APPLICATION_X_TEXINFO
  MIME_TYPE_APPLICATION_X_TROFF string = C.MIME_TYPE_APPLICATION_X_TROFF
  MIME_TYPE_APPLICATION_X_USTAR string = C.MIME_TYPE_APPLICATION_X_USTAR
  MIME_TYPE_APPLICATION_ZIP string = C.MIME_TYPE_APPLICATION_ZIP
  MIME_TYPE_AUDIO_BASIC string = C.MIME_TYPE_AUDIO_BASIC
  MIME_TYPE_AUDIO_MID string = C.MIME_TYPE_AUDIO_MID
  MIME_TYPE_AUDIO_MPEG string = C.MIME_TYPE_AUDIO_MPEG
  MIME_TYPE_AUDIO_X_AIFF string = C.MIME_TYPE_AUDIO_X_AIFF
  MIME_TYPE_AUDIO_X_MPEGURL string = C.MIME_TYPE_AUDIO_X_MPEGURL
  MIME_TYPE_AUDIO_X_WAV string = C.MIME_TYPE_AUDIO_X_WAV
  MIME_TYPE_IMAGE_BMP string = C.MIME_TYPE_IMAGE_BMP
  MIME_TYPE_IMAGE_CIS_COD string = C.MIME_TYPE_IMAGE_CIS_COD
  MIME_TYPE_IMAGE_GIF string = C.MIME_TYPE_IMAGE_GIF
  MIME_TYPE_IMAGE_IEF string = C.MIME_TYPE_IMAGE_IEF
  MIME_TYPE_IMAGE_JPEG string = C.MIME_TYPE_IMAGE_JPEG
  MIME_TYPE_IMAGE_PIPEG string = C.MIME_TYPE_IMAGE_PIPEG
  MIME_TYPE_IMAGE_SVG_XML string = C.MIME_TYPE_IMAGE_SVG_XML
  MIME_TYPE_IMAGE_TIFF string = C.MIME_TYPE_IMAGE_TIFF
  MIME_TYPE_IMAGE_X_CMX string = C.MIME_TYPE_IMAGE_X_CMX
  MIME_TYPE_IMAGE_X_ICON string = C.MIME_TYPE_IMAGE_X_ICON
  MIME_TYPE_IMAGE_X_RGB string = C.MIME_TYPE_IMAGE_X_RGB
  MIME_TYPE_IMAGE_X_XBITMAP string = C.MIME_TYPE_IMAGE_X_XBITMAP
  MIME_TYPE_IMAGE_X_XPIXMAP string = C.MIME_TYPE_IMAGE_X_XPIXMAP
  MIME_TYPE_IMAGE_X_XWINDOWDUMP string = C.MIME_TYPE_IMAGE_X_XWINDOWDUMP
  MIME_TYPE_MESSAGE_RFC822 string = C.MIME_TYPE_MESSAGE_RFC822
  MIME_TYPE_TEXT_CSS string = C.MIME_TYPE_TEXT_CSS
  MIME_TYPE_TEXT_H323 string = C.MIME_TYPE_TEXT_H323
  MIME_TYPE_TEXT_HTML string = C.MIME_TYPE_TEXT_HTML
  MIME_TYPE_TEXT_IULS string = C.MIME_TYPE_TEXT_IULS
  MIME_TYPE_TEXT_PLAIN string = C.MIME_TYPE_TEXT_PLAIN
  MIME_TYPE_TEXT_RICHTEXT string = C.MIME_TYPE_TEXT_RICHTEXT
  MIME_TYPE_TEXT_SCRIPTLET string = C.MIME_TYPE_TEXT_SCRIPTLET
  MIME_TYPE_TEXT_WEBVIEWHTML string = C.MIME_TYPE_TEXT_WEBVIEWHTML
  MIME_TYPE_TEXT_X_COMPONENT string = C.MIME_TYPE_TEXT_X_COMPONENT
  MIME_TYPE_TEXT_X_SETEXT string = C.MIME_TYPE_TEXT_X_SETEXT
  MIME_TYPE_TEXT_X_VCARD string = C.MIME_TYPE_TEXT_X_VCARD
  MIME_TYPE_VIDEO_MPEG string = C.MIME_TYPE_VIDEO_MPEG
  MIME_TYPE_VIDEO_QUICKTIME string = C.MIME_TYPE_VIDEO_QUICKTIME
  MIME_TYPE_VIDEO_X_MSVIDEO string = C.MIME_TYPE_VIDEO_X_MSVIDEO
)
type TMledit struct {
  TWidget
}

func TMleditCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TMledit{}
  retObj.handle = unsafe.Pointer(C.mledit_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TMledit) SetReadonly(readonly bool) TRet {
  return TRet(C.mledit_set_readonly((*C.widget_t)(this.handle), (C.bool_t)(readonly)));
}

func (this TMledit) SetCancelable(cancelable bool) TRet {
  return TRet(C.mledit_set_cancelable((*C.widget_t)(this.handle), (C.bool_t)(cancelable)));
}

func (this TMledit) SetFocus(focus bool) TRet {
  return TRet(C.mledit_set_focus((*C.widget_t)(this.handle), (C.bool_t)(focus)));
}

func (this TMledit) SetWrapWord(wrap_word bool) TRet {
  return TRet(C.mledit_set_wrap_word((*C.widget_t)(this.handle), (C.bool_t)(wrap_word)));
}

func (this TMledit) SetOverwrite(overwrite bool) TRet {
  return TRet(C.mledit_set_overwrite((*C.widget_t)(this.handle), (C.bool_t)(overwrite)));
}

func (this TMledit) SetMaxLines(max_lines uint32) TRet {
  return TRet(C.mledit_set_max_lines((*C.widget_t)(this.handle), (C.uint32_t)(max_lines)));
}

func (this TMledit) SetMaxChars(max_chars uint32) TRet {
  return TRet(C.mledit_set_max_chars((*C.widget_t)(this.handle), (C.uint32_t)(max_chars)));
}

func (this TMledit) SetTips(tips string) TRet {
  atips := C.CString(tips)
  defer C.free(unsafe.Pointer(atips))
  return TRet(C.mledit_set_tips((*C.widget_t)(this.handle), atips));
}

func (this TMledit) SetTrTips(tr_tips string) TRet {
  atr_tips := C.CString(tr_tips)
  defer C.free(unsafe.Pointer(atr_tips))
  return TRet(C.mledit_set_tr_tips((*C.widget_t)(this.handle), atr_tips));
}

func (this TMledit) SetKeyboard(keyboard string) TRet {
  akeyboard := C.CString(keyboard)
  defer C.free(unsafe.Pointer(akeyboard))
  return TRet(C.mledit_set_keyboard((*C.widget_t)(this.handle), akeyboard));
}

func (this TMledit) SetCursor(cursor uint32) TRet {
  return TRet(C.mledit_set_cursor((*C.widget_t)(this.handle), (C.uint32_t)(cursor)));
}

func (this TMledit) GetCursor() uint32 {
  return (uint32)(C.mledit_get_cursor((*C.widget_t)(this.handle)));
}

func (this TMledit) SetScrollLine(scroll_line uint32) TRet {
  return TRet(C.mledit_set_scroll_line((*C.widget_t)(this.handle), (C.uint32_t)(scroll_line)));
}

func (this TMledit) ScrollToOffset(offset uint32) TRet {
  return TRet(C.mledit_scroll_to_offset((*C.widget_t)(this.handle), (C.uint32_t)(offset)));
}

func (this TMledit) SetOpenImWhenFocused(open_im_when_focused bool) TRet {
  return TRet(C.mledit_set_open_im_when_focused((*C.widget_t)(this.handle), (C.bool_t)(open_im_when_focused)));
}

func (this TMledit) SetCloseImWhenBlured(close_im_when_blured bool) TRet {
  return TRet(C.mledit_set_close_im_when_blured((*C.widget_t)(this.handle), (C.bool_t)(close_im_when_blured)));
}

func (this TMledit) SetSelect(start uint32, end uint32) TRet {
  return TRet(C.mledit_set_select((*C.widget_t)(this.handle), (C.uint32_t)(start), (C.uint32_t)(end)));
}

func (this TMledit) GetSelectedText() string {
  return C.GoString(C.mledit_get_selected_text((*C.widget_t)(this.handle)));
}

func (this TMledit) GetCurrentLineIndex() uint32 {
  return (uint32)(C.mledit_get_current_line_index((*C.widget_t)(this.handle)));
}

func (this TMledit) GetCurrentRowIndex() uint32 {
  return (uint32)(C.mledit_get_current_row_index((*C.widget_t)(this.handle)));
}

func (this TMledit) InsertText(offset uint32, text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.mledit_insert_text((*C.widget_t)(this.handle), (C.uint32_t)(offset), atext));
}

func TMleditCast(widget TWidget) TMledit {
  retObj := TMledit{}
  retObj.handle = unsafe.Pointer(C.mledit_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TMledit) GetTips() string {
  return C.GoString((*C.mledit_t)(unsafe.Pointer(this.handle)).tips);
}

func (this TMledit) GetTrTips() string {
  return C.GoString((*C.mledit_t)(unsafe.Pointer(this.handle)).tr_tips);
}

func (this TMledit) GetKeyboard() string {
  return C.GoString((*C.mledit_t)(unsafe.Pointer(this.handle)).keyboard);
}

func (this TMledit) GetMaxLines() uint32 {
  return (uint32)((*C.mledit_t)(unsafe.Pointer(this.handle)).max_lines);
}

func (this TMledit) GetMaxChars() uint32 {
  return (uint32)((*C.mledit_t)(unsafe.Pointer(this.handle)).max_chars);
}

func (this TMledit) GetScrollLine() uint32 {
  return (uint32)((*C.mledit_t)(unsafe.Pointer(this.handle)).scroll_line);
}

func (this TMledit) GetOverwrite() bool {
  return (bool)((*C.mledit_t)(unsafe.Pointer(this.handle)).overwrite);
}

func (this TMledit) GetWrapWord() bool {
  return (bool)((*C.mledit_t)(unsafe.Pointer(this.handle)).wrap_word);
}

func (this TMledit) GetReadonly() bool {
  return (bool)((*C.mledit_t)(unsafe.Pointer(this.handle)).readonly);
}

func (this TMledit) GetCancelable() bool {
  return (bool)((*C.mledit_t)(unsafe.Pointer(this.handle)).cancelable);
}

func (this TMledit) GetOpenImWhenFocused() bool {
  return (bool)((*C.mledit_t)(unsafe.Pointer(this.handle)).open_im_when_focused);
}

func (this TMledit) GetCloseImWhenBlured() bool {
  return (bool)((*C.mledit_t)(unsafe.Pointer(this.handle)).close_im_when_blured);
}

type TModelEvent struct {
  TEvent
}

func TModelEventCast(event TEvent) TModelEvent {
  retObj := TModelEvent{}
  retObj.handle = unsafe.Pointer(C.model_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TModelEvent) GetName() string {
  return C.GoString((*C.model_event_t)(unsafe.Pointer(this.handle)).name);
}

func (this TModelEvent) GetChangeType() string {
  return C.GoString((*C.model_event_t)(unsafe.Pointer(this.handle)).change_type);
}

func (this TModelEvent) GetModel() TObject {
  retObj := TObject{}
  retObj.handle = unsafe.Pointer((*C.model_event_t)(unsafe.Pointer(this.handle)).model)
  return retObj
}

type TMultiGestureEvent struct {
  TEvent
}

func TMultiGestureEventCast(event TEvent) TMultiGestureEvent {
  retObj := TMultiGestureEvent{}
  retObj.handle = unsafe.Pointer(C.multi_gesture_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TMultiGestureEvent) GetX() int {
  return (int)((*C.multi_gesture_event_t)(unsafe.Pointer(this.handle)).x);
}

func (this TMultiGestureEvent) GetY() int {
  return (int)((*C.multi_gesture_event_t)(unsafe.Pointer(this.handle)).y);
}

func (this TMultiGestureEvent) GetRotation() float64 {
  return (float64)((*C.multi_gesture_event_t)(unsafe.Pointer(this.handle)).rotation);
}

func (this TMultiGestureEvent) GetDistance() float64 {
  return (float64)((*C.multi_gesture_event_t)(unsafe.Pointer(this.handle)).distance);
}

type TMutableImage struct {
  TImageBase
}

func TMutableImageCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TMutableImage{}
  retObj.handle = unsafe.Pointer(C.mutable_image_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

type TNamedValue struct {
  handle unsafe.Pointer
}

func TNamedValueCreate() TNamedValue {
  retObj := TNamedValue{}
  retObj.handle = unsafe.Pointer(C.named_value_create())
  return retObj
}

func TNamedValueCast(nv TNamedValue) TNamedValue {
  retObj := TNamedValue{}
  retObj.handle = unsafe.Pointer(C.named_value_cast((*C.named_value_t)(nv.handle)))
  return retObj
}

func (this TNamedValue) SetName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.named_value_set_name((*C.named_value_t)(this.handle), aname));
}

func (this TNamedValue) SetValue(value TValue) TRet {
  return TRet(C.named_value_set_value((*C.named_value_t)(this.handle), (*C.value_t)(value.handle)));
}

func (this TNamedValue) GetValue() TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.named_value_get_value((*C.named_value_t)(this.handle)))
  return retObj
}

func (this TNamedValue) Destroy() TRet {
  return TRet(C.named_value_destroy((*C.named_value_t)(this.handle)));
}

func (this TNamedValue) GetName() string {
  return C.GoString((*C.named_value_t)(unsafe.Pointer(this.handle)).name);
}

type TNamedValueHash struct {
  TNamedValue
}

func TNamedValueHashCreate() TNamedValueHash {
  retObj := TNamedValueHash{}
  retObj.handle = unsafe.Pointer(C.named_value_hash_create())
  return retObj
}

func (this TNamedValueHash) SetName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.named_value_hash_set_name((*C.named_value_hash_t)(this.handle), aname));
}

func (this TNamedValueHash) Destroy() TRet {
  return TRet(C.named_value_hash_destroy((*C.named_value_hash_t)(this.handle)));
}

func (this TNamedValueHash) Clone() TNamedValueHash {
  retObj := TNamedValueHash{}
  retObj.handle = unsafe.Pointer(C.named_value_hash_clone((*C.named_value_hash_t)(this.handle)))
  return retObj
}

func TNamedValueHashGetHashFromStr(str string) int64 {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  return (int64)(C.named_value_hash_get_hash_from_str(astr));
}

type TNativeWindow struct {
  TObject
}

func (this TNativeWindow) Move(x int, y int, force bool) TRet {
  return TRet(C.native_window_move((*C.native_window_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.bool_t)(force)));
}

func (this TNativeWindow) Resize(w int, h int, force bool) TRet {
  return TRet(C.native_window_resize((*C.native_window_t)(this.handle), (C.wh_t)(w), (C.wh_t)(h), (C.bool_t)(force)));
}

func (this TNativeWindow) SetOrientation(old_orientation int64, new_orientation int64) TRet {
  return TRet(C.native_window_set_orientation((*C.native_window_t)(this.handle), (C.lcd_orientation_t)(old_orientation), (C.lcd_orientation_t)(new_orientation)));
}

func (this TNativeWindow) Minimize() TRet {
  return TRet(C.native_window_minimize((*C.native_window_t)(this.handle)));
}

func (this TNativeWindow) Maximize() TRet {
  return TRet(C.native_window_maximize((*C.native_window_t)(this.handle)));
}

func (this TNativeWindow) Restore() TRet {
  return TRet(C.native_window_restore((*C.native_window_t)(this.handle)));
}

func (this TNativeWindow) Center() TRet {
  return TRet(C.native_window_center((*C.native_window_t)(this.handle)));
}

func (this TNativeWindow) ShowBorder(show bool) TRet {
  return TRet(C.native_window_show_border((*C.native_window_t)(this.handle), (C.bool_t)(show)));
}

func (this TNativeWindow) SetFullscreen(fullscreen bool) TRet {
  return TRet(C.native_window_set_fullscreen((*C.native_window_t)(this.handle), (C.bool_t)(fullscreen)));
}

func (this TNativeWindow) SetCursor(name string, img TBitmap) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.native_window_set_cursor((*C.native_window_t)(this.handle), aname, (*C.bitmap_t)(img.handle)));
}

func (this TNativeWindow) SetTitle(app_name string) TRet {
  aapp_name := C.CString(app_name)
  defer C.free(unsafe.Pointer(aapp_name))
  return TRet(C.native_window_set_title((*C.native_window_t)(this.handle), aapp_name));
}

type TObject struct {
  TEmitter
}

func (this TObject) Unref() TRet {
  return TRet(C.object_unref((*C.object_t)(this.handle)));
}

func TObjectRef(obj TObject) TObject {
  retObj := TObject{}
  retObj.handle = unsafe.Pointer(C.object_ref((*C.object_t)(obj.handle)))
  return retObj
}

func (this TObject) GetType() string {
  return C.GoString(C.object_get_type((*C.object_t)(this.handle)));
}

func (this TObject) GetDesc() string {
  return C.GoString(C.object_get_desc((*C.object_t)(this.handle)));
}

func (this TObject) GetSize() uint32 {
  return (uint32)(C.object_get_size((*C.object_t)(this.handle)));
}

func (this TObject) IsCollection() bool {
  return (bool)(C.object_is_collection((*C.object_t)(this.handle)));
}

func (this TObject) SetName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_name((*C.object_t)(this.handle), aname));
}

func (this TObject) Compare(other TObject) int {
  return (int)(C.object_compare((*C.object_t)(this.handle), (*C.object_t)(other.handle)));
}

func (this TObject) GetProp(name string, v TValue) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_get_prop((*C.object_t)(this.handle), aname, (*C.value_t)(v.handle)));
}

func (this TObject) GetPropStr(name string) string {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return C.GoString(C.object_get_prop_str((*C.object_t)(this.handle), aname));
}

func (this TObject) GetPropPointer(name string) unsafe.Pointer {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (unsafe.Pointer)(C.object_get_prop_pointer((*C.object_t)(this.handle), aname));
}

func (this TObject) GetPropObject(name string) TObject {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TObject{}
  retObj.handle = unsafe.Pointer(C.object_get_prop_object((*C.object_t)(this.handle), aname))
  return retObj
}

func (this TObject) GetPropInt(name string, defval int32) int32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int32)(C.object_get_prop_int((*C.object_t)(this.handle), aname, (C.int32_t)(defval)));
}

func (this TObject) GetPropBool(name string, defval bool) bool {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (bool)(C.object_get_prop_bool((*C.object_t)(this.handle), aname, (C.bool_t)(defval)));
}

func (this TObject) GetPropFloat(name string, defval float64) float64 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (float64)(C.object_get_prop_float((*C.object_t)(this.handle), aname, (C.float_t)(defval)));
}

func (this TObject) GetPropDouble(name string, defval float64) float64 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (float64)(C.object_get_prop_double((*C.object_t)(this.handle), aname, (C.double)(defval)));
}

func (this TObject) RemoveProp(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_remove_prop((*C.object_t)(this.handle), aname));
}

func (this TObject) SetProp(name string, value TValue) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop((*C.object_t)(this.handle), aname, (*C.value_t)(value.handle)));
}

func (this TObject) SetPropStr(name string, value string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.object_set_prop_str((*C.object_t)(this.handle), aname, avalue));
}

func (this TObject) SetPropObject(name string, value TObject) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_object((*C.object_t)(this.handle), aname, (*C.object_t)(value.handle)));
}

func (this TObject) SetPropInt(name string, value int32) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_int((*C.object_t)(this.handle), aname, (C.int32_t)(value)));
}

func (this TObject) SetPropBool(name string, value bool) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_bool((*C.object_t)(this.handle), aname, (C.bool_t)(value)));
}

func (this TObject) SetPropFloat(name string, value float64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_float((*C.object_t)(this.handle), aname, (C.float_t)(value)));
}

func (this TObject) SetPropDouble(name string, value float64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_double((*C.object_t)(this.handle), aname, (C.double)(value)));
}

func (this TObject) CopyProp(src TObject, name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_copy_prop((*C.object_t)(this.handle), (*C.object_t)(src.handle), aname));
}

func (this TObject) CopyProps(src TObject, overwrite bool) TRet {
  return TRet(C.object_copy_props((*C.object_t)(this.handle), (*C.object_t)(src.handle), (C.bool_t)(overwrite)));
}

func (this TObject) HasProp(name string) bool {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (bool)(C.object_has_prop((*C.object_t)(this.handle), aname));
}

func (this TObject) Eval(expr string, v TValue) TRet {
  aexpr := C.CString(expr)
  defer C.free(unsafe.Pointer(aexpr))
  return TRet(C.object_eval((*C.object_t)(this.handle), aexpr, (*C.value_t)(v.handle)));
}

func (this TObject) CanExec(name string, args string) bool {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  aargs := C.CString(args)
  defer C.free(unsafe.Pointer(aargs))
  return (bool)(C.object_can_exec((*C.object_t)(this.handle), aname, aargs));
}

func (this TObject) Execute(name string, args string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  aargs := C.CString(args)
  defer C.free(unsafe.Pointer(aargs))
  return TRet(C.object_exec((*C.object_t)(this.handle), aname, aargs));
}

func (this TObject) NotifyChanged() TRet {
  return TRet(C.object_notify_changed((*C.object_t)(this.handle)));
}

func (this TObject) HasPropByPath(path string) bool {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return (bool)(C.object_has_prop_by_path((*C.object_t)(this.handle), apath));
}

func (this TObject) GetPropStrByPath(path string) string {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return C.GoString(C.object_get_prop_str_by_path((*C.object_t)(this.handle), apath));
}

func (this TObject) GetPropPointerByPath(path string) unsafe.Pointer {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return (unsafe.Pointer)(C.object_get_prop_pointer_by_path((*C.object_t)(this.handle), apath));
}

func (this TObject) GetPropObjectByPath(path string) TObject {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  retObj := TObject{}
  retObj.handle = unsafe.Pointer(C.object_get_prop_object_by_path((*C.object_t)(this.handle), apath))
  return retObj
}

func (this TObject) GetPropIntByPath(path string, defval int32) int32 {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return (int32)(C.object_get_prop_int_by_path((*C.object_t)(this.handle), apath, (C.int32_t)(defval)));
}

func (this TObject) GetPropBoolByPath(path string, defval bool) bool {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return (bool)(C.object_get_prop_bool_by_path((*C.object_t)(this.handle), apath, (C.bool_t)(defval)));
}

func (this TObject) GetPropFloatByPath(path string, defval float64) float64 {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return (float64)(C.object_get_prop_float_by_path((*C.object_t)(this.handle), apath, (C.float_t)(defval)));
}

func (this TObject) SetPropByPath(path string, value TValue) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return TRet(C.object_set_prop_by_path((*C.object_t)(this.handle), apath, (*C.value_t)(value.handle)));
}

func (this TObject) SetPropStrByPath(path string, value string) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.object_set_prop_str_by_path((*C.object_t)(this.handle), apath, avalue));
}

func (this TObject) SetPropObjectByPath(path string, value TObject) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return TRet(C.object_set_prop_object_by_path((*C.object_t)(this.handle), apath, (*C.object_t)(value.handle)));
}

func (this TObject) SetPropIntByPath(path string, value int32) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return TRet(C.object_set_prop_int_by_path((*C.object_t)(this.handle), apath, (C.int32_t)(value)));
}

func (this TObject) SetPropBoolByPath(path string, value bool) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return TRet(C.object_set_prop_bool_by_path((*C.object_t)(this.handle), apath, (C.bool_t)(value)));
}

func (this TObject) SetPropFloatByPath(path string, value float64) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  return TRet(C.object_set_prop_float_by_path((*C.object_t)(this.handle), apath, (C.float_t)(value)));
}

func (this TObject) CanExecByPath(path string, args string) bool {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  aargs := C.CString(args)
  defer C.free(unsafe.Pointer(aargs))
  return (bool)(C.object_can_exec_by_path((*C.object_t)(this.handle), apath, aargs));
}

func (this TObject) ExecuteByPath(path string, args string) TRet {
  apath := C.CString(path)
  defer C.free(unsafe.Pointer(apath))
  aargs := C.CString(args)
  defer C.free(unsafe.Pointer(aargs))
  return TRet(C.object_exec_by_path((*C.object_t)(this.handle), apath, aargs));
}

func (this TObject) GetPropInt8(name string, defval int8) int8 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int8)(C.object_get_prop_int8((*C.object_t)(this.handle), aname, (C.int8_t)(defval)));
}

func (this TObject) SetPropInt8(name string, value int8) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_int8((*C.object_t)(this.handle), aname, (C.int8_t)(value)));
}

func (this TObject) GetPropUint8(name string, defval uint8) uint8 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (uint8)(C.object_get_prop_uint8((*C.object_t)(this.handle), aname, (C.uint8_t)(defval)));
}

func (this TObject) SetPropUint8(name string, value uint8) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_uint8((*C.object_t)(this.handle), aname, (C.uint8_t)(value)));
}

func (this TObject) GetPropInt16(name string, defval int16) int16 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int16)(C.object_get_prop_int16((*C.object_t)(this.handle), aname, (C.int16_t)(defval)));
}

func (this TObject) SetPropInt16(name string, value int16) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_int16((*C.object_t)(this.handle), aname, (C.int16_t)(value)));
}

func (this TObject) GetPropUint16(name string, defval uint16) uint16 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (uint16)(C.object_get_prop_uint16((*C.object_t)(this.handle), aname, (C.uint16_t)(defval)));
}

func (this TObject) SetPropUint16(name string, value uint16) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_uint16((*C.object_t)(this.handle), aname, (C.uint16_t)(value)));
}

func (this TObject) GetPropInt32(name string, defval int32) int32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int32)(C.object_get_prop_int32((*C.object_t)(this.handle), aname, (C.int32_t)(defval)));
}

func (this TObject) SetPropInt32(name string, value int32) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_int32((*C.object_t)(this.handle), aname, (C.int32_t)(value)));
}

func (this TObject) GetPropUint32(name string, defval uint32) uint32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (uint32)(C.object_get_prop_uint32((*C.object_t)(this.handle), aname, (C.uint32_t)(defval)));
}

func (this TObject) SetPropUint32(name string, value uint32) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_uint32((*C.object_t)(this.handle), aname, (C.uint32_t)(value)));
}

func (this TObject) GetPropInt64(name string, defval int64) int64 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int64)(C.object_get_prop_int64((*C.object_t)(this.handle), aname, (C.int64_t)(defval)));
}

func (this TObject) SetPropInt64(name string, value int64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_int64((*C.object_t)(this.handle), aname, (C.int64_t)(value)));
}

func (this TObject) GetPropUint64(name string, defval int64) int64 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int64)(C.object_get_prop_uint64((*C.object_t)(this.handle), aname, (C.uint64_t)(defval)));
}

func (this TObject) SetPropUint64(name string, value int64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.object_set_prop_uint64((*C.object_t)(this.handle), aname, (C.uint64_t)(value)));
}

func (this TObject) ClearProps() TRet {
  return TRet(C.object_clear_props((*C.object_t)(this.handle)));
}

func (this TObject) GetRefCount() int32 {
  return (int32)((*C.object_t)(unsafe.Pointer(this.handle)).ref_count);
}

func (this TObject) GetName() string {
  return C.GoString((*C.object_t)(unsafe.Pointer(this.handle)).name);
}

type TObjectArray struct {
  TObject
}

func TObjectArrayCreate() TObject {
  retObj := TObjectArray{}
  retObj.handle = unsafe.Pointer(C.object_array_create())
  return retObj.TObject
}

func (this TObjectArray) Unref() TRet {
  return TRet(C.object_array_unref((*C.object_t)(this.handle)));
}

func (this TObjectArray) ClearProps() TRet {
  return TRet(C.object_array_clear_props((*C.object_t)(this.handle)));
}

func (this TObjectArray) Insert(index uint32, v TValue) TRet {
  return TRet(C.object_array_insert((*C.object_t)(this.handle), (C.uint32_t)(index), (*C.value_t)(v.handle)));
}

func (this TObjectArray) Push(v TValue) TRet {
  return TRet(C.object_array_push((*C.object_t)(this.handle), (*C.value_t)(v.handle)));
}

func (this TObjectArray) IndexOf(v TValue) int32 {
  return (int32)(C.object_array_index_of((*C.object_t)(this.handle), (*C.value_t)(v.handle)));
}

func (this TObjectArray) LastIndexOf(v TValue) int32 {
  return (int32)(C.object_array_last_index_of((*C.object_t)(this.handle), (*C.value_t)(v.handle)));
}

func (this TObjectArray) Remove(index uint32) TRet {
  return TRet(C.object_array_remove((*C.object_t)(this.handle), (C.uint32_t)(index)));
}

func (this TObjectArray) RemoveValue(v TValue) TRet {
  return TRet(C.object_array_remove_value((*C.object_t)(this.handle), (*C.value_t)(v.handle)));
}

func (this TObjectArray) GetAndRemove(index uint32, v TValue) TRet {
  return TRet(C.object_array_get_and_remove((*C.object_t)(this.handle), (C.uint32_t)(index), (*C.value_t)(v.handle)));
}

func (this TObjectArray) GetSize() uint32 {
  return (uint32)((*C.object_array_t)(unsafe.Pointer(this.handle)).size);
}

type TObjectCmd string
const (
  OBJECT_CMD_SAVE string = C.OBJECT_CMD_SAVE
  OBJECT_CMD_RELOAD string = C.OBJECT_CMD_RELOAD
  OBJECT_CMD_MOVE_UP string = C.OBJECT_CMD_MOVE_UP
  OBJECT_CMD_MOVE_DOWN string = C.OBJECT_CMD_MOVE_DOWN
  OBJECT_CMD_REMOVE string = C.OBJECT_CMD_REMOVE
  OBJECT_CMD_REMOVE_CHECKED string = C.OBJECT_CMD_REMOVE_CHECKED
  OBJECT_CMD_CLEAR string = C.OBJECT_CMD_CLEAR
  OBJECT_CMD_ADD string = C.OBJECT_CMD_ADD
  OBJECT_CMD_DETAIL string = C.OBJECT_CMD_DETAIL
  OBJECT_CMD_EDIT string = C.OBJECT_CMD_EDIT
)
type TObjectDefault struct {
  TObject
}

func TObjectDefaultCreate() TObject {
  retObj := TObjectDefault{}
  retObj.handle = unsafe.Pointer(C.object_default_create())
  return retObj.TObject
}

func TObjectDefaultCreateEx(enable_path bool) TObject {
  retObj := TObjectDefault{}
  retObj.handle = unsafe.Pointer(C.object_default_create_ex((C.bool_t)(enable_path)))
  return retObj.TObject
}

func (this TObjectDefault) Unref() TRet {
  return TRet(C.object_default_unref((*C.object_t)(this.handle)));
}

func (this TObjectDefault) ClearProps() TRet {
  return TRet(C.object_default_clear_props((*C.object_t)(this.handle)));
}

func (this TObjectDefault) SetKeepPropType(keep_prop_type bool) TRet {
  return TRet(C.object_default_set_keep_prop_type((*C.object_t)(this.handle), (C.bool_t)(keep_prop_type)));
}

func (this TObjectDefault) SetNameCaseInsensitive(name_case_insensitive bool) TRet {
  return TRet(C.object_default_set_name_case_insensitive((*C.object_t)(this.handle), (C.bool_t)(name_case_insensitive)));
}

type TObjectHash struct {
  TObject
}

func TObjectHashCreate() TObject {
  retObj := TObjectHash{}
  retObj.handle = unsafe.Pointer(C.object_hash_create())
  return retObj.TObject
}

func TObjectHashCreateEx(enable_path bool) TObject {
  retObj := TObjectHash{}
  retObj.handle = unsafe.Pointer(C.object_hash_create_ex((C.bool_t)(enable_path)))
  return retObj.TObject
}

func (this TObjectHash) SetKeepPropType(keep_prop_type bool) TRet {
  return TRet(C.object_hash_set_keep_prop_type((*C.object_t)(this.handle), (C.bool_t)(keep_prop_type)));
}

type TObjectProp string
const (
  OBJECT_PROP_SIZE string = C.OBJECT_PROP_SIZE
  OBJECT_PROP_CHECKED string = C.OBJECT_PROP_CHECKED
  OBJECT_PROP_SELECTED_INDEX string = C.OBJECT_PROP_SELECTED_INDEX
)
type TOffsetChangeEvent struct {
  TEvent
}

func TOffsetChangeEventCast(event TEvent) TOffsetChangeEvent {
  retObj := TOffsetChangeEvent{}
  retObj.handle = unsafe.Pointer(C.offset_change_event_cast((*C.event_t)(event.handle)))
  return retObj
}

type TOrientationEvent struct {
  TEvent
}

func TOrientationEventCast(event TEvent) TOrientationEvent {
  retObj := TOrientationEvent{}
  retObj.handle = unsafe.Pointer(C.orientation_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TOrientationEvent) GetOrientation() int64 {
  return (int64)((*C.orientation_event_t)(unsafe.Pointer(this.handle)).orientation);
}

func (this TOrientationEvent) GetOldOrientation() int64 {
  return (int64)((*C.orientation_event_t)(unsafe.Pointer(this.handle)).old_orientation);
}

type TOverlay struct {
  TWindowBase
}

func TOverlayCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TOverlay{}
  retObj.handle = unsafe.Pointer(C.overlay_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TOverlay) SetClickThrough(click_through bool) TRet {
  return TRet(C.overlay_set_click_through((*C.widget_t)(this.handle), (C.bool_t)(click_through)));
}

func (this TOverlay) SetAlwaysOnTop(always_on_top bool) TRet {
  return TRet(C.overlay_set_always_on_top((*C.widget_t)(this.handle), (C.bool_t)(always_on_top)));
}

func (this TOverlay) SetModeless(modeless bool) TRet {
  return TRet(C.overlay_set_modeless((*C.widget_t)(this.handle), (C.bool_t)(modeless)));
}

func TOverlayCast(widget TWidget) TOverlay {
  retObj := TOverlay{}
  retObj.handle = unsafe.Pointer(C.overlay_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TOverlay) GetClickThrough() bool {
  return (bool)((*C.overlay_t)(unsafe.Pointer(this.handle)).click_through);
}

func (this TOverlay) GetAlwaysOnTop() bool {
  return (bool)((*C.overlay_t)(unsafe.Pointer(this.handle)).always_on_top);
}

func (this TOverlay) GetModeless() bool {
  return (bool)((*C.overlay_t)(unsafe.Pointer(this.handle)).modeless);
}

type TPages struct {
  TWidget
}

func TPagesCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TPages{}
  retObj.handle = unsafe.Pointer(C.pages_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TPagesCast(widget TWidget) TPages {
  retObj := TPages{}
  retObj.handle = unsafe.Pointer(C.pages_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TPages) SetActive(index uint32) TRet {
  return TRet(C.pages_set_active((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TPages) SetAutoFocused(auto_focused bool) TRet {
  return TRet(C.pages_set_auto_focused((*C.widget_t)(this.handle), (C.bool_t)(auto_focused)));
}

func (this TPages) SetActiveByName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.pages_set_active_by_name((*C.widget_t)(this.handle), aname));
}

func (this TPages) GetActive() uint32 {
  return (uint32)((*C.pages_t)(unsafe.Pointer(this.handle)).active);
}

func (this TPages) GetAutoFocused() bool {
  return (bool)((*C.pages_t)(unsafe.Pointer(this.handle)).auto_focused);
}

type TPaintEvent struct {
  TEvent
}

func TPaintEventCast(event TEvent) TPaintEvent {
  retObj := TPaintEvent{}
  retObj.handle = unsafe.Pointer(C.paint_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TPaintEvent) GetC() TCanvas {
  retObj := TCanvas{}
  retObj.handle = unsafe.Pointer((*C.paint_event_t)(unsafe.Pointer(this.handle)).c)
  return retObj
}

type TPoint struct {
  handle unsafe.Pointer
}

type TPointerEvent struct {
  TEvent
}

func TPointerEventCast(event TEvent) TPointerEvent {
  retObj := TPointerEvent{}
  retObj.handle = unsafe.Pointer(C.pointer_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TPointerEvent) GetX() int {
  return (int)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).x);
}

func (this TPointerEvent) GetY() int {
  return (int)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).y);
}

func (this TPointerEvent) GetButton() int {
  return (int)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).button);
}

func (this TPointerEvent) GetPressed() bool {
  return (bool)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).pressed);
}

func (this TPointerEvent) GetAlt() bool {
  return (bool)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).alt);
}

func (this TPointerEvent) GetCtrl() bool {
  return (bool)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).ctrl);
}

func (this TPointerEvent) GetCmd() bool {
  return (bool)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).cmd);
}

func (this TPointerEvent) GetMenu() bool {
  return (bool)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).menu);
}

func (this TPointerEvent) GetShift() bool {
  return (bool)((*C.pointer_event_t)(unsafe.Pointer(this.handle)).shift);
}

type TPointf struct {
  handle unsafe.Pointer
}

type TPopup struct {
  TWindowBase
}

func TPopupCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TPopup{}
  retObj.handle = unsafe.Pointer(C.popup_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TPopupCast(widget TWidget) TPopup {
  retObj := TPopup{}
  retObj.handle = unsafe.Pointer(C.popup_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TPopup) SetCloseWhenClick(close_when_click bool) TRet {
  return TRet(C.popup_set_close_when_click((*C.widget_t)(this.handle), (C.bool_t)(close_when_click)));
}

func (this TPopup) SetCloseWhenClickOutside(close_when_click_outside bool) TRet {
  return TRet(C.popup_set_close_when_click_outside((*C.widget_t)(this.handle), (C.bool_t)(close_when_click_outside)));
}

func (this TPopup) SetCloseWhenTimeout(close_when_timeout uint32) TRet {
  return TRet(C.popup_set_close_when_timeout((*C.widget_t)(this.handle), (C.uint32_t)(close_when_timeout)));
}

func (this TPopup) GetCloseWhenClick() bool {
  return (bool)((*C.popup_t)(unsafe.Pointer(this.handle)).close_when_click);
}

func (this TPopup) GetCloseWhenClickOutside() bool {
  return (bool)((*C.popup_t)(unsafe.Pointer(this.handle)).close_when_click_outside);
}

func (this TPopup) GetCloseWhenTimeout() uint32 {
  return (uint32)((*C.popup_t)(unsafe.Pointer(this.handle)).close_when_timeout);
}

type TProgressBar struct {
  TWidget
}

func TProgressBarCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TProgressBar{}
  retObj.handle = unsafe.Pointer(C.progress_bar_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TProgressBarCast(widget TWidget) TProgressBar {
  retObj := TProgressBar{}
  retObj.handle = unsafe.Pointer(C.progress_bar_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TProgressBar) SetValue(value float64) TRet {
  return TRet(C.progress_bar_set_value((*C.widget_t)(this.handle), (C.double)(value)));
}

func (this TProgressBar) SetMax(max float64) TRet {
  return TRet(C.progress_bar_set_max((*C.widget_t)(this.handle), (C.double)(max)));
}

func (this TProgressBar) SetFormat(format string) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.progress_bar_set_format((*C.widget_t)(this.handle), aformat));
}

func (this TProgressBar) SetVertical(vertical bool) TRet {
  return TRet(C.progress_bar_set_vertical((*C.widget_t)(this.handle), (C.bool_t)(vertical)));
}

func (this TProgressBar) SetShowText(show_text bool) TRet {
  return TRet(C.progress_bar_set_show_text((*C.widget_t)(this.handle), (C.bool_t)(show_text)));
}

func (this TProgressBar) SetReverse(reverse bool) TRet {
  return TRet(C.progress_bar_set_reverse((*C.widget_t)(this.handle), (C.bool_t)(reverse)));
}

func (this TProgressBar) GetPercent() uint32 {
  return (uint32)(C.progress_bar_get_percent((*C.widget_t)(this.handle)));
}

func (this TProgressBar) GetMax() float64 {
  return (float64)((*C.progress_bar_t)(unsafe.Pointer(this.handle)).max);
}

func (this TProgressBar) GetFormat() string {
  return C.GoString((*C.progress_bar_t)(unsafe.Pointer(this.handle)).format);
}

func (this TProgressBar) GetVertical() bool {
  return (bool)((*C.progress_bar_t)(unsafe.Pointer(this.handle)).vertical);
}

func (this TProgressBar) GetShowText() bool {
  return (bool)((*C.progress_bar_t)(unsafe.Pointer(this.handle)).show_text);
}

func (this TProgressBar) GetReverse() bool {
  return (bool)((*C.progress_bar_t)(unsafe.Pointer(this.handle)).reverse);
}

type TProgressCircle struct {
  TWidget
}

func TProgressCircleCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TProgressCircle{}
  retObj.handle = unsafe.Pointer(C.progress_circle_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TProgressCircleCast(widget TWidget) TProgressCircle {
  retObj := TProgressCircle{}
  retObj.handle = unsafe.Pointer(C.progress_circle_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TProgressCircle) SetValue(value float64) TRet {
  return TRet(C.progress_circle_set_value((*C.widget_t)(this.handle), (C.float_t)(value)));
}

func (this TProgressCircle) SetMax(max uint32) TRet {
  return TRet(C.progress_circle_set_max((*C.widget_t)(this.handle), (C.uint32_t)(max)));
}

func (this TProgressCircle) SetFormat(format string) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.progress_circle_set_format((*C.widget_t)(this.handle), aformat));
}

func (this TProgressCircle) SetLineWidth(line_width uint32) TRet {
  return TRet(C.progress_circle_set_line_width((*C.widget_t)(this.handle), (C.uint32_t)(line_width)));
}

func (this TProgressCircle) SetStartAngle(start_angle int32) TRet {
  return TRet(C.progress_circle_set_start_angle((*C.widget_t)(this.handle), (C.int32_t)(start_angle)));
}

func (this TProgressCircle) SetLineCap(line_cap string) TRet {
  aline_cap := C.CString(line_cap)
  defer C.free(unsafe.Pointer(aline_cap))
  return TRet(C.progress_circle_set_line_cap((*C.widget_t)(this.handle), aline_cap));
}

func (this TProgressCircle) SetShowText(show_text bool) TRet {
  return TRet(C.progress_circle_set_show_text((*C.widget_t)(this.handle), (C.bool_t)(show_text)));
}

func (this TProgressCircle) SetCounterClockWise(counter_clock_wise bool) TRet {
  return TRet(C.progress_circle_set_counter_clock_wise((*C.widget_t)(this.handle), (C.bool_t)(counter_clock_wise)));
}

func (this TProgressCircle) GetMax() float64 {
  return (float64)((*C.progress_circle_t)(unsafe.Pointer(this.handle)).max);
}

func (this TProgressCircle) GetFormat() string {
  return C.GoString((*C.progress_circle_t)(unsafe.Pointer(this.handle)).format);
}

func (this TProgressCircle) GetStartAngle() int32 {
  return (int32)((*C.progress_circle_t)(unsafe.Pointer(this.handle)).start_angle);
}

func (this TProgressCircle) GetLineWidth() uint32 {
  return (uint32)((*C.progress_circle_t)(unsafe.Pointer(this.handle)).line_width);
}

func (this TProgressCircle) GetLineCap() string {
  return C.GoString((*C.progress_circle_t)(unsafe.Pointer(this.handle)).line_cap);
}

func (this TProgressCircle) GetCounterClockWise() bool {
  return (bool)((*C.progress_circle_t)(unsafe.Pointer(this.handle)).counter_clock_wise);
}

func (this TProgressCircle) GetShowText() bool {
  return (bool)((*C.progress_circle_t)(unsafe.Pointer(this.handle)).show_text);
}

type TProgressEvent struct {
  TEvent
}

func TProgressEventCast(event TEvent) TProgressEvent {
  retObj := TProgressEvent{}
  retObj.handle = unsafe.Pointer(C.progress_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TProgressEvent) GetPercent() uint32 {
  return (uint32)((*C.progress_event_t)(unsafe.Pointer(this.handle)).percent);
}

type TPropChangeEvent struct {
  TEvent
}

func TPropChangeEventCast(event TEvent) TPropChangeEvent {
  retObj := TPropChangeEvent{}
  retObj.handle = unsafe.Pointer(C.prop_change_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TPropChangeEvent) GetName() string {
  return C.GoString((*C.prop_change_event_t)(unsafe.Pointer(this.handle)).name);
}

type TRect struct {
  handle unsafe.Pointer
}

func TRectCreate(x int, y int, w int, h int) TRect {
  retObj := TRect{}
  retObj.handle = unsafe.Pointer(C.rect_create((C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj
}

func (this TRect) Set(x int, y int, w int, h int) TRect {
  retObj := TRect{}
  retObj.handle = unsafe.Pointer(C.rect_set((*C.rect_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj
}

func TRectCast(rect TRect) TRect {
  retObj := TRect{}
  retObj.handle = unsafe.Pointer(C.rect_cast((*C.rect_t)(rect.handle)))
  return retObj
}

func (this TRect) Destroy() TRet {
  return TRet(C.rect_destroy((*C.rect_t)(this.handle)));
}

func (this TRect) GetX() int {
  return (int)((*C.rect_t)(unsafe.Pointer(this.handle)).x);
}

func (this TRect) GetY() int {
  return (int)((*C.rect_t)(unsafe.Pointer(this.handle)).y);
}

func (this TRect) GetW() int {
  return (int)((*C.rect_t)(unsafe.Pointer(this.handle)).w);
}

func (this TRect) GetH() int {
  return (int)((*C.rect_t)(unsafe.Pointer(this.handle)).h);
}

type TRectf struct {
  handle unsafe.Pointer
}

func (this TRectf) GetX() float64 {
  return (float64)((*C.rectf_t)(unsafe.Pointer(this.handle)).x);
}

func (this TRectf) GetY() float64 {
  return (float64)((*C.rectf_t)(unsafe.Pointer(this.handle)).y);
}

func (this TRectf) GetW() float64 {
  return (float64)((*C.rectf_t)(unsafe.Pointer(this.handle)).w);
}

func (this TRectf) GetH() float64 {
  return (float64)((*C.rectf_t)(unsafe.Pointer(this.handle)).h);
}

type TRet int
const (
  RET_OK TRet = C.RET_OK
  RET_OOM TRet = C.RET_OOM
  RET_FAIL TRet = C.RET_FAIL
  RET_NOT_IMPL TRet = C.RET_NOT_IMPL
  RET_QUIT TRet = C.RET_QUIT
  RET_FOUND TRet = C.RET_FOUND
  RET_BUSY TRet = C.RET_BUSY
  RET_REMOVE TRet = C.RET_REMOVE
  RET_REPEAT TRet = C.RET_REPEAT
  RET_NOT_FOUND TRet = C.RET_NOT_FOUND
  RET_DONE TRet = C.RET_DONE
  RET_STOP TRet = C.RET_STOP
  RET_SKIP TRet = C.RET_SKIP
  RET_CONTINUE TRet = C.RET_CONTINUE
  RET_OBJECT_CHANGED TRet = C.RET_OBJECT_CHANGED
  RET_ITEMS_CHANGED TRet = C.RET_ITEMS_CHANGED
  RET_BAD_PARAMS TRet = C.RET_BAD_PARAMS
  RET_TIMEOUT TRet = C.RET_TIMEOUT
  RET_CRC TRet = C.RET_CRC
  RET_IO TRet = C.RET_IO
  RET_EOS TRet = C.RET_EOS
  RET_NOT_MODIFIED TRet = C.RET_NOT_MODIFIED
  RET_NO_PERMISSION TRet = C.RET_NO_PERMISSION
  RET_INVALID_ADDR TRet = C.RET_INVALID_ADDR
  RET_EXCEED_RANGE TRet = C.RET_EXCEED_RANGE
  RET_MAX_NR TRet = C.RET_MAX_NR
)
type TRichText struct {
  TWidget
}

func TRichTextCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TRichText{}
  retObj.handle = unsafe.Pointer(C.rich_text_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TRichText) SetText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.rich_text_set_text((*C.widget_t)(this.handle), atext));
}

func (this TRichText) SetYslidable(yslidable bool) TRet {
  return TRet(C.rich_text_set_yslidable((*C.widget_t)(this.handle), (C.bool_t)(yslidable)));
}

func TRichTextCast(widget TWidget) TRichText {
  retObj := TRichText{}
  retObj.handle = unsafe.Pointer(C.rich_text_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TRichText) GetLineGap() uint32 {
  return (uint32)((*C.rich_text_t)(unsafe.Pointer(this.handle)).line_gap);
}

func (this TRichText) GetYslidable() bool {
  return (bool)((*C.rich_text_t)(unsafe.Pointer(this.handle)).yslidable);
}

type TRichTextView struct {
  TWidget
}

func TRichTextViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TRichTextView{}
  retObj.handle = unsafe.Pointer(C.rich_text_view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TRichTextViewCast(widget TWidget) TRichTextView {
  retObj := TRichTextView{}
  retObj.handle = unsafe.Pointer(C.rich_text_view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TRlog struct {
  handle unsafe.Pointer
}

func TRlogCreate(filename_pattern string, max_size uint32, buff_size uint32) TRlog {
  afilename_pattern := C.CString(filename_pattern)
  defer C.free(unsafe.Pointer(afilename_pattern))
  retObj := TRlog{}
  retObj.handle = unsafe.Pointer(C.rlog_create(afilename_pattern, (C.uint32_t)(max_size), (C.uint32_t)(buff_size)))
  return retObj
}

func (this TRlog) Write(str string) TRet {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  return TRet(C.rlog_write((*C.rlog_t)(this.handle), astr));
}

type TRow struct {
  TWidget
}

func TRowCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TRow{}
  retObj.handle = unsafe.Pointer(C.row_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TRowCast(widget TWidget) TRow {
  retObj := TRow{}
  retObj.handle = unsafe.Pointer(C.row_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TScrollBar struct {
  TWidget
}

func TScrollBarCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TScrollBar{}
  retObj.handle = unsafe.Pointer(C.scroll_bar_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TScrollBarCast(widget TWidget) TScrollBar {
  retObj := TScrollBar{}
  retObj.handle = unsafe.Pointer(C.scroll_bar_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func TScrollBarCreateMobile(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TScrollBar{}
  retObj.handle = unsafe.Pointer(C.scroll_bar_create_mobile((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TScrollBarCreateDesktop(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TScrollBar{}
  retObj.handle = unsafe.Pointer(C.scroll_bar_create_desktop((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TScrollBar) SetParams(virtual_size int32, row int32) TRet {
  return TRet(C.scroll_bar_set_params((*C.widget_t)(this.handle), (C.int32_t)(virtual_size), (C.int32_t)(row)));
}

func (this TScrollBar) ScrollTo(value int32, duration int32) TRet {
  return TRet(C.scroll_bar_scroll_to((*C.widget_t)(this.handle), (C.int32_t)(value), (C.int32_t)(duration)));
}

func (this TScrollBar) SetValue(value int32) TRet {
  return TRet(C.scroll_bar_set_value((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TScrollBar) AddDelta(delta int32) TRet {
  return TRet(C.scroll_bar_add_delta((*C.widget_t)(this.handle), (C.int32_t)(delta)));
}

func (this TScrollBar) SetValueOnly(value int32) TRet {
  return TRet(C.scroll_bar_set_value_only((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TScrollBar) SetAutoHide(auto_hide bool) TRet {
  return TRet(C.scroll_bar_set_auto_hide((*C.widget_t)(this.handle), (C.bool_t)(auto_hide)));
}

func (this TScrollBar) IsMobile() bool {
  return (bool)(C.scroll_bar_is_mobile((*C.widget_t)(this.handle)));
}

func (this TScrollBar) SetAnimatorTime(animator_time uint32) TRet {
  return TRet(C.scroll_bar_set_animator_time((*C.widget_t)(this.handle), (C.uint32_t)(animator_time)));
}

func (this TScrollBar) HideByOpacityAnimation(duration int32, delay int32) TRet {
  return TRet(C.scroll_bar_hide_by_opacity_animation((*C.widget_t)(this.handle), (C.int32_t)(duration), (C.int32_t)(delay)));
}

func (this TScrollBar) ShowByOpacityAnimation(duration int32, delay int32) TRet {
  return TRet(C.scroll_bar_show_by_opacity_animation((*C.widget_t)(this.handle), (C.int32_t)(duration), (C.int32_t)(delay)));
}

func (this TScrollBar) SetWheelScroll(scroll bool) TRet {
  return TRet(C.scroll_bar_set_wheel_scroll((*C.widget_t)(this.handle), (C.bool_t)(scroll)));
}

func (this TScrollBar) SetScrollDelta(scroll_delta uint32) TRet {
  return TRet(C.scroll_bar_set_scroll_delta((*C.widget_t)(this.handle), (C.uint32_t)(scroll_delta)));
}

func (this TScrollBar) GetVirtualSize() int32 {
  return (int32)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).virtual_size);
}

func (this TScrollBar) GetRow() int32 {
  return (int32)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).row);
}

func (this TScrollBar) GetAnimatorTime() uint32 {
  return (uint32)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).animator_time);
}

func (this TScrollBar) GetScrollDelta() uint32 {
  return (uint32)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).scroll_delta);
}

func (this TScrollBar) GetAnimatable() bool {
  return (bool)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).animatable);
}

func (this TScrollBar) GetAutoHide() bool {
  return (bool)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).auto_hide);
}

func (this TScrollBar) GetWheelScroll() bool {
  return (bool)((*C.scroll_bar_t)(unsafe.Pointer(this.handle)).wheel_scroll);
}

type TScrollView struct {
  TWidget
}

func TScrollViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TScrollView{}
  retObj.handle = unsafe.Pointer(C.scroll_view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TScrollViewCast(widget TWidget) TScrollView {
  retObj := TScrollView{}
  retObj.handle = unsafe.Pointer(C.scroll_view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TScrollView) SetVirtualW(w int) TRet {
  return TRet(C.scroll_view_set_virtual_w((*C.widget_t)(this.handle), (C.wh_t)(w)));
}

func (this TScrollView) SetVirtualH(h int) TRet {
  return TRet(C.scroll_view_set_virtual_h((*C.widget_t)(this.handle), (C.wh_t)(h)));
}

func (this TScrollView) SetXslidable(xslidable bool) TRet {
  return TRet(C.scroll_view_set_xslidable((*C.widget_t)(this.handle), (C.bool_t)(xslidable)));
}

func (this TScrollView) SetYslidable(yslidable bool) TRet {
  return TRet(C.scroll_view_set_yslidable((*C.widget_t)(this.handle), (C.bool_t)(yslidable)));
}

func (this TScrollView) SetSnapToPage(snap_to_page bool) TRet {
  return TRet(C.scroll_view_set_snap_to_page((*C.widget_t)(this.handle), (C.bool_t)(snap_to_page)));
}

func (this TScrollView) SetMoveToPage(move_to_page bool) TRet {
  return TRet(C.scroll_view_set_move_to_page((*C.widget_t)(this.handle), (C.bool_t)(move_to_page)));
}

func (this TScrollView) SetRecursive(recursive bool) TRet {
  return TRet(C.scroll_view_set_recursive((*C.widget_t)(this.handle), (C.bool_t)(recursive)));
}

func (this TScrollView) SetRecursiveOnly(recursive bool) TRet {
  return TRet(C.scroll_view_set_recursive_only((*C.widget_t)(this.handle), (C.bool_t)(recursive)));
}

func (this TScrollView) SetOffset(xoffset int32, yoffset int32) TRet {
  return TRet(C.scroll_view_set_offset((*C.widget_t)(this.handle), (C.int32_t)(xoffset), (C.int32_t)(yoffset)));
}

func (this TScrollView) SetSpeedScale(xspeed_scale float64, yspeed_scale float64) TRet {
  return TRet(C.scroll_view_set_speed_scale((*C.widget_t)(this.handle), (C.float_t)(xspeed_scale), (C.float_t)(yspeed_scale)));
}

func (this TScrollView) SetSlideLimitRatio(slide_limit_ratio float64) TRet {
  return TRet(C.scroll_view_set_slide_limit_ratio((*C.widget_t)(this.handle), (C.float_t)(slide_limit_ratio)));
}

func (this TScrollView) ScrollTo(xoffset_end int32, yoffset_end int32, duration int32) TRet {
  return TRet(C.scroll_view_scroll_to((*C.widget_t)(this.handle), (C.int32_t)(xoffset_end), (C.int32_t)(yoffset_end), (C.int32_t)(duration)));
}

func (this TScrollView) ScrollDeltaTo(xoffset_delta int32, yoffset_delta int32, duration int32) TRet {
  return TRet(C.scroll_view_scroll_delta_to((*C.widget_t)(this.handle), (C.int32_t)(xoffset_delta), (C.int32_t)(yoffset_delta), (C.int32_t)(duration)));
}

func (this TScrollView) GetVirtualW() int {
  return (int)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).virtual_w);
}

func (this TScrollView) GetVirtualH() int {
  return (int)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).virtual_h);
}

func (this TScrollView) GetXoffset() int32 {
  return (int32)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).xoffset);
}

func (this TScrollView) GetYoffset() int32 {
  return (int32)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).yoffset);
}

func (this TScrollView) GetXspeedScale() float64 {
  return (float64)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).xspeed_scale);
}

func (this TScrollView) GetYspeedScale() float64 {
  return (float64)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).yspeed_scale);
}

func (this TScrollView) GetXslidable() bool {
  return (bool)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).xslidable);
}

func (this TScrollView) GetYslidable() bool {
  return (bool)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).yslidable);
}

func (this TScrollView) GetSnapToPage() bool {
  return (bool)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).snap_to_page);
}

func (this TScrollView) GetMoveToPage() bool {
  return (bool)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).move_to_page);
}

func (this TScrollView) GetRecursive() bool {
  return (bool)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).recursive);
}

func (this TScrollView) GetSlideLimitRatio() float64 {
  return (float64)((*C.scroll_view_t)(unsafe.Pointer(this.handle)).slide_limit_ratio);
}

type TSerialWidget struct {
  TWidget
}

func TSerialWidgetCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSerialWidget{}
  retObj.handle = unsafe.Pointer(C.serial_widget_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSerialWidgetCast(widget TWidget) TSerialWidget {
  retObj := TSerialWidget{}
  retObj.handle = unsafe.Pointer(C.serial_widget_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSerialWidget) SetBaudrate(baudrate uint32) TRet {
  return TRet(C.serial_widget_set_baudrate((*C.widget_t)(this.handle), (C.uint32_t)(baudrate)));
}

func (this TSerialWidget) SetDevice(device string) TRet {
  adevice := C.CString(device)
  defer C.free(unsafe.Pointer(adevice))
  return TRet(C.serial_widget_set_device((*C.widget_t)(this.handle), adevice));
}

func (this TSerialWidget) SetBytesize(bytesize uint32) TRet {
  return TRet(C.serial_widget_set_bytesize((*C.widget_t)(this.handle), (C.uint32_t)(bytesize)));
}

func (this TSerialWidget) SetParity(parity uint32) TRet {
  return TRet(C.serial_widget_set_parity((*C.widget_t)(this.handle), (C.uint32_t)(parity)));
}

func (this TSerialWidget) SetStopbits(stopbits uint32) TRet {
  return TRet(C.serial_widget_set_stopbits((*C.widget_t)(this.handle), (C.uint32_t)(stopbits)));
}

func (this TSerialWidget) SetFlowcontrol(flowcontrol uint32) TRet {
  return TRet(C.serial_widget_set_flowcontrol((*C.widget_t)(this.handle), (C.uint32_t)(flowcontrol)));
}

func (this TSerialWidget) SetCheckInterval(check_interval uint32) TRet {
  return TRet(C.serial_widget_set_check_interval((*C.widget_t)(this.handle), (C.uint32_t)(check_interval)));
}

func (this TSerialWidget) GetDevice() string {
  return C.GoString((*C.serial_widget_t)(unsafe.Pointer(this.handle)).device);
}

func (this TSerialWidget) GetBaudrate() uint32 {
  return (uint32)((*C.serial_widget_t)(unsafe.Pointer(this.handle)).baudrate);
}

func (this TSerialWidget) GetBytesize() uint32 {
  return (uint32)((*C.serial_widget_t)(unsafe.Pointer(this.handle)).bytesize);
}

func (this TSerialWidget) GetParity() uint32 {
  return (uint32)((*C.serial_widget_t)(unsafe.Pointer(this.handle)).parity);
}

func (this TSerialWidget) GetStopbits() uint32 {
  return (uint32)((*C.serial_widget_t)(unsafe.Pointer(this.handle)).stopbits);
}

func (this TSerialWidget) GetFlowcontrol() uint32 {
  return (uint32)((*C.serial_widget_t)(unsafe.Pointer(this.handle)).flowcontrol);
}

func (this TSerialWidget) GetCheckInterval() uint32 {
  return (uint32)((*C.serial_widget_t)(unsafe.Pointer(this.handle)).check_interval);
}

type TSlideIndicator struct {
  TWidget
}

func TSlideIndicatorCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSlideIndicator{}
  retObj.handle = unsafe.Pointer(C.slide_indicator_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSlideIndicatorCreateLinear(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSlideIndicator{}
  retObj.handle = unsafe.Pointer(C.slide_indicator_create_linear((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSlideIndicatorCreateArc(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSlideIndicator{}
  retObj.handle = unsafe.Pointer(C.slide_indicator_create_arc((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSlideIndicatorCast(widget TWidget) TSlideIndicator {
  retObj := TSlideIndicator{}
  retObj.handle = unsafe.Pointer(C.slide_indicator_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSlideIndicator) SetValue(value uint32) TRet {
  return TRet(C.slide_indicator_set_value((*C.widget_t)(this.handle), (C.uint32_t)(value)));
}

func (this TSlideIndicator) SetMax(max uint32) TRet {
  return TRet(C.slide_indicator_set_max((*C.widget_t)(this.handle), (C.uint32_t)(max)));
}

func (this TSlideIndicator) SetDefaultPaint(default_paint TIndicatorDefaultPaint) TRet {
  return TRet(C.slide_indicator_set_default_paint((*C.widget_t)(this.handle), (C.indicator_default_paint_t)(default_paint)));
}

func (this TSlideIndicator) SetAutoHide(auto_hide uint16) TRet {
  return TRet(C.slide_indicator_set_auto_hide((*C.widget_t)(this.handle), (C.uint16_t)(auto_hide)));
}

func (this TSlideIndicator) SetMargin(margin int32) TRet {
  return TRet(C.slide_indicator_set_margin((*C.widget_t)(this.handle), (C.int32_t)(margin)));
}

func (this TSlideIndicator) SetSpacing(spacing float64) TRet {
  return TRet(C.slide_indicator_set_spacing((*C.widget_t)(this.handle), (C.float_t)(spacing)));
}

func (this TSlideIndicator) SetSize(size uint32) TRet {
  return TRet(C.slide_indicator_set_size((*C.widget_t)(this.handle), (C.uint32_t)(size)));
}

func (this TSlideIndicator) SetAnchor(anchor_x string, anchor_y string) TRet {
  aanchor_x := C.CString(anchor_x)
  defer C.free(unsafe.Pointer(aanchor_x))
  aanchor_y := C.CString(anchor_y)
  defer C.free(unsafe.Pointer(aanchor_y))
  return TRet(C.slide_indicator_set_anchor((*C.widget_t)(this.handle), aanchor_x, aanchor_y));
}

func (this TSlideIndicator) SetIndicatedTarget(target_name string) TRet {
  atarget_name := C.CString(target_name)
  defer C.free(unsafe.Pointer(atarget_name))
  return TRet(C.slide_indicator_set_indicated_target((*C.widget_t)(this.handle), atarget_name));
}

func (this TSlideIndicator) SetTransition(transition bool) TRet {
  return TRet(C.slide_indicator_set_transition((*C.widget_t)(this.handle), (C.bool_t)(transition)));
}

func (this TSlideIndicator) GetMax() uint32 {
  return (uint32)((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).max);
}

func (this TSlideIndicator) GetDefaultPaint() TIndicatorDefaultPaint {
  return TIndicatorDefaultPaint((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).default_paint);
}

func (this TSlideIndicator) GetAutoHide() uint16 {
  return (uint16)((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).auto_hide);
}

func (this TSlideIndicator) GetMargin() int32 {
  return (int32)((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).margin);
}

func (this TSlideIndicator) GetSpacing() float64 {
  return (float64)((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).spacing);
}

func (this TSlideIndicator) GetSize() uint32 {
  return (uint32)((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).size);
}

func (this TSlideIndicator) GetAnchorX() string {
  return C.GoString((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).anchor_x);
}

func (this TSlideIndicator) GetAnchorY() string {
  return C.GoString((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).anchor_y);
}

func (this TSlideIndicator) GetIndicatedTarget() string {
  return C.GoString((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).indicated_target);
}

func (this TSlideIndicator) GetTransition() bool {
  return (bool)((*C.slide_indicator_t)(unsafe.Pointer(this.handle)).transition);
}

type TSlideMenu struct {
  TWidget
}

func TSlideMenuCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSlideMenu{}
  retObj.handle = unsafe.Pointer(C.slide_menu_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSlideMenuCast(widget TWidget) TSlideMenu {
  retObj := TSlideMenu{}
  retObj.handle = unsafe.Pointer(C.slide_menu_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSlideMenu) SetValue(value uint32) TRet {
  return TRet(C.slide_menu_set_value((*C.widget_t)(this.handle), (C.uint32_t)(value)));
}

func (this TSlideMenu) SetAlignV(align_v TAlignV) TRet {
  return TRet(C.slide_menu_set_align_v((*C.widget_t)(this.handle), (C.align_v_t)(align_v)));
}

func (this TSlideMenu) SetMinScale(min_scale float64) TRet {
  return TRet(C.slide_menu_set_min_scale((*C.widget_t)(this.handle), (C.float_t)(min_scale)));
}

func (this TSlideMenu) SetSpacer(spacer int32) TRet {
  return TRet(C.slide_menu_set_spacer((*C.widget_t)(this.handle), (C.int32_t)(spacer)));
}

func (this TSlideMenu) SetMenuW(menu_w string) TRet {
  amenu_w := C.CString(menu_w)
  defer C.free(unsafe.Pointer(amenu_w))
  return TRet(C.slide_menu_set_menu_w((*C.widget_t)(this.handle), amenu_w));
}

func (this TSlideMenu) SetClip(clip bool) TRet {
  return TRet(C.slide_menu_set_clip((*C.widget_t)(this.handle), (C.bool_t)(clip)));
}

func (this TSlideMenu) ScrollToPrev() TRet {
  return TRet(C.slide_menu_scroll_to_prev((*C.widget_t)(this.handle)));
}

func (this TSlideMenu) ScrollToNext() TRet {
  return TRet(C.slide_menu_scroll_to_next((*C.widget_t)(this.handle)));
}

func (this TSlideMenu) GetAlignV() TAlignV {
  return TAlignV((*C.slide_menu_t)(unsafe.Pointer(this.handle)).align_v);
}

func (this TSlideMenu) GetMinScale() float64 {
  return (float64)((*C.slide_menu_t)(unsafe.Pointer(this.handle)).min_scale);
}

func (this TSlideMenu) GetSpacer() int32 {
  return (int32)((*C.slide_menu_t)(unsafe.Pointer(this.handle)).spacer);
}

func (this TSlideMenu) GetMenuW() string {
  return C.GoString((*C.slide_menu_t)(unsafe.Pointer(this.handle)).menu_w);
}

func (this TSlideMenu) GetClip() bool {
  return (bool)((*C.slide_menu_t)(unsafe.Pointer(this.handle)).clip);
}

type TSlideView struct {
  TWidget
}

func TSlideViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSlideView{}
  retObj.handle = unsafe.Pointer(C.slide_view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSlideViewCast(widget TWidget) TSlideView {
  retObj := TSlideView{}
  retObj.handle = unsafe.Pointer(C.slide_view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSlideView) SetAutoPlay(auto_play uint16) TRet {
  return TRet(C.slide_view_set_auto_play((*C.widget_t)(this.handle), (C.uint16_t)(auto_play)));
}

func (this TSlideView) SetActive(index uint32) TRet {
  return TRet(C.slide_view_set_active((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TSlideView) SetActiveEx(index uint32, animate bool) TRet {
  return TRet(C.slide_view_set_active_ex((*C.widget_t)(this.handle), (C.uint32_t)(index), (C.bool_t)(animate)));
}

func (this TSlideView) SetVertical(vertical bool) TRet {
  return TRet(C.slide_view_set_vertical((*C.widget_t)(this.handle), (C.bool_t)(vertical)));
}

func (this TSlideView) SetAnimHint(anim_hint string) TRet {
  aanim_hint := C.CString(anim_hint)
  defer C.free(unsafe.Pointer(aanim_hint))
  return TRet(C.slide_view_set_anim_hint((*C.widget_t)(this.handle), aanim_hint));
}

func (this TSlideView) SetLoop(loop bool) TRet {
  return TRet(C.slide_view_set_loop((*C.widget_t)(this.handle), (C.bool_t)(loop)));
}

func (this TSlideView) SetDragThreshold(drag_threshold uint32) TRet {
  return TRet(C.slide_view_set_drag_threshold((*C.widget_t)(this.handle), (C.uint32_t)(drag_threshold)));
}

func (this TSlideView) SetAnimatingTime(animating_time uint32) TRet {
  return TRet(C.slide_view_set_animating_time((*C.widget_t)(this.handle), (C.uint32_t)(animating_time)));
}

func (this TSlideView) RemoveIndex(index uint32) TRet {
  return TRet(C.slide_view_remove_index((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TSlideView) GetVertical() bool {
  return (bool)((*C.slide_view_t)(unsafe.Pointer(this.handle)).vertical);
}

func (this TSlideView) GetAutoPlay() uint16 {
  return (uint16)((*C.slide_view_t)(unsafe.Pointer(this.handle)).auto_play);
}

func (this TSlideView) GetLoop() bool {
  return (bool)((*C.slide_view_t)(unsafe.Pointer(this.handle)).loop);
}

func (this TSlideView) GetAnimHint() string {
  return C.GoString((*C.slide_view_t)(unsafe.Pointer(this.handle)).anim_hint);
}

func (this TSlideView) GetDragThreshold() uint32 {
  return (uint32)((*C.slide_view_t)(unsafe.Pointer(this.handle)).drag_threshold);
}

func (this TSlideView) GetAnimatingTime() uint32 {
  return (uint32)((*C.slide_view_t)(unsafe.Pointer(this.handle)).animating_time);
}

type TSlider struct {
  TWidget
}

func TSliderCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSlider{}
  retObj.handle = unsafe.Pointer(C.slider_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSliderCast(widget TWidget) TSlider {
  retObj := TSlider{}
  retObj.handle = unsafe.Pointer(C.slider_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSlider) SetValue(value float64) TRet {
  return TRet(C.slider_set_value((*C.widget_t)(this.handle), (C.double)(value)));
}

func (this TSlider) SetMin(min float64) TRet {
  return TRet(C.slider_set_min((*C.widget_t)(this.handle), (C.double)(min)));
}

func (this TSlider) SetMax(max float64) TRet {
  return TRet(C.slider_set_max((*C.widget_t)(this.handle), (C.double)(max)));
}

func (this TSlider) SetLineCap(line_cap string) TRet {
  aline_cap := C.CString(line_cap)
  defer C.free(unsafe.Pointer(aline_cap))
  return TRet(C.slider_set_line_cap((*C.widget_t)(this.handle), aline_cap));
}

func (this TSlider) SetStep(step float64) TRet {
  return TRet(C.slider_set_step((*C.widget_t)(this.handle), (C.double)(step)));
}

func (this TSlider) SetBarSize(bar_size uint32) TRet {
  return TRet(C.slider_set_bar_size((*C.widget_t)(this.handle), (C.uint32_t)(bar_size)));
}

func (this TSlider) SetVertical(vertical bool) TRet {
  return TRet(C.slider_set_vertical((*C.widget_t)(this.handle), (C.bool_t)(vertical)));
}

func (this TSlider) SetDragThreshold(drag_threshold uint32) TRet {
  return TRet(C.slider_set_drag_threshold((*C.widget_t)(this.handle), (C.uint32_t)(drag_threshold)));
}

func (this TSlider) GetMin() float64 {
  return (float64)((*C.slider_t)(unsafe.Pointer(this.handle)).min);
}

func (this TSlider) GetMax() float64 {
  return (float64)((*C.slider_t)(unsafe.Pointer(this.handle)).max);
}

func (this TSlider) GetStep() float64 {
  return (float64)((*C.slider_t)(unsafe.Pointer(this.handle)).step);
}

func (this TSlider) GetBarSize() uint32 {
  return (uint32)((*C.slider_t)(unsafe.Pointer(this.handle)).bar_size);
}

func (this TSlider) GetDraggerSize() uint32 {
  return (uint32)((*C.slider_t)(unsafe.Pointer(this.handle)).dragger_size);
}

func (this TSlider) GetLineCap() string {
  return C.GoString((*C.slider_t)(unsafe.Pointer(this.handle)).line_cap);
}

func (this TSlider) GetVertical() bool {
  return (bool)((*C.slider_t)(unsafe.Pointer(this.handle)).vertical);
}

func (this TSlider) GetDraggerAdaptToIcon() bool {
  return (bool)((*C.slider_t)(unsafe.Pointer(this.handle)).dragger_adapt_to_icon);
}

func (this TSlider) GetSlideWithBar() bool {
  return (bool)((*C.slider_t)(unsafe.Pointer(this.handle)).slide_with_bar);
}

func (this TSlider) GetDragThreshold() uint32 {
  return (uint32)((*C.slider_t)(unsafe.Pointer(this.handle)).drag_threshold);
}

type TSpinBox struct {
  TEdit
}

func TSpinBoxCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSpinBox{}
  retObj.handle = unsafe.Pointer(C.spin_box_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSpinBoxCast(widget TWidget) TSpinBox {
  retObj := TSpinBox{}
  retObj.handle = unsafe.Pointer(C.spin_box_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSpinBox) SetEasyTouchMode(easy_touch_mode bool) TRet {
  return TRet(C.spin_box_set_easy_touch_mode((*C.widget_t)(this.handle), (C.bool_t)(easy_touch_mode)));
}

func (this TSpinBox) SetButtonPosition(button_position string) TRet {
  abutton_position := C.CString(button_position)
  defer C.free(unsafe.Pointer(abutton_position))
  return TRet(C.spin_box_set_button_position((*C.widget_t)(this.handle), abutton_position));
}

func (this TSpinBox) SpinSetRepeat(repeat int32) TRet {
  return TRet(C.spin_set_repeat((*C.widget_t)(this.handle), (C.int32_t)(repeat)));
}

func (this TSpinBox) GetEasyTouchMode() bool {
  return (bool)((*C.spin_box_t)(unsafe.Pointer(this.handle)).easy_touch_mode);
}

func (this TSpinBox) GetButtonPosition() string {
  return C.GoString((*C.spin_box_t)(unsafe.Pointer(this.handle)).button_position);
}

type TStyle struct {
  handle unsafe.Pointer
}

func (this TStyle) NotifyWidgetStateChanged(widget TWidget) TRet {
  return TRet(C.style_notify_widget_state_changed((*C.style_t)(this.handle), (*C.widget_t)(widget.handle)));
}

func (this TStyle) IsValid() bool {
  return (bool)(C.style_is_valid((*C.style_t)(this.handle)));
}

func (this TStyle) GetInt(name string, defval int32) int32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int32)(C.style_get_int((*C.style_t)(this.handle), aname, (C.int32_t)(defval)));
}

func (this TStyle) GetUint(name string, defval uint32) uint32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (uint32)(C.style_get_uint((*C.style_t)(this.handle), aname, (C.uint32_t)(defval)));
}

func (this TStyle) GetStr(name string, defval string) string {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  adefval := C.CString(defval)
  defer C.free(unsafe.Pointer(adefval))
  return C.GoString(C.style_get_str((*C.style_t)(this.handle), aname, adefval));
}

func (this TStyle) Get(state string, name string, value TValue) TRet {
  astate := C.CString(state)
  defer C.free(unsafe.Pointer(astate))
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.style_get((*C.style_t)(this.handle), astate, aname, (*C.value_t)(value.handle)));
}

func (this TStyle) Set(state string, name string, value TValue) TRet {
  astate := C.CString(state)
  defer C.free(unsafe.Pointer(astate))
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.style_set((*C.style_t)(this.handle), astate, aname, (*C.value_t)(value.handle)));
}

func (this TStyle) UpdateState(theme TTheme, widget_type string, style_name string, widget_state string) TRet {
  awidget_type := C.CString(widget_type)
  defer C.free(unsafe.Pointer(awidget_type))
  astyle_name := C.CString(style_name)
  defer C.free(unsafe.Pointer(astyle_name))
  awidget_state := C.CString(widget_state)
  defer C.free(unsafe.Pointer(awidget_state))
  return TRet(C.style_update_state((*C.style_t)(this.handle), (*C.theme_t)(theme.handle), awidget_type, astyle_name, awidget_state));
}

func (this TStyle) GetStyleState() string {
  return C.GoString(C.style_get_style_state((*C.style_t)(this.handle)));
}

func (this TStyle) IsMutable() bool {
  return (bool)(C.style_is_mutable((*C.style_t)(this.handle)));
}

func (this TStyle) GetStyleType() string {
  return C.GoString(C.style_get_style_type((*C.style_t)(this.handle)));
}

type TStyleId string
const (
  STYLE_ID_BG_COLOR string = C.STYLE_ID_BG_COLOR
  STYLE_ID_FG_COLOR string = C.STYLE_ID_FG_COLOR
  STYLE_ID_DRAGGER_COLOR string = C.STYLE_ID_DRAGGER_COLOR
  STYLE_ID_MASK_COLOR string = C.STYLE_ID_MASK_COLOR
  STYLE_ID_FONT_NAME string = C.STYLE_ID_FONT_NAME
  STYLE_ID_FONT_SIZE string = C.STYLE_ID_FONT_SIZE
  STYLE_ID_FONT_STYLE string = C.STYLE_ID_FONT_STYLE
  STYLE_ID_TEXT_COLOR string = C.STYLE_ID_TEXT_COLOR
  STYLE_ID_HIGHLIGHT_FONT_NAME string = C.STYLE_ID_HIGHLIGHT_FONT_NAME
  STYLE_ID_HIGHLIGHT_FONT_SIZE string = C.STYLE_ID_HIGHLIGHT_FONT_SIZE
  STYLE_ID_HIGHLIGHT_TEXT_COLOR string = C.STYLE_ID_HIGHLIGHT_TEXT_COLOR
  STYLE_ID_TIPS_TEXT_COLOR string = C.STYLE_ID_TIPS_TEXT_COLOR
  STYLE_ID_TEXT_ALIGN_H string = C.STYLE_ID_TEXT_ALIGN_H
  STYLE_ID_TEXT_ALIGN_V string = C.STYLE_ID_TEXT_ALIGN_V
  STYLE_ID_BORDER_COLOR string = C.STYLE_ID_BORDER_COLOR
  STYLE_ID_BORDER_WIDTH string = C.STYLE_ID_BORDER_WIDTH
  STYLE_ID_BORDER string = C.STYLE_ID_BORDER
  STYLE_ID_BG_IMAGE string = C.STYLE_ID_BG_IMAGE
  STYLE_ID_BG_IMAGE_DRAW_TYPE string = C.STYLE_ID_BG_IMAGE_DRAW_TYPE
  STYLE_ID_ICON string = C.STYLE_ID_ICON
  STYLE_ID_FG_IMAGE string = C.STYLE_ID_FG_IMAGE
  STYLE_ID_FG_IMAGE_DRAW_TYPE string = C.STYLE_ID_FG_IMAGE_DRAW_TYPE
  STYLE_ID_SPACER string = C.STYLE_ID_SPACER
  STYLE_ID_MARGIN string = C.STYLE_ID_MARGIN
  STYLE_ID_MARGIN_LEFT string = C.STYLE_ID_MARGIN_LEFT
  STYLE_ID_MARGIN_RIGHT string = C.STYLE_ID_MARGIN_RIGHT
  STYLE_ID_MARGIN_TOP string = C.STYLE_ID_MARGIN_TOP
  STYLE_ID_MARGIN_BOTTOM string = C.STYLE_ID_MARGIN_BOTTOM
  STYLE_ID_ICON_AT string = C.STYLE_ID_ICON_AT
  STYLE_ID_ACTIVE_ICON string = C.STYLE_ID_ACTIVE_ICON
  STYLE_ID_X_OFFSET string = C.STYLE_ID_X_OFFSET
  STYLE_ID_Y_OFFSET string = C.STYLE_ID_Y_OFFSET
  STYLE_ID_SELECTED_BG_COLOR string = C.STYLE_ID_SELECTED_BG_COLOR
  STYLE_ID_SELECTED_FG_COLOR string = C.STYLE_ID_SELECTED_FG_COLOR
  STYLE_ID_SELECTED_TEXT_COLOR string = C.STYLE_ID_SELECTED_TEXT_COLOR
  STYLE_ID_ROUND_RADIUS string = C.STYLE_ID_ROUND_RADIUS
  STYLE_ID_ROUND_RADIUS_TOP_LEFT string = C.STYLE_ID_ROUND_RADIUS_TOP_LEFT
  STYLE_ID_ROUND_RADIUS_TOP_RIGHT string = C.STYLE_ID_ROUND_RADIUS_TOP_RIGHT
  STYLE_ID_ROUND_RADIUS_BOTTOM_LEFT string = C.STYLE_ID_ROUND_RADIUS_BOTTOM_LEFT
  STYLE_ID_ROUND_RADIUS_BOTTOM_RIGHT string = C.STYLE_ID_ROUND_RADIUS_BOTTOM_RIGHT
  STYLE_ID_CHILDREN_LAYOUT string = C.STYLE_ID_CHILDREN_LAYOUT
  STYLE_ID_SELF_LAYOUT string = C.STYLE_ID_SELF_LAYOUT
  STYLE_ID_FOCUSABLE string = C.STYLE_ID_FOCUSABLE
  STYLE_ID_FEEDBACK string = C.STYLE_ID_FEEDBACK
  STYLE_ID_CLEAR_BG string = C.STYLE_ID_CLEAR_BG
  STYLE_ID_GRID_COLOR string = C.STYLE_ID_GRID_COLOR
  STYLE_ID_EVEN_BG_COLOR string = C.STYLE_ID_EVEN_BG_COLOR
  STYLE_ID_ODD_BG_COLOR string = C.STYLE_ID_ODD_BG_COLOR
)
type TStyleMutable struct {
  TStyle
}

func (this TStyleMutable) SetName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.style_mutable_set_name((*C.style_t)(this.handle), aname));
}

func (this TStyleMutable) SetInt(state string, name string, val uint32) TRet {
  astate := C.CString(state)
  defer C.free(unsafe.Pointer(astate))
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.style_mutable_set_int((*C.style_t)(this.handle), astate, aname, (C.uint32_t)(val)));
}

func TStyleMutableCast(s TStyle) TStyleMutable {
  retObj := TStyleMutable{}
  retObj.handle = unsafe.Pointer(C.style_mutable_cast((*C.style_t)(s.handle)))
  return retObj
}

func TStyleMutableCreate(default_style TStyle) TStyle {
  retObj := TStyleMutable{}
  retObj.handle = unsafe.Pointer(C.style_mutable_create((*C.style_t)(default_style.handle)))
  return retObj.TStyle
}

func (this TStyleMutable) GetName() string {
  return C.GoString((*C.style_mutable_t)(unsafe.Pointer(this.handle)).name);
}

type TSvgImage struct {
  TImageBase
}

func TSvgImageCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSvgImage{}
  retObj.handle = unsafe.Pointer(C.svg_image_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TSvgImage) SetImage(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.svg_image_set_image((*C.widget_t)(this.handle), aname));
}

func (this TSvgImage) SetCacheMode(is_cache_mode bool) TRet {
  return TRet(C.svg_image_set_cache_mode((*C.widget_t)(this.handle), (C.bool_t)(is_cache_mode)));
}

func (this TSvgImage) SetDrawType(draw_type TImageDrawType) TRet {
  return TRet(C.svg_image_set_draw_type((*C.widget_t)(this.handle), (C.image_draw_type_t)(draw_type)));
}

func TSvgImageCast(widget TWidget) TSvgImage {
  retObj := TSvgImage{}
  retObj.handle = unsafe.Pointer(C.svg_image_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSvgImage) GetIsCacheMode() bool {
  return (bool)((*C.svg_image_t)(unsafe.Pointer(this.handle)).is_cache_mode);
}

func (this TSvgImage) GetDrawType() TImageDrawType {
  return TImageDrawType((*C.svg_image_t)(unsafe.Pointer(this.handle)).draw_type);
}

type TSwitch struct {
  TWidget
}

func TSwitchCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSwitch{}
  retObj.handle = unsafe.Pointer(C.switch_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TSwitch) SetValue(value bool) TRet {
  return TRet(C.switch_set_value((*C.widget_t)(this.handle), (C.bool_t)(value)));
}

func TSwitchCast(widget TWidget) TSwitch {
  retObj := TSwitch{}
  retObj.handle = unsafe.Pointer(C.switch_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TSwitch) GetMaxXoffsetRatio() float64 {
  return (float64)((*C.switch_t)(unsafe.Pointer(this.handle)).max_xoffset_ratio);
}

type TSystemBar struct {
  TWindowBase
}

func TSystemBarCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TSystemBar{}
  retObj.handle = unsafe.Pointer(C.system_bar_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TSystemBarCast(widget TWidget) TSystemBar {
  retObj := TSystemBar{}
  retObj.handle = unsafe.Pointer(C.system_bar_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TSystemEvent struct {
  TEvent
}

func TSystemEventCast(event TEvent) TSystemEvent {
  retObj := TSystemEvent{}
  retObj.handle = unsafe.Pointer(C.system_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TSystemEvent) GetSdlEvent() unsafe.Pointer {
  return (unsafe.Pointer)((*C.system_event_t)(unsafe.Pointer(this.handle)).sdl_event);
}

type TSystemInfoFlag int
const (
  SYSTEM_INFO_FLAG_NONE TSystemInfoFlag = C.SYSTEM_INFO_FLAG_NONE
  SYSTEM_INFO_FLAG_FAST_LCD_PORTRAIT TSystemInfoFlag = C.SYSTEM_INFO_FLAG_FAST_LCD_PORTRAIT
)
type TTabButton struct {
  TWidget
}

func TTabButtonCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TTabButton{}
  retObj.handle = unsafe.Pointer(C.tab_button_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TTabButtonCast(widget TWidget) TTabButton {
  retObj := TTabButton{}
  retObj.handle = unsafe.Pointer(C.tab_button_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TTabButton) SetValue(value bool) TRet {
  return TRet(C.tab_button_set_value((*C.widget_t)(this.handle), (C.bool_t)(value)));
}

func (this TTabButton) SetIcon(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.tab_button_set_icon((*C.widget_t)(this.handle), aname));
}

func (this TTabButton) SetActiveIcon(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.tab_button_set_active_icon((*C.widget_t)(this.handle), aname));
}

func (this TTabButton) SetMaxW(max_w int32) TRet {
  return TRet(C.tab_button_set_max_w((*C.widget_t)(this.handle), (C.int32_t)(max_w)));
}

func (this TTabButton) Restack(index uint32) TRet {
  return TRet(C.tab_button_restack((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TTabButton) SetLoadUi(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.tab_button_set_load_ui((*C.widget_t)(this.handle), aname));
}

func (this TTabButton) GetLoadUi() string {
  return C.GoString((*C.tab_button_t)(unsafe.Pointer(this.handle)).load_ui);
}

func (this TTabButton) GetActiveIcon() string {
  return C.GoString((*C.tab_button_t)(unsafe.Pointer(this.handle)).active_icon);
}

func (this TTabButton) GetIcon() string {
  return C.GoString((*C.tab_button_t)(unsafe.Pointer(this.handle)).icon);
}

func (this TTabButton) GetMaxW() int32 {
  return (int32)((*C.tab_button_t)(unsafe.Pointer(this.handle)).max_w);
}

type TTabButtonGroup struct {
  TWidget
}

func TTabButtonGroupCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TTabButtonGroup{}
  retObj.handle = unsafe.Pointer(C.tab_button_group_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TTabButtonGroup) SetCompact(compact bool) TRet {
  return TRet(C.tab_button_group_set_compact((*C.widget_t)(this.handle), (C.bool_t)(compact)));
}

func (this TTabButtonGroup) SetScrollable(scrollable bool) TRet {
  return TRet(C.tab_button_group_set_scrollable((*C.widget_t)(this.handle), (C.bool_t)(scrollable)));
}

func (this TTabButtonGroup) SetDragChild(drag_child bool) TRet {
  return TRet(C.tab_button_group_set_drag_child((*C.widget_t)(this.handle), (C.bool_t)(drag_child)));
}

func TTabButtonGroupCast(widget TWidget) TTabButtonGroup {
  retObj := TTabButtonGroup{}
  retObj.handle = unsafe.Pointer(C.tab_button_group_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TTabButtonGroup) GetCompact() bool {
  return (bool)((*C.tab_button_group_t)(unsafe.Pointer(this.handle)).compact);
}

func (this TTabButtonGroup) GetScrollable() bool {
  return (bool)((*C.tab_button_group_t)(unsafe.Pointer(this.handle)).scrollable);
}

func (this TTabButtonGroup) GetDragChild() bool {
  return (bool)((*C.tab_button_group_t)(unsafe.Pointer(this.handle)).drag_child);
}

type TTabControl struct {
  TWidget
}

func TTabControlCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TTabControl{}
  retObj.handle = unsafe.Pointer(C.tab_control_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TTabControlCast(widget TWidget) TTabControl {
  retObj := TTabControl{}
  retObj.handle = unsafe.Pointer(C.tab_control_cast((*C.widget_t)(widget.handle)))
  return retObj
}

type TTextSelector struct {
  TWidget
}

func TTextSelectorCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TTextSelector{}
  retObj.handle = unsafe.Pointer(C.text_selector_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TTextSelectorCast(widget TWidget) TTextSelector {
  retObj := TTextSelector{}
  retObj.handle = unsafe.Pointer(C.text_selector_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TTextSelector) ResetOptions() TRet {
  return TRet(C.text_selector_reset_options((*C.widget_t)(this.handle)));
}

func (this TTextSelector) CountOptions() int32 {
  return (int32)(C.text_selector_count_options((*C.widget_t)(this.handle)));
}

func (this TTextSelector) AppendOption(value int32, text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.text_selector_append_option((*C.widget_t)(this.handle), (C.int32_t)(value), atext));
}

func (this TTextSelector) SetOptions(options string) TRet {
  aoptions := C.CString(options)
  defer C.free(unsafe.Pointer(aoptions))
  return TRet(C.text_selector_set_options((*C.widget_t)(this.handle), aoptions));
}

func (this TTextSelector) SetRangeOptionsEx(start int32, nr uint32, step int32, format string) TRet {
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.text_selector_set_range_options_ex((*C.widget_t)(this.handle), (C.int32_t)(start), (C.uint32_t)(nr), (C.int32_t)(step), aformat));
}

func (this TTextSelector) SetRangeOptions(start int32, nr uint32, step int32) TRet {
  return TRet(C.text_selector_set_range_options((*C.widget_t)(this.handle), (C.int32_t)(start), (C.uint32_t)(nr), (C.int32_t)(step)));
}

func (this TTextSelector) GetValueInt() int32 {
  return (int32)(C.text_selector_get_value((*C.widget_t)(this.handle)));
}

func (this TTextSelector) SetValue(value int32) TRet {
  return TRet(C.text_selector_set_value((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TTextSelector) GetTextValue() string {
  return C.GoString(C.text_selector_get_text((*C.widget_t)(this.handle)));
}

func (this TTextSelector) SetText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.text_selector_set_text((*C.widget_t)(this.handle), atext));
}

func (this TTextSelector) SetSelectedIndex(index uint32) TRet {
  return TRet(C.text_selector_set_selected_index((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TTextSelector) SetVisibleNr(visible_nr uint32) TRet {
  return TRet(C.text_selector_set_visible_nr((*C.widget_t)(this.handle), (C.uint32_t)(visible_nr)));
}

func (this TTextSelector) SetLocalizeOptions(localize_options bool) TRet {
  return TRet(C.text_selector_set_localize_options((*C.widget_t)(this.handle), (C.bool_t)(localize_options)));
}

func (this TTextSelector) SetLoopOptions(loop_options bool) TRet {
  return TRet(C.text_selector_set_loop_options((*C.widget_t)(this.handle), (C.bool_t)(loop_options)));
}

func (this TTextSelector) SetYspeedScale(yspeed_scale float64) TRet {
  return TRet(C.text_selector_set_yspeed_scale((*C.widget_t)(this.handle), (C.float_t)(yspeed_scale)));
}

func (this TTextSelector) SetAnimatingTime(animating_time uint32) TRet {
  return TRet(C.text_selector_set_animating_time((*C.widget_t)(this.handle), (C.uint32_t)(animating_time)));
}

func (this TTextSelector) SetEnableValueAnimator(enable_value_animator bool) TRet {
  return TRet(C.text_selector_set_enable_value_animator((*C.widget_t)(this.handle), (C.bool_t)(enable_value_animator)));
}

func (this TTextSelector) SetMaskEasing(mask_easing TEasingType) TRet {
  return TRet(C.text_selector_set_mask_easing((*C.widget_t)(this.handle), (C.easing_type_t)(mask_easing)));
}

func (this TTextSelector) SetMaskAreaScale(mask_area_scale float64) TRet {
  return TRet(C.text_selector_set_mask_area_scale((*C.widget_t)(this.handle), (C.float_t)(mask_area_scale)));
}

func (this TTextSelector) SetEllipses(ellipses bool) TRet {
  return TRet(C.text_selector_set_ellipses((*C.widget_t)(this.handle), (C.bool_t)(ellipses)));
}

func (this TTextSelector) GetVisibleNr() uint32 {
  return (uint32)((*C.text_selector_t)(unsafe.Pointer(this.handle)).visible_nr);
}

func (this TTextSelector) GetSelectedIndex() int32 {
  return (int32)((*C.text_selector_t)(unsafe.Pointer(this.handle)).selected_index);
}

func (this TTextSelector) GetOptions() string {
  return C.GoString((*C.text_selector_t)(unsafe.Pointer(this.handle)).options);
}

func (this TTextSelector) GetYspeedScale() float64 {
  return (float64)((*C.text_selector_t)(unsafe.Pointer(this.handle)).yspeed_scale);
}

func (this TTextSelector) GetAnimatingTime() uint32 {
  return (uint32)((*C.text_selector_t)(unsafe.Pointer(this.handle)).animating_time);
}

func (this TTextSelector) GetLocalizeOptions() bool {
  return (bool)((*C.text_selector_t)(unsafe.Pointer(this.handle)).localize_options);
}

func (this TTextSelector) GetLoopOptions() bool {
  return (bool)((*C.text_selector_t)(unsafe.Pointer(this.handle)).loop_options);
}

func (this TTextSelector) GetEnableValueAnimator() bool {
  return (bool)((*C.text_selector_t)(unsafe.Pointer(this.handle)).enable_value_animator);
}

func (this TTextSelector) GetEllipses() bool {
  return (bool)((*C.text_selector_t)(unsafe.Pointer(this.handle)).ellipses);
}

func (this TTextSelector) GetMaskEasing() TEasingType {
  return TEasingType((*C.text_selector_t)(unsafe.Pointer(this.handle)).mask_easing);
}

func (this TTextSelector) GetMaskAreaScale() float64 {
  return (float64)((*C.text_selector_t)(unsafe.Pointer(this.handle)).mask_area_scale);
}

type TTheme struct {
  handle unsafe.Pointer
}

func TThemeInstance() TTheme {
  retObj := TTheme{}
  retObj.handle = unsafe.Pointer(C.theme())
  return retObj
}

type TThemeChangeEvent struct {
  TEvent
}

func TThemeChangeEventCast(event TEvent) TThemeChangeEvent {
  retObj := TThemeChangeEvent{}
  retObj.handle = unsafe.Pointer(C.theme_change_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TThemeChangeEvent) GetName() string {
  return C.GoString((*C.theme_change_event_t)(unsafe.Pointer(this.handle)).name);
}

type TTimeClock struct {
  TWidget
}

func TTimeClockCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TTimeClock{}
  retObj.handle = unsafe.Pointer(C.time_clock_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TTimeClockCast(widget TWidget) TTimeClock {
  retObj := TTimeClock{}
  retObj.handle = unsafe.Pointer(C.time_clock_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TTimeClock) SetHour(hour int32) TRet {
  return TRet(C.time_clock_set_hour((*C.widget_t)(this.handle), (C.int32_t)(hour)));
}

func (this TTimeClock) SetMinute(minute int32) TRet {
  return TRet(C.time_clock_set_minute((*C.widget_t)(this.handle), (C.int32_t)(minute)));
}

func (this TTimeClock) SetSecond(second int32) TRet {
  return TRet(C.time_clock_set_second((*C.widget_t)(this.handle), (C.int32_t)(second)));
}

func (this TTimeClock) SetHourImage(hour string) TRet {
  ahour := C.CString(hour)
  defer C.free(unsafe.Pointer(ahour))
  return TRet(C.time_clock_set_hour_image((*C.widget_t)(this.handle), ahour));
}

func (this TTimeClock) SetMinuteImage(minute_image string) TRet {
  aminute_image := C.CString(minute_image)
  defer C.free(unsafe.Pointer(aminute_image))
  return TRet(C.time_clock_set_minute_image((*C.widget_t)(this.handle), aminute_image));
}

func (this TTimeClock) SetSecondImage(second_image string) TRet {
  asecond_image := C.CString(second_image)
  defer C.free(unsafe.Pointer(asecond_image))
  return TRet(C.time_clock_set_second_image((*C.widget_t)(this.handle), asecond_image));
}

func (this TTimeClock) SetBgImage(bg_image string) TRet {
  abg_image := C.CString(bg_image)
  defer C.free(unsafe.Pointer(abg_image))
  return TRet(C.time_clock_set_bg_image((*C.widget_t)(this.handle), abg_image));
}

func (this TTimeClock) SetImage(image string) TRet {
  aimage := C.CString(image)
  defer C.free(unsafe.Pointer(aimage))
  return TRet(C.time_clock_set_image((*C.widget_t)(this.handle), aimage));
}

func (this TTimeClock) SetHourAnchor(anchor_x string, anchor_y string) TRet {
  aanchor_x := C.CString(anchor_x)
  defer C.free(unsafe.Pointer(aanchor_x))
  aanchor_y := C.CString(anchor_y)
  defer C.free(unsafe.Pointer(aanchor_y))
  return TRet(C.time_clock_set_hour_anchor((*C.widget_t)(this.handle), aanchor_x, aanchor_y));
}

func (this TTimeClock) SetMinuteAnchor(anchor_x string, anchor_y string) TRet {
  aanchor_x := C.CString(anchor_x)
  defer C.free(unsafe.Pointer(aanchor_x))
  aanchor_y := C.CString(anchor_y)
  defer C.free(unsafe.Pointer(aanchor_y))
  return TRet(C.time_clock_set_minute_anchor((*C.widget_t)(this.handle), aanchor_x, aanchor_y));
}

func (this TTimeClock) SetSecondAnchor(anchor_x string, anchor_y string) TRet {
  aanchor_x := C.CString(anchor_x)
  defer C.free(unsafe.Pointer(aanchor_x))
  aanchor_y := C.CString(anchor_y)
  defer C.free(unsafe.Pointer(aanchor_y))
  return TRet(C.time_clock_set_second_anchor((*C.widget_t)(this.handle), aanchor_x, aanchor_y));
}

func (this TTimeClock) GetHour() int32 {
  return (int32)((*C.time_clock_t)(unsafe.Pointer(this.handle)).hour);
}

func (this TTimeClock) GetMinute() int32 {
  return (int32)((*C.time_clock_t)(unsafe.Pointer(this.handle)).minute);
}

func (this TTimeClock) GetSecond() int32 {
  return (int32)((*C.time_clock_t)(unsafe.Pointer(this.handle)).second);
}

func (this TTimeClock) GetImage() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).image);
}

func (this TTimeClock) GetBgImage() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).bg_image);
}

func (this TTimeClock) GetHourImage() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).hour_image);
}

func (this TTimeClock) GetMinuteImage() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).minute_image);
}

func (this TTimeClock) GetSecondImage() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).second_image);
}

func (this TTimeClock) GetHourAnchorX() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).hour_anchor_x);
}

func (this TTimeClock) GetHourAnchorY() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).hour_anchor_y);
}

func (this TTimeClock) GetMinuteAnchorX() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).minute_anchor_x);
}

func (this TTimeClock) GetMinuteAnchorY() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).minute_anchor_y);
}

func (this TTimeClock) GetSecondAnchorX() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).second_anchor_x);
}

func (this TTimeClock) GetSecondAnchorY() string {
  return C.GoString((*C.time_clock_t)(unsafe.Pointer(this.handle)).second_anchor_y);
}

func TTimeNowS() int64 {
  return (int64)(C.time_now_s());
}

func TTimeNowMs() int64 {
  return (int64)(C.time_now_ms());
}

func TTimeNowUs() int64 {
  return (int64)(C.time_now_us());
}

func TTimerRemove(timer_id uint32) TRet {
  return TRet(C.timer_remove((C.uint32_t)(timer_id)));
}

func TTimerRemoveAllByCtx(ctx unsafe.Pointer) TRet {
  return TRet(C.timer_remove_all_by_ctx((unsafe.Pointer)(ctx)));
}

func TTimerReset(timer_id uint32) TRet {
  return TRet(C.timer_reset((C.uint32_t)(timer_id)));
}

func TTimerSuspend(timer_id uint32) TRet {
  return TRet(C.timer_suspend((C.uint32_t)(timer_id)));
}

func TTimerResume(timer_id uint32) TRet {
  return TRet(C.timer_resume((C.uint32_t)(timer_id)));
}

func TTimerModify(timer_id uint32, duration uint32) TRet {
  return TRet(C.timer_modify((C.uint32_t)(timer_id), (C.uint32_t)(duration)));
}

type TTimerInfo struct {
  TObject
}

func TTimerInfoCast(timer TTimerInfo) TTimerInfo {
  retObj := TTimerInfo{}
  retObj.handle = unsafe.Pointer(C.timer_info_cast((*C.timer_info_t)(timer.handle)))
  return retObj
}

func (this TTimerInfo) GetCtx() unsafe.Pointer {
  return (unsafe.Pointer)((*C.timer_info_t)(unsafe.Pointer(this.handle)).ctx);
}

func (this TTimerInfo) GetExtraCtx() unsafe.Pointer {
  return (unsafe.Pointer)((*C.timer_info_t)(unsafe.Pointer(this.handle)).extra_ctx);
}

func (this TTimerInfo) GetId() uint32 {
  return (uint32)((*C.timer_info_t)(unsafe.Pointer(this.handle)).id);
}

func (this TTimerInfo) GetNow() int64 {
  return (int64)((*C.timer_info_t)(unsafe.Pointer(this.handle)).now);
}

type TTimerManager struct {
  handle unsafe.Pointer
}

type TTimerWidget struct {
  TWidget
}

func TTimerWidgetCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TTimerWidget{}
  retObj.handle = unsafe.Pointer(C.timer_widget_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TTimerWidgetCast(widget TWidget) TTimerWidget {
  retObj := TTimerWidget{}
  retObj.handle = unsafe.Pointer(C.timer_widget_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TTimerWidget) SetDuration(duration uint32) TRet {
  return TRet(C.timer_widget_set_duration((*C.widget_t)(this.handle), (C.uint32_t)(duration)));
}

func (this TTimerWidget) GetDuration() uint32 {
  return (uint32)((*C.timer_widget_t)(unsafe.Pointer(this.handle)).duration);
}

type TUiLoadEvent struct {
  TEvent
}

func TUiLoadEventCast(event TEvent) TUiLoadEvent {
  retObj := TUiLoadEvent{}
  retObj.handle = unsafe.Pointer(C.ui_load_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TUiLoadEvent) GetRoot() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer((*C.ui_load_event_t)(unsafe.Pointer(this.handle)).root)
  return retObj
}

func (this TUiLoadEvent) GetName() string {
  return C.GoString((*C.ui_load_event_t)(unsafe.Pointer(this.handle)).name);
}

type TValue struct {
  handle unsafe.Pointer
}

func (this TValue) SetBool(value bool) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_bool((*C.value_t)(this.handle), (C.bool_t)(value)))
  return retObj
}

func (this TValue) Bool() bool {
  return (bool)(C.value_bool((*C.value_t)(this.handle)));
}

func (this TValue) SetInt8(value int8) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_int8((*C.value_t)(this.handle), (C.int8_t)(value)))
  return retObj
}

func (this TValue) Int8() int8 {
  return (int8)(C.value_int8((*C.value_t)(this.handle)));
}

func (this TValue) SetUint8(value uint8) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_uint8((*C.value_t)(this.handle), (C.uint8_t)(value)))
  return retObj
}

func (this TValue) Uint8() uint8 {
  return (uint8)(C.value_uint8((*C.value_t)(this.handle)));
}

func (this TValue) SetInt16(value int16) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_int16((*C.value_t)(this.handle), (C.int16_t)(value)))
  return retObj
}

func (this TValue) Int16() int16 {
  return (int16)(C.value_int16((*C.value_t)(this.handle)));
}

func (this TValue) SetUint16(value uint16) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_uint16((*C.value_t)(this.handle), (C.uint16_t)(value)))
  return retObj
}

func (this TValue) Uint16() uint16 {
  return (uint16)(C.value_uint16((*C.value_t)(this.handle)));
}

func (this TValue) SetInt32(value int32) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_int32((*C.value_t)(this.handle), (C.int32_t)(value)))
  return retObj
}

func (this TValue) Int32() int32 {
  return (int32)(C.value_int32((*C.value_t)(this.handle)));
}

func (this TValue) SetUint32(value uint32) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_uint32((*C.value_t)(this.handle), (C.uint32_t)(value)))
  return retObj
}

func (this TValue) SetInt64(value int64) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_int64((*C.value_t)(this.handle), (C.int64_t)(value)))
  return retObj
}

func (this TValue) Int64() int64 {
  return (int64)(C.value_int64((*C.value_t)(this.handle)));
}

func (this TValue) SetUint64(value int64) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_uint64((*C.value_t)(this.handle), (C.uint64_t)(value)))
  return retObj
}

func (this TValue) Uint64() int64 {
  return (int64)(C.value_uint64((*C.value_t)(this.handle)));
}

func (this TValue) SetFloat(value float64) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_float((*C.value_t)(this.handle), (C.float_t)(value)))
  return retObj
}

func (this TValue) Float32() float64 {
  return (float64)(C.value_float32((*C.value_t)(this.handle)));
}

func (this TValue) SetFloat64(value float64) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_double((*C.value_t)(this.handle), (C.double)(value)))
  return retObj
}

func (this TValue) Float64() float64 {
  return (float64)(C.value_double((*C.value_t)(this.handle)));
}

func (this TValue) SetStr(value string) TValue {
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_dup_str((*C.value_t)(this.handle), avalue))
  return retObj
}

func (this TValue) Str() string {
  return C.GoString(C.value_str((*C.value_t)(this.handle)));
}

func (this TValue) StrEx(buff string, size uint32) string {
  abuff := C.CString(buff)
  defer C.free(unsafe.Pointer(abuff))
  return C.GoString(C.value_str_ex((*C.value_t)(this.handle), abuff, (C.uint32_t)(size)));
}

func (this TValue) IsNull() bool {
  return (bool)(C.value_is_null((*C.value_t)(this.handle)));
}

func (this TValue) Equal(other TValue) bool {
  return (bool)(C.value_equal((*C.value_t)(this.handle), (*C.value_t)(other.handle)));
}

func (this TValue) SetInt(value int32) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_int((*C.value_t)(this.handle), (C.int32_t)(value)))
  return retObj
}

func (this TValue) SetObject(value TObject) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_object((*C.value_t)(this.handle), (*C.object_t)(value.handle)))
  return retObj
}

func (this TValue) Object() TObject {
  retObj := TObject{}
  retObj.handle = unsafe.Pointer(C.value_object((*C.value_t)(this.handle)))
  return retObj
}

func (this TValue) SetToken(value uint32) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_set_token((*C.value_t)(this.handle), (C.uint32_t)(value)))
  return retObj
}

func (this TValue) Token() uint32 {
  return (uint32)(C.value_token((*C.value_t)(this.handle)));
}

func TValueCreate() TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_create())
  return retObj
}

func (this TValue) Destroy() TRet {
  return TRet(C.value_destroy((*C.value_t)(this.handle)));
}

func (this TValue) Reset() TRet {
  return TRet(C.value_reset((*C.value_t)(this.handle)));
}

func TValueCast(value TValue) TValue {
  retObj := TValue{}
  retObj.handle = unsafe.Pointer(C.value_cast((*C.value_t)(value.handle)))
  return retObj
}

func (this TValue) Id() string {
  return C.GoString(C.value_id((*C.value_t)(this.handle)));
}

func (this TValue) Func() unsafe.Pointer {
  return (unsafe.Pointer)(C.value_func((*C.value_t)(this.handle)));
}

func (this TValue) FuncDef() unsafe.Pointer {
  return (unsafe.Pointer)(C.value_func_def((*C.value_t)(this.handle)));
}

func (this TValue) Bitmap() unsafe.Pointer {
  return (unsafe.Pointer)(C.value_bitmap((*C.value_t)(this.handle)));
}

func (this TValue) Rect() TRect {
  retObj := TRect{}
  retObj.handle = unsafe.Pointer(C.value_rect((*C.value_t)(this.handle)))
  return retObj
}

type TValueChangeEvent struct {
  TEvent
}

func TValueChangeEventCast(event TEvent) TValueChangeEvent {
  retObj := TValueChangeEvent{}
  retObj.handle = unsafe.Pointer(C.value_change_event_cast((*C.event_t)(event.handle)))
  return retObj
}

type TValueType int
const (
  VALUE_TYPE_INVALID TValueType = C.VALUE_TYPE_INVALID
  VALUE_TYPE_BOOL TValueType = C.VALUE_TYPE_BOOL
  VALUE_TYPE_INT8 TValueType = C.VALUE_TYPE_INT8
  VALUE_TYPE_UINT8 TValueType = C.VALUE_TYPE_UINT8
  VALUE_TYPE_INT16 TValueType = C.VALUE_TYPE_INT16
  VALUE_TYPE_UINT16 TValueType = C.VALUE_TYPE_UINT16
  VALUE_TYPE_INT32 TValueType = C.VALUE_TYPE_INT32
  VALUE_TYPE_UINT32 TValueType = C.VALUE_TYPE_UINT32
  VALUE_TYPE_INT64 TValueType = C.VALUE_TYPE_INT64
  VALUE_TYPE_UINT64 TValueType = C.VALUE_TYPE_UINT64
  VALUE_TYPE_POINTER TValueType = C.VALUE_TYPE_POINTER
  VALUE_TYPE_FLOAT TValueType = C.VALUE_TYPE_FLOAT
  VALUE_TYPE_FLOAT32 TValueType = C.VALUE_TYPE_FLOAT32
  VALUE_TYPE_DOUBLE TValueType = C.VALUE_TYPE_DOUBLE
  VALUE_TYPE_STRING TValueType = C.VALUE_TYPE_STRING
  VALUE_TYPE_WSTRING TValueType = C.VALUE_TYPE_WSTRING
  VALUE_TYPE_OBJECT TValueType = C.VALUE_TYPE_OBJECT
  VALUE_TYPE_SIZED_STRING TValueType = C.VALUE_TYPE_SIZED_STRING
  VALUE_TYPE_BINARY TValueType = C.VALUE_TYPE_BINARY
  VALUE_TYPE_UBJSON TValueType = C.VALUE_TYPE_UBJSON
  VALUE_TYPE_TOKEN TValueType = C.VALUE_TYPE_TOKEN
  VALUE_TYPE_GRADIENT TValueType = C.VALUE_TYPE_GRADIENT
  VALUE_TYPE_ID TValueType = C.VALUE_TYPE_ID
  VALUE_TYPE_FUNC TValueType = C.VALUE_TYPE_FUNC
  VALUE_TYPE_FUNC_DEF TValueType = C.VALUE_TYPE_FUNC_DEF
  VALUE_TYPE_POINTER_REF TValueType = C.VALUE_TYPE_POINTER_REF
  VALUE_TYPE_BITMAP TValueType = C.VALUE_TYPE_BITMAP
  VALUE_TYPE_RECT TValueType = C.VALUE_TYPE_RECT
)
type TVgcanvas struct {
  handle unsafe.Pointer
}

func TVgcanvasCast(vg TVgcanvas) TVgcanvas {
  retObj := TVgcanvas{}
  retObj.handle = unsafe.Pointer(C.vgcanvas_cast((*C.vgcanvas_t)(vg.handle)))
  return retObj
}

func (this TVgcanvas) Flush() TRet {
  return TRet(C.vgcanvas_flush((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) BeginPath() TRet {
  return TRet(C.vgcanvas_begin_path((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) MoveTo(x float64, y float64) TRet {
  return TRet(C.vgcanvas_move_to((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) LineTo(x float64, y float64) TRet {
  return TRet(C.vgcanvas_line_to((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) QuadTo(cpx float64, cpy float64, x float64, y float64) TRet {
  return TRet(C.vgcanvas_quad_to((*C.vgcanvas_t)(this.handle), (C.float_t)(cpx), (C.float_t)(cpy), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) BezierTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) TRet {
  return TRet(C.vgcanvas_bezier_to((*C.vgcanvas_t)(this.handle), (C.float_t)(cp1x), (C.float_t)(cp1y), (C.float_t)(cp2x), (C.float_t)(cp2y), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, r float64) TRet {
  return TRet(C.vgcanvas_arc_to((*C.vgcanvas_t)(this.handle), (C.float_t)(x1), (C.float_t)(y1), (C.float_t)(x2), (C.float_t)(y2), (C.float_t)(r)));
}

func (this TVgcanvas) Arc(x float64, y float64, r float64, start_angle float64, end_angle float64, ccw bool) TRet {
  return TRet(C.vgcanvas_arc((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y), (C.float_t)(r), (C.float_t)(start_angle), (C.float_t)(end_angle), (C.bool_t)(ccw)));
}

func (this TVgcanvas) IsPointInPath(x float64, y float64) bool {
  return (bool)(C.vgcanvas_is_point_in_path((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) Rect(x float64, y float64, w float64, h float64) TRet {
  return TRet(C.vgcanvas_rect((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y), (C.float_t)(w), (C.float_t)(h)));
}

func (this TVgcanvas) RoundedRect(x float64, y float64, w float64, h float64, r float64) TRet {
  return TRet(C.vgcanvas_rounded_rect((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y), (C.float_t)(w), (C.float_t)(h), (C.float_t)(r)));
}

func (this TVgcanvas) Ellipse(x float64, y float64, rx float64, ry float64) TRet {
  return TRet(C.vgcanvas_ellipse((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y), (C.float_t)(rx), (C.float_t)(ry)));
}

func (this TVgcanvas) ClosePath() TRet {
  return TRet(C.vgcanvas_close_path((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) PathWinding(dir bool) TRet {
  return TRet(C.vgcanvas_path_winding((*C.vgcanvas_t)(this.handle), (C.bool_t)(dir)));
}

func (this TVgcanvas) Rotate(rad float64) TRet {
  return TRet(C.vgcanvas_rotate((*C.vgcanvas_t)(this.handle), (C.float_t)(rad)));
}

func (this TVgcanvas) Scale(x float64, y float64) TRet {
  return TRet(C.vgcanvas_scale((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) Translate(x float64, y float64) TRet {
  return TRet(C.vgcanvas_translate((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y)));
}

func (this TVgcanvas) Transform(a float64, b float64, c float64, d float64, e float64, f float64) TRet {
  return TRet(C.vgcanvas_transform((*C.vgcanvas_t)(this.handle), (C.float_t)(a), (C.float_t)(b), (C.float_t)(c), (C.float_t)(d), (C.float_t)(e), (C.float_t)(f)));
}

func (this TVgcanvas) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) TRet {
  return TRet(C.vgcanvas_set_transform((*C.vgcanvas_t)(this.handle), (C.float_t)(a), (C.float_t)(b), (C.float_t)(c), (C.float_t)(d), (C.float_t)(e), (C.float_t)(f)));
}

func (this TVgcanvas) ClipPath() TRet {
  return TRet(C.vgcanvas_clip_path((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) ClipRect(x float64, y float64, w float64, h float64) TRet {
  return TRet(C.vgcanvas_clip_rect((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y), (C.float_t)(w), (C.float_t)(h)));
}

func (this TVgcanvas) IsRectfInClipRect(left float64, top float64, right float64, bottom float64) bool {
  return (bool)(C.vgcanvas_is_rectf_in_clip_rect((*C.vgcanvas_t)(this.handle), (C.float_t)(left), (C.float_t)(top), (C.float_t)(right), (C.float_t)(bottom)));
}

func (this TVgcanvas) IntersectClipRect(x float64, y float64, w float64, h float64) TRet {
  return TRet(C.vgcanvas_intersect_clip_rect((*C.vgcanvas_t)(this.handle), (C.float_t)(x), (C.float_t)(y), (C.float_t)(w), (C.float_t)(h)));
}

func (this TVgcanvas) Fill() TRet {
  return TRet(C.vgcanvas_fill((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) Stroke() TRet {
  return TRet(C.vgcanvas_stroke((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) Paint(stroke bool, img TBitmap) TRet {
  return TRet(C.vgcanvas_paint((*C.vgcanvas_t)(this.handle), (C.bool_t)(stroke), (*C.bitmap_t)(img.handle)));
}

func (this TVgcanvas) SetFont(font string) TRet {
  afont := C.CString(font)
  defer C.free(unsafe.Pointer(afont))
  return TRet(C.vgcanvas_set_font((*C.vgcanvas_t)(this.handle), afont));
}

func (this TVgcanvas) SetFontSize(size float64) TRet {
  return TRet(C.vgcanvas_set_font_size((*C.vgcanvas_t)(this.handle), (C.float_t)(size)));
}

func (this TVgcanvas) SetTextAlign(value string) TRet {
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.vgcanvas_set_text_align((*C.vgcanvas_t)(this.handle), avalue));
}

func (this TVgcanvas) SetTextBaseline(value string) TRet {
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.vgcanvas_set_text_baseline((*C.vgcanvas_t)(this.handle), avalue));
}

func (this TVgcanvas) FillText(text string, x float64, y float64, max_width float64) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.vgcanvas_fill_text((*C.vgcanvas_t)(this.handle), atext, (C.float_t)(x), (C.float_t)(y), (C.float_t)(max_width)));
}

func (this TVgcanvas) MeasureText(text string) float64 {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return (float64)(C.vgcanvas_measure_text((*C.vgcanvas_t)(this.handle), atext));
}

func (this TVgcanvas) DrawImage(img TBitmap, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) TRet {
  return TRet(C.vgcanvas_draw_image((*C.vgcanvas_t)(this.handle), (*C.bitmap_t)(img.handle), (C.float_t)(sx), (C.float_t)(sy), (C.float_t)(sw), (C.float_t)(sh), (C.float_t)(dx), (C.float_t)(dy), (C.float_t)(dw), (C.float_t)(dh)));
}

func (this TVgcanvas) DrawImageRepeat(img TBitmap, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64, dst_w float64, dst_h float64) TRet {
  return TRet(C.vgcanvas_draw_image_repeat((*C.vgcanvas_t)(this.handle), (*C.bitmap_t)(img.handle), (C.float_t)(sx), (C.float_t)(sy), (C.float_t)(sw), (C.float_t)(sh), (C.float_t)(dx), (C.float_t)(dy), (C.float_t)(dw), (C.float_t)(dh), (C.float_t)(dst_w), (C.float_t)(dst_h)));
}

func (this TVgcanvas) DrawIcon(img TBitmap, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) TRet {
  return TRet(C.vgcanvas_draw_icon((*C.vgcanvas_t)(this.handle), (*C.bitmap_t)(img.handle), (C.float_t)(sx), (C.float_t)(sy), (C.float_t)(sw), (C.float_t)(sh), (C.float_t)(dx), (C.float_t)(dy), (C.float_t)(dw), (C.float_t)(dh)));
}

func (this TVgcanvas) SetAntialias(value bool) TRet {
  return TRet(C.vgcanvas_set_antialias((*C.vgcanvas_t)(this.handle), (C.bool_t)(value)));
}

func (this TVgcanvas) SetGlobalAlpha(alpha float64) TRet {
  return TRet(C.vgcanvas_set_global_alpha((*C.vgcanvas_t)(this.handle), (C.float_t)(alpha)));
}

func (this TVgcanvas) SetLineWidth(value float64) TRet {
  return TRet(C.vgcanvas_set_line_width((*C.vgcanvas_t)(this.handle), (C.float_t)(value)));
}

func (this TVgcanvas) SetFillColor(color string) TRet {
  acolor := C.CString(color)
  defer C.free(unsafe.Pointer(acolor))
  return TRet(C.vgcanvas_set_fill_color_str((*C.vgcanvas_t)(this.handle), acolor));
}

func (this TVgcanvas) SetStrokeColor(str string) TRet {
  astr := C.CString(str)
  defer C.free(unsafe.Pointer(astr))
  return TRet(C.vgcanvas_set_stroke_color_str((*C.vgcanvas_t)(this.handle), astr));
}

func (this TVgcanvas) SetLineCap(value string) TRet {
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.vgcanvas_set_line_cap((*C.vgcanvas_t)(this.handle), avalue));
}

func (this TVgcanvas) SetLineJoin(value string) TRet {
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.vgcanvas_set_line_join((*C.vgcanvas_t)(this.handle), avalue));
}

func (this TVgcanvas) SetMiterLimit(value float64) TRet {
  return TRet(C.vgcanvas_set_miter_limit((*C.vgcanvas_t)(this.handle), (C.float_t)(value)));
}

func (this TVgcanvas) Save() TRet {
  return TRet(C.vgcanvas_save((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) Restore() TRet {
  return TRet(C.vgcanvas_restore((*C.vgcanvas_t)(this.handle)));
}

func (this TVgcanvas) GetW() uint32 {
  return (uint32)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).w);
}

func (this TVgcanvas) GetH() uint32 {
  return (uint32)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).h);
}

func (this TVgcanvas) GetStride() uint32 {
  return (uint32)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).stride);
}

func (this TVgcanvas) GetRatio() float64 {
  return (float64)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).ratio);
}

func (this TVgcanvas) GetAntiAlias() bool {
  return (bool)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).anti_alias);
}

func (this TVgcanvas) GetLineWidth() float64 {
  return (float64)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).line_width);
}

func (this TVgcanvas) GetGlobalAlpha() float64 {
  return (float64)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).global_alpha);
}

func (this TVgcanvas) GetMiterLimit() float64 {
  return (float64)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).miter_limit);
}

func (this TVgcanvas) GetLineCap() string {
  return C.GoString((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).line_cap);
}

func (this TVgcanvas) GetLineJoin() string {
  return C.GoString((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).line_join);
}

func (this TVgcanvas) GetFont() string {
  return C.GoString((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).font);
}

func (this TVgcanvas) GetFontSize() float64 {
  return (float64)((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).font_size);
}

func (this TVgcanvas) GetTextAlign() string {
  return C.GoString((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).text_align);
}

func (this TVgcanvas) GetTextBaseline() string {
  return C.GoString((*C.vgcanvas_t)(unsafe.Pointer(this.handle)).text_baseline);
}

type TVgcanvasLineCap string
const (
  VGCANVAS_LINE_CAP_ROUND string = C.VGCANVAS_LINE_CAP_ROUND
  VGCANVAS_LINE_CAP_SQUARE string = C.VGCANVAS_LINE_CAP_SQUARE
  VGCANVAS_LINE_CAP_BUTT string = C.VGCANVAS_LINE_CAP_BUTT
)
type TVgcanvasLineJoin string
const (
  VGCANVAS_LINE_JOIN_ROUND string = C.VGCANVAS_LINE_JOIN_ROUND
  VGCANVAS_LINE_JOIN_BEVEL string = C.VGCANVAS_LINE_JOIN_BEVEL
  VGCANVAS_LINE_JOIN_MITTER string = C.VGCANVAS_LINE_JOIN_MITTER
)
type TView struct {
  TWidget
}

func TViewCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TView{}
  retObj.handle = unsafe.Pointer(C.view_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func (this TView) SetDefaultFocusedChild(default_focused_child string) TRet {
  adefault_focused_child := C.CString(default_focused_child)
  defer C.free(unsafe.Pointer(adefault_focused_child))
  return TRet(C.view_set_default_focused_child((*C.widget_t)(this.handle), adefault_focused_child));
}

func TViewCast(widget TWidget) TView {
  retObj := TView{}
  retObj.handle = unsafe.Pointer(C.view_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TView) GetDefaultFocusedChild() string {
  return C.GoString((*C.view_t)(unsafe.Pointer(this.handle)).default_focused_child);
}

type TVpage struct {
  TWidget
}

func TVpageCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TVpage{}
  retObj.handle = unsafe.Pointer(C.vpage_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TVpageCast(widget TWidget) TVpage {
  retObj := TVpage{}
  retObj.handle = unsafe.Pointer(C.vpage_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TVpage) SetUiAsset(ui_asset string) TRet {
  aui_asset := C.CString(ui_asset)
  defer C.free(unsafe.Pointer(aui_asset))
  return TRet(C.vpage_set_ui_asset((*C.widget_t)(this.handle), aui_asset));
}

func (this TVpage) SetAnimHint(anim_hint string) TRet {
  aanim_hint := C.CString(anim_hint)
  defer C.free(unsafe.Pointer(aanim_hint))
  return TRet(C.vpage_set_anim_hint((*C.widget_t)(this.handle), aanim_hint));
}

func (this TVpage) GetUiAsset() string {
  return C.GoString((*C.vpage_t)(unsafe.Pointer(this.handle)).ui_asset);
}

func (this TVpage) GetAnimHint() string {
  return C.GoString((*C.vpage_t)(unsafe.Pointer(this.handle)).anim_hint);
}

type TVpageEvent int
const (
  EVT_VPAGE_WILL_OPEN TVpageEvent = C.EVT_VPAGE_WILL_OPEN
  EVT_VPAGE_OPEN TVpageEvent = C.EVT_VPAGE_OPEN
  EVT_VPAGE_CLOSE TVpageEvent = C.EVT_VPAGE_CLOSE
)
type TWheelEvent struct {
  TEvent
}

func TWheelEventCast(event TEvent) TWheelEvent {
  retObj := TWheelEvent{}
  retObj.handle = unsafe.Pointer(C.wheel_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TWheelEvent) GetX() int {
  return (int)((*C.wheel_event_t)(unsafe.Pointer(this.handle)).x);
}

func (this TWheelEvent) GetY() int {
  return (int)((*C.wheel_event_t)(unsafe.Pointer(this.handle)).y);
}

func (this TWheelEvent) GetDy() int32 {
  return (int32)((*C.wheel_event_t)(unsafe.Pointer(this.handle)).dy);
}

func (this TWheelEvent) GetAlt() bool {
  return (bool)((*C.wheel_event_t)(unsafe.Pointer(this.handle)).alt);
}

func (this TWheelEvent) GetCtrl() bool {
  return (bool)((*C.wheel_event_t)(unsafe.Pointer(this.handle)).ctrl);
}

func (this TWheelEvent) GetShift() bool {
  return (bool)((*C.wheel_event_t)(unsafe.Pointer(this.handle)).shift);
}

type TWidget struct {
  handle unsafe.Pointer
}

func (this TWidget) CountChildren() int32 {
  return (int32)(C.widget_count_children((*C.widget_t)(this.handle)));
}

func (this TWidget) GetChild(index int32) TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_get_child((*C.widget_t)(this.handle), (C.int32_t)(index)))
  return retObj
}

func (this TWidget) FindParentByName(name string) TWidget {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_find_parent_by_name((*C.widget_t)(this.handle), aname))
  return retObj
}

func (this TWidget) FindParentByType(typex string) TWidget {
  atypex := C.CString(typex)
  defer C.free(unsafe.Pointer(atypex))
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_find_parent_by_type((*C.widget_t)(this.handle), atypex))
  return retObj
}

func (this TWidget) GetFocusedWidget() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_get_focused_widget((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWidget) GetNativeWindow() TNativeWindow {
  retObj := TNativeWindow{}
  retObj.handle = unsafe.Pointer(C.widget_get_native_window((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWidget) IndexOf() int32 {
  return (int32)(C.widget_index_of((*C.widget_t)(this.handle)));
}

func (this TWidget) CloseWindow() TRet {
  return TRet(C.widget_close_window((*C.widget_t)(this.handle)));
}

func (this TWidget) CloseWindowForce() TRet {
  return TRet(C.widget_close_window_force((*C.widget_t)(this.handle)));
}

func (this TWidget) Back() TRet {
  return TRet(C.widget_back((*C.widget_t)(this.handle)));
}

func (this TWidget) BackToHome() TRet {
  return TRet(C.widget_back_to_home((*C.widget_t)(this.handle)));
}

func (this TWidget) Move(x int, y int) TRet {
  return TRet(C.widget_move((*C.widget_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y)));
}

func (this TWidget) MoveToCenter() TRet {
  return TRet(C.widget_move_to_center((*C.widget_t)(this.handle)));
}

func (this TWidget) Resize(w int, h int) TRet {
  return TRet(C.widget_resize((*C.widget_t)(this.handle), (C.wh_t)(w), (C.wh_t)(h)));
}

func (this TWidget) MoveResize(x int, y int, w int, h int) TRet {
  return TRet(C.widget_move_resize((*C.widget_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)));
}

func (this TWidget) MoveResizeEx(x int, y int, w int, h int, update_layout bool) TRet {
  return TRet(C.widget_move_resize_ex((*C.widget_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h), (C.bool_t)(update_layout)));
}

func (this TWidget) GetValue() float64 {
  return (float64)(C.widget_get_value((*C.widget_t)(this.handle)));
}

func (this TWidget) SetValue(value float64) TRet {
  return TRet(C.widget_set_value((*C.widget_t)(this.handle), (C.float_t)(value)));
}

func (this TWidget) AddValue(delta float64) TRet {
  return TRet(C.widget_add_value((*C.widget_t)(this.handle), (C.float_t)(delta)));
}

func (this TWidget) GetValueInt() int32 {
  return (int32)(C.widget_get_value_int((*C.widget_t)(this.handle)));
}

func (this TWidget) SetValueInt(value int32) TRet {
  return TRet(C.widget_set_value_int((*C.widget_t)(this.handle), (C.int32_t)(value)));
}

func (this TWidget) AddValueInt(delta int32) TRet {
  return TRet(C.widget_add_value_int((*C.widget_t)(this.handle), (C.int32_t)(delta)));
}

func (this TWidget) AnimateValueTo(value float64, duration uint32) TRet {
  return TRet(C.widget_animate_value_to((*C.widget_t)(this.handle), (C.float_t)(value), (C.uint32_t)(duration)));
}

func (this TWidget) IsStyleExist(style_name string, state_name string) bool {
  astyle_name := C.CString(style_name)
  defer C.free(unsafe.Pointer(astyle_name))
  astate_name := C.CString(state_name)
  defer C.free(unsafe.Pointer(astate_name))
  return (bool)(C.widget_is_style_exist((*C.widget_t)(this.handle), astyle_name, astate_name));
}

func (this TWidget) IsSupportHighlighter() bool {
  return (bool)(C.widget_is_support_highlighter((*C.widget_t)(this.handle)));
}

func (this TWidget) HasHighlighter() bool {
  return (bool)(C.widget_has_highlighter((*C.widget_t)(this.handle)));
}

func (this TWidget) UseStyle(style string) TRet {
  astyle := C.CString(style)
  defer C.free(unsafe.Pointer(astyle))
  return TRet(C.widget_use_style((*C.widget_t)(this.handle), astyle));
}

func (this TWidget) SetText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.widget_set_text_utf8((*C.widget_t)(this.handle), atext));
}

func (this TWidget) SetTextEx(text string, check_diff bool) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.widget_set_text_utf8_ex((*C.widget_t)(this.handle), atext, (C.bool_t)(check_diff)));
}

func (this TWidget) SetChildText(name string, text string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.widget_set_child_text_utf8((*C.widget_t)(this.handle), aname, atext));
}

func (this TWidget) SetChildTextWithDouble(name string, format string, value float64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.widget_set_child_text_with_double((*C.widget_t)(this.handle), aname, aformat, (C.double)(value)));
}

func (this TWidget) SetChildTextWithInt(name string, format string, value int) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  aformat := C.CString(format)
  defer C.free(unsafe.Pointer(aformat))
  return TRet(C.widget_set_child_text_with_int((*C.widget_t)(this.handle), aname, aformat, (C.int)(value)));
}

func (this TWidget) SetTrText(text string) TRet {
  atext := C.CString(text)
  defer C.free(unsafe.Pointer(atext))
  return TRet(C.widget_set_tr_text((*C.widget_t)(this.handle), atext));
}

func (this TWidget) GetEnable() bool {
  return (bool)(C.widget_get_enable((*C.widget_t)(this.handle)));
}

func (this TWidget) GetFloating() bool {
  return (bool)(C.widget_get_floating((*C.widget_t)(this.handle)));
}

func (this TWidget) GetAutoAdjustSize() bool {
  return (bool)(C.widget_get_auto_adjust_size((*C.widget_t)(this.handle)));
}

func (this TWidget) GetWithFocusState() bool {
  return (bool)(C.widget_get_with_focus_state((*C.widget_t)(this.handle)));
}

func (this TWidget) GetFocusable() bool {
  return (bool)(C.widget_get_focusable((*C.widget_t)(this.handle)));
}

func (this TWidget) GetSensitive() bool {
  return (bool)(C.widget_get_sensitive((*C.widget_t)(this.handle)));
}

func (this TWidget) GetVisible() bool {
  return (bool)(C.widget_get_visible((*C.widget_t)(this.handle)));
}

func (this TWidget) GetFeedback() bool {
  return (bool)(C.widget_get_feedback((*C.widget_t)(this.handle)));
}

func (this TWidget) GetText() string {
  tstr := C.tk_utf8_dup_utf16(C.widget_get_text((*C.widget_t)(this.handle)), -1)
  ret := C.GoString(tstr)
  C.tk_free(unsafe.Pointer(tstr))
  return ret
}

func (this TWidget) SetName(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_name((*C.widget_t)(this.handle), aname));
}

func (this TWidget) SetTheme(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_theme((*C.widget_t)(this.handle), aname));
}

func (this TWidget) GetThemeName() string {
  return C.GoString(C.widget_get_theme_name((*C.widget_t)(this.handle)));
}

func (this TWidget) SetPointerCursor(cursor string) TRet {
  acursor := C.CString(cursor)
  defer C.free(unsafe.Pointer(acursor))
  return TRet(C.widget_set_pointer_cursor((*C.widget_t)(this.handle), acursor));
}

func (this TWidget) SetAnimation(animation string) TRet {
  aanimation := C.CString(animation)
  defer C.free(unsafe.Pointer(aanimation))
  return TRet(C.widget_set_animation((*C.widget_t)(this.handle), aanimation));
}

func (this TWidget) CreateAnimator(animation string) TRet {
  aanimation := C.CString(animation)
  defer C.free(unsafe.Pointer(aanimation))
  return TRet(C.widget_create_animator((*C.widget_t)(this.handle), aanimation));
}

func (this TWidget) StartAnimator(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_start_animator((*C.widget_t)(this.handle), aname));
}

func (this TWidget) SetAnimatorTimeScale(name string, time_scale float64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_animator_time_scale((*C.widget_t)(this.handle), aname, (C.float_t)(time_scale)));
}

func (this TWidget) PauseAnimator(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_pause_animator((*C.widget_t)(this.handle), aname));
}

func (this TWidget) StopAnimator(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_stop_animator((*C.widget_t)(this.handle), aname));
}

func (this TWidget) DestroyAnimator(name string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_destroy_animator((*C.widget_t)(this.handle), aname));
}

func (this TWidget) SetEnable(enable bool) TRet {
  return TRet(C.widget_set_enable((*C.widget_t)(this.handle), (C.bool_t)(enable)));
}

func (this TWidget) SetFeedback(feedback bool) TRet {
  return TRet(C.widget_set_feedback((*C.widget_t)(this.handle), (C.bool_t)(feedback)));
}

func (this TWidget) SetAutoAdjustSize(auto_adjust_size bool) TRet {
  return TRet(C.widget_set_auto_adjust_size((*C.widget_t)(this.handle), (C.bool_t)(auto_adjust_size)));
}

func (this TWidget) SetFloating(floating bool) TRet {
  return TRet(C.widget_set_floating((*C.widget_t)(this.handle), (C.bool_t)(floating)));
}

func (this TWidget) SetFocused(focused bool) TRet {
  return TRet(C.widget_set_focused((*C.widget_t)(this.handle), (C.bool_t)(focused)));
}

func (this TWidget) SetFocusable(focusable bool) TRet {
  return TRet(C.widget_set_focusable((*C.widget_t)(this.handle), (C.bool_t)(focusable)));
}

func (this TWidget) SetState(state string) TRet {
  astate := C.CString(state)
  defer C.free(unsafe.Pointer(astate))
  return TRet(C.widget_set_state((*C.widget_t)(this.handle), astate));
}

func (this TWidget) SetOpacity(opacity uint8) TRet {
  return TRet(C.widget_set_opacity((*C.widget_t)(this.handle), (C.uint8_t)(opacity)));
}

func (this TWidget) SetDirtyRectTolerance(dirty_rect_tolerance uint16) TRet {
  return TRet(C.widget_set_dirty_rect_tolerance((*C.widget_t)(this.handle), (C.uint16_t)(dirty_rect_tolerance)));
}

func (this TWidget) DestroyChildren() TRet {
  return TRet(C.widget_destroy_children((*C.widget_t)(this.handle)));
}

func (this TWidget) AddChild(child TWidget) TRet {
  return TRet(C.widget_add_child((*C.widget_t)(this.handle), (*C.widget_t)(child.handle)));
}

func (this TWidget) RemoveChild(child TWidget) TRet {
  return TRet(C.widget_remove_child((*C.widget_t)(this.handle), (*C.widget_t)(child.handle)));
}

func (this TWidget) InsertChild(index uint32, child TWidget) TRet {
  return TRet(C.widget_insert_child((*C.widget_t)(this.handle), (C.uint32_t)(index), (*C.widget_t)(child.handle)));
}

func (this TWidget) Restack(index uint32) TRet {
  return TRet(C.widget_restack((*C.widget_t)(this.handle), (C.uint32_t)(index)));
}

func (this TWidget) Child(name string) TWidget {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_child((*C.widget_t)(this.handle), aname))
  return retObj
}

func (this TWidget) Lookup(name string, recursive bool) TWidget {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_lookup((*C.widget_t)(this.handle), aname, (C.bool_t)(recursive)))
  return retObj
}

func (this TWidget) LookupByType(typex string, recursive bool) TWidget {
  atypex := C.CString(typex)
  defer C.free(unsafe.Pointer(atypex))
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_lookup_by_type((*C.widget_t)(this.handle), atypex, (C.bool_t)(recursive)))
  return retObj
}

func (this TWidget) SetVisible(visible bool) TRet {
  return TRet(C.widget_set_visible((*C.widget_t)(this.handle), (C.bool_t)(visible)));
}

func (this TWidget) SetVisibleOnly(visible bool) TRet {
  return TRet(C.widget_set_visible_only((*C.widget_t)(this.handle), (C.bool_t)(visible)));
}

func (this TWidget) SetSensitive(sensitive bool) TRet {
  return TRet(C.widget_set_sensitive((*C.widget_t)(this.handle), (C.bool_t)(sensitive)));
}

func (this TWidget) Off(id uint32) TRet {
  return TRet(C.widget_off((*C.widget_t)(this.handle), (C.uint32_t)(id)));
}

func (this TWidget) InvalidateForce(r TRect) TRet {
  return TRet(C.widget_invalidate_force((*C.widget_t)(this.handle), (*C.rect_t)(r.handle)));
}

func (this TWidget) GetProp(name string, v TValue) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_get_prop((*C.widget_t)(this.handle), aname, (*C.value_t)(v.handle)));
}

func (this TWidget) SetProp(name string, v TValue) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_prop((*C.widget_t)(this.handle), aname, (*C.value_t)(v.handle)));
}

func (this TWidget) SetProps(params string) TRet {
  aparams := C.CString(params)
  defer C.free(unsafe.Pointer(aparams))
  return TRet(C.widget_set_props((*C.widget_t)(this.handle), aparams));
}

func (this TWidget) SetPropStr(name string, v string) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  av := C.CString(v)
  defer C.free(unsafe.Pointer(av))
  return TRet(C.widget_set_prop_str((*C.widget_t)(this.handle), aname, av));
}

func (this TWidget) GetPropStr(name string, defval string) string {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  adefval := C.CString(defval)
  defer C.free(unsafe.Pointer(adefval))
  return C.GoString(C.widget_get_prop_str((*C.widget_t)(this.handle), aname, adefval));
}

func (this TWidget) SetPropPointer(name string, v unsafe.Pointer) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_prop_pointer((*C.widget_t)(this.handle), aname, (unsafe.Pointer)(v)));
}

func (this TWidget) GetPropPointer(name string) unsafe.Pointer {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (unsafe.Pointer)(C.widget_get_prop_pointer((*C.widget_t)(this.handle), aname));
}

func (this TWidget) SetPropFloat(name string, v float64) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_prop_float((*C.widget_t)(this.handle), aname, (C.float_t)(v)));
}

func (this TWidget) GetPropFloat(name string, defval float64) float64 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (float64)(C.widget_get_prop_float((*C.widget_t)(this.handle), aname, (C.float_t)(defval)));
}

func (this TWidget) SetPropInt(name string, v int32) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_prop_int((*C.widget_t)(this.handle), aname, (C.int32_t)(v)));
}

func (this TWidget) GetPropInt(name string, defval int32) int32 {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (int32)(C.widget_get_prop_int((*C.widget_t)(this.handle), aname, (C.int32_t)(defval)));
}

func (this TWidget) SetPropBool(name string, v bool) TRet {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return TRet(C.widget_set_prop_bool((*C.widget_t)(this.handle), aname, (C.bool_t)(v)));
}

func (this TWidget) GetPropBool(name string, defval bool) bool {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  return (bool)(C.widget_get_prop_bool((*C.widget_t)(this.handle), aname, (C.bool_t)(defval)));
}

func (this TWidget) IsWindowOpened() bool {
  return (bool)(C.widget_is_window_opened((*C.widget_t)(this.handle)));
}

func (this TWidget) IsWindowCreated() bool {
  return (bool)(C.widget_is_window_created((*C.widget_t)(this.handle)));
}

func (this TWidget) IsParentOf(child TWidget) bool {
  return (bool)(C.widget_is_parent_of((*C.widget_t)(this.handle), (*C.widget_t)(child.handle)));
}

func (this TWidget) IsDirectParentOf(child TWidget) bool {
  return (bool)(C.widget_is_direct_parent_of((*C.widget_t)(this.handle), (*C.widget_t)(child.handle)));
}

func (this TWidget) IsWindow() bool {
  return (bool)(C.widget_is_window((*C.widget_t)(this.handle)));
}

func (this TWidget) IsSystemBar() bool {
  return (bool)(C.widget_is_system_bar((*C.widget_t)(this.handle)));
}

func (this TWidget) IsNormalWindow() bool {
  return (bool)(C.widget_is_normal_window((*C.widget_t)(this.handle)));
}

func (this TWidget) IsFullscreenWindow() bool {
  return (bool)(C.widget_is_fullscreen_window((*C.widget_t)(this.handle)));
}

func (this TWidget) IsDialog() bool {
  return (bool)(C.widget_is_dialog((*C.widget_t)(this.handle)));
}

func (this TWidget) IsPopup() bool {
  return (bool)(C.widget_is_popup((*C.widget_t)(this.handle)));
}

func (this TWidget) IsOverlay() bool {
  return (bool)(C.widget_is_overlay((*C.widget_t)(this.handle)));
}

func (this TWidget) IsAlwaysOnTop() bool {
  return (bool)(C.widget_is_always_on_top((*C.widget_t)(this.handle)));
}

func (this TWidget) IsOpenedDialog() bool {
  return (bool)(C.widget_is_opened_dialog((*C.widget_t)(this.handle)));
}

func (this TWidget) IsOpenedPopup() bool {
  return (bool)(C.widget_is_opened_popup((*C.widget_t)(this.handle)));
}

func (this TWidget) IsKeyboard() bool {
  return (bool)(C.widget_is_keyboard((*C.widget_t)(this.handle)));
}

func (this TWidget) IsDesigningWindow() bool {
  return (bool)(C.widget_is_designing_window((*C.widget_t)(this.handle)));
}

func (this TWidget) IsWindowManager() bool {
  return (bool)(C.widget_is_window_manager((*C.widget_t)(this.handle)));
}

func (this TWidget) GetWindow() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_get_window((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWidget) GetWindowManager() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_get_window_manager((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWidget) GetType() string {
  return C.GoString(C.widget_get_type((*C.widget_t)(this.handle)));
}

func (this TWidget) Clone(parent TWidget) TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_clone((*C.widget_t)(this.handle), (*C.widget_t)(parent.handle)))
  return retObj
}

func (this TWidget) Equal(other TWidget) bool {
  return (bool)(C.widget_equal((*C.widget_t)(this.handle), (*C.widget_t)(other.handle)));
}

func TWidgetCast(widget TWidget) TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.widget_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TWidget) Destroy() TRet {
  return TRet(C.widget_destroy((*C.widget_t)(this.handle)));
}

func (this TWidget) DestroyAsync() TRet {
  return TRet(C.widget_destroy_async((*C.widget_t)(this.handle)));
}

func (this TWidget) Unref() TRet {
  return TRet(C.widget_unref((*C.widget_t)(this.handle)));
}

func (this TWidget) StrokeBorderRect(c TCanvas, r TRect) TRet {
  return TRet(C.widget_stroke_border_rect((*C.widget_t)(this.handle), (*C.canvas_t)(c.handle), (*C.rect_t)(r.handle)));
}

func (this TWidget) FillBgRect(c TCanvas, r TRect, draw_type TImageDrawType) TRet {
  return TRet(C.widget_fill_bg_rect((*C.widget_t)(this.handle), (*C.canvas_t)(c.handle), (*C.rect_t)(r.handle), (C.image_draw_type_t)(draw_type)));
}

func (this TWidget) FillFgRect(c TCanvas, r TRect, draw_type TImageDrawType) TRet {
  return TRet(C.widget_fill_fg_rect((*C.widget_t)(this.handle), (*C.canvas_t)(c.handle), (*C.rect_t)(r.handle), (C.image_draw_type_t)(draw_type)));
}

func (this TWidget) DispatchToTarget(e TEvent) TRet {
  return TRet(C.widget_dispatch_to_target((*C.widget_t)(this.handle), (*C.event_t)(e.handle)));
}

func (this TWidget) DispatchToKeyTarget(e TEvent) TRet {
  return TRet(C.widget_dispatch_to_key_target((*C.widget_t)(this.handle), (*C.event_t)(e.handle)));
}

func (this TWidget) GetStyleType() string {
  return C.GoString(C.widget_get_style_type((*C.widget_t)(this.handle)));
}

func (this TWidget) UpdateStyle() TRet {
  return TRet(C.widget_update_style((*C.widget_t)(this.handle)));
}

func (this TWidget) UpdateStyleRecursive() TRet {
  return TRet(C.widget_update_style_recursive((*C.widget_t)(this.handle)));
}

func (this TWidget) SetAsKeyTarget() TRet {
  return TRet(C.widget_set_as_key_target((*C.widget_t)(this.handle)));
}

func (this TWidget) FocusNext() TRet {
  return TRet(C.widget_focus_next((*C.widget_t)(this.handle)));
}

func (this TWidget) FocusPrev() TRet {
  return TRet(C.widget_focus_prev((*C.widget_t)(this.handle)));
}

func (this TWidget) GetStateForStyle(active bool, checked bool) string {
  return C.GoString(C.widget_get_state_for_style((*C.widget_t)(this.handle), (C.bool_t)(active), (C.bool_t)(checked)));
}

func (this TWidget) Layout() TRet {
  return TRet(C.widget_layout((*C.widget_t)(this.handle)));
}

func (this TWidget) SetSelfLayout(params string) TRet {
  aparams := C.CString(params)
  defer C.free(unsafe.Pointer(aparams))
  return TRet(C.widget_set_self_layout((*C.widget_t)(this.handle), aparams));
}

func (this TWidget) SetChildrenLayout(params string) TRet {
  aparams := C.CString(params)
  defer C.free(unsafe.Pointer(aparams))
  return TRet(C.widget_set_children_layout((*C.widget_t)(this.handle), aparams));
}

func (this TWidget) SetSelfLayoutParams(x string, y string, w string, h string) TRet {
  ax := C.CString(x)
  defer C.free(unsafe.Pointer(ax))
  ay := C.CString(y)
  defer C.free(unsafe.Pointer(ay))
  aw := C.CString(w)
  defer C.free(unsafe.Pointer(aw))
  ah := C.CString(h)
  defer C.free(unsafe.Pointer(ah))
  return TRet(C.widget_set_self_layout_params((*C.widget_t)(this.handle), ax, ay, aw, ah));
}

func (this TWidget) SetStyleInt(state_and_name string, value int32) TRet {
  astate_and_name := C.CString(state_and_name)
  defer C.free(unsafe.Pointer(astate_and_name))
  return TRet(C.widget_set_style_int((*C.widget_t)(this.handle), astate_and_name, (C.int32_t)(value)));
}

func (this TWidget) SetStyleStr(state_and_name string, value string) TRet {
  astate_and_name := C.CString(state_and_name)
  defer C.free(unsafe.Pointer(astate_and_name))
  avalue := C.CString(value)
  defer C.free(unsafe.Pointer(avalue))
  return TRet(C.widget_set_style_str((*C.widget_t)(this.handle), astate_and_name, avalue));
}

func (this TWidget) SetStyleColor(state_and_name string, value uint32) TRet {
  astate_and_name := C.CString(state_and_name)
  defer C.free(unsafe.Pointer(astate_and_name))
  return TRet(C.widget_set_style_color((*C.widget_t)(this.handle), astate_and_name, (C.uint32_t)(value)));
}

func (this TWidget) AddChildDefault(child TWidget) TRet {
  return TRet(C.widget_add_child_default((*C.widget_t)(this.handle), (*C.widget_t)(child.handle)));
}

func (this TWidget) GetX() int {
  return (int)((*C.widget_t)(unsafe.Pointer(this.handle)).x);
}

func (this TWidget) GetY() int {
  return (int)((*C.widget_t)(unsafe.Pointer(this.handle)).y);
}

func (this TWidget) GetW() int {
  return (int)((*C.widget_t)(unsafe.Pointer(this.handle)).w);
}

func (this TWidget) GetH() int {
  return (int)((*C.widget_t)(unsafe.Pointer(this.handle)).h);
}

func (this TWidget) GetName() string {
  return C.GoString((*C.widget_t)(unsafe.Pointer(this.handle)).name);
}

func (this TWidget) GetPointerCursor() string {
  return C.GoString((*C.widget_t)(unsafe.Pointer(this.handle)).pointer_cursor);
}

func (this TWidget) GetTrText() string {
  return C.GoString((*C.widget_t)(unsafe.Pointer(this.handle)).tr_text);
}

func (this TWidget) GetStyle() string {
  return C.GoString((*C.widget_t)(unsafe.Pointer(this.handle)).style);
}

func (this TWidget) GetAnimation() string {
  return C.GoString((*C.widget_t)(unsafe.Pointer(this.handle)).animation);
}

func (this TWidget) GetOpacity() uint8 {
  return (uint8)((*C.widget_t)(unsafe.Pointer(this.handle)).opacity);
}

func (this TWidget) GetDirtyRectTolerance() uint16 {
  return (uint16)((*C.widget_t)(unsafe.Pointer(this.handle)).dirty_rect_tolerance);
}

func (this TWidget) GetParent() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer((*C.widget_t)(unsafe.Pointer(this.handle)).parent)
  return retObj
}

type TWidgetAnimatorEvent struct {
  TEvent
}

func TWidgetAnimatorEventCast(event TEvent) TWidgetAnimatorEvent {
  retObj := TWidgetAnimatorEvent{}
  retObj.handle = unsafe.Pointer(C.widget_animator_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TWidgetAnimatorEvent) GetWidget() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer((*C.widget_animator_event_t)(unsafe.Pointer(this.handle)).widget)
  return retObj
}

func (this TWidgetAnimatorEvent) GetAnimator() unsafe.Pointer {
  return (unsafe.Pointer)((*C.widget_animator_event_t)(unsafe.Pointer(this.handle)).animator);
}

type TWidgetCursor string
const (
  WIDGET_CURSOR_DEFAULT string = C.WIDGET_CURSOR_DEFAULT
  WIDGET_CURSOR_EDIT string = C.WIDGET_CURSOR_EDIT
  WIDGET_CURSOR_HAND string = C.WIDGET_CURSOR_HAND
  WIDGET_CURSOR_WAIT string = C.WIDGET_CURSOR_WAIT
  WIDGET_CURSOR_CROSS string = C.WIDGET_CURSOR_CROSS
  WIDGET_CURSOR_NO string = C.WIDGET_CURSOR_NO
  WIDGET_CURSOR_SIZENWSE string = C.WIDGET_CURSOR_SIZENWSE
  WIDGET_CURSOR_SIZENESW string = C.WIDGET_CURSOR_SIZENESW
  WIDGET_CURSOR_SIZEWE string = C.WIDGET_CURSOR_SIZEWE
  WIDGET_CURSOR_SIZENS string = C.WIDGET_CURSOR_SIZENS
  WIDGET_CURSOR_SIZEALL string = C.WIDGET_CURSOR_SIZEALL
)
type TWidgetProp string
const (
  WIDGET_PROP_EXEC string = C.WIDGET_PROP_EXEC
  WIDGET_PROP_X string = C.WIDGET_PROP_X
  WIDGET_PROP_Y string = C.WIDGET_PROP_Y
  WIDGET_PROP_W string = C.WIDGET_PROP_W
  WIDGET_PROP_H string = C.WIDGET_PROP_H
  WIDGET_PROP_MAX_H string = C.WIDGET_PROP_MAX_H
  WIDGET_PROP_DESIGN_W string = C.WIDGET_PROP_DESIGN_W
  WIDGET_PROP_DESIGN_H string = C.WIDGET_PROP_DESIGN_H
  WIDGET_PROP_AUTO_SCALE_CHILDREN_X string = C.WIDGET_PROP_AUTO_SCALE_CHILDREN_X
  WIDGET_PROP_AUTO_SCALE_CHILDREN_Y string = C.WIDGET_PROP_AUTO_SCALE_CHILDREN_Y
  WIDGET_PROP_AUTO_SCALE_CHILDREN_W string = C.WIDGET_PROP_AUTO_SCALE_CHILDREN_W
  WIDGET_PROP_AUTO_SCALE_CHILDREN_H string = C.WIDGET_PROP_AUTO_SCALE_CHILDREN_H
  WIDGET_PROP_INPUTING string = C.WIDGET_PROP_INPUTING
  WIDGET_PROP_ALWAYS_ON_TOP string = C.WIDGET_PROP_ALWAYS_ON_TOP
  WIDGET_PROP_CARET_X string = C.WIDGET_PROP_CARET_X
  WIDGET_PROP_CARET_Y string = C.WIDGET_PROP_CARET_Y
  WIDGET_PROP_LINE_HEIGHT string = C.WIDGET_PROP_LINE_HEIGHT
  WIDGET_PROP_DIRTY_RECT_TOLERANCE string = C.WIDGET_PROP_DIRTY_RECT_TOLERANCE
  WIDGET_PROP_BIDI string = C.WIDGET_PROP_BIDI
  WIDGET_PROP_CANVAS string = C.WIDGET_PROP_CANVAS
  WIDGET_PROP_LOCALIZE_OPTIONS string = C.WIDGET_PROP_LOCALIZE_OPTIONS
  WIDGET_PROP_NATIVE_WINDOW string = C.WIDGET_PROP_NATIVE_WINDOW
  WIDGET_PROP_HIGHLIGHT string = C.WIDGET_PROP_HIGHLIGHT
  WIDGET_PROP_BAR_SIZE string = C.WIDGET_PROP_BAR_SIZE
  WIDGET_PROP_OPACITY string = C.WIDGET_PROP_OPACITY
  WIDGET_PROP_MIN_W string = C.WIDGET_PROP_MIN_W
  WIDGET_PROP_MAX_W string = C.WIDGET_PROP_MAX_W
  WIDGET_PROP_AUTO_ADJUST_SIZE string = C.WIDGET_PROP_AUTO_ADJUST_SIZE
  WIDGET_PROP_SINGLE_INSTANCE string = C.WIDGET_PROP_SINGLE_INSTANCE
  WIDGET_PROP_STRONGLY_FOCUS string = C.WIDGET_PROP_STRONGLY_FOCUS
  WIDGET_PROP_CHILDREN_LAYOUT string = C.WIDGET_PROP_CHILDREN_LAYOUT
  WIDGET_PROP_LAYOUT string = C.WIDGET_PROP_LAYOUT
  WIDGET_PROP_SELF_LAYOUT string = C.WIDGET_PROP_SELF_LAYOUT
  WIDGET_PROP_LAYOUT_W string = C.WIDGET_PROP_LAYOUT_W
  WIDGET_PROP_LAYOUT_H string = C.WIDGET_PROP_LAYOUT_H
  WIDGET_PROP_VIRTUAL_W string = C.WIDGET_PROP_VIRTUAL_W
  WIDGET_PROP_VIRTUAL_H string = C.WIDGET_PROP_VIRTUAL_H
  WIDGET_PROP_LOADING string = C.WIDGET_PROP_LOADING
  WIDGET_PROP_NAME string = C.WIDGET_PROP_NAME
  WIDGET_PROP_TYPE string = C.WIDGET_PROP_TYPE
  WIDGET_PROP_CLOSABLE string = C.WIDGET_PROP_CLOSABLE
  WIDGET_PROP_POINTER_CURSOR string = C.WIDGET_PROP_POINTER_CURSOR
  WIDGET_PROP_VALUE string = C.WIDGET_PROP_VALUE
  WIDGET_PROP_EASY_TOUCH_MODE string = C.WIDGET_PROP_EASY_TOUCH_MODE
  WIDGET_PROP_RADIO string = C.WIDGET_PROP_RADIO
  WIDGET_PROP_REVERSE string = C.WIDGET_PROP_REVERSE
  WIDGET_PROP_LENGTH string = C.WIDGET_PROP_LENGTH
  WIDGET_PROP_LINE_WRAP string = C.WIDGET_PROP_LINE_WRAP
  WIDGET_PROP_WORD_WRAP string = C.WIDGET_PROP_WORD_WRAP
  WIDGET_PROP_ELLIPSES string = C.WIDGET_PROP_ELLIPSES
  WIDGET_PROP_VISIBLE_REVEAL_IN_SCROLL string = C.WIDGET_PROP_VISIBLE_REVEAL_IN_SCROLL
  WIDGET_PROP_TEXT string = C.WIDGET_PROP_TEXT
  WIDGET_PROP_TR_TEXT string = C.WIDGET_PROP_TR_TEXT
  WIDGET_PROP_STYLE string = C.WIDGET_PROP_STYLE
  WIDGET_PROP_STATE string = C.WIDGET_PROP_STATE
  WIDGET_PROP_ENABLE string = C.WIDGET_PROP_ENABLE
  WIDGET_PROP_FEEDBACK string = C.WIDGET_PROP_FEEDBACK
  WIDGET_PROP_FLOATING string = C.WIDGET_PROP_FLOATING
  WIDGET_PROP_MARGIN string = C.WIDGET_PROP_MARGIN
  WIDGET_PROP_SPACING string = C.WIDGET_PROP_SPACING
  WIDGET_PROP_LEFT_MARGIN string = C.WIDGET_PROP_LEFT_MARGIN
  WIDGET_PROP_RIGHT_MARGIN string = C.WIDGET_PROP_RIGHT_MARGIN
  WIDGET_PROP_TOP_MARGIN string = C.WIDGET_PROP_TOP_MARGIN
  WIDGET_PROP_BOTTOM_MARGIN string = C.WIDGET_PROP_BOTTOM_MARGIN
  WIDGET_PROP_STEP string = C.WIDGET_PROP_STEP
  WIDGET_PROP_VISIBLE string = C.WIDGET_PROP_VISIBLE
  WIDGET_PROP_SENSITIVE string = C.WIDGET_PROP_SENSITIVE
  WIDGET_PROP_APPLET_NAME string = C.WIDGET_PROP_APPLET_NAME
  WIDGET_PROP_ANIMATION string = C.WIDGET_PROP_ANIMATION
  WIDGET_PROP_ANIM_HINT string = C.WIDGET_PROP_ANIM_HINT
  WIDGET_PROP_FULLSCREEN string = C.WIDGET_PROP_FULLSCREEN
  WIDGET_PROP_DISABLE_ANIM string = C.WIDGET_PROP_DISABLE_ANIM
  WIDGET_PROP_OPEN_ANIM_HINT string = C.WIDGET_PROP_OPEN_ANIM_HINT
  WIDGET_PROP_CLOSE_ANIM_HINT string = C.WIDGET_PROP_CLOSE_ANIM_HINT
  WIDGET_PROP_MIN string = C.WIDGET_PROP_MIN
  WIDGET_PROP_ACTION_TEXT string = C.WIDGET_PROP_ACTION_TEXT
  WIDGET_PROP_TIPS string = C.WIDGET_PROP_TIPS
  WIDGET_PROP_TR_TIPS string = C.WIDGET_PROP_TR_TIPS
  WIDGET_PROP_INPUT_TYPE string = C.WIDGET_PROP_INPUT_TYPE
  WIDGET_PROP_KEYBOARD string = C.WIDGET_PROP_KEYBOARD
  WIDGET_PROP_DEFAULT_FOCUSED_CHILD string = C.WIDGET_PROP_DEFAULT_FOCUSED_CHILD
  WIDGET_PROP_READONLY string = C.WIDGET_PROP_READONLY
  WIDGET_PROP_CANCELABLE string = C.WIDGET_PROP_CANCELABLE
  WIDGET_PROP_PASSWORD_VISIBLE string = C.WIDGET_PROP_PASSWORD_VISIBLE
  WIDGET_PROP_ACTIVE string = C.WIDGET_PROP_ACTIVE
  WIDGET_PROP_CURR_PAGE string = C.WIDGET_PROP_CURR_PAGE
  WIDGET_PROP_PAGE_MAX_NUMBER string = C.WIDGET_PROP_PAGE_MAX_NUMBER
  WIDGET_PROP_VERTICAL string = C.WIDGET_PROP_VERTICAL
  WIDGET_PROP_SHOW_TEXT string = C.WIDGET_PROP_SHOW_TEXT
  WIDGET_PROP_XOFFSET string = C.WIDGET_PROP_XOFFSET
  WIDGET_PROP_YOFFSET string = C.WIDGET_PROP_YOFFSET
  WIDGET_PROP_ALIGN_V string = C.WIDGET_PROP_ALIGN_V
  WIDGET_PROP_ALIGN_H string = C.WIDGET_PROP_ALIGN_H
  WIDGET_PROP_AUTO_PLAY string = C.WIDGET_PROP_AUTO_PLAY
  WIDGET_PROP_LOOP string = C.WIDGET_PROP_LOOP
  WIDGET_PROP_RUNNING string = C.WIDGET_PROP_RUNNING
  WIDGET_PROP_AUTO_FIX string = C.WIDGET_PROP_AUTO_FIX
  WIDGET_PROP_SELECT_NONE_WHEN_FOCUSED string = C.WIDGET_PROP_SELECT_NONE_WHEN_FOCUSED
  WIDGET_PROP_OPEN_IM_WHEN_FOCUSED string = C.WIDGET_PROP_OPEN_IM_WHEN_FOCUSED
  WIDGET_PROP_CLOSE_IM_WHEN_BLURED string = C.WIDGET_PROP_CLOSE_IM_WHEN_BLURED
  WIDGET_PROP_X_MIN string = C.WIDGET_PROP_X_MIN
  WIDGET_PROP_X_MAX string = C.WIDGET_PROP_X_MAX
  WIDGET_PROP_Y_MIN string = C.WIDGET_PROP_Y_MIN
  WIDGET_PROP_Y_MAX string = C.WIDGET_PROP_Y_MAX
  WIDGET_PROP_MAX string = C.WIDGET_PROP_MAX
  WIDGET_PROP_GRAB_KEYS string = C.WIDGET_PROP_GRAB_KEYS
  WIDGET_PROP_ROW string = C.WIDGET_PROP_ROW
  WIDGET_PROP_STATE_FOR_STYLE string = C.WIDGET_PROP_STATE_FOR_STYLE
  WIDGET_PROP_THEME string = C.WIDGET_PROP_THEME
  WIDGET_PROP_STAGE string = C.WIDGET_PROP_STAGE
  WIDGET_PROP_IMAGE_MANAGER string = C.WIDGET_PROP_IMAGE_MANAGER
  WIDGET_PROP_ASSETS_MANAGER string = C.WIDGET_PROP_ASSETS_MANAGER
  WIDGET_PROP_LOCALE_INFO string = C.WIDGET_PROP_LOCALE_INFO
  WIDGET_PROP_FONT_MANAGER string = C.WIDGET_PROP_FONT_MANAGER
  WIDGET_PROP_THEME_OBJ string = C.WIDGET_PROP_THEME_OBJ
  WIDGET_PROP_DEFAULT_THEME_OBJ string = C.WIDGET_PROP_DEFAULT_THEME_OBJ
  WIDGET_PROP_ITEM_WIDTH string = C.WIDGET_PROP_ITEM_WIDTH
  WIDGET_PROP_ITEM_HEIGHT string = C.WIDGET_PROP_ITEM_HEIGHT
  WIDGET_PROP_DEFAULT_ITEM_HEIGHT string = C.WIDGET_PROP_DEFAULT_ITEM_HEIGHT
  WIDGET_PROP_XSLIDABLE string = C.WIDGET_PROP_XSLIDABLE
  WIDGET_PROP_YSLIDABLE string = C.WIDGET_PROP_YSLIDABLE
  WIDGET_PROP_REPEAT string = C.WIDGET_PROP_REPEAT
  WIDGET_PROP_LONG_PRESS_TIME string = C.WIDGET_PROP_LONG_PRESS_TIME
  WIDGET_PROP_ENABLE_LONG_PRESS string = C.WIDGET_PROP_ENABLE_LONG_PRESS
  WIDGET_PROP_ENABLE_PREVIEW string = C.WIDGET_PROP_ENABLE_PREVIEW
  WIDGET_PROP_CLICK_THROUGH string = C.WIDGET_PROP_CLICK_THROUGH
  WIDGET_PROP_ANIMATABLE string = C.WIDGET_PROP_ANIMATABLE
  WIDGET_PROP_AUTO_HIDE string = C.WIDGET_PROP_AUTO_HIDE
  WIDGET_PROP_AUTO_HIDE_SCROLL_BAR string = C.WIDGET_PROP_AUTO_HIDE_SCROLL_BAR
  WIDGET_PROP_IMAGE string = C.WIDGET_PROP_IMAGE
  WIDGET_PROP_FORMAT string = C.WIDGET_PROP_FORMAT
  WIDGET_PROP_DRAW_TYPE string = C.WIDGET_PROP_DRAW_TYPE
  WIDGET_PROP_SELECTABLE string = C.WIDGET_PROP_SELECTABLE
  WIDGET_PROP_CLICKABLE string = C.WIDGET_PROP_CLICKABLE
  WIDGET_PROP_SCALE_X string = C.WIDGET_PROP_SCALE_X
  WIDGET_PROP_SCALE_Y string = C.WIDGET_PROP_SCALE_Y
  WIDGET_PROP_ANCHOR_X string = C.WIDGET_PROP_ANCHOR_X
  WIDGET_PROP_ANCHOR_Y string = C.WIDGET_PROP_ANCHOR_Y
  WIDGET_PROP_ROTATION string = C.WIDGET_PROP_ROTATION
  WIDGET_PROP_COMPACT string = C.WIDGET_PROP_COMPACT
  WIDGET_PROP_SCROLLABLE string = C.WIDGET_PROP_SCROLLABLE
  WIDGET_PROP_ICON string = C.WIDGET_PROP_ICON
  WIDGET_PROP_OPTIONS string = C.WIDGET_PROP_OPTIONS
  WIDGET_PROP_SELECTED string = C.WIDGET_PROP_SELECTED
  WIDGET_PROP_CHECKED string = C.WIDGET_PROP_CHECKED
  WIDGET_PROP_ACTIVE_ICON string = C.WIDGET_PROP_ACTIVE_ICON
  WIDGET_PROP_LOAD_UI string = C.WIDGET_PROP_LOAD_UI
  WIDGET_PROP_OPEN_WINDOW string = C.WIDGET_PROP_OPEN_WINDOW
  WIDGET_PROP_THEME_OF_POPUP string = C.WIDGET_PROP_THEME_OF_POPUP
  WIDGET_PROP_SELECTED_INDEX string = C.WIDGET_PROP_SELECTED_INDEX
  WIDGET_PROP_CLOSE_WHEN_CLICK string = C.WIDGET_PROP_CLOSE_WHEN_CLICK
  WIDGET_PROP_CLOSE_WHEN_CLICK_OUTSIDE string = C.WIDGET_PROP_CLOSE_WHEN_CLICK_OUTSIDE
  WIDGET_PROP_CLOSE_WHEN_TIMEOUT string = C.WIDGET_PROP_CLOSE_WHEN_TIMEOUT
  WIDGET_PROP_LINE_GAP string = C.WIDGET_PROP_LINE_GAP
  WIDGET_PROP_BG_COLOR string = C.WIDGET_PROP_BG_COLOR
  WIDGET_PROP_BORDER_COLOR string = C.WIDGET_PROP_BORDER_COLOR
  WIDGET_PROP_DELAY string = C.WIDGET_PROP_DELAY
  WIDGET_PROP_IS_KEYBOARD string = C.WIDGET_PROP_IS_KEYBOARD
  WIDGET_PROP_FOCUSED string = C.WIDGET_PROP_FOCUSED
  WIDGET_PROP_FOCUS string = C.WIDGET_PROP_FOCUS
  WIDGET_PROP_FOCUSABLE string = C.WIDGET_PROP_FOCUSABLE
  WIDGET_PROP_WITH_FOCUS_STATE string = C.WIDGET_PROP_WITH_FOCUS_STATE
  WIDGET_PROP_MOVE_FOCUS_PREV_KEY string = C.WIDGET_PROP_MOVE_FOCUS_PREV_KEY
  WIDGET_PROP_MOVE_FOCUS_NEXT_KEY string = C.WIDGET_PROP_MOVE_FOCUS_NEXT_KEY
  WIDGET_PROP_MOVE_FOCUS_UP_KEY string = C.WIDGET_PROP_MOVE_FOCUS_UP_KEY
  WIDGET_PROP_MOVE_FOCUS_DOWN_KEY string = C.WIDGET_PROP_MOVE_FOCUS_DOWN_KEY
  WIDGET_PROP_MOVE_FOCUS_LEFT_KEY string = C.WIDGET_PROP_MOVE_FOCUS_LEFT_KEY
  WIDGET_PROP_MOVE_FOCUS_RIGHT_KEY string = C.WIDGET_PROP_MOVE_FOCUS_RIGHT_KEY
  WIDGET_PROP_ROWS string = C.WIDGET_PROP_ROWS
  WIDGET_PROP_SHOW_GRID string = C.WIDGET_PROP_SHOW_GRID
  WIDGET_PROP_COLUMNS_DEFINITION string = C.WIDGET_PROP_COLUMNS_DEFINITION
  WIDGET_PROP_DRAG_THRESHOLD string = C.WIDGET_PROP_DRAG_THRESHOLD
  WIDGET_PROP_ANIMATING_TIME string = C.WIDGET_PROP_ANIMATING_TIME
  WIDGET_PROP_ANIMATE_PREFIX string = C.WIDGET_PROP_ANIMATE_PREFIX
  WIDGET_PROP_ANIMATE_ANIMATING_TIME string = C.WIDGET_PROP_ANIMATE_ANIMATING_TIME
  WIDGET_PROP_DIRTY_RECT string = C.WIDGET_PROP_DIRTY_RECT
  WIDGET_PROP_SCREEN_SAVER_TIME string = C.WIDGET_PROP_SCREEN_SAVER_TIME
  WIDGET_PROP_SHOW_FPS string = C.WIDGET_PROP_SHOW_FPS
  WIDGET_PROP_MAX_FPS string = C.WIDGET_PROP_MAX_FPS
  WIDGET_PROP_VALIDATOR string = C.WIDGET_PROP_VALIDATOR
)
type TWidgetState string
const (
  WIDGET_STATE_NONE string = C.WIDGET_STATE_NONE
  WIDGET_STATE_NORMAL string = C.WIDGET_STATE_NORMAL
  WIDGET_STATE_ACTIVATED string = C.WIDGET_STATE_ACTIVATED
  WIDGET_STATE_CHANGED string = C.WIDGET_STATE_CHANGED
  WIDGET_STATE_PRESSED string = C.WIDGET_STATE_PRESSED
  WIDGET_STATE_OVER string = C.WIDGET_STATE_OVER
  WIDGET_STATE_DISABLE string = C.WIDGET_STATE_DISABLE
  WIDGET_STATE_FOCUSED string = C.WIDGET_STATE_FOCUSED
  WIDGET_STATE_CHECKED string = C.WIDGET_STATE_CHECKED
  WIDGET_STATE_UNCHECKED string = C.WIDGET_STATE_UNCHECKED
  WIDGET_STATE_EMPTY string = C.WIDGET_STATE_EMPTY
  WIDGET_STATE_EMPTY_FOCUS string = C.WIDGET_STATE_EMPTY_FOCUS
  WIDGET_STATE_EMPTY_OVER string = C.WIDGET_STATE_EMPTY_OVER
  WIDGET_STATE_ERROR string = C.WIDGET_STATE_ERROR
  WIDGET_STATE_SELECTED string = C.WIDGET_STATE_SELECTED
  WIDGET_STATE_NORMAL_OF_CHECKED string = C.WIDGET_STATE_NORMAL_OF_CHECKED
  WIDGET_STATE_PRESSED_OF_CHECKED string = C.WIDGET_STATE_PRESSED_OF_CHECKED
  WIDGET_STATE_OVER_OF_CHECKED string = C.WIDGET_STATE_OVER_OF_CHECKED
  WIDGET_STATE_DISABLE_OF_CHECKED string = C.WIDGET_STATE_DISABLE_OF_CHECKED
  WIDGET_STATE_FOCUSED_OF_CHECKED string = C.WIDGET_STATE_FOCUSED_OF_CHECKED
  WIDGET_STATE_NORMAL_OF_ACTIVE string = C.WIDGET_STATE_NORMAL_OF_ACTIVE
  WIDGET_STATE_PRESSED_OF_ACTIVE string = C.WIDGET_STATE_PRESSED_OF_ACTIVE
  WIDGET_STATE_OVER_OF_ACTIVE string = C.WIDGET_STATE_OVER_OF_ACTIVE
  WIDGET_STATE_DISABLE_OF_ACTIVE string = C.WIDGET_STATE_DISABLE_OF_ACTIVE
  WIDGET_STATE_FOCUSED_OF_ACTIVE string = C.WIDGET_STATE_FOCUSED_OF_ACTIVE
  WIDGET_STATE_NORMAL_OF_INDETERMINATE string = C.WIDGET_STATE_NORMAL_OF_INDETERMINATE
  WIDGET_STATE_PRESSED_OF_INDETERMINATE string = C.WIDGET_STATE_PRESSED_OF_INDETERMINATE
  WIDGET_STATE_OVER_OF_INDETERMINATE string = C.WIDGET_STATE_OVER_OF_INDETERMINATE
  WIDGET_STATE_DISABLE_OF_INDETERMINATE string = C.WIDGET_STATE_DISABLE_OF_INDETERMINATE
  WIDGET_STATE_FOCUSED_OF_INDETERMINATE string = C.WIDGET_STATE_FOCUSED_OF_INDETERMINATE
)
type TWidgetType string
const (
  WIDGET_TYPE_NONE string = C.WIDGET_TYPE_NONE
  WIDGET_TYPE_WINDOW_MANAGER string = C.WIDGET_TYPE_WINDOW_MANAGER
  WIDGET_TYPE_NORMAL_WINDOW string = C.WIDGET_TYPE_NORMAL_WINDOW
  WIDGET_TYPE_OVERLAY string = C.WIDGET_TYPE_OVERLAY
  WIDGET_TYPE_TOOL_BAR string = C.WIDGET_TYPE_TOOL_BAR
  WIDGET_TYPE_DIALOG string = C.WIDGET_TYPE_DIALOG
  WIDGET_TYPE_POPUP string = C.WIDGET_TYPE_POPUP
  WIDGET_TYPE_SYSTEM_BAR string = C.WIDGET_TYPE_SYSTEM_BAR
  WIDGET_TYPE_SYSTEM_BAR_BOTTOM string = C.WIDGET_TYPE_SYSTEM_BAR_BOTTOM
  WIDGET_TYPE_SPRITE string = C.WIDGET_TYPE_SPRITE
  WIDGET_TYPE_KEYBOARD string = C.WIDGET_TYPE_KEYBOARD
  WIDGET_TYPE_DND string = C.WIDGET_TYPE_DND
  WIDGET_TYPE_LABEL string = C.WIDGET_TYPE_LABEL
  WIDGET_TYPE_BUTTON string = C.WIDGET_TYPE_BUTTON
  WIDGET_TYPE_IMAGE string = C.WIDGET_TYPE_IMAGE
  WIDGET_TYPE_ICON string = C.WIDGET_TYPE_ICON
  WIDGET_TYPE_EDIT string = C.WIDGET_TYPE_EDIT
  WIDGET_TYPE_PROGRESS_BAR string = C.WIDGET_TYPE_PROGRESS_BAR
  WIDGET_TYPE_GROUP_BOX string = C.WIDGET_TYPE_GROUP_BOX
  WIDGET_TYPE_CHECK_BUTTON string = C.WIDGET_TYPE_CHECK_BUTTON
  WIDGET_TYPE_RADIO_BUTTON string = C.WIDGET_TYPE_RADIO_BUTTON
  WIDGET_TYPE_DIALOG_TITLE string = C.WIDGET_TYPE_DIALOG_TITLE
  WIDGET_TYPE_DIALOG_CLIENT string = C.WIDGET_TYPE_DIALOG_CLIENT
  WIDGET_TYPE_SLIDER string = C.WIDGET_TYPE_SLIDER
  WIDGET_TYPE_VIEW string = C.WIDGET_TYPE_VIEW
  WIDGET_TYPE_PAGE string = C.WIDGET_TYPE_PAGE
  WIDGET_TYPE_COMBO_BOX string = C.WIDGET_TYPE_COMBO_BOX
  WIDGET_TYPE_COMBO_BOX_ITEM string = C.WIDGET_TYPE_COMBO_BOX_ITEM
  WIDGET_TYPE_SLIDE_VIEW string = C.WIDGET_TYPE_SLIDE_VIEW
  WIDGET_TYPE_SLIDE_INDICATOR string = C.WIDGET_TYPE_SLIDE_INDICATOR
  WIDGET_TYPE_SLIDE_INDICATOR_ARC string = C.WIDGET_TYPE_SLIDE_INDICATOR_ARC
  WIDGET_TYPE_PAGES string = C.WIDGET_TYPE_PAGES
  WIDGET_TYPE_TAB_BUTTON string = C.WIDGET_TYPE_TAB_BUTTON
  WIDGET_TYPE_TAB_CONTROL string = C.WIDGET_TYPE_TAB_CONTROL
  WIDGET_TYPE_TAB_BUTTON_GROUP string = C.WIDGET_TYPE_TAB_BUTTON_GROUP
  WIDGET_TYPE_BUTTON_GROUP string = C.WIDGET_TYPE_BUTTON_GROUP
  WIDGET_TYPE_CANDIDATES string = C.WIDGET_TYPE_CANDIDATES
  WIDGET_TYPE_SPIN_BOX string = C.WIDGET_TYPE_SPIN_BOX
  WIDGET_TYPE_DRAGGER string = C.WIDGET_TYPE_DRAGGER
  WIDGET_TYPE_SCROLL_BAR string = C.WIDGET_TYPE_SCROLL_BAR
  WIDGET_TYPE_SCROLL_BAR_DESKTOP string = C.WIDGET_TYPE_SCROLL_BAR_DESKTOP
  WIDGET_TYPE_SCROLL_BAR_MOBILE string = C.WIDGET_TYPE_SCROLL_BAR_MOBILE
  WIDGET_TYPE_SCROLL_VIEW string = C.WIDGET_TYPE_SCROLL_VIEW
  WIDGET_TYPE_LIST_VIEW string = C.WIDGET_TYPE_LIST_VIEW
  WIDGET_TYPE_LIST_VIEW_H string = C.WIDGET_TYPE_LIST_VIEW_H
  WIDGET_TYPE_LIST_ITEM string = C.WIDGET_TYPE_LIST_ITEM
  WIDGET_TYPE_COLOR_PICKER string = C.WIDGET_TYPE_COLOR_PICKER
  WIDGET_TYPE_COLOR_COMPONENT string = C.WIDGET_TYPE_COLOR_COMPONENT
  WIDGET_TYPE_COLOR_TILE string = C.WIDGET_TYPE_COLOR_TILE
  WIDGET_TYPE_CLIP_VIEW string = C.WIDGET_TYPE_CLIP_VIEW
  WIDGET_TYPE_RICH_TEXT string = C.WIDGET_TYPE_RICH_TEXT
  WIDGET_TYPE_APP_BAR string = C.WIDGET_TYPE_APP_BAR
  WIDGET_TYPE_GRID string = C.WIDGET_TYPE_GRID
  WIDGET_TYPE_GRID_ITEM string = C.WIDGET_TYPE_GRID_ITEM
  WIDGET_TYPE_ROW string = C.WIDGET_TYPE_ROW
  WIDGET_TYPE_COLUMN string = C.WIDGET_TYPE_COLUMN
  WIDGET_TYPE_CALIBRATION_WIN string = C.WIDGET_TYPE_CALIBRATION_WIN
)
type TWindow struct {
  TWindowBase
}

func TWindowCreate(parent TWidget, x int, y int, w int, h int) TWidget {
  retObj := TWindow{}
  retObj.handle = unsafe.Pointer(C.window_create((*C.widget_t)(parent.handle), (C.xy_t)(x), (C.xy_t)(y), (C.wh_t)(w), (C.wh_t)(h)))
  return retObj.TWidget
}

func TWindowCreateDefault() TWidget {
  retObj := TWindow{}
  retObj.handle = unsafe.Pointer(C.window_create_default())
  return retObj.TWidget
}

func (this TWindow) SetFullscreen(fullscreen bool) TRet {
  return TRet(C.window_set_fullscreen((*C.widget_t)(this.handle), (C.bool_t)(fullscreen)));
}

func (this TWindow) SetAutoScaleChildren(design_w uint32, design_h uint32) TRet {
  return TRet(C.window_set_auto_scale_children((*C.widget_t)(this.handle), (C.uint32_t)(design_w), (C.uint32_t)(design_h)));
}

func TWindowOpen(name string) TWidget {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TWindow{}
  retObj.handle = unsafe.Pointer(C.window_open(aname))
  return retObj.TWidget
}

func TWindowOpenAndClose(name string, to_close TWidget) TWidget {
  aname := C.CString(name)
  defer C.free(unsafe.Pointer(aname))
  retObj := TWindow{}
  retObj.handle = unsafe.Pointer(C.window_open_and_close(aname, (*C.widget_t)(to_close.handle)))
  return retObj.TWidget
}

func (this TWindow) Close() TRet {
  return TRet(C.window_close((*C.widget_t)(this.handle)));
}

func (this TWindow) CloseForce() TRet {
  return TRet(C.window_close_force((*C.widget_t)(this.handle)));
}

func TWindowCast(widget TWidget) TWindow {
  retObj := TWindow{}
  retObj.handle = unsafe.Pointer(C.window_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TWindow) GetFullscreen() bool {
  return (bool)((*C.window_t)(unsafe.Pointer(this.handle)).fullscreen);
}

type TWindowBase struct {
  TWidget
}

func TWindowBaseCast(widget TWidget) TWindowBase {
  retObj := TWindowBase{}
  retObj.handle = unsafe.Pointer(C.window_base_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TWindowBase) GetTheme() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).theme);
}

func (this TWindowBase) GetDesignW() uint16 {
  return (uint16)((*C.window_base_t)(unsafe.Pointer(this.handle)).design_w);
}

func (this TWindowBase) GetDesignH() uint16 {
  return (uint16)((*C.window_base_t)(unsafe.Pointer(this.handle)).design_h);
}

func (this TWindowBase) GetAutoScaleChildrenX() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).auto_scale_children_x);
}

func (this TWindowBase) GetAutoScaleChildrenY() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).auto_scale_children_y);
}

func (this TWindowBase) GetAutoScaleChildrenW() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).auto_scale_children_w);
}

func (this TWindowBase) GetAutoScaleChildrenH() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).auto_scale_children_h);
}

func (this TWindowBase) GetDisableAnim() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).disable_anim);
}

func (this TWindowBase) GetClosable() TWindowClosable {
  return TWindowClosable((*C.window_base_t)(unsafe.Pointer(this.handle)).closable);
}

func (this TWindowBase) GetOpenAnimHint() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).open_anim_hint);
}

func (this TWindowBase) GetCloseAnimHint() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).close_anim_hint);
}

func (this TWindowBase) GetMoveFocusPrevKey() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).move_focus_prev_key);
}

func (this TWindowBase) GetMoveFocusNextKey() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).move_focus_next_key);
}

func (this TWindowBase) GetMoveFocusUpKey() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).move_focus_up_key);
}

func (this TWindowBase) GetMoveFocusDownKey() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).move_focus_down_key);
}

func (this TWindowBase) GetMoveFocusLeftKey() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).move_focus_left_key);
}

func (this TWindowBase) GetMoveFocusRightKey() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).move_focus_right_key);
}

func (this TWindowBase) GetAppletName() string {
  return C.GoString((*C.window_base_t)(unsafe.Pointer(this.handle)).applet_name);
}

func (this TWindowBase) GetSingleInstance() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).single_instance);
}

func (this TWindowBase) GetStronglyFocus() bool {
  return (bool)((*C.window_base_t)(unsafe.Pointer(this.handle)).strongly_focus);
}

type TWindowClosable int
const (
  WINDOW_CLOSABLE_YES TWindowClosable = C.WINDOW_CLOSABLE_YES
  WINDOW_CLOSABLE_NO TWindowClosable = C.WINDOW_CLOSABLE_NO
  WINDOW_CLOSABLE_CONFIRM TWindowClosable = C.WINDOW_CLOSABLE_CONFIRM
)
type TWindowEvent struct {
  TEvent
}

func TWindowEventCast(event TEvent) TWindowEvent {
  retObj := TWindowEvent{}
  retObj.handle = unsafe.Pointer(C.window_event_cast((*C.event_t)(event.handle)))
  return retObj
}

func (this TWindowEvent) GetWindow() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer((*C.window_event_t)(unsafe.Pointer(this.handle)).window)
  return retObj
}

type TWindowManager struct {
  TWidget
}

func TWindowManagerInstance() TWindowManager {
  retObj := TWindowManager{}
  retObj.handle = unsafe.Pointer(C.window_manager())
  return retObj
}

func TWindowManagerCast(widget TWidget) TWindowManager {
  retObj := TWindowManager{}
  retObj.handle = unsafe.Pointer(C.window_manager_cast((*C.widget_t)(widget.handle)))
  return retObj
}

func (this TWindowManager) GetTopMainWindow() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.window_manager_get_top_main_window((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWindowManager) GetTopWindow() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.window_manager_get_top_window((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWindowManager) GetForegroundWindow() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.window_manager_get_foreground_window((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWindowManager) GetPrevWindow() TWidget {
  retObj := TWidget{}
  retObj.handle = unsafe.Pointer(C.window_manager_get_prev_window((*C.widget_t)(this.handle)))
  return retObj
}

func (this TWindowManager) GetPointerX() int {
  return (int)(C.window_manager_get_pointer_x((*C.widget_t)(this.handle)));
}

func (this TWindowManager) GetPointerY() int {
  return (int)(C.window_manager_get_pointer_y((*C.widget_t)(this.handle)));
}

func (this TWindowManager) GetPointerPressed() bool {
  return (bool)(C.window_manager_get_pointer_pressed((*C.widget_t)(this.handle)));
}

func (this TWindowManager) IsAnimating() bool {
  return (bool)(C.window_manager_is_animating((*C.widget_t)(this.handle)));
}

func (this TWindowManager) SetShowFps(show_fps bool) TRet {
  return TRet(C.window_manager_set_show_fps((*C.widget_t)(this.handle), (C.bool_t)(show_fps)));
}

func (this TWindowManager) SetShowFpsPosition(x int, y int) TRet {
  return TRet(C.window_manager_set_show_fps_position((*C.widget_t)(this.handle), (C.xy_t)(x), (C.xy_t)(y)));
}

func (this TWindowManager) SetMaxFps(max_fps uint32) TRet {
  return TRet(C.window_manager_set_max_fps((*C.widget_t)(this.handle), (C.uint32_t)(max_fps)));
}

func (this TWindowManager) SetIgnoreInputEvents(ignore_input_events bool) TRet {
  return TRet(C.window_manager_set_ignore_input_events((*C.widget_t)(this.handle), (C.bool_t)(ignore_input_events)));
}

func (this TWindowManager) SetScreenSaverTime(screen_saver_time uint32) TRet {
  return TRet(C.window_manager_set_screen_saver_time((*C.widget_t)(this.handle), (C.uint32_t)(screen_saver_time)));
}

func (this TWindowManager) SetCursor(cursor string) TRet {
  acursor := C.CString(cursor)
  defer C.free(unsafe.Pointer(acursor))
  return TRet(C.window_manager_set_cursor((*C.widget_t)(this.handle), acursor));
}

func (this TWindowManager) Back() TRet {
  return TRet(C.window_manager_back((*C.widget_t)(this.handle)));
}

func (this TWindowManager) BackToHome() TRet {
  return TRet(C.window_manager_back_to_home((*C.widget_t)(this.handle)));
}

func (this TWindowManager) BackTo(target string) TRet {
  atarget := C.CString(target)
  defer C.free(unsafe.Pointer(atarget))
  return TRet(C.window_manager_back_to((*C.widget_t)(this.handle), atarget));
}

func (this TWindowManager) Resize(w int, h int) TRet {
  return TRet(C.window_manager_resize((*C.widget_t)(this.handle), (C.wh_t)(w), (C.wh_t)(h)));
}

func (this TWindowManager) SetFullscreen(fullscreen bool) TRet {
  return TRet(C.window_manager_set_fullscreen((*C.widget_t)(this.handle), (C.bool_t)(fullscreen)));
}

func (this TWindowManager) CloseAll() TRet {
  return TRet(C.window_manager_close_all((*C.widget_t)(this.handle)));
}

type TWindowStage int
const (
  WINDOW_STAGE_NONE TWindowStage = C.WINDOW_STAGE_NONE
  WINDOW_STAGE_LOADED TWindowStage = C.WINDOW_STAGE_LOADED
  WINDOW_STAGE_CREATED TWindowStage = C.WINDOW_STAGE_CREATED
  WINDOW_STAGE_OPENED TWindowStage = C.WINDOW_STAGE_OPENED
  WINDOW_STAGE_CLOSED TWindowStage = C.WINDOW_STAGE_CLOSED
  WINDOW_STAGE_SUSPEND TWindowStage = C.WINDOW_STAGE_SUSPEND
)
