#include "awtk.h"

extern int GoOnTimer(void* ctx);
extern int GoOnEvent(void* ctx, void* e);
extern int GoReleasePointer(void* ctx);

static ret_t emitter_item_on_destroy(void* data) {
  emitter_item_t* item = (emitter_item_t*)data;
  GoReleasePointer(item->ctx);

  return RET_OK;
}

static ret_t wrap_on_event(void* ctx, event_t* e) {
  return (ret_t)GoOnEvent(ctx, e);
}

static uint32_t wrap_widget_on(widget_t* widget, uint32_t etype, void* ctx) {
  uint32_t ret = widget_on(widget, etype, (event_func_t)wrap_on_event, ctx);
  if (ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    emitter_set_on_destroy(widget->emitter, ret, emitter_item_on_destroy, NULL);
  }

  return ret;
}

static uint32_t wrap_emitter_on(emitter_t* emitter, uint32_t etype, void* ctx) {
  uint32_t ret = emitter_on(emitter, etype, (event_func_t)wrap_on_event, ctx);
  if (ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    emitter_set_on_destroy(emitter, ret, emitter_item_on_destroy, NULL);
  }

  return ret;
}

static ret_t timer_info_on_destroy(void *data) {
  timer_info_t *item = (timer_info_t *)data;
  GoReleasePointer(item->ctx);
  return RET_OK;
}

static ret_t wrap_on_timer(const timer_info_t* info) {
  return (ret_t)GoOnTimer(info->ctx);
}

static uint32_t wrap_add_timer(void* ctx, uint32_t duration) {
  uint32_t ret = timer_add(wrap_on_timer, ctx, duration);
  if(ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    timer_set_on_destroy(ret, timer_info_on_destroy, NULL); 
  }

  return ret;
}

