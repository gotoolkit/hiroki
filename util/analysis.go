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

// GetJiOu return the int slice jiou string
func GetJiOu(n []int) string {
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
