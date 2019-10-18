package pushes

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("smtp.163.com", 25, "?@163.com", "?@kindle.cn", "书", "?", "万古神帝第2446章.txt", "d:/D/2436.txt")
}

func TestSmtpClient_SendMessage(t *testing.T) {
	sc := SmtpClient{}
	sc.Host = "smtp.163.com"
	sc.User = "?@163.com"
	sc.Pwd = "?"
	sc.Port = 25

	msg := EmailMessage{}
	msg.From = sc.User
	msg.To = []string{"?"}
	msg.Subject = "网络书"
	msg.Body = "万古神帝第2445章"
	msg.AddAttachment(Attachment{CT_TEXT, "万古神帝第2445章.txt", "d:/D/2435.txt"})

	sc.SendMessage(&msg)

}
