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

// PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat AvailableSeatStatus
//
// swagger:model PTX.API.Rail.Model.V2THSRAvailableSeatStatusOldWrapper[PTX.Service.DTO.Rail.Specification.V2.THSR.Old.AvailableSeat]
type PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat struct {

	// Array
	//
	// 對號座位狀態資訊(依高鐵規定若營運狀態有異常狀況時，剩餘座位資訊將停留在最後正常運行時間之狀態不做更新，實際資料請參考高鐵各車站現場對號座剩餘座位資訊看板)
	// Required: true
	AvailableSeats []*PTXServiceDTORailSpecificationV2THSROldAvailableSeat "json:\"AvailableSeats\" xml:\"List`1\""

	// 資料總筆數
	Count int64 `json:"Count,omitempty"`

	// DateTime
	//
	// 更新日期時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	UpdateTime string `json:"UpdateTime,omitempty"`
}

// Validate validates this p t x API rail model v2 t h s r available seat status old wrapper p t x service d t o rail specification v2 t h s r old available seat
func (m *PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableSeats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat) validateAvailableSeats(formats strfmt.Registry) error {

	if err := validate.Required("AvailableSeats", "body", m.AvailableSeats); err != nil {
		return err
	}

	for i := 0; i < len(m.AvailableSeats); i++ {
		if swag.IsZero(m.AvailableSeats[i]) { // not required
			continue
		}

		if m.AvailableSeats[i] != nil {
			if err := m.AvailableSeats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AvailableSeats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p t x API rail model v2 t h s r available seat status old wrapper p t x service d t o rail specification v2 t h s r old available seat based on the context it is used
func (m *PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableSeats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat) contextValidateAvailableSeats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableSeats); i++ {

		if m.AvailableSeats[i] != nil {
			if err := m.AvailableSeats[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AvailableSeats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat) UnmarshalBinary(b []byte) error {
	var res PTXAPIRailModelV2THSRAvailableSeatStatusOldWrapperPTXServiceDTORailSpecificationV2THSROldAvailableSeat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
