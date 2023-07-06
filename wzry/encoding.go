package wzry

import "golang.org/x/text/encoding/simplifiedchinese"

// func ConvertStr2GBK(str string) string {
// 	//将utf-8编码的字符串转换为GBK编码
// 	ret, err := simplifiedchinese.GBK.NewEncoder().String(str)
// 	return ret //如果转换失败返回空字符串

// 	//如果是[]byte格式的字符串，可以使用Bytes方法
// 	b, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
// 	return string(b)

// }

// func ConvertGBK2Str(gbkStr string) string {
// 	//将GBK编码的字符串转换为utf-8编码
// 	// ret, err := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
// 	// return ret //如果转换失败返回空字符串

// 	//如果是[]byte格式的字符串，可以使用Bytes方法
// 	b, err := simplifiedchinese.GBK.NewDecoder().Bytes([]byte(gbkStr))
// 	return string(b)
// }

func convertGbkToUtf8(b []byte) string {
	s, _ := simplifiedchinese.GBK.NewDecoder().Bytes(b)
	return string(s)
}
