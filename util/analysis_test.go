package util

import (
    "testing"
    "fmt"
)

var (
    n = []int{5, 8, 9, 19, 24, 27}
)

func TestGetAC(t *testing.T) {
    r := GetAC(n)

    if r != 8 {
        t.Error(fmt.Sprintf("%v的AC值应为8,计算结果为%d", n, r))
    }
}

func TestGetSum(t *testing.T) {
    r := GetSum(n)

    if r != 92 {
        t.Error(fmt.Sprintf("%v的和值应为92,计算结果为%d", n, r))
    }
}

func TestGetJiOu(t *testing.T) {
    r := GetJiOu(n)

    if r != "33" {
        t.Error(fmt.Sprintf("%v的奇偶值应该33,计算结果为%s", n, r))
    }
}
