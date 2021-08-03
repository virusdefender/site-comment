package api

import (
	"encoding/binary"
	"strings"
)

type Comment struct {
	ID        int    `json:"id"`
	ArticleID string `json:"article_id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	Email     string `json:"email"`
	TimeStamp int64  `json:"timestamp"`
	IP        string `json:"ip"`
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func maskEmail(s string) string {
	split := strings.SplitN(s, "@", 2)
	if len(split) != 2 {
		return s
	}
	prefix := split[0]
	plen := len(prefix)
	if plen <= 4 {
		prefix = strings.Repeat("*", plen)
	} else {
		prefix = prefix[:2] + strings.Repeat("*", len(prefix)-3) + prefix[plen-1:]
	}
	return prefix + "@" + split[1]
}

func respBody(data interface{}) interface{} {
	return map[string]interface{}{"data": data}
}
