package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIOAuthGetURLParams creates a new PostRemoteAPIOAuthGetURLParams object
// with the default values initialized.
func NewPostRemoteAPIOAuthGetURLParams() *PostRemoteAPIOAuthGetURLParams {
	var ()
	return &PostRemoteAPIOAuthGetURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIOAuthGetURLParamsWithTimeout creates a new PostRemoteAPIOAuthGetURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIOAuthGetURLParamsWithTimeout(timeout time.Duration) *PostRemoteAPIOAuthGetURLParams {
	var ()
	return &PostRemoteAPIOAuthGetURLParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIOAuthGetURLParamsWithContext creates a new PostRemoteAPIOAuthGetURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIOAuthGetURLParamsWithContext(ctx context.Context) *PostRemoteAPIOAuthGetURLParams {
	var ()
	return &PostRemoteAPIOAuthGetURLParams{

		Context: ctx,
	}
}

/*PostRemoteAPIOAuthGetURLParams contains all the parameters to send to the API endpoint
for the post remote API o auth get URL operation typically these are written to a http.Request
*/
type PostRemoteAPIOAuthGetURLParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API o auth get URL params
func (o *PostRemoteAPIOAuthGetURLParams) WithTimeout(timeout time.Duration) *PostRemoteAPIOAuthGetURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API o auth get URL params
func (o *PostRemoteAPIOAuthGetURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API o auth get URL params
func (o *PostRemoteAPIOAuthGetURLParams) WithContext(ctx context.Context) *PostRemoteAPIOAuthGetURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API o auth get URL params
func (o *PostRemoteAPIOAuthGetURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API o auth get URL params
func (o *PostRemoteAPIOAuthGetURLParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIOAuthGetURLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API o auth get URL params
func (o *PostRemoteAPIOAuthGetURLParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIOAuthGetURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
