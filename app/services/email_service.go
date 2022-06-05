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
	"net"
	"net/smtp"
)

var (
	host     = config.Conf.Mail.Host
	name     = config.Conf.Mail.Name
	password = config.Conf.Mail.Password
	port     = config.Conf.Mail.Port
)

type EmailServiceInterface interface {
	// 发送邮件方法
	SendEmail(to string, subject string, body string) error
	// 获取html模版内容
	GetHtmlTemplate(text string) []byte
}

type EmailService struct{}

func (s EmailService) SendEmail(to string, subject string, body string) error {

	header := make(map[string]string)

	header["From"] = "GO-IM:" + "<" + name + ">"
	header["To"] = to
	header["Subject"] = subject
	header["Content-Type"] = "text/html;chartset=UTF-8"

	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	message += "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		name,
		password,
		host,
	)
	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		name,
		[]string{to},
		[]byte(message),
	)

	if err != nil {
		panic(err)
	}
	return err
}

func (m EmailService) GetHtmlTemplate(text string) []byte {
	return []byte(text)
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
