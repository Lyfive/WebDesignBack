/**
 @author: 李映飞
 @date:   2022/5/29
 @note:
**/
package detection

func CheckSensitiveWord(str string) bool {
	for _, v := range str {
		if v == ' ' || v == ';' || v == '\n' || v == '\r' || v == '\t' || v == '&' || v == '|' || v == '!' {
			return false
		}
	}
	return true
}
