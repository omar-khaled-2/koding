package collaboration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPICollaborationStopReader is a Reader for the PostRemoteAPICollaborationStop structure.
type PostRemoteAPICollaborationStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPICollaborationStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPICollaborationStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPICollaborationStopUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPICollaborationStopOK creates a PostRemoteAPICollaborationStopOK with default headers values
func NewPostRemoteAPICollaborationStopOK() *PostRemoteAPICollaborationStopOK {
	return &PostRemoteAPICollaborationStopOK{}
}

/*PostRemoteAPICollaborationStopOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPICollaborationStopOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPICollaborationStopOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Collaboration.stop][%d] postRemoteApiCollaborationStopOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPICollaborationStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPICollaborationStopUnauthorized creates a PostRemoteAPICollaborationStopUnauthorized with default headers values
func NewPostRemoteAPICollaborationStopUnauthorized() *PostRemoteAPICollaborationStopUnauthorized {
	return &PostRemoteAPICollaborationStopUnauthorized{}
}

/*PostRemoteAPICollaborationStopUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPICollaborationStopUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPICollaborationStopUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Collaboration.stop][%d] postRemoteApiCollaborationStopUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPICollaborationStopUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
