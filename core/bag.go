package core

/**
  打包
*/
type BagMan struct {
	Size   int    //包的大小
	OutDir string //输出目录
}

type TxtBagMan struct {
	BagMan
}

func (this *TxtBagMan) makeBag(filename []string) []string {

}
