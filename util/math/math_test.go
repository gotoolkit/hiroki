package math

import (
	"math/big"
	"testing"
)

func TestCombination(t *testing.T) {
	iterable := make([]int, 33)
	for i := 0; i < 33; i++ {
		iterable[i] = i + 1
	}
	result := Combination(iterable, 6)

	if len(result) != 1107568 {
		t.Errorf("1-33选出6个数字全组合为%d,计算结果为%d", 1107568, len(result))
	}
}

func TestFactorial(t *testing.T) {
	n := big.NewInt(33)
	r := big.NewInt(0)
	r.SetString("8683317618811886495518194401280000000", 10)
	ar := Factorial(n)

	if ar.Cmp(r) != 0 {
		t.Errorf("33的阶乘为%v,计算结果为%v", r, ar)
	}
}
