package core

import "testing"

func TestTxtBagMan_MakeBag(t *testing.T) {
	bag := TxtBagMan{}
	bag.Size = 5
	bag.OutDir = "D:/D/out"
	f := bag.MakeBag([]string{"D:/D/xs_wgsd/2850.txt", "D:/D/xs_wgsd/2851.txt", "D:/D/xs_wgsd/2852.txt", "D:/D/xs_wgsd/2853.txt", "D:/D/xs_wgsd/2854.txt", "D:/D/xs_wgsd/2855.txt", "D:/D/xs_wgsd/2856.txt", "D:/D/xs_wgsd/2857.txt"})
	t.Log(f)
}
