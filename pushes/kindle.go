package pushes

import "github.com/aosfather/kingreading/profiles"

//kindle
type KindlePusher struct {
}

//推送
func (this *KindlePusher) Push(profile *profiles.Profile) {
	// 检查是否有更新
	// 根据更新内容打包
	// 发送到目标地址
}
