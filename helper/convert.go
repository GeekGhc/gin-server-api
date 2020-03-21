package helper

import "strconv"

func String2Int64(arg *string) int64 {
	argInt64, _ := strconv.ParseInt(*arg, 10, 64)
	return argInt64
}

func String2Uint64(arg *string) uint64 {
	argUint64, _ := strconv.ParseUint(*arg, 10, 64)
	return argUint64
}

func String2Int(arg *string) int {
	ints, _ := strconv.Atoi(*arg)
	return ints
}

func Int2String(arg *int) string {
	return strconv.Itoa(*arg)
}

func Int642String(arg *int64) string {
	return strconv.FormatInt(*arg, 10)
}

func Uint642String(arg *uint64) string {
	return strconv.FormatUint(*arg, 10)
}