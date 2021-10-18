package utils

import (
	"context"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func StrToInt64(str string, defaultValue int64) int64 {
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		value = defaultValue
	}
	return value
}

func StrToInt(str string, defaultValue int) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		value = defaultValue
	}
	return value
}

func StrToInt16(str string, defaultValue int16) int16 {
	value, err := strconv.ParseInt(str, 10, 16)
	if err == nil {
		return int16(value)
	}
	return defaultValue
}

func StrToInt32(str string, defaultValue int32) int32 {
	value, err := strconv.ParseInt(str, 10, 32)
	if err == nil {
		return int32(value)
	}
	return defaultValue
}

func StrToFloat32(str string, defaultValue float32) float32 {
	value, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return float32(value)
	}
	return defaultValue
}

func StrToFloat64(str string, defaultValue float64) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return value
	}
	return defaultValue
}

func Uint64ToStr(ui64 uint64) string {
	return strconv.FormatUint(ui64, 10)
}

func Strval(it interface{}) string {
	switch v := it.(type) {
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case bool:
		if v == false {
			return "false"
		}
		return "true"
	case []byte:
		return *((*string)(unsafe.Pointer(&v)))
	default:
		return ""
	}
}

func ParseInterType(it interface{}) string {
	switch it.(type) {
	case int:
		return "int"
	case int32:
		return "int32"
	case int64:
		return "int64"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case []byte:
		return "[]byte"
	case nil:
		return "null"
	case map[string]string:
		return "map[string]string"
	case map[interface{}]interface{}:
		return "map[interface{}]interface{}"
	default:
		return "unknown"
	}
}

func Uint32ToTimeMillisecond(timeout uint32) time.Duration {
	return time.Duration(timeout) * time.Millisecond
}

func Uint32ToTimeSecond(timeout uint32) time.Duration {
	return time.Duration(timeout) * time.Second
}

// 获取第一个非零int值
func FirstNotZeroUInt32(vals ...uint32) uint32 {
	for _, val := range vals {
		if val != uint32(0) {
			return val
		}
	}
	return 0
}

func GetTimeout(ctx context.Context, configTimeout uint32, defaultTimeout uint32) (time.Time, time.Duration) {
	m := FirstNotZeroUInt32(configTimeout, defaultTimeout)
	d := Uint32ToTimeMillisecond(m)
	n := time.Now()
	t := n.Add(d)
	if deadline, ok := ctx.Deadline(); ok && deadline.Before(t) {
		d = deadline.Sub(n)
		t = deadline
	}
	return t, d
}

func AddrToHost(addr string) string {
	tokens := strings.Split(addr, ":")
	if len(tokens) == 2 {
		return tokens[0]
	}
	if addr != "" {
		return addr
	}
	return "-"
}

func Contains(val string, whiteList []string) bool {
	for _, v := range whiteList {
		if v == val {
			return true
		}
	}
	return false
}

func BoolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func Int64Join(arr []int64, sep string) string {
	res := ""
	for i, v := range arr {
		if i > 0 {
			res += sep
		}
		res += Strval(v)
	}
	return res
}

func Int32Join(arr []int32, sep string) string {
	res := ""
	for i, v := range arr {
		if i > 0 {
			res += sep
		}
		res += Strval(v)
	}
	return res
}
