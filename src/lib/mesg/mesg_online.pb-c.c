/* Generated by the protocol buffer compiler.  DO NOT EDIT! */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C_NO_DEPRECATED
#define PROTOBUF_C_NO_DEPRECATED
#endif

#include "mesg_online.pb-c.h"
void   mesg_online__init
                     (MesgOnline         *message)
{
  static MesgOnline init_value = MESG_ONLINE__INIT;
  *message = init_value;
}
size_t mesg_online__get_packed_size
                     (const MesgOnline *message)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_online__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t mesg_online__pack
                     (const MesgOnline *message,
                      uint8_t       *out)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_online__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t mesg_online__pack_to_buffer
                     (const MesgOnline *message,
                      ProtobufCBuffer *buffer)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_online__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
MesgOnline *
       mesg_online__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (MesgOnline *)
     protobuf_c_message_unpack (&mesg_online__descriptor,
                                allocator, len, data);
}
void   mesg_online__free_unpacked
                     (MesgOnline *message,
                      ProtobufCAllocator *allocator)
{
  PROTOBUF_C_ASSERT (message->base.descriptor == &mesg_online__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor mesg_online__field_descriptors[5] =
{
  {
    "uid",
    1,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT64,
    PROTOBUF_C_OFFSETOF(MesgOnline, has_uid),
    PROTOBUF_C_OFFSETOF(MesgOnline, uid),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "token",
    2,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    PROTOBUF_C_OFFSETOF(MesgOnline, token),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "app",
    3,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    PROTOBUF_C_OFFSETOF(MesgOnline, app),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "version",
    4,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    PROTOBUF_C_OFFSETOF(MesgOnline, version),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "terminal",
    5,
    PROTOBUF_C_LABEL_OPTIONAL,
    PROTOBUF_C_TYPE_UINT32,
    PROTOBUF_C_OFFSETOF(MesgOnline, has_terminal),
    PROTOBUF_C_OFFSETOF(MesgOnline, terminal),
    NULL,
    NULL,
    0,            /* packed */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned mesg_online__field_indices_by_name[] = {
  2,   /* field[2] = app */
  4,   /* field[4] = terminal */
  1,   /* field[1] = token */
  0,   /* field[0] = uid */
  3,   /* field[3] = version */
};
static const ProtobufCIntRange mesg_online__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 5 }
};
const ProtobufCMessageDescriptor mesg_online__descriptor =
{
  PROTOBUF_C_MESSAGE_DESCRIPTOR_MAGIC,
  "mesg_online",
  "MesgOnline",
  "MesgOnline",
  "",
  sizeof(MesgOnline),
  5,
  mesg_online__field_descriptors,
  mesg_online__field_indices_by_name,
  1,  mesg_online__number_ranges,
  (ProtobufCMessageInit) mesg_online__init,
  NULL,NULL,NULL    /* reserved[123] */
};