package pushes

import "strings"

func SendMail(host string, port int, from, to, subject, pwd string, body string, filename ...string) {
	sc := SmtpClient{}
	sc.Host = host
	sc.User = from
	sc.Pwd = pwd
	sc.Port = port

	msg := EmailMessage{}
	msg.From = sc.User
	msg.To = []string{to}
	msg.Subject = subject
	msg.Body = body

	names := strings.Split(body, ";")
	for index, f := range filename {
		msg.AddAttachment(Attachment{CT_TEXT, names[index], f})
	}

	sc.SendMessage(&msg)
}
