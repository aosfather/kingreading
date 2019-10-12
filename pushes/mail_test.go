package pushes

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("smtp.qq.com", 465, "xiongxiaopeng@ehomepay.com.cn", "aosfather@kindle.cn", "书", "123456", "test", "")
}
