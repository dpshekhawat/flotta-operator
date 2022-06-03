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
)

// HardwareInfo Hardware information
//
// swagger:model hardware-info
type HardwareInfo struct {

	// boot
	Boot *Boot `json:"boot,omitempty"`

	// cpu
	CPU *CPU `json:"cpu,omitempty"`

	// disks
	Disks []*Disk `json:"disks"`

	// gpus
	Gpus []*Gpu `json:"gpus"`

	// host devices
	HostDevices []*HostDevice `json:"host_devices"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// interfaces
	Interfaces []*Interface `json:"interfaces"`

	// memory
	Memory *Memory `json:"memory,omitempty"`

	// system vendor
	SystemVendor *SystemVendor `json:"system_vendor,omitempty"`
}

// Validate validates this hardware info
func (m *HardwareInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemVendor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HardwareInfo) validateBoot(formats strfmt.Registry) error {
	if swag.IsZero(m.Boot) { // not required
		return nil
	}

	if m.Boot != nil {
		if err := m.Boot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("boot")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfo) validateCPU(formats strfmt.Registry) error {
	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {
		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfo) validateDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) validateGpus(formats strfmt.Registry) error {
	if swag.IsZero(m.Gpus) { // not required
		return nil
	}

	for i := 0; i < len(m.Gpus); i++ {
		if swag.IsZero(m.Gpus[i]) { // not required
			continue
		}

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) validateHostDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.HostDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.HostDevices); i++ {
		if swag.IsZero(m.HostDevices[i]) { // not required
			continue
		}

		if m.HostDevices[i] != nil {
			if err := m.HostDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("host_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) validateMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {
		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfo) validateSystemVendor(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemVendor) { // not required
		return nil
	}

	if m.SystemVendor != nil {
		if err := m.SystemVendor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_vendor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system_vendor")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hardware info based on the context it is used
func (m *HardwareInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCPU(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGpus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemVendor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HardwareInfo) contextValidateBoot(ctx context.Context, formats strfmt.Registry) error {

	if m.Boot != nil {
		if err := m.Boot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("boot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("boot")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfo) contextValidateCPU(ctx context.Context, formats strfmt.Registry) error {

	if m.CPU != nil {
		if err := m.CPU.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfo) contextValidateDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Disks); i++ {

		if m.Disks[i] != nil {
			if err := m.Disks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) contextValidateGpus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Gpus); i++ {

		if m.Gpus[i] != nil {
			if err := m.Gpus[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("gpus" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("gpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HardwareInfo) contextValidateMemory(ctx context.Context, formats strfmt.Registry) error {

	if m.Memory != nil {
		if err := m.Memory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

func (m *HardwareInfo) contextValidateSystemVendor(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemVendor != nil {
		if err := m.SystemVendor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system_vendor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("system_vendor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HardwareInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwareInfo) UnmarshalBinary(b []byte) error {
	var res HardwareInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
