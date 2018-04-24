package main

import (
    "fmt"
    "net/http"
    "hiroki/config"
    "hiroki/database"
    "hiroki/rpc/types"

    "github.com/hprose/hprose-golang/rpc"
    "github.com/hprose/hprose-golang/rpc/filter/jsonrpc"
    "gopkg.in/doug-martin/goqu.v4"
    _ "gopkg.in/doug-martin/goqu.v4/adapters/mysql"
)

func main() {
    database.Init()
    cfg := config.GetInstance()

    defer database.Close()

    service := rpc.NewHTTPService()
    service.Debug = cfg.AppMode == "development"
    service.AddFunction("getResult", getResult)

    service.AddBeforeFilterHandler(
        statFilter{"Server: BeforeFilter"}.handler,
    ).AddFilter(
        jsonrpc.ServiceFilter{},
        LogFilter{"Server"},
    ).AddAfterFilterHandler(
        statFilter{"Server: AfterFilter"}.handler,
    )

    http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.ServerHost, cfg.ServerPort), service)
}

func getResult(params map[string]interface{}) map[string]interface{} {
    resp := types.CreateResponseData()
    filterParams := &types.FilterParams{}
    resp.Data = make(map[string]interface{})

    filters := params["filters"].(map[interface{}]interface{})
    filterParams.QuJian = conversionParameter(filters["quJian"].([]interface{}))
    filterParams.AC = conversionParameter(filters["ac"].([]interface{}))
    filterParams.JiOu = conversionParameter(filters["jiOu"].([]interface{}))
    filterParams.DaXiao = conversionParameter(filters["daXiao"].([]interface{}))
    filterParams.SumRange = conversionParameter(filters["sumRange"].([]interface{}))
    filterParams.LianHao = conversionParameter(filters["lianHao"].([]interface{}))
    filterParams.MaxLianHao = conversionParameter(filters["maxLianHao"].([]interface{}))
    filterParams.KuaJu = conversionParameter(filters["kuaJu"].([]interface{}))

    whereStmts := make([]goqu.Expression, 0)
    whereStmts = appendStmt(whereStmts, "qu_jian", filterParams.QuJian)

    result, err := getFilterQueryResult(filterParams)
    checkError(err)

    resp.Data.(map[string]interface{})["count"] = result.Count
    resp.Data.(map[string]interface{})["results"] = result.Result

    return resp.ToMap()
}

func appendStmt(stmt []goqu.Expression, field string, c []string) []goqu.Expression {
    if len(c) > 0 {
        return append(stmt, goqu.I(field).In(c))
    }
    return stmt
}

func conversionParameter(params []interface{}) []string {
    data := make([]string, len(params))
    for i := range params {
        data[i] = params[i].(string)
    }
    return data
}

func checkError(err error) {
    if err != nil {
        panic(err.Error())
    }
}
