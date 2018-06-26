// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIPublicV1ProfileUsernamePoliciesStepParams creates a new GetAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized.
func NewGetAPIPublicV1ProfileUsernamePoliciesStepParams() *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &GetAPIPublicV1ProfileUsernamePoliciesStepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepParamsWithTimeout creates a new GetAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIPublicV1ProfileUsernamePoliciesStepParamsWithTimeout(timeout time.Duration) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &GetAPIPublicV1ProfileUsernamePoliciesStepParams{

		timeout: timeout,
	}
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepParamsWithContext creates a new GetAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIPublicV1ProfileUsernamePoliciesStepParamsWithContext(ctx context.Context) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &GetAPIPublicV1ProfileUsernamePoliciesStepParams{

		Context: ctx,
	}
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepParamsWithHTTPClient creates a new GetAPIPublicV1ProfileUsernamePoliciesStepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIPublicV1ProfileUsernamePoliciesStepParamsWithHTTPClient(client *http.Client) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	var ()
	return &GetAPIPublicV1ProfileUsernamePoliciesStepParams{
		HTTPClient: client,
	}
}

/*GetAPIPublicV1ProfileUsernamePoliciesStepParams contains all the parameters to send to the API endpoint
for the get API public v1 profile username policies step operation typically these are written to a http.Request
*/
type GetAPIPublicV1ProfileUsernamePoliciesStepParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Step
	  Paging policy step

	*/
	Step float64
	/*Username
	  Your username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithTimeout(timeout time.Duration) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithContext(ctx context.Context) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithHTTPClient(client *http.Client) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithXVOAPIID(xVOAPIID string) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithXVOAPIKey(xVOAPIKey string) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithStep adds the step to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithStep(step float64) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetStep(step float64) {
	o.Step = step
}

// WithUsername adds the username to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WithUsername(username string) *GetAPIPublicV1ProfileUsernamePoliciesStepParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get API public v1 profile username policies step params
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-VO-Api-Id
	if err := r.SetHeaderParam("X-VO-Api-Id", o.XVOAPIID); err != nil {
		return err
	}

	// header param X-VO-Api-Key
	if err := r.SetHeaderParam("X-VO-Api-Key", o.XVOAPIKey); err != nil {
		return err
	}

	// path param step
	if err := r.SetPathParam("step", swag.FormatFloat64(o.Step)); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}