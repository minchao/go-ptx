// Code generated by go-swagger; DO NOT EDIT.

package basic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new basic API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for basic API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
BasicAPIAuthority basic Api authority API
*/
func (a *Client) BasicAPIAuthority(params *BasicAPIAuthorityParams) (*BasicAPIAuthorityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBasicAPIAuthorityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BasicApi_Authority",
		Method:             "GET",
		PathPattern:        "/v2/Basic/Authority",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BasicAPIAuthorityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BasicAPIAuthorityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BasicApi_Authority: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BasicAPIOperator basic Api operator API
*/
func (a *Client) BasicAPIOperator(params *BasicAPIOperatorParams) (*BasicAPIOperatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBasicAPIOperatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BasicApi_Operator",
		Method:             "GET",
		PathPattern:        "/v2/Basic/Operator",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BasicAPIOperatorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BasicAPIOperatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BasicApi_Operator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
BasicAPIProvider basic Api provider API
*/
func (a *Client) BasicAPIProvider(params *BasicAPIProviderParams) (*BasicAPIProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBasicAPIProviderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "BasicApi_Provider",
		Method:             "GET",
		PathPattern:        "/v2/Basic/Provider",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BasicAPIProviderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BasicAPIProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for BasicApi_Provider: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WebSiteAPINews web site Api news API
*/
func (a *Client) WebSiteAPINews(params *WebSiteAPINewsParams) (*WebSiteAPINewsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWebSiteAPINewsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "WebSiteApi_News",
		Method:             "GET",
		PathPattern:        "/v2/PTX/Web/News",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &WebSiteAPINewsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WebSiteAPINewsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for WebSiteApi_News: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
