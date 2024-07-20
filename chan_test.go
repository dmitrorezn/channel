package channel

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleChan(t *testing.T) {
	var chans = make([]chan int, 10)
	for i := range chans {
		chans[i] = make(chan int, 1)
	}

	done := SingleChan(chans...)

	i := rand.Intn(len(chans))
	chans[i] <- i

	select {
	case v := <-done:
		assert.Equal(t, i, v)
	}
}
