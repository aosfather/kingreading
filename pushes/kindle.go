package pushes

import (
	"fmt"
	"github.com/aosfather/kingreading/collect"
	"github.com/aosfather/kingreading/core"
	"github.com/aosfather/kingreading/profiles"
	"log"
	"strconv"
	"strings"
)

//kindle
type KindlePusher struct {
	Pm         *PusherManager         `Inject:""`
	Captions   collect.CaptionManager `Inject:""`
	PortStr    string                 `Value:"kindle.port"`
	Out        string                 `Value:"out.dir"`
	Size       int                    `Value:"out.size"`
	port       int
	Host       string `Value:"kindle.host"`
	Account    string `Value:"kindle.account"` //mail 账户
	Accountpwd string `Value:"kindle.pwd"`     //发送密码
	client     *SmtpClient
	bagman     *core.TxtBagMan //打包文件
}

func (this *KindlePusher) Init() {
	this.port = 25
	if this.PortStr != "" {
		p, e := strconv.Atoi(this.PortStr)
		if e == nil {
			this.port = p
		}
	}

	this.client = &SmtpClient{Port: this.port, Host: this.Host, User: this.Account, Pwd: this.Accountpwd}
	//注册到管理器中
	this.Pm.Add("kindle", this)
	this.bagman = new(core.TxtBagMan)
	if this.Size == 0 {
		this.Size = 1
	}
	this.bagman.Size = this.Size
	this.bagman.OutDir = this.Out

}

//推送
func (this *KindlePusher) Execute(profile *profiles.Profile) {
	// 检查是否有更新
	if profile.OnTrigger() {
		if this.Captions == nil {
			log.Println("caption manager is nil!")
			return
		}
		//检查内容主题是否存在
		caption := this.Captions.GetCaption(profile.Caption)
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
				files = append(files, this.Captions.GetFileName(caption.Name, profile.LastSendIndex+i))

			}
			log.Println(files)
			//打包
			outfiles := this.bagman.MakeBag(files)
			for _, fname := range outfiles {
				body += fmt.Sprintf("%s_%s;", caption.Title, fname[strings.LastIndex(fname, "/"):])
			}

			log.Println("kindle use send mail ")
			// 发送到目标地址
			if this.sendmail(profile.GetProperty("EMAIL"), profile.Title, body, outfiles...) {
				profile.LastSendIndex += max
				log.Println("kindle  send mail success")
				//保存
				profile.ProfileMan.Save(profile)
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
	msg.From = this.Account
	msg.Subject = title
	msg.Body = body
	names := strings.Split(body, ";")
	for index, f := range files {
		msg.AddAttachment(Attachment{CT_TEXT, names[index], f})
	}

	return this.client.SendMessage(&msg)
}
