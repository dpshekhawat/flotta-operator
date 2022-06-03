// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistrationInfo registration info
//
// swagger:model registrationInfo
type RegistrationInfo struct {

	// Certificate Signing Request to be signed by flotta-operator CA
	CertificateRequest string `json:"certificate_request,omitempty"`

	// hardware
	Hardware *HardwareInfo `json:"hardware,omitempty"`
}

// Validate validates this registration info
func (m *RegistrationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHardware(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationInfo) validateHardware(formats strfmt.Registry) error {
	if swag.IsZero(m.Hardware) { // not required
		return nil
	}

	if m.Hardware != nil {
		if err := m.Hardware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardware")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registration info based on the context it is used
func (m *RegistrationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHardware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationInfo) contextValidateHardware(ctx context.Context, formats strfmt.Registry) error {

	if m.Hardware != nil {
		if err := m.Hardware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hardware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hardware")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationInfo) UnmarshalBinary(b []byte) error {
	var res RegistrationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
