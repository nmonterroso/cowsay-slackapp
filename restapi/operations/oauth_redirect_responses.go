package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/nmonterroso/cowsay-slackapp/models"
)

/*OauthRedirectOK proper processing of an oauth accept/deny

swagger:response oauthRedirectOK
*/
type OauthRedirectOK struct {

	// In: body
	Payload *models.OauthComplete `json:"body,omitempty"`
}

// NewOauthRedirectOK creates OauthRedirectOK with default headers values
func NewOauthRedirectOK() *OauthRedirectOK {
	return &OauthRedirectOK{}
}

// WithPayload adds the payload to the oauth redirect o k response
func (o *OauthRedirectOK) WithPayload(payload *models.OauthComplete) *OauthRedirectOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *OauthRedirectOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*OauthRedirectDefault Unexpected error

swagger:response oauthRedirectDefault
*/
type OauthRedirectDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewOauthRedirectDefault creates OauthRedirectDefault with default headers values
func NewOauthRedirectDefault(code int) *OauthRedirectDefault {
	if code <= 0 {
		code = 500
	}

	return &OauthRedirectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the oauth redirect default response
func (o *OauthRedirectDefault) WithStatusCode(code int) *OauthRedirectDefault {
	o._statusCode = code
	return o
}

// WithPayload adds the payload to the oauth redirect default response
func (o *OauthRedirectDefault) WithPayload(payload *models.Error) *OauthRedirectDefault {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *OauthRedirectDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
