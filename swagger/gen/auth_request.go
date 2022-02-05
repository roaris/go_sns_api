// Code generated by go-swagger; DO NOT EDIT.

package gen

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthRequest AuthRequest
//
// swagger:model AuthRequest
type AuthRequest struct {

	// email
	Email string `json:"email,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this auth request
func (m *AuthRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auth request based on context it is used
func (m *AuthRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthRequest) UnmarshalBinary(b []byte) error {
	var res AuthRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
