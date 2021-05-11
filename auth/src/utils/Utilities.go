package utils;

import(
	"fmt"
	"time"
	base64 "encoding/base64"
	json "encoding/json"
)
func Base64encode(obj interface{}) string {
	var str string
	switch obj.(type) {
	case string:
		str = obj.(string)
	case []byte:
		str = string(obj.([]byte))
	case int64:
		str = fmt.Sprintf("%d", obj.(int64))
	}
	return base64.RawStdEncoding.EncodeToString([]byte(str))
}
func ToString(val interface{}) string {
	switch val.(type) {
	case int64:
		return fmt.Sprintf("%d", val.(int64))		
	default:
		return fmt.Sprintf("%d", val.(int32))
	}
}
func DeleteIndex(values interface{}, index int) interface{} {
	switch values.(type){
	case []int:
		vals := values.([]int)
		return append(vals[:index], vals[index+1:]...)
	default:
		vals := values.([]int)
		return append(vals[:index], vals[index+1:]...)
	}
}
func GetCurrentUnixNano() int64 {
	return time.Now().UnixNano()
}
func JsonEncode(value interface{}) []byte {
	jsonbody, _ := json.Marshal(value)
	return jsonbody
}