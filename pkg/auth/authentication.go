package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type Authentication struct {
	AppId  string
	AppKey string
}

func NewAuthentication(id, key string) *Authentication {
	return &Authentication{
		AppId:  id,
		AppKey: key,
	}
}

func (a Authentication) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	date := date()

	_ = req.SetHeaderParam("Authorization", authorization(a.AppId, signature(a.AppKey, date)))
	_ = req.SetHeaderParam("x-date", date)
	return nil
}
