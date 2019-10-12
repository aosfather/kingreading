package pushes

import (
	"github.com/go-gomail/gomail"
	"log"
)

func SendMail(host string, port int, from, to, subject, pwd string, body string, filename string) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", from /*"发件人地址"*/, "发件人") // 发件人
	m.SetHeader("To", m.FormatAddress(to, "收件人"))       // 收件人
	m.SetHeader("Subject", subject)                     // 主题

	//m.SetBody("text/html",xxxxx ") // 可以放html..还有其他的
	m.SetBody("text/plain", body) // 正文
	if filename != "" {
		m.Attach(filename) //添加附件
	}

	d := gomail.NewDialer(host, port, from, pwd) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Println("发送失败", err.Error())
		return
	}

	log.Println("done.发送成功")

}
