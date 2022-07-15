package test

import (
	"fmt"
	"github.com/BANKEX/plasma-research/src/node/plasmautils/plasmacrypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestRSAAccumulator(t *testing.T) {
	// TODO: test package not working

	a := new(plasmacrypto.Accumulator).SetInt(big.NewInt(17))
	a.Accumulate(19)
	res, _ := new(big.Int).SetString("239072435685151324847153", 10)
	fmt.Println("", res)
	a.Value().Cmp(res)
	fmt.Println("", res)

	assert.Equal(t, a.Value(), res, "Right accumulator value")

	a = new(plasmacrypto.Accumulator).SetInt(big.NewInt(17))
	a.BatchAccumulate([]uint32{3, 5})
	res, _ = new(big.Int).SetString("2862423051509815793", 10)
	fmt.Println("", res)
	assert.Equal(t, a.Value(), res, "Right accumulator value")
}
