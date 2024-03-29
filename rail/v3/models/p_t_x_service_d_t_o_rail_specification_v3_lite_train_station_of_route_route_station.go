// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation RouteStation
//
// 路線車站資訊
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.LiteTrain.StationOfRoute.RouteStation
type PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation struct {

	// Single
	//
	// 已累積之里程距離(公里)
	// Required: true
	CumulativeDistance *float32 `json:"CumulativeDistance"`

	// Int32
	//
	// 站序
	// Required: true
	Sequence *int32 `json:"Sequence"`

	// String
	//
	// 車站代碼
	// Required: true
	StationID *string `json:"StationID" xml:"StationID"`

	// NameType
	//
	// 車站名稱
	// Required: true
	StationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"StationName" xml:"NameType"`
}

// Validate validates this p t x service d t o rail specification v3 lite train station of route route station
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCumulativeDistance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStationName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) validateCumulativeDistance(formats strfmt.Registry) error {

	if err := validate.Required("CumulativeDistance", "body", m.CumulativeDistance); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("Sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) validateStationID(formats strfmt.Registry) error {

	if err := validate.Required("StationID", "body", m.StationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) validateStationName(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 lite train station of route route station based on the context it is used
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) contextValidateStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3LiteTrainStationOfRouteRouteStation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
