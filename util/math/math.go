package math

import (
    "math/big"
)

func Combination(iterable []int, r int) [][]int {
    pool := iterable
    n := len(iterable)

    if r > n {
        return nil
    }

    biLen := big.NewInt(int64(n))
    total := big.NewInt(0).Div(Factorial(biLen), big.NewInt(0).Mul(Factorial(big.NewInt(int64(r))), Factorial(big.NewInt(int64(n-r)))))

    result := make([][]int, 0, total.Uint64())
    for i := range result {
        result[i] = make([]int, r)
    }

    indices := make([]int, r)
    for i := range indices {
        indices[i] = i
    }

    item := make([]int, r)

    result = append(result, make([]int, r))
    for i, el := range indices {
        item[i] = pool[el]
    }
    copy(result[len(result)-1], item)

    for {
        i := r - 1
        for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
        }

        if i < 0 {
            return result
        }

        indices[i] += 1
        for j := i + 1; j < r; j += 1 {
            indices[j] = indices[j-1] + 1
        }

        result = append(result, make([]int, r))
        for ; i < len(indices); i += 1 {
            item[i] = pool[indices[i]]
        }
        copy(result[len(result)-1], item)
    }
}

// Factorial x阶乘
func Factorial(n *big.Int) (result *big.Int) {
    b := big.NewInt(0)
    c := big.NewInt(1)

    if n.Cmp(b) == 0 || n.Cmp(b) == -1 {
        result = big.NewInt(1)
    } else {
        result = new(big.Int)
        result.Set(n)
        result.Mul(result, Factorial(n.Sub(n, c)))
    }
    return
}
