package services

import (
	"fmt"
	"im-services/internal/api/requests"
	"im-services/internal/config"
	"im-services/internal/dao/messsage_dao"
	"im-services/internal/enum"
	"im-services/internal/helpers"
	"im-services/internal/models/user"
	"im-services/pkg/date"
	"im-services/pkg/hash"
	"im-services/pkg/model"
	"strings"
	"sync"
)

var (
	messagesServices ImMessageService
	botData          = map[string]string{} // 存储指令
	lock             sync.RWMutex
	BOT_NOE          = 1
	messageDao       messsage_dao.MessageDao
)

// 初始化机器人信息数据
func InitChatBot() {
	var count int64
	model.DB.Table("im_users").Where("id=?", BOT_NOE).Count(&count)
	if count == 0 {
		createdAt := date.NewDate()
		model.DB.Table("im_users").Create(&user.ImUsers{
			ID:            int64(BOT_NOE),
			Email:         config.Conf.GoBot.Email,
			Password:      hash.BcryptHash(config.Conf.GoBot.Password),
			Name:          config.Conf.GoBot.Name,
			CreatedAt:     createdAt,
			UpdatedAt:     createdAt,
			Avatar:        config.Conf.GoBot.Avatar,
			LastLoginTime: createdAt,
			Uid:           helpers.GetUuid(),
			UserJson:      "{}",
			UserType:      1,
		})
	}
}

func GetMessage(key string) string {

	if strings.Contains(key, ":") {
		arr := strings.Split(key, ":")
		if len(arr) == 2 {
			lock.Lock()
			botData[arr[0]] = arr[1]
			lock.Unlock()
			return "很不错就是这样~"
		}
		if len(arr) > 2 {
			return "格式不对呀~"
		}
	}

	if value, ok := botData[key]; ok {
		return value
	} else {
		return "没明白您的意思-暂时还不知道说啥~~~ 你可以通过 xxx:xxx 指令定义消息😊"
	}
}

func InitChatBotMessage(formID int64, toID int64) {

	params := requests.PrivateMessageRequest{
		MsgId:       date.TimeUnixNano(),
		MsgCode:     enum.WsChantMessage,
		MsgClientId: date.TimeUnixNano(),
		FormID:      formID,
		ToID:        toID,
		ChannelType: 1,
		MsgType:     1,
		Message:     fmt.Sprintf("您好呀~ 我是%s~🥰", config.Conf.GoBot.Name),
		SendTime:    date.NewDate(),
		Data:        "",
	}
	messageDao.CreateMessage(params)
	messagesServices.SendPrivateMessage(params)
	params.Message = "我们来玩个游戏吧！你问我答~！👋"
	messageDao.CreateMessage(params)
	messagesServices.SendPrivateMessage(params)
}
