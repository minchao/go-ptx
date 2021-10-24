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

// PTXServiceDTOShipSpecificationV3LivePosition LivePosition
//
// swagger:model PTX.Service.DTO.Ship.Specification.V3.LivePosition
type PTXServiceDTOShipSpecificationV3LivePosition struct {

	// Single
	//
	// 對地航向(Course Over Ground，單位: degree(度，真北為零))，精度至 1/10 degree
	COG float32 `json:"COG,omitempty"`

	// DateTime
	//
	// GPS 記錄時間(ISO8601格式:yyyy-MM-ddTHH:mm:sszzz)
	// Required: true
	// Format: date-time
	GPSTime *strfmt.DateTime `json:"GPSTime"`

	// String
	//
	// AIS 船舶編號
	// Required: true
	MMSI *string `json:"MMSI" xml:"String"`

	// Int32
	//
	// 航行狀態 : [0:'under way using engine',1:'at anchor',2:'not under command',3:'restricted maneuverability',4:'constrained by her draught',5:'moored',6:'aground',7:'engaged in fishing',8:'under way sailing',9:'reserved for future amendment of navigational status for ships carrying DG, HS, or MP, or IMO hazard or pollutant category C, high speed craft (HSC)',10:'reserved for future amendment of navigational status for ships carrying dangerous goods (DG), harmful substances (HS) or marine pollutants (MP), or IMO hazard or pollutant category A, wing in ground (WIG)',11:'power-driven vessel towing astern (regional use)',12:'power-driven vessel pushing ahead or towing alongside (regional use)',13:'reserved for future use',14:'AIS-SART (active), MOB-AIS, EPIRB-AIS',15:'undefined = default (also used by AIS-SART, MOB-AIS and EPIRB-AIS under test)']
	// Required: true
	NAVSTAT *int64 `json:"NAVSTAT"`

	// Single
	//
	// 對地航速(Speed Over Ground，單位: knot(節，海里/小時))，精度至 1/10 knot
	SOG float32 `json:"SOG,omitempty"`

	// String
	//
	// 船舶代碼
	// Required: true
	VesselID *string `json:"VesselID" xml:"String"`

	// NameType
	//
	// 船舶名稱
	VesselName struct {
		PTXServiceDTOSharedSpecificationV3BaseNameType
	} `json:"VesselName,omitempty" xml:"NameType,omitempty"`

	// PointType
	//
	// 船舶位置資訊
	// Required: true
	VesselPosition struct {
		PTXServiceDTOShipSpecificationV3PointType
	} `json:"VesselPosition" xml:"PointType"`

	// Int32
	//
	// 船舶種類 : [0:'N/A',20:'WIG',30:'Vessel-Fishing',31:'not under command',32:'Vessel-Tow>200m,breadth>25m',33:'Vessel-Dredge,Underwater OP',34:'Vessel-Diving OP',35:'Vessel-Military OP',36:'Vessel-Sailing',37:'Vessel-Pleasure craft',40:'HSC',50:'Pilot vessel',51:'Search & rescue vessel',52:'Tug',53:'Port tender',54:'Anti-pollution vessel',55:'Law enforcement vessel',58:'Medical transport',59:'Resolution 18(Mob-83)',60:'Passenger ship',70:'Cargo ship',80:'Tanker',90:'Other',101:'Undefined']
	VesselType int64 `json:"VesselType,omitempty"`
}

// Validate validates this p t x service d t o ship specification v3 live position
func (m *PTXServiceDTOShipSpecificationV3LivePosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGPSTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMMSI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNAVSTAT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVesselID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVesselName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVesselPosition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) validateGPSTime(formats strfmt.Registry) error {

	if err := validate.Required("GPSTime", "body", m.GPSTime); err != nil {
		return err
	}

	if err := validate.FormatOf("GPSTime", "body", "date-time", m.GPSTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) validateMMSI(formats strfmt.Registry) error {

	if err := validate.Required("MMSI", "body", m.MMSI); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) validateNAVSTAT(formats strfmt.Registry) error {

	if err := validate.Required("NAVSTAT", "body", m.NAVSTAT); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) validateVesselID(formats strfmt.Registry) error {

	if err := validate.Required("VesselID", "body", m.VesselID); err != nil {
		return err
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) validateVesselName(formats strfmt.Registry) error {
	if swag.IsZero(m.VesselName) { // not required
		return nil
	}

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) validateVesselPosition(formats strfmt.Registry) error {

	return nil
}

// ContextValidate validate this p t x service d t o ship specification v3 live position based on the context it is used
func (m *PTXServiceDTOShipSpecificationV3LivePosition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVesselName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVesselPosition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) contextValidateVesselName(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *PTXServiceDTOShipSpecificationV3LivePosition) contextValidateVesselPosition(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3LivePosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PTXServiceDTOShipSpecificationV3LivePosition) UnmarshalBinary(b []byte) error {
	var res PTXServiceDTOShipSpecificationV3LivePosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
