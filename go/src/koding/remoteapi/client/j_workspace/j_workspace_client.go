package j_workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j workspace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j workspace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJWorkspaceCreate post remote API j workspace create API
*/
func (a *Client) PostRemoteAPIJWorkspaceCreate(params *PostRemoteAPIJWorkspaceCreateParams) (*PostRemoteAPIJWorkspaceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceCreate",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceCreateReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceCreateOK), nil

}

/*
PostRemoteAPIJWorkspaceCreateDefault post remote API j workspace create default API
*/
func (a *Client) PostRemoteAPIJWorkspaceCreateDefault(params *PostRemoteAPIJWorkspaceCreateDefaultParams) (*PostRemoteAPIJWorkspaceCreateDefaultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceCreateDefaultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceCreateDefault",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.createDefault",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceCreateDefaultReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceCreateDefaultOK), nil

}

/*
PostRemoteAPIJWorkspaceDeleteByID post remote API j workspace delete by ID API
*/
func (a *Client) PostRemoteAPIJWorkspaceDeleteByID(params *PostRemoteAPIJWorkspaceDeleteByIDParams) (*PostRemoteAPIJWorkspaceDeleteByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceDeleteByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceDeleteByID",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.deleteById",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceDeleteByIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceDeleteByIDOK), nil

}

/*
PostRemoteAPIJWorkspaceDeleteByUID post remote API j workspace delete by UID API
*/
func (a *Client) PostRemoteAPIJWorkspaceDeleteByUID(params *PostRemoteAPIJWorkspaceDeleteByUIDParams) (*PostRemoteAPIJWorkspaceDeleteByUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceDeleteByUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceDeleteByUID",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.deleteByUid",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceDeleteByUIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceDeleteByUIDOK), nil

}

/*
PostRemoteAPIJWorkspaceDeleteID post remote API j workspace delete ID API
*/
func (a *Client) PostRemoteAPIJWorkspaceDeleteID(params *PostRemoteAPIJWorkspaceDeleteIDParams) (*PostRemoteAPIJWorkspaceDeleteIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceDeleteIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceDeleteID",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.delete/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceDeleteIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceDeleteIDOK), nil

}

/*
PostRemoteAPIJWorkspaceFetchByMachines post remote API j workspace fetch by machines API
*/
func (a *Client) PostRemoteAPIJWorkspaceFetchByMachines(params *PostRemoteAPIJWorkspaceFetchByMachinesParams) (*PostRemoteAPIJWorkspaceFetchByMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceFetchByMachinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceFetchByMachines",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.fetchByMachines",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceFetchByMachinesReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceFetchByMachinesOK), nil

}

/*
PostRemoteAPIJWorkspaceUpdate post remote API j workspace update API
*/
func (a *Client) PostRemoteAPIJWorkspaceUpdate(params *PostRemoteAPIJWorkspaceUpdateParams) (*PostRemoteAPIJWorkspaceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJWorkspaceUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJWorkspaceUpdate",
		Method:             "POST",
		PathPattern:        "/remote.api/JWorkspace.update",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJWorkspaceUpdateReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJWorkspaceUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
