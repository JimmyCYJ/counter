package counter

import (
	"sync/atomic"
)

var count uint64

// Count safely gets a unique count number.
func Count() uint64 {
	atomic.AddUint64(&count, 1)
	return count
}

// Count2 safely gets a unique count number.
func Count2() uint64 {
	atomic.AddUint64(&count, 2)
	return count
}

func SetCount(c uint64) uint64 {
	count = atomic.LoadUint64(&c)
	return count
}
