/* Generated by the protocol buffer compiler.  DO NOT EDIT! */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C_NO_DEPRECATED
#define PROTOBUF_C_NO_DEPRECATED
#endif

#include "mesg_join.pb-c.h"
void   mesg_join_req__init
                     (MesgJoinReq         *message)
{
  static MesgJoinReq init_value = MESG_JOIN_REQ__INIT;
  *message = init_value;
}
size_t mesg_join_req__get_packed_size
                     (const MesgJoinReq *message)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_join_req__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t mesg_join_req__pack
                     (const MesgJoinReq *message,
                      uint8_t       *out)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_join_req__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t mesg_join_req__pack_to_buffer
                     (const MesgJoinReq *message,
                      ProtobufCBuffer *buffer)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_join_req__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
MesgJoinReq *
       mesg_join_req__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (MesgJoinReq *)
     protobuf_c_message_unpack (&mesg_join_req__descriptor,
                                allocator, len, data);
}
void   mesg_join_req__free_unpacked
                     (MesgJoinReq *message,
                      ProtobufCAllocator *allocator)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_join_req__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor mesg_join_req__field_descriptors[3] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT64,
    PROTOBUF_C_OFFSETOF(MesgJoinReq, has_uid),
    PROTOBUF_C_OFFSETOF(MesgJoinReq, uid),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "rid",
    2,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT64,
    PROTOBUF_C_OFFSETOF(MesgJoinReq, has_rid),
    PROTOBUF_C_OFFSETOF(MesgJoinReq, rid),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "token",
    3,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    PROTOBUF_C_OFFSETOF(MesgJoinReq, token),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned mesg_join_req__field_indices_by_name[] = {
  1,   /* field[1] = rid */
  2,   /* field[2] = token */
  0,   /* field[0] = uid */
};
static const ProtobufCIntRange mesg_join_req__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor mesg_join_req__descriptor =
{
  PROTOBUF_C_MESSAGE_DESCRIPTOR_MAGIC,
  "mesg_join_req",
  "MesgJoinReq",
  "MesgJoinReq",
  "",
  sizeof(MesgJoinReq),
  3,
  mesg_join_req__field_descriptors,
  mesg_join_req__field_indices_by_name,
  1,  mesg_join_req__number_ranges,
  (ProtobufCMessageInit) mesg_join_req__init,
  NULL,NULL,NULL    /* reserved[123] */
};
