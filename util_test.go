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

	rounded = util.Round(1.23454, 4)
	assert.Equal(t, rounded, 1.2345)

	rounded = util.Round(1.23454, 5)
	assert.Equal(t, rounded, 1.23454)
}

//revive:disable:empty-lines
func TestSlices(t *testing.T) {
	assert := assert.New(t)
	t.Run("Map function", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		out := slice.Map(in, func(i int) int { return i * i })
		assert.Equal(out, []int{1, 4, 9, 16})
	})

	t.Run("Fold function", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		out := slice.Fold(in, 0, func(x int, y int) int { return x + y })
		assert.Equal(out, 10)
	})

	t.Run("Deduplicate function", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 1, 2, 3, 4}
		out := slice.Deduplicate(in)
		assert.Equal(out, []int{1, 2, 3, 4})
	})

	t.Run("Remove function", func(t *testing.T) {
		in := []int{1, 1, 1, 2, 3, 4, 1, 1, 1}
		out := slice.Remove(in, 1)
		assert.Equal(out, []int{1, 1, 2, 3, 4, 1, 1, 1})
	})

	t.Run("RemoveAll function", func(t *testing.T) {
		in := []int{1, 1, 1, 2, 3, 4, 1, 1, 1}
		out := slice.RemoveAll(in, 1)
		assert.Equal(out, []int{2, 3, 4})
	})
}
