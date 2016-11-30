package j_account

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

// NewPostRemoteAPIJAccountVerifyEmailByUsernameParams creates a new PostRemoteAPIJAccountVerifyEmailByUsernameParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountVerifyEmailByUsernameParams() *PostRemoteAPIJAccountVerifyEmailByUsernameParams {
	var ()
	return &PostRemoteAPIJAccountVerifyEmailByUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountVerifyEmailByUsernameParamsWithTimeout creates a new PostRemoteAPIJAccountVerifyEmailByUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountVerifyEmailByUsernameParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountVerifyEmailByUsernameParams {
	var ()
	return &PostRemoteAPIJAccountVerifyEmailByUsernameParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountVerifyEmailByUsernameParamsWithContext creates a new PostRemoteAPIJAccountVerifyEmailByUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountVerifyEmailByUsernameParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountVerifyEmailByUsernameParams {
	var ()
	return &PostRemoteAPIJAccountVerifyEmailByUsernameParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountVerifyEmailByUsernameParams contains all the parameters to send to the API endpoint
for the post remote API j account verify email by username operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountVerifyEmailByUsernameParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j account verify email by username params
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountVerifyEmailByUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account verify email by username params
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account verify email by username params
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountVerifyEmailByUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account verify email by username params
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j account verify email by username params
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJAccountVerifyEmailByUsernameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j account verify email by username params
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountVerifyEmailByUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
