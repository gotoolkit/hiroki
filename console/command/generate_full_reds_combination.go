package command

import (
    "github.com/urfave/cli"
    "hiroki/database"
    "hiroki/util/math"
    "hiroki/model"
    "fmt"
    "strings"
    "github.com/jinzhu/gorm"
    "hiroki/util"
)

func GenerateFullRedsCombination(c *cli.Context) error {
    var tx *gorm.DB
    db := database.DB

    iterable := make([]int, 33)
    for i := 1; i <= 33; i++ {
        iterable[i-1] = i
    }

    // 全组合数据
    numbers := math.Combination(iterable, 6)

    i := 0
    for _, item := range numbers {
        if i == 0 {
            tx = db.Begin()
        }

        entity := &model.Number{}
        reds := [6]string{}
        for i, red := range item {
            reds[i] = fmt.Sprintf("%02d", red)
        }

        entity.Reds = strings.Join(reds[:], ",")
        entity.Beginning = reds[0]
        entity.Ending = reds[5]
        entity.AC = uint8(util.GetAC(item))
        entity.Sum = uint8(util.GetSum(item))
        entity.JiOu = util.GetJiOu(item)

        db.Create(entity)
        i++

        if i == 10000 {
            tx.Commit()
            i = 0
        }
    }

    return nil
}
