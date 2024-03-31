package demo2

import (
	"testing"

	"demo2/spider"

	"github.com/golang/mock/gomock"
)

func TestGetGoVersion(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockSpider := spider.NewMockSpider(mockCtl)
	mockSpider.EXPECT().GetBody().Return("go1.13.5")
	goVer := GetGoVersion(mockSpider)

	if goVer != "go1.13.5" {
		t.Errorf("Get wrong version %s", goVer)
	}
}
