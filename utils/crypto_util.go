package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

func getKeys(m map[string]interface{}) []string {
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func JoinStr(params map[string]interface{}) string {
	keys := getKeys(params)
	sort.Strings(keys)
	var result []string
	for _, key := range keys {
		val := params[key]
		if val == nil {
			val = ""
		}
		result = append(result, fmt.Sprintf("%s=%s", key, fmt.Sprintf("%v", val)))
	}
	return strings.Join(result, "&")
}

func Sign(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
