// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/minchao/go-ptx/rail/v2/client/common"
	"github.com/minchao/go-ptx/rail/v2/client/metro"
	"github.com/minchao/go-ptx/rail/v2/client/t_h_s_r"
	"github.com/minchao/go-ptx/rail/v2/client/t_r_a"
)

// Default m o t c transport API v2 HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "ptx.transportdata.tw"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/MOTC"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new m o t c transport API v2 HTTP client.
func NewHTTPClient(formats strfmt.Registry) *MOTCTransportAPIV2 {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new m o t c transport API v2 HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *MOTCTransportAPIV2 {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new m o t c transport API v2 client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *MOTCTransportAPIV2 {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(MOTCTransportAPIV2)
	cli.Transport = transport
	cli.Common = common.New(transport, formats)
	cli.Metro = metro.New(transport, formats)
	cli.Thsr = t_h_s_r.New(transport, formats)
	cli.Tra = t_r_a.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// MOTCTransportAPIV2 is a client for m o t c transport API v2
type MOTCTransportAPIV2 struct {
	Common common.ClientService

	Metro metro.ClientService

	Thsr t_h_s_r.ClientService

	Tra t_r_a.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *MOTCTransportAPIV2) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Common.SetTransport(transport)
	c.Metro.SetTransport(transport)
	c.Thsr.SetTransport(transport)
	c.Tra.SetTransport(transport)
}
