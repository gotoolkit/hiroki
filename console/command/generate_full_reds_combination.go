package command

import (
    "github.com/urfave/cli"
    "hiroki/database"
    "hiroki/util/math"
    "fmt"
    "gopkg.in/cheggaaa/pb.v2"
    "strings"
    "hiroki/util"
)

func GenerateFullRedsCombination(c *cli.Context) error {
    db := database.DB.DB()

    iterable := make([]int, 33)
    for i := 1; i <= 33; i++ {
        iterable[i-1] = i
    }

    // 全组合数据
    numbers := math.Combination(iterable, 6)

    buffer := 10000
    bar := pb.StartNew(len(numbers))
    insertStmts := make([]string, 0, buffer)
    execstring := "INSERT INTO number (reds, beginning, ending, ac, qu_jian, ji_ou, da_xiao, sum, tong_wei, lian_hao, max_lian_hao, kua_ju) VALUES "
    for _, item := range numbers {
        reds := [6]string{}
        for i, red := range item {
            reds[i] = fmt.Sprintf("%02d", red)
        }

        stmt := fmt.Sprintf(
            "(\"%s\", \"%s\", \"%s\", %d, \"%s\", \"%s\", \"%s\", %d, %d, %d, %d, %d)",
            strings.Join(reds[:], ","),
            reds[0],
            reds[5],
            uint8(util.GetAC(item)),
            util.GetQuJian(item),
            util.GetJiOu(item),
            util.GetDaXiao(item),
            uint(util.GetSum(item)),
            uint8(util.GetTongWei(item)),
            uint8(util.GetLianHao(item)),
            uint8(util.GetMaxLianHao(item)),
            uint8(util.GetKuaJu(item)),
        )
        insertStmts = append(insertStmts, stmt)

        if len(insertStmts) == cap(insertStmts) {
            _, err := db.Exec(execstring + strings.Join(insertStmts, ","))
            if err != nil {
                return err
            }
            insertStmts = nil
            insertStmts = make([]string, 0, buffer)
        }

        bar.Increment()
    }

    if len(insertStmts) > 0 {
        _, err := db.Exec(execstring + strings.Join(insertStmts, ","))
        if err != nil {
            return err
        }
    }

    bar.Finish()
    return nil
}
