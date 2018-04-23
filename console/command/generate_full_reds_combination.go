package command

import (
    "github.com/urfave/cli"
    "hiroki/util/math"
    "fmt"
    "hiroki/model"
    "strings"
    "hiroki/database"
)

func GenerateFullRedsCombination(c *cli.Context) error {
    db := database.DB
    iterable := make([]int, 33)
    for i := 1; i <= 33; i++ {
        iterable[i-1] = i
    }

    // 全组合数据
    numbers := math.Combination(iterable, 6)

    for _, item := range numbers {
        entity := &model.Number{}
        reds := [6]string{}
        for i, red := range item {
            reds[i] = fmt.Sprintf("%02d", red)
        }
        entity.Reds = strings.Join(reds[:], ",")
        entity.Beginning = reds[0]
        entity.Ending = reds[5]
        db.Create(entity)
    }

    return nil
}
