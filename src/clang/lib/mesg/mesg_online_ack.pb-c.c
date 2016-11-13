/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: mesg_online_ack.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "mesg_online_ack.pb-c.h"
void   mesg_online_ack__init
                     (MesgOnlineAck         *message)
{
  static MesgOnlineAck init_value = MESG_ONLINE_ACK__INIT;
  *message = init_value;
}
size_t mesg_online_ack__get_packed_size
                     (const MesgOnlineAck *message)
{
  assert(message->base.descriptor == &mesg_online_ack__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t mesg_online_ack__pack
                     (const MesgOnlineAck *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &mesg_online_ack__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t mesg_online_ack__pack_to_buffer
                     (const MesgOnlineAck *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &mesg_online_ack__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
MesgOnlineAck *
       mesg_online_ack__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (MesgOnlineAck *)
     protobuf_c_message_unpack (&mesg_online_ack__descriptor,
                                allocator, len, data);
}
void   mesg_online_ack__free_unpacked
                     (MesgOnlineAck *message,
                      ProtobufCAllocator *allocator)
{
  assert(message->base.descriptor == &mesg_online_ack__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor mesg_online_ack__field_descriptors[7] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(MesgOnlineAck, uid),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "sid",
    2,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(MesgOnlineAck, sid),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "app",
    3,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(MesgOnlineAck, app),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "version",
    4,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(MesgOnlineAck, version),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "terminal",
    5,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT32,
    offsetof(MesgOnlineAck, has_terminal),
    offsetof(MesgOnlineAck, terminal),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "errnum",
    6,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT32,
    offsetof(MesgOnlineAck, has_errnum),
    offsetof(MesgOnlineAck, errnum),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "errmsg",
    7,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(MesgOnlineAck, errmsg),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned mesg_online_ack__field_indices_by_name[] = {
  2,   /* field[2] = app */
  6,   /* field[6] = errmsg */
  5,   /* field[5] = errnum */
  1,   /* field[1] = sid */
  4,   /* field[4] = terminal */
  0,   /* field[0] = uid */
  3,   /* field[3] = version */
};
static const ProtobufCIntRange mesg_online_ack__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 7 }
};
const ProtobufCMessageDescriptor mesg_online_ack__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "mesg_online_ack",
  "MesgOnlineAck",
  "MesgOnlineAck",
  "",
  sizeof(MesgOnlineAck),
  7,
  mesg_online_ack__field_descriptors,
  mesg_online_ack__field_indices_by_name,
  1,  mesg_online_ack__number_ranges,
  (ProtobufCMessageInit) mesg_online_ack__init,
  NULL,NULL,NULL    /* reserved[123] */
};
