package util_test

import (
	"testing"
	"time"

	"github.com/linuxfreak003/util/cache"
	"github.com/linuxfreak003/util/maps"
	"github.com/linuxfreak003/util/number"
	"github.com/linuxfreak003/util/slice"
	"github.com/stretchr/testify/assert"
)

func TestUtils(t *testing.T) {
	assert := assert.New(t)
	t.Run("Round", func(t *testing.T) {
		rounded := number.Round(1.23454, 3)
		assert.Equal(rounded, 1.235)

		rounded = number.Round(1.23454, 4)
		assert.Equal(rounded, 1.2345)

		rounded = number.Round(1.23454, 5)
		assert.Equal(rounded, 1.23454)
	})

	t.Run("Sum", func(t *testing.T) {
		sum := number.Sum(1, 2, 3, 4)
		assert.Equal(sum, 10)

		sumS := number.Sum("a", "b", "c")
		assert.Equal(sumS, "abc")
	})

	t.Run("Min", func(t *testing.T) {
		m := number.Min(1, 2, 3, 3, 4, 5)
		assert.Equal(m, 1)
	})

	t.Run("Max", func(t *testing.T) {
		m := number.Max(1, 2, 3, 3, 4, 5)
		assert.Equal(m, 5)
	})
}

//revive:disable:empty-lines
func TestSlices(t *testing.T) {
	assert := assert.New(t)
	t.Run("Map", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		out := slice.Map(in, func(i int) int { return i * i })
		assert.Equal(out, []int{1, 4, 9, 16})
	})

	t.Run("ToMap", func(t *testing.T) {
		type MyType struct {
			Id        int64
			FirstName string
			LastName  string
		}
		in := []MyType{
			{Id: 1, FirstName: "Bob", LastName: "Ross"},
			{Id: 2, FirstName: "Bob", LastName: "Marley"},
		}
		out := slice.ToMap(in, func(x MyType) (int64, string) {
			return x.Id, x.FirstName + " " + x.LastName
		})
		assert.Len(out, 2)
		assert.Equal(out[1], "Bob Ross")
		assert.Equal(out[2], "Bob Marley")
	})

	t.Run("Reduce/Fold", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		out := slice.Reduce(in, func(x int, y int) int { return x + y }, 0)
		assert.Equal(out, 10)
	})

	t.Run("Deduplicate", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 1, 2, 3, 4}
		out := slice.Deduplicate(in)
		assert.Equal(out, []int{1, 2, 3, 4})
	})

	t.Run("Remove", func(t *testing.T) {
		in := []int{1, 1, 1, 2, 3, 4, 1, 1, 1}
		out := slice.Remove(in, 1)
		assert.Equal(out, []int{1, 1, 2, 3, 4, 1, 1, 1})
	})

	t.Run("RemoveAll", func(t *testing.T) {
		in := []int{1, 1, 1, 2, 3, 4, 1, 1, 1}
		out := slice.RemoveAll(in, 1)
		assert.Equal(out, []int{2, 3, 4})
	})

	t.Run("Reverse", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		out := slice.Reverse(in)
		assert.Equal(out, []int{5, 4, 3, 2, 1})

		in = []int{1, 2, 3, 4}
		out = slice.Reverse(in)
		assert.Equal(out, []int{4, 3, 2, 1})

		in = []int{1, 2}
		out = slice.Reverse(in)
		assert.Equal(out, []int{2, 1})
	})

	t.Run("Intersect", func(t *testing.T) {
		in1 := []int{1, 2, 3, 4}
		in2 := []int{3, 4, 5, 6}
		out := slice.Intersect(in1, in2)
		assert.Equal(out, []int{3, 4})
	})

	t.Run("Index", func(t *testing.T) {
		in := []int{1, 2, 3, 4}
		i := slice.Index(in, 3)
		assert.Equal(i, 2)
	})
}

func TestMaps(t *testing.T) {
	assert := assert.New(t)
	t.Run("ToSlice", func(t *testing.T) {
		in := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		out := maps.ToSlice(in)
		assert.Len(out, 4)
		for k, v := range []maps.Pair[string, int]{
			maps.Pair[string, int]{Key: "a", Value: 1},
			maps.Pair[string, int]{Key: "b", Value: 2},
			maps.Pair[string, int]{Key: "c", Value: 3},
			maps.Pair[string, int]{Key: "d", Value: 4},
		} {
			assert.Equal(out[k], v)
		}
	})
}

func TestSort(t *testing.T) {
	assert := assert.New(t)
	t.Run("Sort", func(t *testing.T) {
		in := []int{1, 4, 3, 2}
		out := slice.Sort(in)
		assert.Equal(out, []int{1, 2, 3, 4})
	})

	t.Run("DumbSort", func(t *testing.T) {
		in := []int{1, 4, 3, 2}
		out := slice.DumbSort(in)
		assert.Equal(out, []int{1, 2, 3, 4})
	})
}

func TestCache(t *testing.T) {
	assert := assert.New(t)
	t.Run("Use cache", func(t *testing.T) {
		c := cache.New("", "")
		worked := c.Set("a", "b", time.Hour)
		assert.True(worked)
		v, b := c.Get("a")
		assert.True(b)
		assert.Equal(v, "b")
	})
	t.Run("Use cache with no duration", func(t *testing.T) {
		c := cache.New("", "")
		worked := c.Set("a", "b", 0)
		assert.True(worked)
		v, b := c.Get("a")
		assert.True(b)
		assert.Equal(v, "b")
	})
}
