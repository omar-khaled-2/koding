package j_compute_stack

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

// NewPostRemoteAPIJComputeStackDeleteAdminMessageIDParams creates a new PostRemoteAPIJComputeStackDeleteAdminMessageIDParams object
// with the default values initialized.
func NewPostRemoteAPIJComputeStackDeleteAdminMessageIDParams() *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams {
	var ()
	return &PostRemoteAPIJComputeStackDeleteAdminMessageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJComputeStackDeleteAdminMessageIDParamsWithTimeout creates a new PostRemoteAPIJComputeStackDeleteAdminMessageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJComputeStackDeleteAdminMessageIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams {
	var ()
	return &PostRemoteAPIJComputeStackDeleteAdminMessageIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJComputeStackDeleteAdminMessageIDParamsWithContext creates a new PostRemoteAPIJComputeStackDeleteAdminMessageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJComputeStackDeleteAdminMessageIDParamsWithContext(ctx context.Context) *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams {
	var ()
	return &PostRemoteAPIJComputeStackDeleteAdminMessageIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJComputeStackDeleteAdminMessageIDParams contains all the parameters to send to the API endpoint
for the post remote API j compute stack delete admin message ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJComputeStackDeleteAdminMessageIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j compute stack delete admin message ID params
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j compute stack delete admin message ID params
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j compute stack delete admin message ID params
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) WithContext(ctx context.Context) *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j compute stack delete admin message ID params
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j compute stack delete admin message ID params
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) WithID(id string) *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j compute stack delete admin message ID params
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJComputeStackDeleteAdminMessageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
