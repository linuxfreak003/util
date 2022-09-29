package util_test

import (
	"testing"

	"github.com/linuxfreak003/util"
	"github.com/linuxfreak003/util/slice"
	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	rounded := util.Round(1.23454, 3)
	assert.Equal(t, rounded, 1.235)
}

func TestSlices(t *testing.T) {
	assert := assert.New(t)
	t.Run("Map function", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		out := slice.Map(in, func(i int) int { return i * i })
		assert.Equal(out, []int{1, 4, 9, 16})
	})
}
