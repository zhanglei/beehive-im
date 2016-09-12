# HTTP接口列表

##数据下发接口
###广播接口
**功能描述**: 用于向全员或某聊天室提交广播消息<br>
**接口类型**: POST<br>
**接口路径**: /chatroom/push?opt=broadcast&rid=${rid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为broadcast.(M)<br>
> rid: 聊天室ID # 当未指定rid时, 则为全员广播消息(O)<br>

**包体内容**: 下发的数据<br>
**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

###点推接口
**功能描述**: 用于指定聊天室的某人下发消息<br>
**接口类型**: POST<br>
**接口路径**: /chatroom/push?opt=p2p&rid=${rid}&uid=${uid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为broadcast.(M)<br>
> rid: 聊天室ID # 当未指定rid时, 则为全员广播消息(O)<br>
> uid: 用户ID(M)<br>

**包体内容**: 下发的数据
**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

##配置接口
###踢人接口
**功能描述**: 用于将某人踢出聊天室<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/config?opt=kick&rid=${rid}&uid=${uid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为kick.(M)<br>
> rid: 聊天室ID(M)<br>
> uid: 用户ID(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

###解除踢人接口
**功能描述**: 用于将某人踢出聊天室<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/config?opt=unkick&rid=${rid}&uid=${uid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为unkick.(M)<br>
> rid: 聊天室ID(M)<br>
> uid: 用户ID(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

###禁言接口
**功能描述**: 禁止某人在聊天室发言<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/config?opt=ban&rid=${rid}&uid=${uid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为ban.(M)<br>
> rid: 聊天室ID(M)<br>
> uid: 用户ID. # 当无uid或uid为0时, 全员禁言; 否则是禁止某人发言.<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

###解除禁言接口
**功能描述**: 禁止某人在聊天室发言<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/config?opt=unban&rid=${rid}&uid=${uid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为unban.(M)<br>
> rid: 聊天室ID(M)<br>
> uid: 用户ID(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

###关闭聊天室接口
**功能描述**: 关闭聊天室<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/config?opt=close&rid=${rid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为close.(M)<br>
> rid: 聊天室ID(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},    // 错误码(M)<br>
>   "errmsg":"${errmsg}" // 错误描述(M)<br>
>}<br>

##查询聊天室状态
###聊天室TOP排行
**功能描述**: 查询各聊天室TOP排行<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/query?opt=top-list&num=${num}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为top-list.(M)<br>
> num: top-${num}排行(O). 如果未设置${num}, 则显示前top-10的排行.<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},       // 错误码(M)<br>
>   "len":${len},           // 列表长度(M)<br>
>   "list":[                // 排行列表(M)<br>
>       {"rid":${rid}, "total":${total}}, // ${rid}:聊天室ID ${total}:聊天室人数<br>
>       {"rid":${rid}, "total":${total}},<br>
>       {"rid":${rid}, "total":${total}},<br>
>       {"rid":${rid}, "total":${total}}],<br>
>   "errmsg":"${errmsg}"    // 错误描述(M)<br>
>}

###查询某聊天室分组列表
**功能描述**: 查询某聊天室分组列表<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/query?opt=group-list&rid=${rid}<br>
**参数描述**:<br>
> opt: 操作选项, 此时为group-list.(M)<br>
> rid: 聊天室ID(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},       // 错误码(M)<br>
>   "rid":${rid},           // 聊天室ID(M)<br>
>   "len":${len},           // 列表长度(M)<br>
>   "list":[                // 分组列表(M)<br>
>       {"gid":${gid}, "total":${total}}, // ${gid}:分组ID ${total}:组人数<br>
>       {"gid":${gid}, "total":${total}},<br>
>       {"gid":${gid}, "total":${total}},<br>
>       {"gid":${gid}, "total":${total}}],<br>
>   "errmsg":"${errmsg}"    // 错误描述(M)<br>
>}

###查询人数分布
**功能描述**: 查询人数分布<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/query?opt=user-dist<br>
**参数描述**:<br>
> opt: 操作选项, 此时为user-dist.(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},       // 错误码(M)<br>
>   "len":${len},           // 列表长度(M)<br>
>   "list":[                // 分组列表(M)<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "total":"${total}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "total":"${total}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "total":"${total}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "total":"${total}"}],<br>
>   "errmsg":"${errmsg}"    // 错误描述(M)<br>
>}

###某用户在线状态
**功能描述**: 查询某用户在线状态<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/query?opt=user-online<br>
**参数描述**:<br>
> opt: 操作选项, 此时为user-online.(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},       // 错误码(M)<br>
>   "uid":"${uid}",         // 用户ID(M)<br>
>   "len":${len},           // 列表长度(M)<br>
>   "list":[                // 当前正登陆聊天室列表(M)<br>
>       {"rid":${rid}},     // ${rid}:聊天室ID<br>
>       {"rid":${rid}}],<br>
>   "errmsg":"${errmsg}"    // 错误描述(M)<br>
>}

##系统维护接口
###查询侦听层状态
**功能描述**: 查询侦听层状态<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/query?opt=listen-list<br>
**参数描述**:<br>
> opt: 操作选项, 此时为listen-list.(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},       // 错误码(M)<br>
>   "len":${len},           // 列表长度(M)<br>
>   "list":[                // 分组列表(M)<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"}],<br>
>   "errmsg":"${errmsg}"    // 错误描述(M)<br>
>}

###查询转发层状态
**功能描述**: 查询转发层状态<br>
**接口类型**: GET<br>
**接口路径**: /chatroom/query?opt=frwder-list<br>
**参数描述**:<br>
> opt: 操作选项, 此时为frwder-list.(M)<br>

**返回结果**:<br>
>{<br>
>   "errno":${errno},       // 错误码(M)<br>
>   "len":${len},           // 列表长度(M)<br>
>   "list":[                // 分组列表(M)<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"},<br>
>       {"nid":${nid}, "ipaddr":"{ipaddr}", "port":${port}, "status":"${status}"}],<br>
>   "errmsg":"${errmsg}"    // 错误描述(M)<br>
>}