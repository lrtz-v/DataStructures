package profiling

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

func TestTiming(T *testing.T) {
	x := new(big.Int).SetUint64(uint64(1000))
	// Arguments to a defer statement is immediately evaluated and stored.
	// The deferred function receives the pre-evaluated values when its invoked.
	defer Duration(time.Now(), "IntFactorial")

	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}

	x.Set(y)
	fmt.Println(x)
}
