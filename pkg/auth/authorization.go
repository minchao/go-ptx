package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func date() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

func signature(key, date string) string {
	mac := hmac.New(sha1.New, []byte(key))
	_, _ = mac.Write([]byte("x-date: " + date))
	hash := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
}

func authorization(username, signature string) string {
	return fmt.Sprintf(`hmac username="%s", algorithm="hmac-sha1", headers="x-date", signature="%s"`,
		username, signature)
}
