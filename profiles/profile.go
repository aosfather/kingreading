package profiles

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

type RateType byte

const (
	RT_HOUR     RateType = 10 //每小时
	RT_DAY      RateType = 11 //每天
	RT_WEEK_MON RateType = 21 //每周一
	RT_WEEK_TUE RateType = 22 //每周二
	RT_WEEK_WED RateType = 23 //每周三
	RT_WEEK_THU RateType = 24 //每周四
	RT_WEEK_FRI RateType = 25 //每周五
	RT_WEEK_SAT RateType = 26 //每周六
	RT_WEEK_SUN RateType = 27 //每周日

)

//推送设置
type Profile struct {
	RemoteType    string            //终端类型
	Caption       string            //内容标题名称，唯一ID
	Rate          RateType          //频率
	LastHour      int               //最晚发出小时
	LastMinutes   int               //最晚发出的分钟
	EmptyEnabled  bool              //空内容是否也发送
	MaxLimit      int               //最大章节限制
	LastSendIndex int               //最后推送的章节索引
	ExtProperties map[string]string //扩展属性
}

/**
  是否满足触发条件
  1、到时间
  2、对应的内容有更新（新的内容）
*/
func (this *Profile) OnTrigger() bool {
	now := time.Now()
	switch this.Rate {
	case RT_HOUR:
		if now.Minute() == this.LastMinutes {
			return true
		}
	case RT_DAY:
		if this.timeOnTrigger(now) {
			return true
		}

	default: //按周的
		if time.Weekday(this.Rate-20) == now.Weekday() && this.timeOnTrigger(now) {
			return true
		}

	}
	return false
}

func (this *Profile) timeOnTrigger(now time.Time) bool {
	if now.Hour() == this.LastHour && now.Minute() == this.LastMinutes {
		return true
	}

	return false
}

func (this *Profile) GetProperty(key string) string {
	if this.ExtProperties != nil {
		return this.ExtProperties[key]
	}

	return ""
}

func (this *Profile) Load(r io.Reader) {
	if r != nil {
		data, e := ioutil.ReadAll(r)
		if e == nil {
			json.Unmarshal(data, this)
		}
	}
}

func (this *Profile) Save(w io.Writer) {
	data, e := json.Marshal(this)
	if e == nil {
		if w != nil {
			w.Write(data)
		}
	}
}
