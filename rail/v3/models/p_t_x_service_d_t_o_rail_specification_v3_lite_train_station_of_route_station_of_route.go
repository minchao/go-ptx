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

// PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute StationOfRoute
//
// 營運路線車站基本資料
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.LiteTrain.StationOfRoute.StationOfRoute
type PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute struct {

	// Int32
	//
	// 營運路線方向描述(0:去程,1:返程)
	// Required: true
	Direction *int32 `json:"Direction"`

	// String
	//
	// 營運路線所屬之路線代碼
	// Required: true
	LineID *string `json:"LineID" xml:"String"`

	// String
	//
	// 營運路線所屬之路線編號
	// Required: true
	LineNo *string `json:"LineNo" xml:"String"`

	// String
	//
	// 營運路線代碼
	// Required: true
	RouteID *string `json:"RouteID" xml:"String"`

	// NameType
	//
	// 營運路線名稱
	// Required: true
	RouteName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"RouteName" xml:"NameType"`

	// Array
	//
	// 營運路線車站資訊
	// Required: true
	Stations []*PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation "json:\"Stations\" xml:\"List`1\""
}

// Validate validates this p t x service d t o rail specification v3 lite train station of route station of route
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRouteName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) validateDirection(formats strfmt.Registry) error {

	if err := validate.Required("Direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) validateLineID(formats strfmt.Registry) error {

	if err := validate.Required("LineID", "body", m.LineID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) validateLineNo(formats strfmt.Registry) error {

	if err := validate.Required("LineNo", "body", m.LineNo); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) validateRouteID(formats strfmt.Registry) error {

	if err := validate.Required("RouteID", "body", m.RouteID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) validateRouteName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) validateStations(formats strfmt.Registry) error {

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

// ContextValidate validate this p t x service d t o rail specification v3 lite train station of route station of route based on the context it is used
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRouteName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) contextValidateRouteName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) contextValidateStations(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteStationOfRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
