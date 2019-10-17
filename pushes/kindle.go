package pushes

import (
	"fmt"
	"github.com/aosfather/bingo"
	"github.com/aosfather/kingreading/collect"
	"github.com/aosfather/kingreading/profiles"
	"log"
	"strconv"
	"strings"
)

//kindle
type KindlePusher struct {
	pm         profiles.ProfileManager
	captions   collect.CaptionManager
	port       int
	host       string
	account    string //mail 账户
	accountpwd string //发送密码
	client     *SmtpClient
}

func (this *KindlePusher) Load(config *bingo.ApplicationContext) {
	this.host = config.GetPropertyFromConfig("kindle.host")
	this.account = config.GetPropertyFromConfig("kindle.account")
	this.accountpwd = config.GetPropertyFromConfig("kindle.pwd") //解密
	this.port = 25
	if config.GetPropertyFromConfig("kindle.port") != "" {
		p, e := strconv.Atoi(config.GetPropertyFromConfig("kindle.port"))
		if e == nil {
			this.port = p
		}
	}

	this.client = &SmtpClient{}
	this.client.Port = this.port
	this.client.Host = this.host
	this.client.User = this.account
	this.client.Pwd = this.accountpwd

	//注册到管理器中
	pm := config.GetService("pushers").(*PusherManager)
	pm.Add("kindle", this)
	this.pm = config.GetService("profiles").(profiles.ProfileManager)
	this.captions = config.GetService("collects").(collect.CaptionManager)

}

//推送
func (this *KindlePusher) Execute(profile *profiles.Profile) {
	// 检查是否有更新
	if profile.OnTrigger() {
		if this.captions == nil {
			log.Println("caption manager is nil!")
			return
		}
		//检查内容主题是否存在
		caption := this.captions.GetCaption(profile.Caption)
		if caption == nil {
			log.Println("caption is not exits! ", profile.Caption)
			return
		}

		//当有新的内容
		if profile.LastSendIndex < caption.Index {
			//检查是否有更新，如果有根据索引和maxlimit进行整理文件列表。
			max := caption.Index - profile.LastSendIndex
			if max > profile.MaxLimit {
				max = profile.MaxLimit
			}

			//获取文件列表
			var files []string
			var body string

			for i := 1; i <= max; i++ {
				files = append(files, this.captions.GetFileName(caption.Name, profile.LastSendIndex+i))
				body += fmt.Sprintf("%s_%d.%s;", caption.Title, profile.LastSendIndex+i, caption.Fix)
			}

			log.Println("kindle use send mail ")
			// 发送到目标地址
			if this.sendmail(profile.GetProperty("EMAIL"), profile.Title, body, files...) {
				profile.LastSendIndex += max
				log.Println("kindle  send mail success")
				//保存
				this.pm.Save(profile)
			}

		}

	}

}

func (this *KindlePusher) sendmail(to string, title string, body string, files ...string) bool {
	if to == "" || strings.Index(to, "@") <= 0 {
		return false //email格式错误
	}

	if title == "" {
		title = "书"
	}

	msg := EmailMessage{}
	msg.To = []string{to}
	msg.From = this.account
	msg.Subject = title
	msg.Body = body
	names := strings.Split(body, ";")
	for index, f := range files {
		msg.AddAttachment(Attachment{CT_TEXT, names[index], f})
	}

	return this.client.SendMessage(&msg)
}
