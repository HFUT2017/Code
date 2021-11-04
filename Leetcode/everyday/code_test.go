package everyday

import "testing"

func TestByteDance(t *testing.T) {
	for i := float64(1); i <= 100; i++ {
		println(ByteDance(i))
	}
}
