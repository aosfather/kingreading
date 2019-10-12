package collect

import (
	"fmt"
	"testing"
)

func TestWuxianSpider_GrabCaption1(t *testing.T) {
	sp := WuxianSpider{}
	fmt.Println(sp.grabIndex("d:/D/index.html"))
}

func TestWuxianSpider_GrabCaption2(t *testing.T) {
	sp := WuxianSpider{}
	sp.grabContent(1, "d:/D/5380339.html")
}

func TestWuxianSpider_GrabCaptionFromUrl(t *testing.T) {
	sp := WuxianSpider{}
	sp.grabContent(1, "http://www.wuxianxs.com/ls10-10050/3026486.html")
}

func TestWuxianSpider_GrabCaption(t *testing.T) {
	sp := WuxianSpider{}
	sp.CaptionPath = "d:/D"
	c := Caption{}

	c.Index = 2436
	c.Url = "http://www.wuxianxs.com/ls10-10050/"
	c.Enabled = true
	sp.GrabCaption(&c)
}
