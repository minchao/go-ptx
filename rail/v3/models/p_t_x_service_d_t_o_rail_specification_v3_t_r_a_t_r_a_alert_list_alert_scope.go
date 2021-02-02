// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope AlertScope
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.TRAAlertList.AlertScope
type PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope struct {

	// AlertScopeLineSection[]
	//
	// 受影響的路線區間
	// Required: true
	LineSections []*PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeLineSection `json:"LineSections"`

	// AlertScopeLine[]
	//
	// 受影響的實體路線
	// Required: true
	Lines []*PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeLine `json:"Lines"`

	// AlertScopeNetwork
	//
	// 受影響的路網
	NetworkList struct {
		PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeNetwork
	} `json:"NetworkList,omitempty"`

	// AlertScopeRoute[]
	//
	// 受影響的營運路線
	// Required: true
	Routes []*PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeRoute `json:"Routes"`

	// AlertScopeStation[]
	//
	// 受影響的車站
	// Required: true
	Stations []*PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeStation `json:"Stations"`

	// AlertScopeTrain[]
	//
	// 受影響的車次
	// Required: true
	Trains []*PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScopeTrain `json:"Trains"`
}

// Validate validates this p t x service d t o rail specification v3 t r a t r a alert list alert scope
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLineSections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) validateLineSections(formats strfmt.Registry) error {

	if err := validate.Required("LineSections", "body", m.LineSections); err != nil {
		return err
	}

	for i := 0; i < len(m.LineSections); i++ {
		if swag.IsZero(m.LineSections[i]) { // not required
			continue
		}

		if m.LineSections[i] != nil {
			if err := m.LineSections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LineSections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) validateLines(formats strfmt.Registry) error {

	if err := validate.Required("Lines", "body", m.Lines); err != nil {
		return err
	}

	for i := 0; i < len(m.Lines); i++ {
		if swag.IsZero(m.Lines[i]) { // not required
			continue
		}

		if m.Lines[i] != nil {
			if err := m.Lines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) validateNetworkList(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkList) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) validateRoutes(formats strfmt.Registry) error {

	if err := validate.Required("Routes", "body", m.Routes); err != nil {
		return err
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) validateStations(formats strfmt.Registry) error {

	if err := validate.Required("Stations", "body", m.Stations); err != nil {
		return err
	}

	for i := 0; i < len(m.Stations); i++ {
		if swag.IsZero(m.Stations[i]) { // not required
			continue
		}

		if m.Stations[i] != nil {
			if err := m.Stations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) validateTrains(formats strfmt.Registry) error {

	if err := validate.Required("Trains", "body", m.Trains); err != nil {
		return err
	}

	for i := 0; i < len(m.Trains); i++ {
		if swag.IsZero(m.Trains[i]) { // not required
			continue
		}

		if m.Trains[i] != nil {
			if err := m.Trains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Trains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a t r a alert list alert scope based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLineSections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) contextValidateLineSections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LineSections); i++ {

		if m.LineSections[i] != nil {
			if err := m.LineSections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("LineSections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) contextValidateLines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Lines); i++ {

		if m.Lines[i] != nil {
			if err := m.Lines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Lines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) contextValidateNetworkList(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) contextValidateRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Routes); i++ {

		if m.Routes[i] != nil {
			if err := m.Routes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) contextValidateStations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Stations); i++ {

		if m.Stations[i] != nil {
			if err := m.Stations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Stations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) contextValidateTrains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Trains); i++ {

		if m.Trains[i] != nil {
			if err := m.Trains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Trains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRATRAAlertListAlertScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
