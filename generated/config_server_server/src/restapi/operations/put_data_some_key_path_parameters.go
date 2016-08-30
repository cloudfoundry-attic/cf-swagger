package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"models"
)

// NewPutDataSomeKeyPathParams creates a new PutDataSomeKeyPathParams object
// with the default values initialized.
func NewPutDataSomeKeyPathParams() PutDataSomeKeyPathParams {
	var ()
	return PutDataSomeKeyPathParams{}
}

// PutDataSomeKeyPathParams contains all the bound params for the put data some key path operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutDataSomeKeyPath
type PutDataSomeKeyPathParams struct {
	/*the data to update
	  Required: true
	  In: body
	*/
	Request *models.Data
	/*the path to the key to retrieve.
	  Required: true
	  In: query
	*/
	SomeKeyPath string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PutDataSomeKeyPathParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	qs := httpkit.Values(r.URL.Query())

	defer r.Body.Close()
	var body models.Data
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		if err == io.EOF {
			res = append(res, errors.Required("request", "body"))
		} else {
			res = append(res, errors.NewParseError("request", "body", "", err))
		}

	} else {
		if err := body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

		if len(res) == 0 {
			o.Request = &body
		}
	}

	qSomeKeyPath, qhkSomeKeyPath, _ := qs.GetOK("someKeyPath")
	if err := o.bindSomeKeyPath(qSomeKeyPath, qhkSomeKeyPath, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutDataSomeKeyPathParams) bindSomeKeyPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("someKeyPath", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("someKeyPath", "query", raw); err != nil {
		return err
	}

	o.SomeKeyPath = raw

	return nil
}
