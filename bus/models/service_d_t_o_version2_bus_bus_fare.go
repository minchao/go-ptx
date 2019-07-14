// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceDTOVersion2BusBusFare BusFare
//
// 票票價種類及費率說明
// swagger:model Service.DTO.Version2.Bus.BusFare
type ServiceDTOVersion2BusBusFare struct {

	// 優惠時段
	DiscountPeriods []*ServiceDTOVersion2BusBusDiscountPeriods `json:"DiscountPeriods"`

	// 費率等級
	// Required: true
	// Enum: [1: 成人 2: 學生 3: 孩童 4: 敬老 5: 愛心 6: 愛心孩童 7: 愛心優待或愛心陪伴 8: 團體 9: 軍警 10: 由各運業者自行定義的半票]
	FareClass *string `json:"FareClass"`

	// 票價名稱
	FareName string `json:"FareName,omitempty"`

	// 計費價格(新臺幣)
	// Required: true
	Price *int32 `json:"Price"`

	// 票種類型
	// Required: true
	// Enum: [1: 一般票 2: 來回票 3: 電子票證 4: 回數票 5: 定期票30天期 6: 定期票60天期 7: 早鳥票 8: 定期票90天期]
	TicketType *string `json:"TicketType"`
}

// Validate validates this service d t o version2 bus bus fare
func (m *ServiceDTOVersion2BusBusFare) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDiscountPeriods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFareClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicketType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusFare) validateDiscountPeriods(formats strfmt.Registry) error {

	if swag.IsZero(m.DiscountPeriods) { // not required
		return nil
	}

	for i := 0; i < len(m.DiscountPeriods); i++ {
		if swag.IsZero(m.DiscountPeriods[i]) { // not required
			continue
		}

		if m.DiscountPeriods[i] != nil {
			if err := m.DiscountPeriods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DiscountPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var serviceDTOVersion2BusBusFareTypeFareClassPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1: 成人","2: 學生","3: 孩童","4: 敬老","5: 愛心","6: 愛心孩童","7: 愛心優待或愛心陪伴","8: 團體","9: 軍警","10: 由各運業者自行定義的半票"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusFareTypeFareClassPropEnum = append(serviceDTOVersion2BusBusFareTypeFareClassPropEnum, v)
	}
}

const (

	// ServiceDTOVersion2BusBusFareFareClassNr1成人 captures enum value "1: 成人"
	ServiceDTOVersion2BusBusFareFareClassNr1成人 string = "1: 成人"

	// ServiceDTOVersion2BusBusFareFareClassNr2學生 captures enum value "2: 學生"
	ServiceDTOVersion2BusBusFareFareClassNr2學生 string = "2: 學生"

	// ServiceDTOVersion2BusBusFareFareClassNr3孩童 captures enum value "3: 孩童"
	ServiceDTOVersion2BusBusFareFareClassNr3孩童 string = "3: 孩童"

	// ServiceDTOVersion2BusBusFareFareClassNr4敬老 captures enum value "4: 敬老"
	ServiceDTOVersion2BusBusFareFareClassNr4敬老 string = "4: 敬老"

	// ServiceDTOVersion2BusBusFareFareClassNr5愛心 captures enum value "5: 愛心"
	ServiceDTOVersion2BusBusFareFareClassNr5愛心 string = "5: 愛心"

	// ServiceDTOVersion2BusBusFareFareClassNr6愛心孩童 captures enum value "6: 愛心孩童"
	ServiceDTOVersion2BusBusFareFareClassNr6愛心孩童 string = "6: 愛心孩童"

	// ServiceDTOVersion2BusBusFareFareClassNr7愛心優待或愛心陪伴 captures enum value "7: 愛心優待或愛心陪伴"
	ServiceDTOVersion2BusBusFareFareClassNr7愛心優待或愛心陪伴 string = "7: 愛心優待或愛心陪伴"

	// ServiceDTOVersion2BusBusFareFareClassNr8團體 captures enum value "8: 團體"
	ServiceDTOVersion2BusBusFareFareClassNr8團體 string = "8: 團體"

	// ServiceDTOVersion2BusBusFareFareClassNr9軍警 captures enum value "9: 軍警"
	ServiceDTOVersion2BusBusFareFareClassNr9軍警 string = "9: 軍警"

	// ServiceDTOVersion2BusBusFareFareClassNr10由各運業者自行定義的半票 captures enum value "10: 由各運業者自行定義的半票"
	ServiceDTOVersion2BusBusFareFareClassNr10由各運業者自行定義的半票 string = "10: 由各運業者自行定義的半票"
)

// prop value enum
func (m *ServiceDTOVersion2BusBusFare) validateFareClassEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusFareTypeFareClassPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusFare) validateFareClass(formats strfmt.Registry) error {

	if err := validate.Required("FareClass", "body", m.FareClass); err != nil {
		return err
	}

	// value enum
	if err := m.validateFareClassEnum("FareClass", "body", *m.FareClass); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDTOVersion2BusBusFare) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("Price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

var serviceDTOVersion2BusBusFareTypeTicketTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1: 一般票","2: 來回票","3: 電子票證","4: 回數票","5: 定期票30天期","6: 定期票60天期","7: 早鳥票","8: 定期票90天期"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceDTOVersion2BusBusFareTypeTicketTypePropEnum = append(serviceDTOVersion2BusBusFareTypeTicketTypePropEnum, v)
	}
}

