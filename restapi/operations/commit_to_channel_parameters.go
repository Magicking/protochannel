// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Magicking/protochannel/models"
)

// NewCommitToChannelParams creates a new CommitToChannelParams object
// with the default values initialized.
func NewCommitToChannelParams() CommitToChannelParams {
	var ()
	return CommitToChannelParams{}
}

// CommitToChannelParams contains all the bound params for the commit to channel operation
// typically these are obtained from a http.Request
//
// swagger:parameters commit_to_channel
type CommitToChannelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: body
	*/
	Message *models.Message
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CommitToChannelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Message
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("message", "body"))
			} else {
				res = append(res, errors.NewParseError("message", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Message = &body
			}
		}

	} else {
		res = append(res, errors.Required("message", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}