#include "awtk.h"

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

////////////////////emitter//////////////////////////////////////

static uint32_t wrap_emitter_on(emitter_t* emitter, uint32_t etype, void* ctx) {
  uint32_t ret = emitter_on(emitter, etype, (event_func_t)wrap_on_event, ctx);
  if (ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    emitter_set_on_destroy(emitter, ret, emitter_item_on_destroy, NULL);
  }

  return ret;
}

////////////////////timer//////////////////////////////////////
extern int GoOnTimer(void* ctx);

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

static ret_t wrap_queue_timer(void* ctx, uint32_t duration) {
  return timer_queue_ex(wrap_on_timer, ctx, duration, timer_info_on_destroy, NULL);
}

////////////////////widget_add_timer//////////////////////////////////////
static uint32_t wrap_widget_add_timer(widget_t* widget, void* ctx, uint32_t duration) {
  uint32_t ret = widget_add_timer(widget, wrap_on_timer, duration); 
  if (ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    timer_info_t* info = (timer_info_t*)timer_find(ret);
    if(info != NULL) {
      info->ctx = ctx;
    }
    timer_set_on_destroy(ret, timer_info_on_destroy, NULL); 
  }

  return ret;
}


////////////////////idle//////////////////////////////////////
extern int GoOnIdle(void* ctx);

static ret_t idle_info_on_destroy(void *data) {
  idle_info_t *item = (idle_info_t *)data;
  GoReleasePointer(item->ctx);
  return RET_OK;
}

static ret_t wrap_on_idle(const idle_info_t* info) {
  return (ret_t)GoOnIdle(info->ctx);
}

static uint32_t wrap_add_idle(void* ctx) {
  uint32_t ret = idle_add(wrap_on_idle, ctx);
  if(ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    idle_set_on_destroy(ret, idle_info_on_destroy, NULL); 
  }

  return ret;
}

static ret_t wrap_queue_idle(void* ctx) {
  return idle_queue_ex(wrap_on_idle, ctx, idle_info_on_destroy, NULL);
}

////////////////////widget_add_idle//////////////////////////////////////
static uint32_t wrap_widget_add_idle(widget_t* widget, void* ctx) {
  uint32_t ret = widget_add_idle(widget, wrap_on_idle); 
  if (ret == TK_INVALID_ID) {
    GoReleasePointer(ctx);
  } else {
    idle_info_t* info = (idle_info_t*)idle_find(ret);
    if(info != NULL) {
      info->ctx = ctx;
    }
    idle_set_on_destroy(ret, idle_info_on_destroy, NULL); 
  }

  return ret;
}
