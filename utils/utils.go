package utils

import (
	"encoding/json"
	"strings"

	"github.com/zapsaang/conf_factory/utils/consts"
)

func MustToIndentedJSON(value interface{}) []byte {
	valueJSON, err := ToIndentedJSON(value)
	if err != nil {
		return []byte{}
	}
	return valueJSON
}

func ToIndentedJSON(value interface{}) ([]byte, error) {
	valueJSON, err := json.MarshalIndent(value, consts.EmptyString, consts.JSONIndent)
	if err != nil {
		return nil, err
	}
	return valueJSON, nil
}

func MustToIndentedJSONString(value interface{}) string {
	valueJSON, err := ToIndentedJSONString(value)
	if err != nil {
		return consts.EmptyString
	}
	return valueJSON
}

func ToIndentedJSONString(value interface{}) (string, error) {
	valueJSON, err := json.MarshalIndent(value, consts.EmptyString, consts.JSONIndent)
	if err != nil {
		return consts.EmptyString, err
	}
	return string(valueJSON), nil
}

func IsComment(line string) bool {
	if len(line) == 0 {
		return false
	}
	switch true {
	case strings.HasPrefix(line, "//"):
		return true
	case line[0] == '#':
		return true
	case line[0] == ';':
		return true
	default:
		return false
	}
}
