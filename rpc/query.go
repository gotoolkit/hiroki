package main

import (
	"fmt"
	"ssq/src/config"
	"ssq/src/database"
	"ssq/src/types"

	"gopkg.in/doug-martin/goqu.v4"
	_ "gopkg.in/doug-martin/goqu.v4/adapters/mysql"
)

type NumberQueryResult struct {
	FilterParams *types.FilterParams
	Count        uint
	Result       []string
}

func getFilterQueryResult(params *types.FilterParams) (*NumberQueryResult, error) {
	db := database.DB
	cfg := config.GetInstance()
	result := &NumberQueryResult{
		Count:        0,
		FilterParams: params,
	}

	whereStmts := make([]goqu.Expression, 0)
	whereStmts = appendStmt(whereStmts, "qu_jian", params.QuJian)

	// 统计数量
	sqlStmt, _, _ := db.From("common_number").Select(goqu.COUNT("*")).Where(whereStmts...).ToSql()
	fmt.Println(sqlStmt)
	rows, err := database.DB.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.Scan(&result.Count)
	}

	// 小于指定数量才返回实际结果
	if result.Count < cfg.MaxResult {
		sqlStmt, _, _ = db.From("common_number").Select("reds").Where(whereStmts...).ToSql()
		fmt.Println(sqlStmt)
		rows, err := database.DB.Query(sqlStmt)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var reds string
			err = rows.Scan(&reds)
			result.Result = append(result.Result, reds)
		}
	}

	return result, nil
}
