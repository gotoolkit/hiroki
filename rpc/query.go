package main

import (
    "hiroki/rpc/types"

    _ "gopkg.in/doug-martin/goqu.v4/adapters/mysql"
)

type NumberQueryResult struct {
    FilterParams *types.FilterParams
    Count        uint
    Result       []string
}

func getFilterQueryResult(params *types.FilterParams) (*NumberQueryResult, error) {
    return nil, nil
}
