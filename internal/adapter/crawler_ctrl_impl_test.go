package adapter

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

var (
	cc *crawlerCtrl
)

func setupTestCrawlerCtrl(t *testing.T) func(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	cc = &crawlerCtrl{}
	return func(t *testing.T) {
		cc = nil
		mockCtrl.Finish()
	}
}

func TestNewCrawlerCtrl(t *testing.T) {
	tearDown := setupTestCrawlerCtrl(t)
	if !reflect.DeepEqual(cc, NewCrawlerCtrl()) {
		t.Fatalf("CrawlerCtrl.NewCrawlerCtrl() returns false")
	}
	tearDown(t)
}