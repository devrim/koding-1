package j_user

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

// NewPostRemoteAPIJUserConvertParams creates a new PostRemoteAPIJUserConvertParams object
// with the default values initialized.
func NewPostRemoteAPIJUserConvertParams() *PostRemoteAPIJUserConvertParams {
	var ()
	return &PostRemoteAPIJUserConvertParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJUserConvertParamsWithTimeout creates a new PostRemoteAPIJUserConvertParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJUserConvertParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJUserConvertParams {
	var ()
	return &PostRemoteAPIJUserConvertParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJUserConvertParamsWithContext creates a new PostRemoteAPIJUserConvertParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJUserConvertParamsWithContext(ctx context.Context) *PostRemoteAPIJUserConvertParams {
	var ()
	return &PostRemoteAPIJUserConvertParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJUserConvertParams contains all the parameters to send to the API endpoint
for the post remote API j user convert operation typically these are written to a http.Request
*/
type PostRemoteAPIJUserConvertParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j user convert params
func (o *PostRemoteAPIJUserConvertParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJUserConvertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j user convert params
func (o *PostRemoteAPIJUserConvertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j user convert params
func (o *PostRemoteAPIJUserConvertParams) WithContext(ctx context.Context) *PostRemoteAPIJUserConvertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j user convert params
func (o *PostRemoteAPIJUserConvertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j user convert params
func (o *PostRemoteAPIJUserConvertParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJUserConvertParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j user convert params
func (o *PostRemoteAPIJUserConvertParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJUserConvertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
