package utils

import "encoding/base64"

func GetBase64Code(password string) string {
	base64Str := base64.StdEncoding.EncodeToString([]byte(password))
	return base64Str
}
