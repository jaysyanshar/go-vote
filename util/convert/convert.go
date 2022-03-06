package convert

import "strconv"

func StrToInt64(value string) (int64, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	return i, err
}
