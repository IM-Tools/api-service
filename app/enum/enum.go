package enum

const (
	ParamError = 1000 // 参数错误
	ApiError   = 1001 // 接口异常 比如调用第三方接口 或者代码异常

	WsSuccess      = 200  // 聊天消息
	WsCreate       = 1000 // 添加好友
	WsFriendOk     = 1001 // 同意好友申请
	WsFriendError  = 1002 // 拒绝
	WsNotFriend    = 1003 // 非好友关系
	WsPing         = 1004 // 心跳
	WsAck          = 1005 // 确认机制
	WsUserOffline  = 2000 // 用户离线
	WsUserOnline   = 2001 // 用户在线
	WsIsUserStatus = 2002 // 前端请求判断用户是否在线

)
