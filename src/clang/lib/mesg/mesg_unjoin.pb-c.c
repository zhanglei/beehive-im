/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: mesg_unjoin.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "mesg_unjoin.pb-c.h"
void   mesg_unjoin_req__init
                     (MesgUnjoinReq         *message)
{
  static MesgUnjoinReq init_value = MESG_UNJOIN_REQ__INIT;
  *message = init_value;
}
size_t mesg_unjoin_req__get_packed_size
                     (const MesgUnjoinReq *message)
{
  assert(message->base.descriptor == &mesg_unjoin_req__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t mesg_unjoin_req__pack
                     (const MesgUnjoinReq *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &mesg_unjoin_req__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t mesg_unjoin_req__pack_to_buffer
                     (const MesgUnjoinReq *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &mesg_unjoin_req__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
MesgUnjoinReq *
       mesg_unjoin_req__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (MesgUnjoinReq *)
     protobuf_c_message_unpack (&mesg_unjoin_req__descriptor,
                                allocator, len, data);
}
void   mesg_unjoin_req__free_unpacked
                     (MesgUnjoinReq *message,
                      ProtobufCAllocator *allocator)
{
  assert(message->base.descriptor == &mesg_unjoin_req__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor mesg_unjoin_req__field_descriptors[2] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(MesgUnjoinReq, uid),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "rid",
    2,
    PROTOBUF_C_LABEL_REQUIRED,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(MesgUnjoinReq, rid),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned mesg_unjoin_req__field_indices_by_name[] = {
  1,   /* field[1] = rid */
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange mesg_unjoin_req__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor mesg_unjoin_req__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "mesg_unjoin_req",
  "MesgUnjoinReq",
  "MesgUnjoinReq",
  "",
  sizeof(MesgUnjoinReq),
  2,
  mesg_unjoin_req__field_descriptors,
  mesg_unjoin_req__field_indices_by_name,
  1,  mesg_unjoin_req__number_ranges,
  (ProtobufCMessageInit) mesg_unjoin_req__init,
  NULL,NULL,NULL    /* reserved[123] */
};
