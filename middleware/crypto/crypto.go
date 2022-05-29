/**
 @author: 李映飞
 @date:   2022/5/29
 @note:
**/
package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Check(content, encrypted string) bool {
	return strings.EqualFold(encrypted, Encrypt(content))
}

func Encrypt(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
