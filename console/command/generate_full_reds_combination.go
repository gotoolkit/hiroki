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
    "gopkg.in/cheggaaa/pb.v2"
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

    bar := pb.StartNew(len(numbers))
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
        entity.QuJian = util.GetQuJian(item)
        entity.JiOu = util.GetJiOu(item)
        entity.DaXiao = util.GetDaXiao(item)
        entity.Sum = uint(util.GetSum(item))
        entity.TongWei = uint8(util.GetTongWei(item))
        entity.LianHao = uint8(util.GetLianHao(item))
        entity.MaxLianHao = uint8(util.GetMaxLianHao(item))
        entity.KuaJu = uint8(util.GetKuaJu(item))

        db.Create(entity)
        i++

        if i == 2000 {
            tx.Commit()
            i = 0
        }

        bar.Increment()
    }
    bar.Finish()
    return nil
}
