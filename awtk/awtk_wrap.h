#include "awtk.h"

extern int GoOnEvent(void* ctx, void* e);
extern int GoOnEventDestroy(void* ctx);

static ret_t emitter_item_on_destroy(void* data) {
  emitter_item_t* item = (emitter_item_t*)data;
  GoOnEventDestroy(item->ctx);

  return RET_OK;
}

static ret_t wrap_on_event(void* ctx, event_t* e) {
  return (ret_t)GoOnEvent(ctx, e);
}

static uint32_t wrap_widget_on(widget_t* widget, uint32_t etype, void* ctx) {
  uint32_t ret = widget_on(widget, etype, (event_func_t)wrap_on_event, ctx);
  if (ret == TK_INVALID_ID) {
    GoOnEventDestroy(ctx);
  } else {
    emitter_set_on_destroy(widget->emitter, ret, emitter_item_on_destroy, NULL);
  }

  return ret;
}
