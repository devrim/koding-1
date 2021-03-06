package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJGroupAddSubscriptionIDParams creates a new PostRemoteAPIJGroupAddSubscriptionIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupAddSubscriptionIDParams() *PostRemoteAPIJGroupAddSubscriptionIDParams {
	var ()
	return &PostRemoteAPIJGroupAddSubscriptionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupAddSubscriptionIDParamsWithTimeout creates a new PostRemoteAPIJGroupAddSubscriptionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupAddSubscriptionIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupAddSubscriptionIDParams {
	var ()
	return &PostRemoteAPIJGroupAddSubscriptionIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupAddSubscriptionIDParamsWithContext creates a new PostRemoteAPIJGroupAddSubscriptionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupAddSubscriptionIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupAddSubscriptionIDParams {
	var ()
	return &PostRemoteAPIJGroupAddSubscriptionIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupAddSubscriptionIDParams contains all the parameters to send to the API endpoint
for the post remote API j group add subscription ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupAddSubscriptionIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j group add subscription ID params
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupAddSubscriptionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group add subscription ID params
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group add subscription ID params
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupAddSubscriptionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group add subscription ID params
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j group add subscription ID params
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) WithID(id string) *PostRemoteAPIJGroupAddSubscriptionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group add subscription ID params
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupAddSubscriptionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
