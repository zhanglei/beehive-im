/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: mesg_join_ack.proto */

#ifndef PROTOBUF_C_mesg_5fjoin_5fack_2eproto__INCLUDED
#define PROTOBUF_C_mesg_5fjoin_5fack_2eproto__INCLUDED

#include <protobuf-c/protobuf-c.h>

PROTOBUF_C__BEGIN_DECLS

#if PROTOBUF_C_VERSION_NUMBER < 1000000
# error This file was generated by a newer version of protoc-c which is incompatible with your libprotobuf-c headers. Please update your headers.
#elif 1000002 < PROTOBUF_C_MIN_COMPILER_VERSION
# error This file was generated by an older version of protoc-c which is incompatible with your libprotobuf-c headers. Please regenerate this file with a newer version of protoc-c.
#endif


typedef struct _MesgJoinAck MesgJoinAck;


/* --- enums --- */


/* --- messages --- */

struct  _MesgJoinAck
{
  ProtobufCMessage base;
  uint64_t uid;
  uint64_t rid;
  uint32_t gid;
  protobuf_c_boolean has_errnum;
  uint32_t errnum;
  char *errmsg;
};
#define MESG_JOIN_ACK__INIT \
 { PROTOBUF_C_MESSAGE_INIT (&mesg_join_ack__descriptor) \
    , 0, 0, 0, 0,0, NULL }


/* MesgJoinAck methods */
void   mesg_join_ack__init
                     (MesgJoinAck         *message);
size_t mesg_join_ack__get_packed_size
                     (const MesgJoinAck   *message);
size_t mesg_join_ack__pack
                     (const MesgJoinAck   *message,
                      uint8_t             *out);
size_t mesg_join_ack__pack_to_buffer
                     (const MesgJoinAck   *message,
                      ProtobufCBuffer     *buffer);
MesgJoinAck *
       mesg_join_ack__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data);
void   mesg_join_ack__free_unpacked
                     (MesgJoinAck *message,
                      ProtobufCAllocator *allocator);
/* --- per-message closures --- */

typedef void (*MesgJoinAck_Closure)
                 (const MesgJoinAck *message,
                  void *closure_data);

/* --- services --- */


/* --- descriptors --- */

extern const ProtobufCMessageDescriptor mesg_join_ack__descriptor;

PROTOBUF_C__END_DECLS


#endif  /* PROTOBUF_C_mesg_5fjoin_5fack_2eproto__INCLUDED */
