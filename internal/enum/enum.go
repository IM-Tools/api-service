package enum

const (
	ParamError = 1000 // 参数错误
	ApiError   = 1001 // 接口异常 比如调用第三方接口 或者代码异常

	WsChantMessage    = 200  // 聊天消息
	VideoChantMessage = 600  // 好友视频请求
	WsCreate          = 1000 // 添加好友
	WsFriendOk        = 1001 // 同意好友申请
	WsFriendError     = 1002 // 拒绝
	WsNotFriend       = 1003 // 非好友关系
	WsPing            = 1004 // 心跳
	WsAck             = 1005 // 确认机制

	WsUserOffline  = 2000 // 用户离线
	WsUserOnline   = 2001 // 用户在线
	WsIsUserStatus = 2002 // 前端请求判断用户是否在线
	WsSession      = 2003 // 会话推送
	WsLoginOut     = 2004 // 异地登录

	WsGroupMessage = 3000 // 入群邀请

	PrivateMessage = 1 // 私聊消息
	GroupMessage   = 2 // 群聊消息

	TEXT         = 1 // 文本
	VOICE        = 2 // 语音
	FILE         = 3 // 文件
	IMAGE        = 4 // 图片
	LOGOUT_GROUP = 5 // 退出群聊
	JOIN_GROUP   = 6 // 加入群聊

)
