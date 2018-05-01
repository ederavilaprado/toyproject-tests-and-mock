package main

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestTwoIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	g := Goblin(t)

	g.Describe("Numbers xxx", func() {
		g.It("Should add two numbers ", func() {
			g.Assert(1 + 1).Equal(2)
		})

		g.Xit("Should do something usefull here", func() {
			g.Assert(1 + 1).Equal(2)
		})
	})
}
