package utils

// 判断请求字段是否是空字符串
func ParamIsNull(s ...string) bool {
	for _, value := range s {
		if len(value) == 0 {
			return true
		}
	}
	return false
}
