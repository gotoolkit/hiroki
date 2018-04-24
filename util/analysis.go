package util

import (
    "math"
    math2 "hiroki/util/math"
    "strings"
    "strconv"
)

// GetAC return the int slice ac value
func GetAC(n []int) uint {
    if len(n) == 0 {
        return 0
    }

    // 从数据里抽拿出2个数字作为一个组合
    cb := math2.Combination(n, 2)
    keys := make(map[uint]bool)
    tmp := make([]uint, 0, len(cb))

    for _, v := range cb {
        abs := uint(math.Abs(float64(v[1] - v[0])))
        if _, v := keys[abs]; !v {
            keys[abs] = true
            tmp = append(tmp, abs)
        }
    }

    return uint(len(tmp) - (len(n) - 1))
}

// GetSum return the int slice sum value
func GetSum(n []int) uint {
    s := 0
    for _, v := range n {
        s += v
    }
    return uint(s)
}

// GetDaXiao return the int slice daxiao string
func GetDaXiao(n []int) string {
    r := make([]string, 2)
    c := [2]int{0, 0}
    for _, v := range n {
        if v >= 17 {
            c[0] += 1
        } else {
            c[1] += 1
        }
    }
    r[0] = strconv.Itoa(c[0])
    r[1] = strconv.Itoa(c[1])

    return strings.Join(r, "")
}

// GetDaXiao return the int slice jiou string
func GetJiOu(n []int) string {
    r := make([]string, 2)
    c := [2]int{0, 0}
    for _, v := range n {
        if v%2 != 0 {
            c[0]++
        } else {
            c[1]++
        }
    }
    r[0] = strconv.Itoa(c[0])
    r[1] = strconv.Itoa(c[1])

    return strings.Join(r, "")
}

// GetQuJian return the int slice qujian string
func GetQuJian(n []int) string {
    r := make([]string, 3)
    c := [3]int{0, 0, 0}
    for _, v := range n {
        if v <= 11 {
            c[0]++
        } else if v > 11 && v <= 22 {
            c[1]++
        } else {
            c[2]++
        }
    }
    r[0] = strconv.Itoa(c[0])
    r[1] = strconv.Itoa(c[1])
    r[2] = strconv.Itoa(c[2])

    return strings.Join(r, "")
}

// GetTongWei return the int slice tongwei count value
func GetTongWei(n []int) uint {
    keys := make(map[uint8]bool)
    for _, v := range n {
        end := uint8(v % 10)
        if _, vv := keys[end]; !vv {
            keys[end] = true
        }
    }
    return uint(len(keys))
}

// GetLianHao 返回一组数字存在连续数字的个量
func GetLianHao(n []int) uint {
    lianHao := 0
    flag := false
    for k, v := range n {
        if k > 0 {
            if uint(math.Abs(float64(v-n[k-1]))) == 1 {
                if !flag {
                    lianHao++
                    flag = true
                }
            } else {
                flag = false
            }
        }
    }
    return uint(lianHao)
}

// GetMaxLianHao 返回一组数字存在的最大连号数量
func GetMaxLianHao(n []int) uint {
    maxLianHao := 1
    tmpMaxLianHao := 1
    for k, v := range n {
        if k > 0 {
            tmp := uint(math.Abs(float64(v - n[k-1])))
            if tmp == 1 {
                tmpMaxLianHao++
                if tmpMaxLianHao > maxLianHao {
                    maxLianHao = tmpMaxLianHao
                }
            } else {
                tmpMaxLianHao = 1
            }
        }
    }
    return uint(maxLianHao)
}

// GetKuaJu 返回一组数据的跨距
func GetKuaJu(n []int) uint {
    var (
        max int = n[0]
        min int = n[0]
    )
    for _, v := range n {
        if v > max {
            max = v
        }
        if v < min {
            min = v
        }
    }
    return uint(max - min)
}
