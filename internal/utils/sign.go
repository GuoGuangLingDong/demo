package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
	"strings"
)

func Sign(appSecret string, reqMap map[string]interface{}) string {
	keys := make([]string, 0)
	for k, _ := range reqMap {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	kvs := make([]string, 0)
	for _, k := range keys {
		v := reqMap[k]
		switch v.(type) {
		case []map[string]interface{}:
			r := ""
			for _, v := range reqMap[k].([]map[string]interface{}) {
				for _, vv := range v {
					r += gconv.String(vv)
				}
			}
			kvs = append(kvs, fmt.Sprintf("%s=%v", k, r))
		default:
			kvs = append(kvs, fmt.Sprintf("%s=%v", k, reqMap[k]))
		}
	}
	hash := strings.Join(kvs, "&")

	key := []byte(appSecret)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(hash))

	str := hex.EncodeToString(mac.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(str))
}
