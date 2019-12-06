package ownerr

import "errors"

var (
	ERR_EDGE_FROM          = errors.New("err edge from")
	ERR_EDGE_TO            = errors.New("err edge to")
	Err_SOURCE_SAME_TARGET = errors.New("same source and target")
	ERR_BELOW_RANGE        = errors.New("index < 0")
)

const MaxWeight = 1000
