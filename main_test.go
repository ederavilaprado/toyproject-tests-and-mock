package main

import (
	"testing"

	"github.com/Smarp/funcmock"
	"github.com/stretchr/testify/assert"
)

func TestMyFunctionUse(t *testing.T) {
	// myFunction = func(a, b int) int {
	// 	return 10
	// }

	c := funcmock.Mock(&myFunc)
	c.SetDefaultReturn(30)
	defer c.Restore()

	assert.Equal(t, "=> 30", myFunctionUse())

	// fmt.Printf("=> %+v\n", c.CallCount())

}
