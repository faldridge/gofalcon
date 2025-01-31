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

// FwmgrAPIFirewallFieldsV1 fwmgr api firewall fields v1
//
// swagger:model fwmgr.api.FirewallFieldsV1
type FwmgrAPIFirewallFieldsV1 struct {

	// default monitor
	// Required: true
	DefaultMonitor *FwmgrDomainMonitoring `json:"default_monitor"`

	// id
	// Required: true
	ID *string `json:"id"`

	// platform
	// Required: true
	Platform *string `json:"platform"`

	// platform fields
	// Required: true
	PlatformFields []*FwmgrDomainField `json:"platform_fields"`
}

// Validate validates this fwmgr api firewall fields v1
func (m *FwmgrAPIFirewallFieldsV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrAPIFirewallFieldsV1) validateDefaultMonitor(formats strfmt.Registry) error {

	if err := validate.Required("default_monitor", "body", m.DefaultMonitor); err != nil {
		return err
	}

	if m.DefaultMonitor != nil {
		if err := m.DefaultMonitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_monitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_monitor")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPIFirewallFieldsV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPIFirewallFieldsV1) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrAPIFirewallFieldsV1) validatePlatformFields(formats strfmt.Registry) error {

	if err := validate.Required("platform_fields", "body", m.PlatformFields); err != nil {
		return err
	}

	for i := 0; i < len(m.PlatformFields); i++ {
		if swag.IsZero(m.PlatformFields[i]) { // not required
			continue
		}

		if m.PlatformFields[i] != nil {
			if err := m.PlatformFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platform_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("platform_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fwmgr api firewall fields v1 based on the context it is used
func (m *FwmgrAPIFirewallFieldsV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDefaultMonitor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatformFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrAPIFirewallFieldsV1) contextValidateDefaultMonitor(ctx context.Context, formats strfmt.Registry) error {

	if m.DefaultMonitor != nil {
		if err := m.DefaultMonitor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("default_monitor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("default_monitor")
			}
			return err
		}
	}

	return nil
}

func (m *FwmgrAPIFirewallFieldsV1) contextValidatePlatformFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlatformFields); i++ {

		if m.PlatformFields[i] != nil {
			if err := m.PlatformFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platform_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("platform_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrAPIFirewallFieldsV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrAPIFirewallFieldsV1) UnmarshalBinary(b []byte) error {
	var res FwmgrAPIFirewallFieldsV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
