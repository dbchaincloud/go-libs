package header

import (
	"net/http"
)

const (
	headerOfXDbcChainid string = "X-Dbc-Chainid"
)

func GetDBCChainID(r *http.Request) string {
	chainID := ""
	if len(r.Header[headerOfXDbcChainid]) == 1 {
		chainID = r.Header[headerOfXDbcChainid][0]
	}
	return chainID
}

