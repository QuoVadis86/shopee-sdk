package shopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// GenerateSignature creates the HMAC-SHA256 signature for a Shopee API request.
// The base string is: partner_id + api_path + timestamp + [access_token] + [shop_id]
// concatenated in that order without separators.
func GenerateSignature(partnerKey string, partnerID int64, apiPath string, timestamp int64, accessToken string, shopID int64) string {
	parts := make([]byte, 0, 256)
	parts = strconv.AppendInt(parts, partnerID, 10)
	parts = append(parts, apiPath...)
	parts = strconv.AppendInt(parts, timestamp, 10)
	if accessToken != "" {
		parts = append(parts, accessToken...)
	}
	if shopID > 0 {
		parts = strconv.AppendInt(parts, shopID, 10)
	}
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write(parts)
	return hex.EncodeToString(mac.Sum(nil))
}

func stringsJoin(elems []string, sep string) string {
	if len(elems) == 0 {
		return ""
	}
	n := len(sep) * (len(elems) - 1)
	for _, e := range elems {
		n += len(e)
	}
	b := make([]byte, n)
	i := 0
	for idx, e := range elems {
		if idx > 0 {
			copy(b[i:], sep)
			i += len(sep)
		}
		copy(b[i:], e)
		i += len(e)
	}
	return string(b)
}
