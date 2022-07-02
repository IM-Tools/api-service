/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package services

import (
	"crypto/tls"
	"fmt"
	"im-services/config"
	"im-services/pkg/logger"
	"im-services/pkg/redis"
	"net"
	"net/smtp"
	"time"
)

const (
	REGISTERED_CODE = 1 // 注册验证码
	RESET_PS_CODE   = 2 // 重置密码
)

type EmailServiceInterface interface {
	// 发送邮件方法
	SendEmail(code string, emailType int, email string, subject string, body string) error
	// 获取html模版内容
	GetHtmlTemplate(text string) []byte
	// 获取缓存key
	getCacheFix(email string, emailType int) string

	CheckCode(email string, code string, emailType int) bool
}

type EmailService struct{}

// 发送邮件方法
// code 验证码 emailType 邮件类型  email 发送邮箱 subject 主题 body 内容
func (s EmailService) SendEmail(code string, emailType int, email string, subject string, body string) error {

	header := make(map[string]string)

	header["From"] = "im-service:" + "<" + config.Conf.Mail.Name + ">"
	header["To"] = email
	header["Subject"] = subject
	header["Content-Type"] = "text/html;chartset=UTF-8"

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	message += "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		config.Conf.Mail.Name,
		config.Conf.Mail.Password,
		config.Conf.Mail.Host,
	)

	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", config.Conf.Mail.Host, config.Conf.Mail.Port),
		auth,
		config.Conf.Mail.Name,
		[]string{email},
		[]byte(message),
	)

	if err != nil {
		return err
	}
	redis.RedisDB.Set(s.getCacheFix(email, emailType), code, time.Minute*5)
	return nil
}

// 获取缓存key
func (s EmailService) getCacheFix(email string, emailType int) string {
	switch emailType {
	case REGISTERED_CODE:
		return fmt.Sprintf("%s.%d", email, REGISTERED_CODE)
	case RESET_PS_CODE:
		return fmt.Sprintf("%s.%d", email, RESET_PS_CODE)
	default:
		return fmt.Sprintf("%s.%d", email, REGISTERED_CODE)
	}
}

func (s EmailService) GetHtmlTemplate(text string) []byte {
	return []byte(text)
}

// 检查邮件是否正确
func (s EmailService) CheckCode(email string, code string, emailType int) bool {
	cacheFix := s.getCacheFix(email, emailType)

	redisCmd := redis.RedisDB.Get(cacheFix)
	val, _ := redisCmd.Result()
	if val != code {
		return false
	}

	return true
}

//return a smtp client
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		logger.Logger.Error("Dialing Error")
		return nil, err
	}
	//分解主机端口字符串
	hosts, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, hosts)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送

func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := dial(addr)
	if err != nil {
		logger.Logger.Error("Create smpt client error:")
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				logger.Logger.Error("Error during AUTH")
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addrs := range to {
		if err = c.Rcpt(addrs); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
