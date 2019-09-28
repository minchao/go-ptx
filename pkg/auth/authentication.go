package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type Authentication struct {
	AppID  string
	AppKey string
}

func NewAuthentication(id, key string) *Authentication {
	return &Authentication{
		AppID:  id,
		AppKey: key,
	}
}

func (a Authentication) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	date := date()

	_ = req.SetHeaderParam("Authorization", authorization(a.AppID, signature(a.AppKey, date)))
	_ = req.SetHeaderParam("x-date", date)
	return nil
}
