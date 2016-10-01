/* Generated by the protocol buffer compiler.  DO NOT EDIT! */

#ifndef PROTOBUF_C_mesg_5fonline_2eproto__INCLUDED
#define PROTOBUF_C_mesg_5fonline_2eproto__INCLUDED

#include <google/protobuf-c/protobuf-c.h>

PROTOBUF_C_BEGIN_DECLS


typedef struct _MesgOnlineReq MesgOnlineReq;


/* --- enums --- */


/* --- messages --- */

struct  _MesgOnlineReq
{
  ProtobufCMessage base;
  protobuf_c_boolean has_uid;
  uint64_t uid;
  char *token;
  char *app;
  char *version;
  protobuf_c_boolean has_terminal;
  uint32_t terminal;
};
#define MESG_ONLINE_REQ__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_online_req__descriptor) \
    , 0,0, NULL, NULL, NULL, 0,0 }


/* MesgOnlineReq methods */
void   mesg_online_req__init
                     (MesgOnlineReq         *message);
size_t mesg_online_req__get_packed_size
                     (const MesgOnlineReq   *message);
size_t mesg_online_req__pack
                     (const MesgOnlineReq   *message,
                      uint8_t             *out);
size_t mesg_online_req__pack_to_buffer
                     (const MesgOnlineReq   *message,
                      ProtobufCBuffer     *buffer);
MesgOnlineReq *
       mesg_online_req__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_online_req__free_unpacked
                     (MesgOnlineReq *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*MesgOnlineReq_Closure)
                 (const MesgOnlineReq *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor mesg_online_req__descriptor;

PROTOBUF_C_END_DECLS


#endif  /* PROTOBUF_mesg_5fonline_2eproto__INCLUDED */
