package helper

import (
	"fmt"
	"strconv"
)

func ConvertToInt(param interface{}) (uint64, error) {
	var numberUint uint64
	paramTo := fmt.Sprintf("%v", param)
	paramUint64, err := strconv.ParseInt(paramTo, 0, 0)
	if err != nil {
		return numberUint, err
	}
	return uint64(paramUint64), nil
}
