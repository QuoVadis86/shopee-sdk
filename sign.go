package shopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

// GenerateSignature creates the HMAC-SHA256 signature for a Shopee API request.
//
// The base string is constructed as follows depending on the API type:
//   - Public API:   partner_id + api_path + timestamp
//   - Shop API:     partner_id + api_path + timestamp + access_token + shop_id
//   - Merchant API: partner_id + api_path + timestamp + access_token + merchant_id
//
// The signature is the lowercase hex-encoded HMAC-SHA256 of the base string
// using partnerKey as the secret.
func GenerateSignature(partnerKey string, partnerID int64, apiPath string, timestamp int64, accessToken string, shopID int64, merchantID int64) string {
	parts := make([]byte, 0, 256)
	parts = strconv.AppendInt(parts, partnerID, 10)
	parts = append(parts, apiPath...)
	parts = strconv.AppendInt(parts, timestamp, 10)
	if accessToken != "" {
		parts = append(parts, accessToken...)
		if merchantID > 0 {
			parts = strconv.AppendInt(parts, merchantID, 10)
		} else if shopID > 0 {
			parts = strconv.AppendInt(parts, shopID, 10)
		}
	} else if shopID > 0 {
		parts = strconv.AppendInt(parts, shopID, 10)
	}
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write(parts)
	return hex.EncodeToString(mac.Sum(nil))
}

// stringsJoin joins a slice of strings with the given separator.
// Deprecated: use strings.Join from the standard library instead.
func stringsJoin(elems []string, sep string) string {
	return strings.Join(elems, sep)
}
