package go_cmq

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

//
// // MakeParamStr 构建提交参数的字符串，考虑动态参数原因，改为使用map
// func MakeParamStr(params interface{}) string {
// 	var paramStr string
// 	v := reflect.ValueOf(params)
// 	elemType := v.Elem().Type()
// 	count := v.NumField()
// 	var keyNames []string
// 	for i := 0; i < count; i++ {
// 		if elemType.Field(i).Name == "Signature" {
// 			continue
// 		}
// 		keyName := strings.Replace(elemType.Field(i).Name, "_", ".", -1)
// 		keyNames = append(keyNames, keyName)
// 	}
// 	// 对键值进行排序
// 	sort.Strings(keyNames)
// 	for n := 0; n < len(keyNames); n++ {
// 		if n == 0 {
// 			paramStr = paramStr + "?"
// 		} else {
// 			paramStr = paramStr + "&"
// 		}
// 		switch v.Field(n).Kind() {
// 		case reflect.String:
// 			paramStr = paramStr + keyNames[n] + "=" + v.Field(n).String()
// 		case reflect.Int64:
// 			paramStr = paramStr + keyNames[n] + "=" + strconv.FormatInt(v.Field(n).Int(), 10)
// 		case reflect.Float64:
// 			paramStr = paramStr + keyNames[n] + "=" + strconv.FormatFloat(v.Field(n).Float(), 'E', -1, 64)
// 		case reflect.Bool:
// 			paramStr = paramStr + keyNames[n] + "=" + strconv.FormatBool(v.Field(n).Bool())
// 		}
// 	}
// 	return paramStr
// }

const (
	signSHA256 = "HmacSHA256"
	signSHA1   = "HmacSHA1"
)

func mapToURLParam(src map[string]interface{}, encoder bool) string {
	var keys []string
	for k, _ := range src {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	param := make([]string, len(keys))

	for i, k := range keys {
		key := strings.Replace(k, "_", ".", -1)
		if s, ok := src[k].(string); ok {
			if encoder {
				param[i] = key + "=" + url.QueryEscape(s)
			} else {
				param[i] = key + "=" + s
			}
		} else if s, ok := src[k].(int); ok {
			param[i] = key + "=" + strconv.Itoa(s)
		} else if s, ok := src[k].(int64); ok {
			param[i] = fmt.Sprintf("%s=%d", key, s)
		} else {
			param[i] = fmt.Sprintf("%s=%T", key, k)
		}
	}
	return strings.Join(param, "&")
}
