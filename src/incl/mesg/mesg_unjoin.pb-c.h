/* Generated by the protocol buffer compiler.  DO NOT EDIT! */

#ifndef PROTOBUF_C_mesg_5funjoin_2eproto__INCLUDED
#define PROTOBUF_C_mesg_5funjoin_2eproto__INCLUDED

#include <google/protobuf-c/protobuf-c.h>

PROTOBUF_C_BEGIN_DECLS


typedef struct _MesgUnjoinReq MesgUnjoinReq;


/* --- enums --- */


/* --- messages --- */

struct  _MesgUnjoinReq
{
  ProtobufCMessage base;
  protobuf_c_boolean has_uid;
  uint64_t uid;
  protobuf_c_boolean has_rid;
  uint64_t rid;
};
#define MESG_UNJOIN_REQ__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_unjoin_req__descriptor) \
    , 0,0, 0,0 }


/* MesgUnjoinReq methods */
void   mesg_unjoin_req__init
                     (MesgUnjoinReq         *message);
size_t mesg_unjoin_req__get_packed_size
                     (const MesgUnjoinReq   *message);
size_t mesg_unjoin_req__pack
                     (const MesgUnjoinReq   *message,
                      uint8_t             *out);
size_t mesg_unjoin_req__pack_to_buffer
                     (const MesgUnjoinReq   *message,
                      ProtobufCBuffer     *buffer);
MesgUnjoinReq *
       mesg_unjoin_req__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_unjoin_req__free_unpacked
                     (MesgUnjoinReq *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*MesgUnjoinReq_Closure)
                 (const MesgUnjoinReq *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor mesg_unjoin_req__descriptor;

PROTOBUF_C_END_DECLS


#endif  /* PROTOBUF_mesg_5funjoin_2eproto__INCLUDED */
