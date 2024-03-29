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

// PTXServiceDTORailSpecificationV3TRAODFareODFare ODFare
//
// swagger:model PTX.Service.DTO.Rail.Specification.V3.TRA.ODFare.ODFare
type PTXServiceDTORailSpecificationV3TRAODFareODFare struct {

	// String
	//
	// 迄點車站代碼
	// Required: true
	DestinationStationID *string `json:"DestinationStationID" xml:"DestinationStationID"`

	// NameType
	//
	// 迄點車站名稱
	// Required: true
	DestinationStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"DestinationStationName" xml:"NameType"`

	// Int32
	//
	// 方向 : [0:'順行',1:'逆行']
	Direction int64 `json:"Direction,omitempty"`

	// Array
	//
	// 票價
	// Required: true
	Fares []*PTXServiceDTORailSpecificationV3TRAODFareFare "json:\"Fares\" xml:\"List`1\""

	// String
	//
	// 起點車站代碼
	// Required: true
	OriginStationID *string `json:"OriginStationID" xml:"OriginStationID"`

	// NameType
	//
	// 起點車站名稱
	// Required: true
	OriginStationName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"OriginStationName" xml:"NameType"`

	// Int32
	//
	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	// Required: true
	TrainType *int32 `json:"TrainType"`

	// 起迄站間乘車距離
	TravelDistance float32 `json:"TravelDistance,omitempty"`
}

// Validate validates this p t x service d t o rail specification v3 t r a o d fare o d fare
func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginStationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) validateDestinationStationID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationID", "body", m.DestinationStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) validateDestinationStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) validateFares(formats strfmt.Registry) error {

	if err := validate.Required("Fares", "body", m.Fares); err != nil {
		return err
	}

	for i := 0; i < len(m.Fares); i++ {
		if swag.IsZero(m.Fares[i]) { // not required
			continue
		}

		if m.Fares[i] != nil {
			if err := m.Fares[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) validateOriginStationID(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationID", "body", m.OriginStationID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) validateOriginStationName(formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) validateTrainType(formats strfmt.Registry) error {

	if err := validate.Required("TrainType", "body", m.TrainType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p t x service d t o rail specification v3 t r a o d fare o d fare based on the context it is used
func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDestinationStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFares(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginStationName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) contextValidateDestinationStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) contextValidateFares(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fares); i++ {

		if m.Fares[i] != nil {
			if err := m.Fares[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Fares" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Fares" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) contextValidateOriginStationName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTORailSpecificationV3TRAODFareODFare) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTORailSpecificationV3TRAODFareODFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
