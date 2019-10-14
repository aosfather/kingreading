package pushes

import (
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

}

func (this *KindlePusher) Load(config map[string]string) {
	this.host = config["kindle.host"]
	this.account = config["kindle.account"]
	this.accountpwd = config["kindle.pwd"] //解密
	this.port = 25
	if config["kindle.port"] != "" {
		p, e := strconv.Atoi(config["kindle.port"])
		if e == nil {
			this.port = p
		}
	}

}

//推送
func (this *KindlePusher) Push(profile *profiles.Profile) {
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
			for i := 1; i <= max; i++ {
				files = append(files, this.captions.GetFileName(caption.Name, profile.LastSendIndex+i))
			}

			// 发送到目标地址
			if this.sendmail(profile.GetProperty("EMAIL"), profile.Title, "", files...) {
				profile.LastSendIndex += max
			}

			//保存
			this.pm.Save(profile)

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

	SendMail(this.host, this.port, this.account, to, title, this.accountpwd, body, files...)
	return true
}
