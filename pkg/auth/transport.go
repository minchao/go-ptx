package auth

import (
	"net/http"
)

type Transport struct {
	AppId  string
	AppKey string

	Transport http.RoundTripper
}

func NewTransport(id, key string) *Transport {
	return &Transport{
		AppId:  id,
		AppKey: key,
	}
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	date := date()

	req2.Header.Set("Authorization", authorization(t.AppId, signature(t.AppKey, date)))
	req2.Header.Set("x-date", date)

	return t.transport().RoundTrip(req2)
}

func (t *Transport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
