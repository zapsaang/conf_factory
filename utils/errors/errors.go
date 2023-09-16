package errors

import (
	"fmt"

	"github.com/zapsaang/conf_factory/conf/env"
	"github.com/zapsaang/conf_factory/utils/consts"
)

var (
	ErrArgumentsLengthMisMatch  = fmt.Errorf("arguments length mismatch")
	ErrArgumentsContentMisMatch = fmt.Errorf("arguments content mismatch")
	ErrArgumentsTypeMisMatch    = fmt.Errorf("arguments type mismatch")

	ErrStatusCode = map[error]int{
		ErrArgumentsLengthMisMatch:  consts.StatusBadRequest,
		ErrArgumentsContentMisMatch: consts.StatusBadRequest,
		ErrArgumentsTypeMisMatch:    consts.StatusBadRequest,
	}
)

func ErrorToResp(err error) (code int, msg string) {
	code, ok := ErrStatusCode[err]
	if !env.IsProduction() {
		msg = err.Error()
	}
	if !ok {
		return consts.StatusBadRequest, msg
	}
	return code, msg
}
