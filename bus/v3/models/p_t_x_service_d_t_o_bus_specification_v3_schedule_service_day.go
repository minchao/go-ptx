// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PTXServiceDTOBusSpecificationV3ScheduleServiceDay ServiceDay
//
// swagger:model PTX.Service.DTO.Bus.Specification.V3.Schedule+ServiceDay
type PTXServiceDTOBusSpecificationV3ScheduleServiceDay struct {

	// integer
	//
	// 國定假日後一日營運與否 : [0:'否',1:'是']
	DayAfterHoliday int32 `json:"DayAfterHoliday,omitempty"`

	// integer
	//
	// 國定假日前一日營運與否 : [0:'否',1:'是']
	DayBeforeHoliday int32 `json:"DayBeforeHoliday,omitempty"`

	// integer
	//
	// 星期五營運與否 : [0:'否',1:'是']
	Friday int32 `json:"Friday,omitempty"`

	// integer
	//
	// 星期一營運與否 : [0:'否',1:'是']
	Monday int32 `json:"Monday,omitempty"`

	// integer
	//
	// 國定假日營運與否 : [0:'否',1:'是']
	NationalHolidays int32 `json:"NationalHolidays,omitempty"`

	// integer
	//
	// 星期六營運與否 : [0:'否',1:'是']
	Saturday int32 `json:"Saturday,omitempty"`

	// String
	//
	// 服務日標籤
	ServiceTag string `json:"ServiceTag,omitempty"`

	// integer
	//
	// 星期日營運與否 : [0:'否',1:'是']
	Sunday int32 `json:"Sunday,omitempty"`

	// integer
	//
	// 星期四營運與否 : [0:'否',1:'是']
	Thursday int32 `json:"Thursday,omitempty"`

	// integer
	//
	// 星期二營運與否 : [0:'否',1:'是']
	Tuesday int32 `json:"Tuesday,omitempty"`

	// integer
	//
	// 颱風停止上班上課期間營運與否 : [0:'否',1:'是']
	TyphoonDay int32 `json:"TyphoonDay,omitempty"`

	// integer
	//
	// 星期三營運與否 : [0:'否',1:'是']
	Wednesday int32 `json:"Wednesday,omitempty"`
}

// Validate validates this p t x service d t o bus specification v3 schedule service day
func (m *PTXServiceDTOBusSpecificationV3ScheduleServiceDay) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p t x service d t o bus specification v3 schedule service day based on context it is used
func (m *PTXServiceDTOBusSpecificationV3ScheduleServiceDay) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3ScheduleServiceDay) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOBusSpecificationV3ScheduleServiceDay) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOBusSpecificationV3ScheduleServiceDay
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
