package transport

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

type AuthTransport struct {
	AppId  string
	AppKey string

	Transport http.RoundTripper
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	date := t.date()
	signature := t.signature(t.AppKey, date)

	req2.Header.Set("Authorization", fmt.Sprintf(`hmac username="%s", algorithm="hmac-sha1", headers="x-date", signature="%s"`,
		t.AppId, signature))
	req2.Header.Set("x-date", date)

	return t.transport().RoundTrip(req2)
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *AuthTransport) signature(key, date string) string {
	mac := hmac.New(sha1.New, []byte(key))
	_, _ = mac.Write([]byte("x-date: " + date))
	hash := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(hash)
}

func (t *AuthTransport) date() string {
	return time.Now().UTC().Format(http.TimeFormat)
}
