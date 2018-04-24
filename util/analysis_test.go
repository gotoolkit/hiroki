package util

import (
    "testing"
    "fmt"
)

var (
    n1 = []int{5, 8, 9, 19, 24, 27}
    n2 = []int{1, 2, 3, 14, 28, 29}
    n3 = []int{1, 10, 15, 20, 25, 30}
)

func TestGetAC(t *testing.T) {
    r := GetAC(n1)

    if r != 8 {
        t.Error(fmt.Sprintf("%v的AC值应为8,计算结果为%d", n1, r))
    }
}

func TestGetSum(t *testing.T) {
    r := GetSum(n1)

    if r != 92 {
        t.Error(fmt.Sprintf("%v的和值应为92,计算结果为%d", n1, r))
    }
}

func TestGetDaXiao(t *testing.T) {
    r := GetDaXiao(n1)

    if r != "33" {
        t.Error(fmt.Sprintf("%v的大小比应该为33,计算结果为%s", n1, r))
    }
}

func TestGetJiOu(t *testing.T) {
    r := GetJiOu(n1)

    if r != "42" {
        t.Error(fmt.Sprintf("%v的奇偶比应该为42,计算结果为%s", n1, r))
    }
}

func TestGetQuJian(t *testing.T) {
    r := GetQuJian(n1)

    if r != "312" {
        t.Error(fmt.Sprintf("%v的区间比应该为312,计算结果为%s", n1, r))
    }
}

func TestGetTongWei(t *testing.T) {
    r := GetTongWei(n1)

    if r != 5 {
        t.Error(fmt.Sprintf("%v的尾号个量应为312,计算结果为%d", n1, r))
    }
}

func TestGetLianHao(t *testing.T) {
    r1 := GetLianHao(n1)
    r2 := GetLianHao(n2)
    r3 := GetLianHao(n3)

    if r1 != 1 {
        t.Error(fmt.Printf("%v的连号个量应为1,计算结果为%d", n1, r1))
    }

    if r2 != 2 {
        t.Error(fmt.Printf("%v的连号个量应为1,计算结果为%d", n2, r2))
    }

    if r3 != 0 {
        t.Error(fmt.Printf("%v的连号个量应为0,计算结果为%d", n2, r3))
    }
}

func TestGetMaxLianHao(t *testing.T) {
    r1 := GetMaxLianHao(n1)
    r2 := GetMaxLianHao(n2)
    r3 := GetMaxLianHao(n3)

    if r1 != 2 {
        t.Error(fmt.Printf("%v的最大连号应为2,计算结果为%d", n1, r1))
    }

    if r2 != 3 {
        t.Error(fmt.Printf("%v的最大连号应为2,计算结果为%d", n1, r2))
    }

    if r3 != 1 {
        t.Error(fmt.Printf("%v的最大连号应为1,计算结果为%d", n1, r3))
    }
}

func TestGetKuaJu(t *testing.T) {
    r1 := GetKuaJu(n1)
    r2 := GetKuaJu(n2)

    if r1 != 22 {
        t.Error(fmt.Printf("%v的跨距应为22,计算结果为%d", n1, r1))
    }

    if r2 != 28 {
        t.Error(fmt.Printf("%v的跨距应为28,计算结果为%d", n2, r2))
    }
}
