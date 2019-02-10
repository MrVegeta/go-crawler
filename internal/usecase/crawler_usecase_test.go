package usecase

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"

	"../mock"
)

var (
	mockCrawlerCtrl *mock.MockCrawlerCtrl
	cu *crawlerUseCase
)


func setupTestCrawlerUseCase(t *testing.T) func(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockCrawlerCtrl = mock.NewMockCrawlerCtrl(mockCtrl)
	cu = &crawlerUseCase{mockCrawlerCtrl}
	return func(t *testing.T) {
		cu = nil
		mockCrawlerCtrl = nil
		mockCtrl.Finish()
	}
}

func TestNewCrawlerUseCase(t *testing.T) {
	tearDown := setupTestCrawlerUseCase(t)
	if !reflect.DeepEqual(cu, NewCrawlerUseCase(mockCrawlerCtrl)) {
		t.Fatal("CrawlerNodeUseCase.NewCrawlerUseCase() returns false")
	}
	tearDown(t)
}

func TestCrawlerUseCase_Start(t *testing.T) {
	tests := []struct {
		name string
		expectedErr error
	} {
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tearDown := setupTestCrawlerUseCase(t)
			tearDown(t)
		})
	}
}

func TestCrawlerUseCase_Stop(t *testing.T) {
	tests := []struct {
		name string
		expectedErr error
	} {
		{},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tearDown := setupTestCrawlerUseCase(t)

			tearDown(t)
		})
	}
}