const (

	// ServiceDTOVersion2BusBusFareTicketTypeNr1一般票 captures enum value "1: 一般票"
	ServiceDTOVersion2BusBusFareTicketTypeNr1一般票 string = "1: 一般票"

	// ServiceDTOVersion2BusBusFareTicketTypeNr2來回票 captures enum value "2: 來回票"
	ServiceDTOVersion2BusBusFareTicketTypeNr2來回票 string = "2: 來回票"

	// ServiceDTOVersion2BusBusFareTicketTypeNr3電子票證 captures enum value "3: 電子票證"
	ServiceDTOVersion2BusBusFareTicketTypeNr3電子票證 string = "3: 電子票證"

	// ServiceDTOVersion2BusBusFareTicketTypeNr4回數票 captures enum value "4: 回數票"
	ServiceDTOVersion2BusBusFareTicketTypeNr4回數票 string = "4: 回數票"

	// ServiceDTOVersion2BusBusFareTicketTypeNr5定期票30天期 captures enum value "5: 定期票30天期"
	ServiceDTOVersion2BusBusFareTicketTypeNr5定期票30天期 string = "5: 定期票30天期"

	// ServiceDTOVersion2BusBusFareTicketTypeNr6定期票60天期 captures enum value "6: 定期票60天期"
	ServiceDTOVersion2BusBusFareTicketTypeNr6定期票60天期 string = "6: 定期票60天期"

	// ServiceDTOVersion2BusBusFareTicketTypeNr7早鳥票 captures enum value "7: 早鳥票"
	ServiceDTOVersion2BusBusFareTicketTypeNr7早鳥票 string = "7: 早鳥票"

	// ServiceDTOVersion2BusBusFareTicketTypeNr8定期票90天期 captures enum value "8: 定期票90天期"
	ServiceDTOVersion2BusBusFareTicketTypeNr8定期票90天期 string = "8: 定期票90天期"
)

// prop value enum
func (m *ServiceDTOVersion2BusBusFare) validateTicketTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, serviceDTOVersion2BusBusFareTypeTicketTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceDTOVersion2BusBusFare) validateTicketType(formats strfmt.Registry) error {

	if err := validate.Required("TicketType", "body", m.TicketType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTicketTypeEnum("TicketType", "body", *m.TicketType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusFare) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDTOVersion2BusBusFare) UnmarshalBinary(b []byte) error {
	var res ServiceDTOVersion2BusBusFare
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
