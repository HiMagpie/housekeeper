package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"hash/crc32"
	"io"
)

/**
 * Md5加密字符串
 */
func Md5Str(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * offset 从0开始
 */
func Md5AndSub(s string, offset, length int) string {
	str := Md5Str(s)
	if len(str) <= offset {
		return ""
	}

	end := offset + length
	if end > len(str) - 1 {
		end = len(str) - 1
	}
	return str[offset: end]
}

/**
 * 字符串转int64
 */
func AtoInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// calculate string's crc32 int value
func Crc32Str(s string) (uint32, error) {
	ieee := crc32.NewIEEE()
	_, err := io.WriteString(ieee, s)
	if err != nil {
		return 0, err
	}

	return ieee.Sum32(), nil
}