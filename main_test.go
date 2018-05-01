package main

import (
	"errors"
	"flag"
	"fmt"
	"testing"

	"github.com/Smarp/funcmock"
	"github.com/ederavilaprado/toyproject-tests-and-mock/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update golden files")

func TestMyFunctionExpec(t *testing.T) {
	assert.Equal(t, 70, myFunction(50, 20))
}

func TestMyFunctionUse(t *testing.T) {
	c := funcmock.Mock(&myFunc)
	c.SetDefaultReturn(30)
	defer c.Restore()

	fmt.Printf("=> update %+v\n", *update)

	assert.Equal(t, "=> 30", myFunctionUse())
	assert.Equal(t, 1, c.CallCount())
}

func TestRunExampleFunction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExample := mock.NewMockExample(ctrl)
	mockExample.EXPECT().Set("my_key", gomock.Any()).Return(nil).Times(2)
	mockExample.EXPECT().Get("my_key").Return("explicit_value", nil).Times(1)
	assert.Equal(t, "explicit_value", runExampleFunction(mockExample))

	mockExample.EXPECT().Get("my_key").Return("", errors.New("Error"))
	assert.Equal(t, "", runExampleFunction(mockExample))

	mockExample.EXPECT().Set(gomock.Any(), gomock.Any()).Return(errors.New("error"))
	assert.Equal(t, "", runExampleFunction(mockExample))

}
