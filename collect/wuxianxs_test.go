package collect

import "testing"

func TestWuxianSpider_GrabCaption(t *testing.T) {
	sp := WuxianSpider{}
	sp.grabIndex("d:/D/index.html")
}

func TestWuxianSpider_GrabCaption2(t *testing.T) {
	sp := WuxianSpider{}
	sp.grabContent("d:/D/3027586.html")
}
