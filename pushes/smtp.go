package pushes

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/smtp"
	"strings"
	"time"
)

type ContentType string

func (this ContentType) ToString() string {
	return "Content-Type: text/" + fmt.Sprintf("%s", this) + "; charset=UTF-8"
}

const (
	CT_HTML ContentType = "html"
	CT_TEXT ContentType = "plain"
)

type SmtpClient struct {
	Host string
	User string
	Pwd  string
	Port int
}

func (this *SmtpClient) SendMessage(m *EmailMessage) bool {
	auth := smtp.PlainAuth("", this.User, this.Pwd, this.Host)
	e := smtp.SendMail(fmt.Sprintf("%s:%d", this.Host, this.Port), auth, m.From, m.To, m.ToByte())
	if e != nil {
		log.Println(e)
		return false
	}
	return true
}

type EmailMessage struct {
	Type    ContentType
	Charset string //字符编码
	From    string
	To      []string //接收者
	Subject string   //标题
	Body    string   //邮件内容
	attachs []Attachment
}

func (this *EmailMessage) toHeader(boundary string) map[string]string {
	Header := make(map[string]string)
	Header["From"] = this.From
	Header["To"] = strings.Join(this.To, ";")
	//Header["Cc"] = strings.Join(message.cc, ";")
	//Header["Bcc"] = strings.Join(message.bcc, ";")
	Header["Subject"] = this.Subject
	Header["Content-Type"] = "multipart/mixed;boundary=" + boundary
	Header["Date"] = time.Now().String()
	return Header

}

func (this *EmailMessage) AddAttachment(a Attachment) {
	if a.FileName != "" {
		this.attachs = append(this.attachs, a)
	}

}

func (this *EmailMessage) ToByte() []byte {
	buffer := bytes.NewBuffer(nil)
	boundary := "GoBoundary"
	//写header头 key：value的形式
	headers := this.toHeader(boundary)
	header := ""
	for key, value := range headers {
		header += key + ":" + value + "\r\n"
	}
	header += "\r\n"
	buffer.WriteString(header)

	//写正文内容
	body := "\r\n--" + boundary + "\r\n"
	body += "Content-Transfer-Encoding:base64\r\n"
	body += "Content-Type: text/html; charset=UTF-8 \r\n"

	buffer.WriteString(body)
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(this.Body)))
	base64.StdEncoding.Encode(payload, []byte(this.Body))

	buffer.WriteString("\r\n")

	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}

	//写附件
	for _, a := range this.attachs {
		a.Write(boundary, buffer)
	}

	return buffer.Bytes()
}

type Attachment struct {
	Type     ContentType
	Name     string
	FileName string
}

func (this *Attachment) Write(boundary string, buffer *bytes.Buffer) {
	filename := mime.BEncoding.Encode("utf-8", this.Name)
	attachment := "\r\n--" + boundary + "\r\n"
	attachment += "Content-Disposition: attachment;filename=\"" + filename + "\"\r\n"
	attachment += "Content-Transfer-Encoding:base64\r\n"
	attachment += this.Type.ToString() + ";name=\"" + filename + "\"\r\n"
	//attachment += "Content-ID: <" + this.Name + "> \r\n\r\n"
	buffer.WriteString(attachment)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(err.(string))
		}
	}()

	this.writeFile(buffer)

}
func (this *Attachment) writeFile(buffer *bytes.Buffer) {
	file, err := ioutil.ReadFile(this.FileName)
	if err != nil {
		panic(err.Error())
	}
	payload := make([]byte, base64.StdEncoding.EncodedLen(len(file)))
	base64.StdEncoding.Encode(payload, file)
	buffer.WriteString("\r\n")
	for index, line := 0, len(payload); index < line; index++ {
		buffer.WriteByte(payload[index])
		if (index+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}
}
