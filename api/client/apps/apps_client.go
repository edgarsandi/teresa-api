package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new apps API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for apps API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateApp creates an app

Create an app
*/
func (a *Client) CreateApp(params *CreateAppParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAppCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createApp",
		Method:             "POST",
		PathPattern:        "/teams/{team_id}/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAppReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAppCreated), nil
}

/*
GetAppDetails gets app details

Get app details
*/
func (a *Client) GetAppDetails(params *GetAppDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAppDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAppDetails",
		Method:             "GET",
		PathPattern:        "/teams/{team_id}/apps/{app_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAppDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppDetailsOK), nil
}

/*
GetApps gets team apps

Get team apps
*/
func (a *Client) GetApps(params *GetAppsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAppsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApps",
		Method:             "GET",
		PathPattern:        "/teams/{team_id}/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAppsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsOK), nil
}

/*
UpdateApp updates app

Update app properties, such as number of replicas and other things.
*/
func (a *Client) UpdateApp(params *UpdateAppParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateApp",
		Method:             "PUT",
		PathPattern:        "/teams/{team_id}/apps/{app_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateAppReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAppOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
