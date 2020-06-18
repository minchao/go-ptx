// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion3RailTRAODFareODFare ODFare
//
// swagger:model Service.DTO.Version3.Rail.TRA.ODFare.ODFare
type ServiceDTOVersion3RailTRAODFareODFare struct {

	// 迄點車站代碼
	// Required: true
	DestinationStationID *string `json:"DestinationStationID"`

	// NameType
	//
	// 迄點車站名稱
	// Required: true
	DestinationStationName *ServiceDTOVersion3BaseNameType `json:"DestinationStationName"`

	// integer
	//
	// 方向 : [0:'順行',1:'逆行']
	Direction int32 `json:"Direction,omitempty"`

	// 票價
	// Required: true
	Fares []*ServiceDTOVersion3RailTRAODFareFare `json:"Fares"`

	// 起點車站代碼
	// Required: true
	OriginStationID *string `json:"OriginStationID"`

	// NameType
	//
	// 起點車站名稱
	// Required: true
	OriginStationName *ServiceDTOVersion3BaseNameType `json:"OriginStationName"`

	// 車種簡碼 = ['1: 太魯閣', '2: 普悠瑪', '3: 自強', '4: 莒光', '5: 復興', '6: 區間', '7: 普快', '10: 區間快']
	// Required: true
	TrainType *int32 `json:"TrainType"`

	// 起迄站間乘車距離
	TravelDistance float32 `json:"TravelDistance,omitempty"`
}

// Validate validates this service d t o version3 rail t r a o d fare o d fare
func (m *ServiceDTOVersion3RailTRAODFareODFare) Validate(formats strfmt.Registry) error {
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

func (m *ServiceDTOVersion3RailTRAODFareODFare) validateDestinationStationID(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationID", "body", m.DestinationStationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAODFareODFare) validateDestinationStationName(formats strfmt.Registry) error {

	if err := validate.Required("DestinationStationName", "body", m.DestinationStationName); err != nil {
		return err
	}

	if m.DestinationStationName != nil {
		if err := m.DestinationStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DestinationStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAODFareODFare) validateFares(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAODFareODFare) validateOriginStationID(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationID", "body", m.OriginStationID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAODFareODFare) validateOriginStationName(formats strfmt.Registry) error {

	if err := validate.Required("OriginStationName", "body", m.OriginStationName); err != nil {
		return err
	}

	if m.OriginStationName != nil {
		if err := m.OriginStationName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OriginStationName")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceDTOVersion3RailTRAODFareODFare) validateTrainType(formats strfmt.Registry) error {

	if err := validate.Required("TrainType", "body", m.TrainType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAODFareODFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion3RailTRAODFareODFare) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion3RailTRAODFareODFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
