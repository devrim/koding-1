package j_machine

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

// NewPostRemoteAPIJMachineUnshareIDParams creates a new PostRemoteAPIJMachineUnshareIDParams object
// with the default values initialized.
func NewPostRemoteAPIJMachineUnshareIDParams() *PostRemoteAPIJMachineUnshareIDParams {
	var ()
	return &PostRemoteAPIJMachineUnshareIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJMachineUnshareIDParamsWithTimeout creates a new PostRemoteAPIJMachineUnshareIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJMachineUnshareIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJMachineUnshareIDParams {
	var ()
	return &PostRemoteAPIJMachineUnshareIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJMachineUnshareIDParamsWithContext creates a new PostRemoteAPIJMachineUnshareIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJMachineUnshareIDParamsWithContext(ctx context.Context) *PostRemoteAPIJMachineUnshareIDParams {
	var ()
	return &PostRemoteAPIJMachineUnshareIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJMachineUnshareIDParams contains all the parameters to send to the API endpoint
for the post remote API j machine unshare ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJMachineUnshareIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j machine unshare ID params
func (o *PostRemoteAPIJMachineUnshareIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJMachineUnshareIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j machine unshare ID params
func (o *PostRemoteAPIJMachineUnshareIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j machine unshare ID params
func (o *PostRemoteAPIJMachineUnshareIDParams) WithContext(ctx context.Context) *PostRemoteAPIJMachineUnshareIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j machine unshare ID params
func (o *PostRemoteAPIJMachineUnshareIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j machine unshare ID params
func (o *PostRemoteAPIJMachineUnshareIDParams) WithID(id string) *PostRemoteAPIJMachineUnshareIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j machine unshare ID params
func (o *PostRemoteAPIJMachineUnshareIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJMachineUnshareIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
