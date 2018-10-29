// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReceiptFile receipt file
// swagger:model ReceiptFile
type ReceiptFile struct {

	// date
	Date int64 `json:"date,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// horodatingaddress
	Horodatingaddress string `json:"horodatingaddress,omitempty"`

	// transactionhash
	Transactionhash string `json:"transactionhash,omitempty"`
}

// Validate validates this receipt file
func (m *ReceiptFile) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReceiptFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceiptFile) UnmarshalBinary(b []byte) error {
	var res ReceiptFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
