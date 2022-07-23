/**
  @author:panliang
  @data:2022/7/2
  @note
**/
package enum

const (
	PARAMS_ERROR = 1000 // 参数错误
	API_ERROR    = 1001 // 接口异常 比如调用第三方接口 或者代码异常

	// ws 消息
	WS_SUCCESS      = 200  // 聊天消息
	WS_CREATE       = 1000 // 添加好友
	WS_FRIEND_OK    = 1001 // 同意好友申请
	WS_FRIEND_ERROR = 1002 // 拒绝好友申请
	WS_NOT_FRIEND   = 1003 // 非好友关系
	WS_PING         = 1004 // 心跳

	WS_USER_OFFLINE   = 2000 // 用户离线
	WS_USER_ONLINE    = 2001 // 用户在线
	WS_IS_USER_STATUS = 2002 // 前端请求判断用户是否在线
)
