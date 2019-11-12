package main

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestMethod(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAPI := NewMockSome(mockCtrl)

	expected := "TEST"

	gomock.InOrder(
		mockAPI.EXPECT().Method("bob").
			Return(expected).
			Times(1),
	)

	actual := mockAPI.Method("bob")

	if actual != expected {
		t.Errorf("expected != actual")
	}
}
