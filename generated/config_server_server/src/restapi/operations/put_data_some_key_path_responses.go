package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"models"
)

/*PutDataSomeKeyPathOK OK

swagger:response putDataSomeKeyPathOK
*/
type PutDataSomeKeyPathOK struct {
}

// NewPutDataSomeKeyPathOK creates PutDataSomeKeyPathOK with default headers values
func NewPutDataSomeKeyPathOK() *PutDataSomeKeyPathOK {
	return &PutDataSomeKeyPathOK{}
}

// WriteResponse to the client
func (o *PutDataSomeKeyPathOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
}

/*PutDataSomeKeyPathDefault Unexpected error

swagger:response putDataSomeKeyPathDefault
*/
type PutDataSomeKeyPathDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutDataSomeKeyPathDefault creates PutDataSomeKeyPathDefault with default headers values
func NewPutDataSomeKeyPathDefault(code int) *PutDataSomeKeyPathDefault {
	if code <= 0 {
		code = 500
	}

	return &PutDataSomeKeyPathDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put data some key path default response
func (o *PutDataSomeKeyPathDefault) WithStatusCode(code int) *PutDataSomeKeyPathDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put data some key path default response
func (o *PutDataSomeKeyPathDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put data some key path default response
func (o *PutDataSomeKeyPathDefault) WithPayload(payload *models.Error) *PutDataSomeKeyPathDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put data some key path default response
func (o *PutDataSomeKeyPathDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutDataSomeKeyPathDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
