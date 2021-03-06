// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EndpointStatusChange Indication of a change of status
// swagger:model EndpointStatusChange

type EndpointStatusChange struct {

	// Code indicate type of status change
	Code string `json:"code,omitempty"`

	// Status message
	Message string `json:"message,omitempty"`

	// Timestamp when status change occurred
	Timestamp string `json:"timestamp,omitempty"`
}

/* polymorph EndpointStatusChange code false */

/* polymorph EndpointStatusChange message false */

/* polymorph EndpointStatusChange timestamp false */

// Validate validates this endpoint status change
func (m *EndpointStatusChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var endpointStatusChangeTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointStatusChangeTypeCodePropEnum = append(endpointStatusChangeTypeCodePropEnum, v)
	}
}

const (
	// EndpointStatusChangeCodeOk captures enum value "ok"
	EndpointStatusChangeCodeOk string = "ok"
	// EndpointStatusChangeCodeFailed captures enum value "failed"
	EndpointStatusChangeCodeFailed string = "failed"
)

// prop value enum
func (m *EndpointStatusChange) validateCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, endpointStatusChangeTypeCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EndpointStatusChange) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointStatusChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointStatusChange) UnmarshalBinary(b []byte) error {
	var res EndpointStatusChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
