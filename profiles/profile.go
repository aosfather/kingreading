package profiles

type RateType byte

const (
	RT_DAY      RateType = 11
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
	RemoteType   string   //终端类型
	Caption      string   //内容标题
	Rate         RateType //频率
	LastTime     string   //最晚发出时间
	EmptyEnabled bool     //空内容是否也发送
}

/**
  是否满足触发条件
  1、到时间
  2、对应的内容有更新（新的内容）
*/
func (this *Profile) OnTrigger() bool {

	return false
}
